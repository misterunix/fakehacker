package main

import (
	"embed"
	"fakehacker/data"
	"fmt"
	"os"

	"github.com/awesome-gocui/gocui"
)

//go:embed hack1.gz
var efile embed.FS

var (
	views   = []string{}
	curView = -1
	idxView = 0
)

func main() {

	data.Init()

	g, err := gocui.NewGui(gocui.Output256, true)
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

	err = skullWindow(g)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = doPopups(g)
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

func initKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyEsc, gocui.ModNone,
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
