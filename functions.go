package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
func cursorDown(g *gocui.Gui, v *gocui.View, n int) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy > n-2 {
			return nil
		}
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			//显示最后项目后停止向下滚动
			if oy > n-cy-2 {
				return nil
			}
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}
func nextView(g *gocui.Gui, v *gocui.View) error {

	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]

	printfLog("Going from view " + v.Name() + " to " + name + "\n")

	if _, err := g.SetCurrentView(name); err != nil {
		return err
	}

	if nextIndex == 0 {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

	active = nextIndex
	return nil
}

func dialogItem(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}
	switch l {
	case "Confirm":
		fallthrough
	case "Close":
		if err := g.DeleteView("dialog"); err != nil {
			return nil
		}
	}
	if _, err = g.SetCurrentView("v3"); err != nil {
		return err
	}
	return nil
}

func itemSelect(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}
	printfLog("get %v\n", l)
	maxX, maxY := g.Size()
	if v, err := g.SetView("dialog", maxX/2-30, maxY/2, maxX/2+30, maxY/2+5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Dialog"
		v.Wrap = true
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorYellow
		fmt.Fprintf(v, "%s\n", "Confirm")
		fmt.Fprintf(v, "%s\n", "Close")
		if _, err := g.SetCurrentView("dialog"); err != nil {
			return err
		}
	}
	return nil
}
