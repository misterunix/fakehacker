package main

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

// skullWindow : Create the view for the skull.
// Width has to be 85 minumn for the window to render correctly.
func skullWindow(g *gocui.Gui) error {
	//maxX, maxY := g.Size()
	// 50 wide from the right side half the height
	x0 := 0
	y0 := 0
	x1 := 34
	y1 := 25

	name := "skull1"
	v, err := g.SetView(name, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	if _, err := g.SetCurrentView(name); err != nil {
		return err
	}

	views = append(views, name)
	curView = len(views) - 1
	idxView += 1
	go skull1(g, v, name)
	return nil
}

// skull1 : Display the skull with matrix going from left to right.
func skull1(g *gocui.Gui, v *gocui.View, name string) {

	var nameFound = false
	for _, n := range views {
		if n == name {
			nameFound = true
			break
		}
	}
	if !nameFound {
		return
	}
	v.Frame = false
	v.Wrap = false
	v.Autoscroll = false

	// srcraw := readFileToString("bksrc.txt")

	slines, err := readLines("skull.txt")
	if err != nil {
		return
	}

	for i, s := range slines {
		v.SetWritePos(0, i)
		fmt.Fprintf(v, "%s", s)
	}

}
