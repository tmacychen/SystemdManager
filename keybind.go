package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

func initKeyBind(g *gocui.Gui) {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	for _, view := range viewArr {
		if err := g.SetKeybinding(view, gocui.KeyTab, gocui.ModNone,
			func(g *gocui.Gui, v *gocui.View) error {
				return nextView(g, v)
			}); err != nil {
			log.Panicln(err)
		}
	}

	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return cursorUp(g, v)
		}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return cursorDown(g, v)
		}); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("v3", gocui.KeyEnter, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return itemSelect(g, v)
		}); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("dialog", gocui.KeyEnter, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return dialogItem(g, v)
		}); err != nil {
		log.Panicln(err)
	}

}
