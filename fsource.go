package main

import (
	"fakehacker/data"
	"fmt"
	"strings"
	"time"

	"github.com/awesome-gocui/gocui"
)

// sourceWindow : creates the gocui view for the fake source screen
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

// source1 : runs the fake source screen
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

	/*
		ss, _ := readLines("hack1.txt")
		for _, k := range ss {
			y := Chunks(k, width)
			lines = append(lines, y...)
		}
	*/

	lines, err := readGzipLines("hack1.gz")
	//lines, err := readLines("hack1.txt")
	if err != nil {
		return
	}

	if !data.SourceWrap {
		for i, j := range lines {
			if len(j) > width {
				lines[i] = j[:width]
			}
		}
	}

	l := len(lines)
	cp := height

	var lastlines = make([]string, height)

	for i := 0; i < height; i++ {
		v.SetWritePos(0, i)
		lastlines[i] = lines[i]
		padSpacingR := width - len(lastlines[i]) // padSpacingR : the number of characters needed to right pad line
		padR := strings.Repeat(" ", padSpacingR) // padR : padding to right
		fmt.Fprintf(v, "%v%s", lastlines[i], padR)
		g.Update(func(g *gocui.Gui) error {
			return nil
		})
		time.Sleep(110 * time.Millisecond)
		//time.Sleep(2 * time.Second)
	}

	for {

		if cp == l {
			cp = 0
		}

		for m := 0; m <= height-2; m++ {
			lastlines[m] = lastlines[m+1]
			v.SetWritePos(0, m)
			padSpace := width - len(lastlines[m])
			fmt.Fprintf(v, "%v%s", lastlines[m], strings.Repeat(" ", padSpace))
		}
		v.SetWritePos(0, height-1)
		lastlines[height-1] = lines[cp]
		padSpace := width - len(lastlines[height-1])
		fmt.Fprintf(v, "%v%s", lastlines[height-1], strings.Repeat(" ", padSpace))
		cp++

		g.Update(func(g *gocui.Gui) error {
			return nil
		})

		for {
			if !data.Pause {
				break
			}
			time.Sleep(250 * time.Millisecond)
		}

		time.Sleep(110 * time.Millisecond)
	}

}
