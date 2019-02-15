package main

import (
	"io/ioutil"

	"github.com/coreos/go-systemd/dbus"
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
func getServiceStatus(unit string) {

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
	path := m["FragmentPath"].(string)
	fc, err := ioutil.ReadFile(path)
	if err != nil {
		printfLog("%v\n", err)
		return ""
	}

	return string(fc)

}
