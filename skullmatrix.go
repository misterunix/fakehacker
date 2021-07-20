package main

import (
	"fakehacker/data"
	"fmt"
	"time"

	"github.com/awesome-gocui/gocui"
)

// skullWindow : Create the view for the skull.
// Width has to be 85 minumn for the window to render correctly.
func skullWindow(g *gocui.Gui) error {
	//maxX, maxY := g.Size()

	x0 := 0
	y0 := 0
	x1 := 28
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

	fakeSource, err := readFileToString("bksrc.txt")
	if err != nil {
		return
	}

	var skullLines []string
	skullLines = append(skullLines, "                            ")
	skullLines = append(skullLines, "          ~~~~~~~~          ")
	skullLines = append(skullLines, "        ~~~~~~~~~~~~        ")
	skullLines = append(skullLines, "       ~~~~~~~~~~~~~~       ")
	skullLines = append(skullLines, "      ~~~~~~~~~~~~~~~~      ")
	skullLines = append(skullLines, "     ~~~~~~~~~~~~~~~~~      ")
	skullLines = append(skullLines, "     ~~~~~~~~~~~~~~~~~~     ")
	skullLines = append(skullLines, "     ~~~~~~~~~~~~~~~~~~     ")
	skullLines = append(skullLines, "     ~~~~   ~~~~   ~~~~     ")
	skullLines = append(skullLines, "     ~~~     ~~~    ~~~     ")
	skullLines = append(skullLines, "     ~~~     ~~~    ~~~     ")
	skullLines = append(skullLines, "      ~~~    ~ ~    ~~      ")
	skullLines = append(skullLines, "      ~~~~~~~~ ~~~~~~~~     ")
	skullLines = append(skullLines, "      ~~~~~~~   ~~~~~~      ")
	skullLines = append(skullLines, "          ~~~   ~~          ")
	skullLines = append(skullLines, "        ~  ~~~~~~~ ~~       ")
	skullLines = append(skullLines, "        ~~ ~ ~ ~ ~ ~        ")
	skullLines = append(skullLines, "         ~         ~        ")
	skullLines = append(skullLines, "         ~ ~     ~ ~        ")
	skullLines = append(skullLines, "         ~~~ ~ ~ ~~         ")
	skullLines = append(skullLines, "          ~~~~~~~~~         ")
	skullLines = append(skullLines, "            ~~~~~           ")
	skullLines = append(skullLines, "                            ")

	//skullLines, err := readLinesRaw("skull.txt")
	//	if err != nil {
	//		return
	//}

	rawLineLength := len(fakeSource)
	numSkullLines := len(skullLines)

	var offset = make([]int, numSkullLines) // offset : Offset into the fakeSource string. Slice with the count of skullLines.

	d := rawLineLength / numSkullLines

	offset[0] = 0
	for i := 1; i < numSkullLines; i++ {
		offset[i] = i * d
	}

	var scroll = make([]string, numSkullLines)
	var cc byte
	var cs string
	var fs, fs1, fs2 byte
	for {

		for ind, skullLine := range skullLines {

			skullLineLen := len(skullLine)

			scroll[ind] = ""
			o := offset[ind]
			for i := 0; i < skullLineLen; i++ {

				cc = skullLine[i]
				if cc == '~' {
					cs = "\033[48;5;16m " // fmt.Sprintf("\033[48;5;16m ")
				} else {
					fs = fakeSource[o]
					if o > 0 {
						fs1 = fakeSource[o-1]
					} else {
						fs1 = fakeSource[o]
					}
					if o+1 < rawLineLength-1 {
						fs2 = fakeSource[o+1]
					} else {
						fs1 = fakeSource[o]
					}

					if fs >= 65 && fs <= 90 {
						cs = fmt.Sprintf("\033[38;5;33m\033[48;5;17m%s", string(fs))
					} else {
						if (fs1 >= 65 && fs1 <= 90) || (fs2 >= 65 && fs2 <= 90) {
							cs = fmt.Sprintf("\033[38;5;25m\033[48;5;17m%s", string(fs))
						} else {
							cs = fmt.Sprintf("\033[38;5;19m\033[48;5;17m%s", string(fs))
						}
					}
				}

				scroll[ind] += cs
				o++
				if o >= rawLineLength {
					o = ind * d
					offset[ind] = o
				}
			}
			offset[ind]++
			if offset[ind] >= rawLineLength {
				offset[ind] = ind * d
			}

		}

		for i, s := range scroll {
			v.SetWritePos(0, i)
			fmt.Fprintf(v, "%s", s)
		}

		for {
			if !data.Pause {
				break
			}
			time.Sleep(250 * time.Millisecond)
		}

		time.Sleep(60 * time.Millisecond)
	}
}
