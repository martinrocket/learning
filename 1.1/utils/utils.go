package utils

import (
	"os"
	"os/exec"
)

//Clear the screen
func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	w := exec.Command("clr")
	w.Stdout = os.Stdout
	w.Run()
}
