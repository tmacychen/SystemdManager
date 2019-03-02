package main

import (
	"io"
	"io/ioutil"
	"time"

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
	printfLog("journal Daemon  start \n")
	isSubRun := false
	doneChan := make(chan struct{}, 1)
	for {
		select {
		case <-done:
			printfLog("Journal daemon done\n")
			return
		case i := <-itemChan:
			printfLog("Get Item :%v\n", i)
			//timer.Reset(time.Duration(200) * time.Microsecond)
			if i != "" && isSubRun {
				isSubRun = false
				doneChan <- struct{}{}
				printfLog("stop the old routine\n")
			}
			printfLog("start to show status \n")
			// get  unit's log
			//go getServiceStatus(g, i, timer)
			printfLog("sub routine is run \n")
			go getServiceStatus(g, i, doneChan)
			isSubRun = true
		}
	}

}
func getServiceStatus(g *gocui.Gui, unit string, done chan struct{}) error {
	v, err := g.View("v3")
	if err != nil {
		printfLog("get view of v3 error:%v\n", err)
		return err
	}
	v.Clear()
	printfLog("v3 clear and tht unit is %v\n", unit)
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
		return err
	}
	if journalReader == nil {
		printfLog("got nil journal reader ! item :%v\n", Item)
		return nil
	}
	defer journalReader.Close()

	buffer := make([]byte, 64*(1<<10))
	for {
		select {
		case <-done:
			printfLog("read sub done!\n")
			return nil
		case <-time.After(time.Duration(100) * time.Millisecond):
			c, err := journalReader.Read(buffer)
			if err != nil && err != io.EOF {
				printfLog("journalReader err:%v\n", err)
				return err
			}
			if c > 0 {
				g.Update(func(g *gocui.Gui) error {
					_, err := v.Write(buffer[:c])
					return err
				})
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
		return ""
	}

	defer c.Close()
	m, err := c.GetUnitProperties(unit)
	if err != nil {
		printfLog("get unitProperties Error :%v\n", err)
		return ""
	}
	var realPath string
	path := m["FragmentPath"]
	if path == nil {
		return "config file does not found!"

	}
	realPath = path.(string)
	fc, err := ioutil.ReadFile(realPath)
	if err != nil {
		printfLog("iounil %v\n", err)
		return "config file read error"
	}

	return string(fc)

}
