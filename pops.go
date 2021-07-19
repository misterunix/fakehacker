package main

import (
	"fmt"
	"os"
	"time"

	"github.com/awesome-gocui/gocui"
)

func doPopups(g *gocui.Gui) error {

	type pop struct {
		name string
		msg  string
		on   bool
	}

	var p []pop
	q := pop{}

	count := 0

	q.name = fmt.Sprintf("pop%d", count)
	q.msg = "FAILURE"
	q.on = false
	p = append(p, q)
	count++

	q.name = fmt.Sprintf("pop%d", count)
	q.msg = "INTRUSION DETECTED"
	q.on = false
	p = append(p, q)
	count++

	q.name = fmt.Sprintf("pop%d", count)
	q.msg = "CRITICAL ERROR"
	q.on = false
	p = append(p, q)
	count++

	q.name = fmt.Sprintf("pop%d", count)
	q.msg = "ERROR"
	q.on = false
	p = append(p, q)
	count++

	q.name = fmt.Sprintf("pop%d", count)
	q.msg = "ANTIVIRUS RUNNING"
	q.on = false
	p = append(p, q)
	count++

	q.name = fmt.Sprintf("pop%d", count)
	q.msg = "EXTERNAL ACCESS DETECTED"
	q.on = false
	p = append(p, q)
	count++

	q.name = fmt.Sprintf("pop%d", count)
	q.msg = "SHUTDOWN Y/n"
	q.on = false
	p = append(p, q)

	l := len(p)

	for i := 0; i < l; i++ {
		go popup(g, p[i].name, p[i].msg)
	}

	time.Sleep(5 * time.Second)
	/*
		for i := 0; i < l; i++ {
			err := g.DeleteView(p[i].name)
			if err != nil {
				return err
			}
		}
	*/
	return nil

}

// popup : Create a popup error.
func popup(g *gocui.Gui, name string, message string) {
	maxX, maxY := g.Size()

	msg := "  " + message + "  "
	msgWidth := len(msg)

	x0 := Roll(1, maxX-msgWidth)
	y0 := Roll(1, maxY-2)
	x1 := x0 + msgWidth
	y1 := y0 + 2

	v, err := g.SetView(name, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return
		}
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "'"+name+"'")
	}
	//	if _, err := g.SetCurrentView(name); err != nil {
	//		return
	//	}

	v.Frame = true
	fmt.Fprintf(v, "%s", msg)
	g.Update(func(g *gocui.Gui) error {
		return err
	})

	views = append(views, name)
	curView = len(views) - 1
	idxView += 1
	//go skull1(g, v, name)
	return
}
