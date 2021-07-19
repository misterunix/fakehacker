package main

import (
	"fmt"

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

	var x0, y0, x1, y1 int
	maxX, maxY := g.Size()

	count := 0

	q.name = fmt.Sprintf("pop%d", count)
	q.msg = "FAILURE"
	q.on = false
	p = append(p, q)
	x0 = rnd(0, maxX-(len(q.msg)+4))
	y0 = rnd(0, maxY-4)
	x1 = x0 + (len(q.msg) + 5)
	y1 = y0 + 2
	v, err := g.SetView(q.name, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		//fmt.Fprintf(os.Stderr, "g.SetView(%s,%d,%d,%d,%d) - %s\n", name, x0, y0, x1, y1, err)
	}
	go popup(g, v, q.name, q.msg)
	count++
	//time.Sleep(500 * time.Millisecond)

	q.name = fmt.Sprintf("pop%d", count)
	q.msg = "INTRUSION DETECTED"
	q.on = false
	p = append(p, q)
	x0 = rnd(0, maxX-(len(q.msg)+4))
	y0 = rnd(0, maxY-4)
	x1 = x0 + (len(q.msg) + 5)
	y1 = y0 + 2
	v, err = g.SetView(q.name, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		//fmt.Fprintf(os.Stderr, "g.SetView(%s,%d,%d,%d,%d) - %s\n", name, x0, y0, x1, y1, err)
	}
	go popup(g, v, q.name, q.msg)
	count++
	//time.Sleep(500 * time.Millisecond)

	q.name = fmt.Sprintf("pop%d", count)
	q.msg = "CRITICAL ERROR"
	q.on = false
	p = append(p, q)
	x0 = rnd(0, maxX-(len(q.msg)+4))
	y0 = rnd(0, maxY-4)
	x1 = x0 + (len(q.msg) + 5)
	y1 = y0 + 2
	v, err = g.SetView(q.name, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		//fmt.Fprintf(os.Stderr, "g.SetView(%s,%d,%d,%d,%d) - %s\n", name, x0, y0, x1, y1, err)
	}
	go popup(g, v, q.name, q.msg)
	count++
	//time.Sleep(500 * time.Millisecond)

	q.name = fmt.Sprintf("pop%d", count)
	q.msg = "ERROR"
	q.on = false
	p = append(p, q)
	x0 = rnd(0, maxX-(len(q.msg)+4))
	y0 = rnd(0, maxY-4)
	x1 = x0 + (len(q.msg) + 5)
	y1 = y0 + 2
	v, err = g.SetView(q.name, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		//fmt.Fprintf(os.Stderr, "g.SetView(%s,%d,%d,%d,%d) - %s\n", name, x0, y0, x1, y1, err)
	}
	go popup(g, v, q.name, q.msg)
	count++
	//time.Sleep(500 * time.Millisecond)

	q.name = fmt.Sprintf("pop%d", count)
	q.msg = "ANTIVIRUS RUNNING"
	q.on = false
	p = append(p, q)
	x0 = rnd(0, maxX-(len(q.msg)+4))
	y0 = rnd(0, maxY-4)
	x1 = x0 + (len(q.msg) + 5)
	y1 = y0 + 2
	v, err = g.SetView(q.name, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		//fmt.Fprintf(os.Stderr, "g.SetView(%s,%d,%d,%d,%d) - %s\n", name, x0, y0, x1, y1, err)
	}
	go popup(g, v, q.name, q.msg)
	count++
	//time.Sleep(500 * time.Millisecond)

	q.name = fmt.Sprintf("pop%d", count)
	q.msg = "EXTERNAL ACCESS DETECTED"
	q.on = false
	p = append(p, q)
	x0 = rnd(0, maxX-(len(q.msg)+4))
	y0 = rnd(0, maxY-4)
	x1 = x0 + (len(q.msg) + 5)
	y1 = y0 + 2
	v, err = g.SetView(q.name, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		//fmt.Fprintf(os.Stderr, "g.SetView(%s,%d,%d,%d,%d) - %s\n", name, x0, y0, x1, y1, err)
	}
	go popup(g, v, q.name, q.msg)
	count++
	//time.Sleep(500 * time.Millisecond)

	q.name = fmt.Sprintf("pop%d", count)
	q.msg = "SHUTDOWN Y/n"
	q.on = false
	p = append(p, q)
	x0 = rnd(0, maxX-(len(q.msg)+4))
	y0 = rnd(0, maxY-4)
	x1 = x0 + (len(q.msg) + 5)
	y1 = y0 + 2
	v, err = g.SetView(q.name, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		//fmt.Fprintf(os.Stderr, "g.SetView(%s,%d,%d,%d,%d) - %s\n", name, x0, y0, x1, y1, err)
	}
	go popup(g, v, q.name, q.msg)
	//time.Sleep(500 * time.Millisecond)
	//l := len(p)

	//for i := 0; i < l; i++ {
	//	go popup(g, p[i].name, p[i].msg)
	//}

	//time.Sleep(5 * time.Second)
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
func popup(g *gocui.Gui, v *gocui.View, name string, message string) {

	msg := "  " + message + "  "

	v.Frame = true
	v.Wrap = false
	v.Autoscroll = false
	fmt.Fprintf(v, "%s", msg)
	g.Update(func(g *gocui.Gui) error {
		//fmt.Fprintf(os.Stderr, "g.update error\n")
		return nil
	})

	views = append(views, name)
	curView = len(views) - 1
	idxView += 1
	//go skull1(g, v, name)
	return
}
