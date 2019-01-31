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
			printfLog("O:(%v,%v),C:(%v,%v) \n", ox, oy, cx, cy)
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
		printfLog("O:(%v,%v),C:(%v,%v) \n", ox, oy, cx, cy)
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func PageUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, maxY := g.Size()
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		printfLog("O:(%v,%v),C:(%v,%v) \n", ox, oy, cx, cy)
		// 配置光标到第一行
		if err := v.SetCursor(cx, 0); err != nil {
			printfLog("%v\n", err)
		}
		// 向上滚动
		if oy > 0 {
			dy := oy - maxY*4/10 + 1
			if dy < 0 {
				dy = 0
			}
			if err := v.SetOrigin(ox, dy); err != nil {
				return err
			}
		}
	}
	return nil
}

func PageDown(g *gocui.Gui, v *gocui.View, n int) error {
	if v != nil {
		_, maxY := g.Size()
		cx, _ := v.Cursor()
		ox, oy := v.Origin()
		//		printfLog("O:(%v,%v),C:(%v,%v) \n", ox, oy, cx, cy)
		dy := maxY*4/10 - 1
		if dy+oy >= n {
			if err := v.SetOrigin(ox, n-dy); err != nil {
				return err
			}
			if err := v.SetCursor(cx, dy-1); err != nil {
				return err
			}
		} else {
			if err := v.SetOrigin(ox, dy+oy); err != nil {
				return err
			}
		}
	}

	return nil
}

func PageHome(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		if err := v.SetOrigin(0, 0); err != nil {
			return err
		}
		if err := v.SetCursor(0, 0); err != nil {
			return err
		}
	}
	return nil
}

func PageEnd(g *gocui.Gui, v *gocui.View, n int) error {

	if v != nil {
		_, maxY := g.Size()
		cy := maxY*4/10 - 2
		oy := n - maxY*4/10 + 1
		//		printfLog("cy %v,oy %v\n", cy, oy)
		if err := v.SetOrigin(0, oy); err != nil {
			//		printfLog("ori %v\n", err)
			return err
		}
		if err := v.SetCursor(0, cy); err != nil {
			//	printfLog("Cur %v\n", err)
			return err
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
	if _, err = g.SetCurrentView("v2"); err != nil {
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
