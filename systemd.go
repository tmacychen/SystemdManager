package main

import (
	"github.com/coreos/go-systemd/dbus"
)

var UnitsNum int

func systemdUnits() []dbus.UnitStatus {
	c, err := dbus.New()
	if err != nil {
		printfLog("%v\n", err)
	}

	units, err := c.ListUnitsByPatterns(nil, []string{"*.service"})
	if err != nil {
		printfLog("%v\n", err)
	}

	UnitsNum = len(units)
	return units
}
