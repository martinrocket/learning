package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// const vLine = "─"
// const uRight = "┐"
// const uLeft = "┌"
// const space = " "
// const hLine = "│"
// const bRight = "┘"
// const bLeft = "└"
// const bl = " "

type format string

type boxDraw struct {
	vLine  string
	uRight string
	uLeft  string
	space  string
	hLine  string
	bRight string
	bLeft  string
	bl     string
}

type B interface {
	box1()
}

type Menu struct {
	myMenu []string
}

func main() {
	s := Menu{}
	t := Menu{}
	s.myMenu = []string{"hello", "bye", "where", "sugar", "jana"}
	t.myMenu = []string{"1. Load Menu", "2. Next Menu", "3. Exit"}
	clearScreen()
	fmt.Printf("Computer OS: %v\n", runtime.GOOS)
	buildBox(40, 10, s, "center")
	buildBox(40, 10, t, "left")
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	w := exec.Command("clr")
	w.Stdout = os.Stdout
	w.Run()

}

func buildBox(x int, y int, m Menu, f format) {
	b := boxDraw{}

	top := b.box1().uLeft + strings.Repeat(b.box1().vLine, x) + b.box1().uRight
	bottom := b.box1().bLeft + strings.Repeat(b.box1().vLine, x) + b.box1().bRight
	blank := b.box1().hLine + strings.Repeat(b.box1().space, x) + b.box1().hLine

	fmt.Print(top + "\n")
	fmt.Print(blank + "\n")

	switch f {
	case "center":
		for i := range m.myMenu {
			fmt.Print(middleCentered(x, m.myMenu[i]))
		}

	case "left":
		for i := range m.myMenu {
			fmt.Print(middleLeftJustify(x, m.myMenu[i]))
		}
	}

	fmt.Print(blank + "\n")
	fmt.Print(bottom + "\n")

}

func middleCentered(x int, s string) string {
	b := boxDraw{}
	var sLen int
	if len(s)%2 == 0 {
		sLen = len(s)
		return b.box1().hLine + strings.Repeat(b.box1().space, x/2-sLen/2) + s + strings.Repeat(b.box1().space, x/2-sLen/2) + b.box1().hLine + "\n"
	} else {
		sLen = len(s) + 1
		return b.box1().hLine + strings.Repeat(b.box1().space, x/2-sLen/2) + s + strings.Repeat(b.box1().space, x/2-sLen/2+1) + b.box1().hLine + "\n"
	}

}

func middleLeftJustify(x int, s string) string {
	b := boxDraw{}
	return b.box1().hLine + b.box1().bl + s + strings.Repeat(b.box1().space, (x-1-len(s))) + b.box1().hLine + "\n"
}

func (m boxDraw) box1() boxDraw {
	m.vLine = "─"
	m.uRight = "┐"
	m.uLeft = "┌"
	m.space = " "
	m.hLine = "│"
	m.bRight = "┘"
	m.bLeft = "└"
	m.bl = " "
	return m

}
