package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const vLine = "─"
const uRight = "┌"
const uLeft = "┐"
const space = " "
const hLine = "│"

func main() {
	clearScreen()
	fmt.Printf("Print OS %v\n", runtime.GOOS)
	vLine10()
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	w := exec.Command("clr")
	w.Stdout = os.Stdout
	w.Run()

}

func vLine10() {
	top := uRight + strings.Repeat(vLine, 10) + uLeft

	fmt.Print(top + "\n" + strings.Repeat(middle(), 10) + "\n")
}

func middle() string {
	return hLine + strings.Repeat(space, 10) + hLine + "\n"
}
