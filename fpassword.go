package main

import (
	"fakehacker/data"
	"fmt"
	"strings"
	"time"

	"github.com/awesome-gocui/gocui"
)

func passwordCrack(g *gocui.Gui) error {

	// bottom right
	maxX, maxY := g.Size()
	x0 := maxX - 37
	y0 := maxY - 4
	x1 := maxX - 1
	y1 := maxY - 1

	name := "pass1"
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
	go pass1(g, v, name)
	return nil

}

func pass1(g *gocui.Gui, v *gocui.View, name string) {

	// password : holds the password and the encrypted password
	type password struct {
		pass  string
		crypt string
	}

	// passwords : slice of type password
	var passwords []password
	passwords = make([]password, 0) // file is 1000 entries so allocate memory

	// read the file in to a slice of strings
	///plines, err := readLines("password.txt")
	plines, err := readGzipLines("password.gz", &pfile)
	if err != nil {
		return
	}

	// loop through the lines from the file and add to slice passwords
	for _, pass := range plines {
		pw := strings.Split(pass, " ")
		p := password{}
		p.pass = strings.TrimSpace(pw[0])
		p.crypt = strings.TrimSpace(pw[1])
		passwords = append(passwords, p)
	}

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

	mx, _ := v.Size()

	width := mx - 1
	for {
		for _, cpass := range passwords {
			cpassLength := len(cpass.pass)         // cpassLength : Length of the password
			wpL := (width / 2) - (cpassLength / 2) // wpL : padding count to center password
			padL := strings.Repeat(" ", wpL)       // padL : left padding of the pass
			wpR := width - len(padL) - cpassLength
			padR := strings.Repeat(" ", wpR) // padR : left padding of the pass

			for pos := cpassLength - 1; pos >= 0; pos-- {

				// print crypt string
				v.SetWritePos(0, 0)
				fmt.Fprintf(v, "\033[48;5;0m\033[38;5;15m%s", cpass.crypt)

				randomCount := Roll(25, 20) // randomCount : Number of fake attempt
				for c := 0; c < randomCount; c++ {

					var hiddenpass string
					rc := Roll(1, (92)) + 32 // rc : random character byte

					hiddenpass = strings.Repeat("*", cpassLength) // hiddenpass : string containing count of *

					v.SetWritePos(0, 1)

					fmt.Fprintf(v, "%s", padL)

					for i := 0; i < cpassLength; i++ {
						switch {
						case i < pos:
							fmt.Fprintf(v, "\033[48;5;0m\033[38;5;15m%s", string(hiddenpass[i]))
						case i == pos:
							fmt.Fprintf(v, "\033[48;5;0m\033[38;5;15m%s", string(byte(rc)))
						case i > pos:
							fmt.Fprintf(v, "\033[48;5;0m\033[38;5;15m%s", string(cpass.pass[i]))
						}
					}

					fmt.Fprintf(v, "%s", padR)

					g.Update(func(g *gocui.Gui) error {
						return nil
					})

					for {
						if !data.Pause {
							break
						}
						time.Sleep(250 * time.Millisecond)
					}

					time.Sleep(20 * time.Millisecond)
				}
			}

			v.Clear()
			v.SetWritePos(0, 0)
			fmt.Fprintf(v, "\033[37;40m%s", cpass.crypt)
			v.SetWritePos(0, 1)
			wp := (width / 2) - (len(cpass.pass) / 2)
			pad := strings.Repeat(" ", wp)
			fmt.Fprintf(v, "%s%s", pad, cpass.pass)
			g.Update(func(g *gocui.Gui) error {
				return nil
			})
			time.Sleep(5 * time.Second)
		}
	}
}
