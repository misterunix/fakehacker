package main

import (
	"fakehacker/ansicolor"
	"fmt"
	"time"

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

	//var skull []string

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

	fakeSource, err := readFileToString("bksrc.txt")
	if err != nil {
		//fmt.Fprintf(os.Stderr, "%s", err)
		return
	}

	skullLines, err := readLinesRaw("skull.txt")
	if err != nil {
		return
	}

	rawLineLength := len(fakeSource)
	numSkullLines := len(skullLines)

	var offset = make([]int, numSkullLines) // offset : Offset into the fakeSource string. Slice with the count of skullLines.

	d := rawLineLength / numSkullLines

	offset[0] = 0
	for i := 1; i < numSkullLines; i++ {
		offset[i] = i * d
	}

	//fmt.Fprintln(os.Stderr, offset)

	var scroll = make([]string, numSkullLines)
	var cc byte
	var cs string
	for {

		for ind, skullLine := range skullLines {

			skullLineLen := len(skullLine)

			scroll[ind] = ""
			o := offset[ind]
			for i := 0; i < skullLineLen; i++ {

				cc = skullLine[i]
				if cc == '~' {
					//c = ' '
					cs = fmt.Sprintf("%s ", ansicolor.BColor[17])
				} else {
					fs := fakeSource[o]
					cs = fmt.Sprintf("%s%s%s", ansicolor.FColor[19], ansicolor.BColor[17], string(fs))
					if fs >= 65 && fs <= 90 {
						cs = fmt.Sprintf("%s%s%s", ansicolor.FColor[21], ansicolor.BColor[17], string(fs))
					} else {
						if o+1 < rawLineLength-1 {
							if fakeSource[o+1] >= 65 && fs <= 90 {
								cs = fmt.Sprintf("%s%s%s", ansicolor.FColor[20], ansicolor.BColor[17], string(fs))
							}
						}
						if o > 0 {
							if fakeSource[o-1] >= 65 && fs <= 90 {
								cs = fmt.Sprintf("%s%s%s", ansicolor.FColor[20], ansicolor.BColor[17], string(fs))
							}
						}

					}
				}

				scroll[ind] += cs
				o++

			}
			offset[ind]++
			if offset[ind] >= rawLineLength {
				offset[ind] = numSkullLines * d
			}

			//offset[ind]++

		}

		//fmt.Fprintln(os.Stderr)

		//v.Clear()
		for i, s := range scroll {
			v.SetWritePos(0, i)
			fmt.Fprintf(v, "%s", s)
		}
		time.Sleep(1 * time.Second)
	}
}
