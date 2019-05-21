package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/martinrocket/learning/1.1/data"
)

func main() {
	c := exec.Command("clear")

	c.Stdout = os.Stdout
	c.Run()
	data.PrintMenu()
	fmt.Println()
	fmt.Printf("Lesson 1.1\n")
	handles()

}
