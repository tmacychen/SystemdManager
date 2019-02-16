package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("v1", 0, 0, maxX-1, maxY/10-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Input Box"
		v.Editable = true
		v.Wrap = true

		if _, err = g.SetCurrentView("v1"); err != nil {
			return err
		}
	}
	if v, err := g.SetView("v2", 0, maxY/10, maxX-1, maxY/2-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Units Status"
		v.Wrap = true
		//v.Autoscroll = true
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorYellow
		for _, i := range systemdUnits() {
			//show service's status  in fix width
			blank := maxX*7/8 - len(i.Name)
			// fix that there is no spacee between the unit's name and status
			if blank <= 5 {
				blank = maxX
			}
			fmt.Fprintf(v, "%s%*s\n", i.Name, blank, i.ActiveState)
		}
	}
	if v, err := g.SetView("v3", 0, maxY/2, maxX/2-1, maxY*9/10-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Units Logs"
		v.Wrap = true
		v.Autoscroll = true
	}
	if v, err := g.SetView("v4", maxX/2, maxY/2, maxX-1, maxY*9/10-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Units Config"
		v.Wrap = true
		v.Highlight = true
		v.Editable = true
	}
	if v, err := g.SetView("v5", 0, maxY*9/10, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Keyboard shortcut Help"
		v.Autoscroll = true
		v.Wrap = true
		fmt.Fprintln(v, "Tab: Switch between the panels")
	}
	return nil
}
