package main

import (
	"fmt"
	"strconv"

	"github.com/jroimartin/gocui"
)

func isdie() bool {
	return die
}

func toggdebug(g *gocui.Gui, v *gocui.View) error {
	debug = !debug
	showdebug(g, v, "Key:D toggdebug")
	return nil
}

func showdebug(g *gocui.Gui, v *gocui.View, log string) error {
	out, err := g.View("v2")
	if err != nil {
		return err
	}
	if debug {
		out.Frame = true
		out.Clear()
		fmt.Fprintf(out, log)
		printbomb(out)
	} else {
		out.Clear()
		out.Frame = false
	}
	return nil
}

func upperView(g *gocui.Gui, v *gocui.View) error {
	if isdie() {
		return nil
	}
	nextIndex := active - 20
	if nextIndex < 0 {
		nextIndex = nextIndex + 260
	}

	name := blockArr[nextIndex]

	showdebug(g, v, "Key:UP Action:Going from view "+v.Name()+" to "+name)

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	active = nextIndex
	return nil
}

func downView(g *gocui.Gui, v *gocui.View) error {
	if isdie() {
		return nil
	}
	nextIndex := active + 20
	if nextIndex > 260 {
		nextIndex = nextIndex - 260
	}

	name := blockArr[nextIndex]

	showdebug(g, v, "Key:Down Action:Going from view "+v.Name()+" to "+name)

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	active = nextIndex
	return nil
}

func leftView(g *gocui.Gui, v *gocui.View) error {
	if isdie() {
		return nil
	}
	nextIndex := active - 1
	if nextIndex%20 == 19 || nextIndex == -1 {
		nextIndex = nextIndex + 20
	}

	name := blockArr[nextIndex]

	showdebug(g, v, "Key:Left Action:Going from view "+v.Name()+" to "+name)

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	active = nextIndex
	return nil
}

func rightView(g *gocui.Gui, v *gocui.View) error {
	if isdie() {
		return nil
	}
	nextIndex := active + 1
	if nextIndex%20 == 0 {
		nextIndex = nextIndex - 20
		if nextIndex > 260 {
			nextIndex = 1
		}
	}

	name := blockArr[nextIndex]

	showdebug(g, v, "Key:Right Action:Going from view "+v.Name()+" to "+name)

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	active = nextIndex
	return nil
}

func toggleView(g *gocui.Gui, v *gocui.View) error {
	//check dis
	if isdie() {
		return nil
	}

	//first step must save
	if stepcount == 0 {
		for boomArr[active] != 0 {
			boomArr = bombinit()
		}
	}

	if blockstatus[active] != 2 {
		name := blockArr[active]
		block, err := g.View(name)
		if err != nil {
			return err
		}
		block.Frame = false
		block.Clear()
		fmt.Fprintf(block, "%2v", strconv.Itoa(boomArr[active]))

		switch boomArr[active] {
		case -1:
			if blockstatus[active] != 2 {
				gameover(g)
			}
		case 0:
			cleanblock(g, findempty(active, boomArr, []int{}))
		default:
			cleanblock(g, flagcheck(active, boomArr, blockstatus))
			if isdie() {
				gameover(g)
			}
		}

		stepcount++
	}

	checkwin(g)
	printstatus(g)

	//debug
	showdebug(g, v, "Key:Space Action:toggle view "+v.Name())

	return nil
}
func printstatus(g *gocui.Gui) error {
	bfine := 0
	for i := 1; i < 260; i++ {
		if blockstatus[i] == 2 {
			bfine++
		}
	}
	out, err := g.View("status")
	if err != nil {
		return err
	}
	out.Clear()
	out.Title = "Game Status"
	fmt.Fprintln(out, "Step:", stepcount)
	fmt.Fprintln(out, "Bomb find:", bfine, "/40")

	return nil
}

func cleanblock(g *gocui.Gui, blocklist []int) error {
	for _, s := range blocklist {
		block, err := g.View(blockArr[s])
		if err != nil {
			return err
		}
		block.Frame = false
		block.Clear()
		fmt.Fprintf(block, "%2v", strconv.Itoa(boomArr[s]))
		blockstatus[s] = 1
	}
	return nil
}

func checkwin(g *gocui.Gui) error {
	//display all boom
	for i := 1; i < 260; i++ {
		if boomArr[i] >= 0 && blockstatus[i] != 1 {
			return nil
		}
	}
	die = true
	if _, err := g.SetViewOnTop("GameOver"); err != nil {
		return err
	}
	out, err := g.View("GameOver")
	if err != nil {
		return err
	}
	out.Clear()
	out.Frame = true
	out.Title = "You Win"
	fmt.Fprintf(out, "Good Job! \nPress R to Start New Game")
	return nil
}

func gameover(g *gocui.Gui) error {
	die = true
	//display all boom
	for i := 1; i < 260; i++ {
		if boomArr[i] == -1 {
			block, err := g.View(blockArr[i])
			if err != nil {
				return err
			}
			block.Frame = false
			block.Clear()
			fmt.Fprintf(block, "%2v", strconv.Itoa(boomArr[i]))
		}
	}
	if _, err := g.SetViewOnTop("GameOver"); err != nil {
		return err
	}
	out, err := g.View("GameOver")
	if err != nil {
		return err
	}
	out.Clear()
	out.Frame = true
	out.Title = "GameOver"
	fmt.Fprintf(out, "Boom hits Game Over ! \nPress R to Restart")
	return nil
}

func ResetView(g *gocui.Gui, v *gocui.View) error {
	if _, err := g.SetViewOnBottom("GameOver"); err != nil {
		return err
	}
	out, err := g.View("GameOver")
	if err != nil {
		return err
	}
	out.Frame = false
	boomArr = bombinit()
	die = false
	for i := 1; i <= 260; i++ {
		name := blockArr[i]
		block, err := g.View(name)
		if err != nil {
			return err
		}
		block.Frame = true
		block.Clear()
	}

	blockstatus = [260]int{}
	stepcount = 0
	printstatus(g)
	return nil
}

func flagView(g *gocui.Gui, v *gocui.View) error {
	if blockstatus[active] != 1 {
		name := blockArr[active]
		block, err := g.View(name)
		if err != nil {
			return err
		}
		block.Clear()

		if blockstatus[active] == 0 {
			blockstatus[active] = 2
			fmt.Fprintf(block, "%2v", "F")
		} else {
			blockstatus[active] = 0
		}

	}
	printstatus(g)
	return nil
}

func printbomb(v *gocui.View) error {
	for i := 20; i <= 260; i = i + 20 {
		fmt.Fprintf(v, "\n%2v%1v", boomArr[i-20:i], blockstatus[i-20:i])
	}
	return nil
}
