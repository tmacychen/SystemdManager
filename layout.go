package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("v1", 0, 0, maxX/2-1, maxY/5-1); err != nil {
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

	if v, err := g.SetView("v2", maxX/2, 0, maxX-1, maxY/2-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Units Logs"
		v.Wrap = true
		v.Autoscroll = true
	}
	if v, err := g.SetView("v3", 0, maxY/5, maxX/2-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Units Status"
		v.Wrap = true
		//v.Autoscroll = true
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorYellow

		for i := 0; i < 50; i++ {
			fmt.Fprintf(v, "Item %v\n", i)
		}

	}
	if v, err := g.SetView("v4", maxX/2, maxY/2, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Units Config"
		v.Autoscroll = true
		v.Wrap = true
	}
	return nil
}
