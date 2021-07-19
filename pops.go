package main

import "github.com/awesome-gocui/gocui"

// popup : Create a popup error.
func popup(g *gocui.Gui, name string) error {
	maxX, maxY := g.Size()

	msg := "  " + name + "  "
	msgWidth := len(msg)

	x0 := Roll(1, msgWidth)
	y0 := 0
	x1 := 28
	y1 := 25

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
