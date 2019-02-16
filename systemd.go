package main

import (
	"io/ioutil"

	"github.com/coreos/go-systemd/dbus"
	"github.com/coreos/go-systemd/sdjournal"
	"github.com/jroimartin/gocui"
)

var UnitsNum int

func systemdUnits() []dbus.UnitStatus {
	c, err := dbus.New()
	if err != nil {
		printfLog("%v\n", err)
	}

	defer c.Close()
	units, err := c.ListUnitsByPatterns(nil, []string{"*.service"})
	if err != nil {
		printfLog("%v\n", err)
	}

	UnitsNum = len(units)
	return units
}

func journalDaemon(g *gocui.Gui, done <-chan struct{}) {
	subDone := make(chan struct{}, 1)
	isSubRun := false
	for {
		select {
		case <-done:
			return
		default:
			i := <-itemChan
			printfLog("%v\n", i)
			if i != "" && isSubRun {
				subDone <- struct{}{}
			}
			go getServiceStatus(g, i, subDone)
			isSubRun = true
		}
	}

}
func getServiceStatus(g *gocui.Gui, unit string, done chan struct{}) {
	journalReader, err := sdjournal.NewJournalReader(sdjournal.JournalReaderConfig{
		NumFromTail: 10,
		Matches: []sdjournal.Match{
			{
				Field: sdjournal.SD_JOURNAL_FIELD_SYSTEMD_UNIT,
				Value: Item,
			},
		},
	})
	if err != nil {
		printfLog("error open journal %v\n", err)
	}
	if journalReader == nil {
		printfLog("got nil journal reader ! item :%v\n", Item)
	}
	defer journalReader.Close()
	v, err := g.View("v3")
	if err != nil {
		printfLog("get view of v4 error:%v\n", err)
	}
	v.Clear()
	for {
		select {
		case <-done:
			return
		default:
			if err := journalReader.Follow(nil, v); err != sdjournal.ErrExpired {
				printfLog("follow err :%v\n", err)
			}
		}
	}
}
func getServiceFiles(unit string) string {
	if unit == "" {
		return ""
	}
	c, err := dbus.New()
	if err != nil {
		printfLog("%v\n", err)
	}

	defer c.Close()
	m, err := c.GetUnitProperties(unit)
	if err != nil {
		printfLog("%v\n", err)
	}
	var realPath string
	path := m["FragmentPath"]
	if path == nil {
		return ""
	}
	realPath = path.(string)
	fc, err := ioutil.ReadFile(realPath)
	if err != nil {
		printfLog("%v\n", err)
		return ""
	}

	return string(fc)

}
