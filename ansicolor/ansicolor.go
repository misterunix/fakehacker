package main

import (
	"fmt"
	"os"

	"github.com/awesome-gocui/gocui"
)

func main() {

	g, err := gocui.NewGui(gocui.Output256, true)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(1)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := initKeybindings(g); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		fmt.Println(err)
		os.Exit(1)
	}

}

func layout(g *gocui.Gui) error {

	maxX, maxY := g.Size()
	v, err := g.SetView("main", 0, 0, maxX-1, maxY-1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	var s string
	var l int
	v.SetWritePos(0, l)
	for i := 0; i < 256; i++ {

		if (i % 16) == 0 {
			fmt.Fprintln(v, s)
			s = ""
			l++
			v.SetWritePos(0, l)
		}
		s += fmt.Sprintf("\033[48;5;%dm   ", i)
	}
	fmt.Fprintln(v, "\033[48;5;0m")
	return nil
}

func initKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}
	return nil
}
