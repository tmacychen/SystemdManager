package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

var (
	active   = 0
	viewArr  = []string{"v1", "v2", "v3", "v4"}
	itemChan = make(chan string)
)

func init() {
	initLog()
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	initKeyBind(g)
	done := make(chan struct{})
	go journalDaemon(g, done)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	defer close(done)
}
