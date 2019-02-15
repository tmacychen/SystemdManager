package main

import (
	"fmt"
	"testing"
)

func Test_systemdUnits(t *testing.T) {
	for _, i := range systemdUnits() {
		fmt.Printf("%v\n", i)
	}
}

func Test_getUnitsFile(t *testing.T) {
	getServiceFiles("sshd.service")
}
