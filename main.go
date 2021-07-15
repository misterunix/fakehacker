package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/awesome-gocui/gocui"
)

var (
	views   = []string{}
	curView = -1
	idxView = 0
)

func Roll(count, sides int) int {
	var t int
	for i := 0; i < count; i++ {
		r := cryptoRandSecure(int64(sides)) + 1
		t = t + int(r)
	}
	return t
}

func cryptoRandSecure(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0
	}
	return nBig.Int64()
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
	plines, err := readLines("password.txt")
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

			for pos := len(cpass.pass) - 1; pos >= 0; pos-- {

				randomCount := Roll(25, 20)
				for c := 0; c < randomCount; c++ {
					var hiddenpass string
					rc := Roll(1, (92)) + 32
					hiddenpass = strings.Repeat("*", len(cpass.pass))
					v.Clear()
					v.SetWritePos(0, 0)
					fmt.Fprintf(v, "\033[37;40m%s", cpass.crypt)

					v.SetWritePos(0, 1)
					wp := (width / 2) - (len(cpass.pass) / 2)
					pad := strings.Repeat(" ", wp)
					fmt.Fprintf(v, "%s", pad)

					for i := 0; i < len(cpass.pass); i++ {
						switch {
						case i < pos:
							fmt.Fprintf(v, "\033[37;40m%s", string(hiddenpass[i]))
						case i == pos:
							fmt.Fprintf(v, "\033[30;47m%s", string(byte(rc)))
						case i > pos:
							fmt.Fprintf(v, "\033[37;40m%s", string(cpass.pass[i]))
						}
					}

					g.Update(func(g *gocui.Gui) error {
						return nil
					})

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

func hack1(g *gocui.Gui, v *gocui.View, name string) {

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

func main() {

	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(1)
	}
	defer g.Close()
	// Set GUI managers and key bindings
	// ...
	g.ASCII = false

	g.SetManagerFunc(layout)

	if err := initKeybindings(g); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = sourceWindow(g)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = passwordCrack(g)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	/*
		if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
			// handle error
		}
	*/

	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		fmt.Println(err)
		os.Exit(1)
	}

}

func layout(g *gocui.Gui) error {
	/*
		maxX, maxY := g.Size()
		v, err := g.SetView("main", maxX-10, maxY-4, maxX-1, maxY-1, 0)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}

			fmt.Fprintln(v, "KEY")

		}
	*/
	return nil
}

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
	go hack1(g, v, name)
	return nil
}

/*
func newView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	name := fmt.Sprintf("v%v", idxView)
	v, err := g.SetView(name, maxX/2-5, maxY/2-5, maxX/2+5, maxY/2+5)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		fmt.Fprintln(v, strings.Repeat(name+" ", 30))
	}
	if _, err := g.SetCurrentView(name); err != nil {
		return err
	}

	views = append(views, name)
	curView = len(views) - 1
	idxView += 1
	return nil
}
*/

func initKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}
	return nil
}

/*
func initKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeySpace, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return newView(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyBackspace, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return delView(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return nextView(g, true)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return moveView(g, v, -delta, 0)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return moveView(g, v, delta, 0)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return moveView(g, v, 0, delta)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return moveView(g, v, 0, -delta)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 't', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			_, err := g.SetViewOnTop(views[curView])
			return err
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'b', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			_, err := g.SetViewOnBottom(views[curView])
			return err
		}); err != nil {
		return err
	}
	return nil
}
*/

/*;
func delView(g *gocui.Gui) error {
	if len(views) <= 1 {
		return nil
	}

	if err := g.DeleteView(views[curView]); err != nil {
		return err
	}
	views = append(views[:curView], views[curView+1:]...)

	return nextView(g, false)
}
*/

// Chunks : Split the string into a slice where each string has the max length of chunkSize.
// Each string in the slice is left justfied and padded with spaces.
func Chunks(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}

	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			//	fmt.Fprintf(os.Stderr, "::%d\n", len(s[currentStart:i]))
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])

	for i, tl := range chunks {
		tt := fmt.Sprintf("%%-%dv", chunkSize)
		ttt := fmt.Sprintf(tt, tl)
		chunks[i] = ttt
	}
	return chunks
}

// readLines : Read file from "path" spliting into lines on the newline.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ss := strings.TrimSpace(scanner.Text())
		lines = append(lines, ss)
	}
	return lines, scanner.Err()
}
