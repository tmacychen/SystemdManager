package main

import (
	"github.com/coreos/go-systemd/dbus"
)

var UnitsNum int

func systemdUnits() []string {
	c, err := dbus.New()
	if err != nil {
		printfLog("%v\n", err)
	}

	units, err := c.ListUnitsByPatterns(nil, []string{"*.service"})
	if err != nil {
		printfLog("%v\n", err)
	}
	unitNames := []string{}
	for _, u := range units {
		unitNames = append(unitNames, u.Name)
	}

	UnitsNum = len(units)
	printfLog("num :%v\nunitNmaes :%v\n", UnitsNum, unitNames)
	return unitNames
}
