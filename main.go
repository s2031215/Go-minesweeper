package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jroimartin/gocui"
)

var (
	blockArr    = []string{}
	boomArr     = bombinit()
	active      = 0
	die         = false
	blockstatus = [260]int{}
	stepcount   = 0
	debug       = false
)

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(blockArr)
	name := blockArr[nextIndex]

	out, err := g.View("v2")
	out.Clear()
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Going from view "+v.Name()+" to "+name)

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	active = nextIndex
	return nil
}

func layout(g *gocui.Gui) error {

	//maxX, maxY := g.Size()
	if v, err := g.SetView("v1", 0, 0, 100, 26); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = ""
		v.Wrap = true
	}

	if v, err := g.SetView("v2", 0, 27, 119, 43); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Debug"
		v.Wrap = true
		v.Frame = false
		v.Autoscroll = true
	}

	if v, err := g.SetView("status", 100, 0, 119, 13); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Game Status"
		fmt.Fprintln(v, "Step:", stepcount)
		fmt.Fprintln(v, "Bomb find: 0/40")
	}

	if v, err := g.SetView("Info", 100, 13, 119, 26); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "GoMine by S2"
		v.Wrap = false
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			fmt.Fprintln(v, "Go Minesweeper\nBuild with gocui")
			fmt.Fprintln(v, "\nControl:")
			fmt.Fprintln(v, "←↑→↓: Move View")
			fmt.Fprintln(v, "Space: ToggleView")
			fmt.Fprintln(v, "F: Flag Mine")
			fmt.Fprintln(v, "R: Reset Game")
			fmt.Fprintln(v, "D: Enable debug")
			fmt.Fprintln(v, "^C: Exit")
		}
		return nil
	}

	if v, err := g.SetView("GameOver", 30, 10, 70, 16); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "GameOver"
		v.Wrap = true
		v.Frame = false
		fmt.Fprintf(v, "Boom hits Game Over ! \nPress R to Restart")

		if _, err := g.SetViewOnBottom("GameOver"); err != nil {
			return err
		}
	}

	x, y := 0, 0
	for i := 1; i <= 260; i++ {
		block := "b" + strconv.Itoa(i)
		blockArr = append(blockArr, block)

		if v, err := g.SetView(block, x, y, x+5, y+2); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			//v.Title = strconv.Itoa(i)
			v.Wrap = true
			v.Autoscroll = true
			if block == "b1" {
				if _, err = setCurrentViewOnTop(g, block); err != nil {
					return err
				}
			}

		}
		x = x + 5
		if i%20 == 0 {
			y = y + 2
			x = 0
		}
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, upperView); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, downView); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone, leftView); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone, rightView); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeySpace, gocui.ModNone, toggleView); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", rune('r'), gocui.ModNone, ResetView); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", rune('f'), gocui.ModNone, flagView); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", rune('d'), gocui.ModNone, toggdebug); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
