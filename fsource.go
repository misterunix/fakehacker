package main

import (
	"fmt"
	"time"

	"github.com/awesome-gocui/gocui"
)

func sourceWindow(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	// 50 wide from the right side half the height
	x0 := maxX - 51
	y0 := 0
	x1 := maxX - 1
	y1 := maxY / 2

	name := "hack1"
	v, err := g.SetView(name, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		//v.Wrap = true
		//v.Autoscroll = true
		//v.FgColor = gocui.ColorGreen

		//fmt.Fprintln(v, strings.Repeat(name+" ", 30))
	}
	if _, err := g.SetCurrentView(name); err != nil {
		return err
	}

	views = append(views, name)
	curView = len(views) - 1
	idxView += 1
	go source1(g, v, name)
	return nil
}

func source1(g *gocui.Gui, v *gocui.View, name string) {

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

	v.Wrap = false
	v.Autoscroll = false

	//v.Overwrite = true
	var lines []string

	mx, my := v.Size()
	//fmt.Fprintln(os.Stderr, mx, my)
	width := mx - 1
	height := my

	ss, _ := readLines("hack1.txt")
	for _, k := range ss {
		y := Chunks(k, width)
		lines = append(lines, y...)
	}

	l := len(lines)
	cp := height

	var lastlines = make([]string, height)

	for i := 0; i < height; i++ {
		v.SetWritePos(0, i)
		lastlines[i] = lines[i]
		//fmt.Fprintf(v, "%-48v", lastlines[i])
		fmt.Fprintf(v, "%v", lastlines[i])
		g.Update(func(g *gocui.Gui) error {
			return nil
		})
		time.Sleep(80 * time.Millisecond)
	}

	for {
		v.Clear()
		//v.SetWritePos(0, li)

		//for viewLine := 0; viewLine < 4; viewLine++ {
		if cp == l {
			cp = 0
		}
		for m := 0; m <= height-2; m++ {
			lastlines[m] = lastlines[m+1]
			v.SetWritePos(0, m)
			fmt.Fprintf(v, "%v", lastlines[m])
		}
		v.SetWritePos(0, height-1)
		lastlines[height-1] = lines[cp]
		fmt.Fprintf(v, "%v", lastlines[height-1])
		cp++

		g.Update(func(g *gocui.Gui) error {
			return nil
		})

		time.Sleep(80 * time.Millisecond)
	}

}
