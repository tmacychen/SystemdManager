package main

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

var UnitsNum int

func systemdUnits() []string {

	units := []string{}
	out, err := exec.Command("systemctl", "list-unit-files", "--no-pager").Output()
	if err != nil {
		printfLog("%v\n", err)
	}
	s := bufio.NewScanner(bytes.NewReader(out))
	for s.Scan() {
		units = append(units, strings.Trim(s.Text(), " "))
	}
	//delete the last tow lines
	units = units[:len(units)-2]

	UnitsNum = len(units)
	return units
}
