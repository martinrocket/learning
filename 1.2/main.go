package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const vLine = "─"
const uRight = "┐"
const uLeft = "┌"
const space = " "
const hLine = "│"
const bRight = "┘"
const bLeft = "└"
const bl = " "

type format string

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
	top := uLeft + strings.Repeat(vLine, x) + uRight
	bottom := bLeft + strings.Repeat(vLine, x) + bRight
	blank := hLine + strings.Repeat(space, x) + hLine

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
	var sLen int
	if len(s)%2 == 0 {
		sLen = len(s)
		return hLine + strings.Repeat(space, x/2-sLen/2) + s + strings.Repeat(space, x/2-sLen/2) + hLine + "\n"
	} else {
		sLen = len(s) + 1
		return hLine + strings.Repeat(space, x/2-sLen/2) + s + strings.Repeat(space, x/2-sLen/2+1) + hLine + "\n"
	}

}

func middleLeftJustify(x int, s string) string {

	return hLine + bl + s + strings.Repeat(space, (x-1-len(s))) + hLine + "\n"
}
