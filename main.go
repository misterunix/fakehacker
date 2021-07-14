package main

import (
	"bufio"
	"fmt"
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
	width := mx - 1
	height := my

	ss, _ := readLines("hack1.txt")
	for _, k := range ss {
		y := Chunks(k, width)
		lines = append(lines, y...)
		//for _, t := range y {
		//	lines = append(lines, t)
		//}
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
			//g.SetCurrentView(name)
			//v, err := g.View(name)
			//if err != nil {
			// handle error
			//}
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
			//fmt.Fprintf(v, "%-48v", lastlines[m])
			fmt.Fprintf(v, "%v", lastlines[m])
		}
		v.SetWritePos(0, height-1)
		lastlines[height-1] = lines[cp]
		//fmt.Fprintf(v, "%-48v", lastlines[3])
		fmt.Fprintf(v, "%v", lastlines[height-1])
		cp++

		g.Update(func(g *gocui.Gui) error {
			//g.SetCurrentView(name)
			//v, err := g.View(name)
			//if err != nil {
			// handle error
			//}
			return nil
		})
		//		fmt.Fprintf(os.Stderr, "%s", string(text[cp]))

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
	maxX, maxY := g.Size()
	v, err := g.SetView("main", maxX-10, maxY-4, maxX-1, maxY-1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		fmt.Fprintln(v, "KEY")
		/*
			fmt.Fprintln(v, "Space: New View")
			fmt.Fprintln(v, "Tab: Next View")
			fmt.Fprintln(v, "← ↑ → ↓: Move View")
			fmt.Fprintln(v, "Backspace: Delete View")
			fmt.Fprintln(v, "t: Set view on top")
			fmt.Fprintln(v, "b: Set view on bottom")
			fmt.Fprintln(v, "^C: Exit")
		*/
	}
	return nil
}

func sourceWindow(g *gocui.Gui) error {
	//maxX, maxY := g.Size()
	//name := fmt.Sprintf("hack1", idxView)
	name := "hack1"
	v, err := g.SetView(name, 0, 0, 50, 12, 0) // 48x10
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
			//tt := fmt.Sprintf("%%-%d", chunkSize)
			//ttt := fmt.Sprintf(tt, s[currentStart:i])
			chunks = append(chunks, s[currentStart:i])
			//chunks = append(chunks, ttt)
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	for i := range chunks {
		chunks[i] = strings.TrimLeft(chunks[i], "\t")
		chunks[i] = strings.TrimSpace(chunks[i])
		chunks[i] = strings.TrimRight(chunks[i], "\n")
		chunks[i] = strings.TrimRight(chunks[i], "\r")
		chunks[i] = strings.TrimRight(chunks[i], "\n")
		tt := fmt.Sprintf("%%-%d", chunkSize)
		ttt := fmt.Sprintf(tt, chunks[i])
		chunks[i] = ttt
	}
	return chunks
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
