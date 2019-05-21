package data

import (
	"fmt"
)

type menu struct {
	a string
}

// PrintMenu exports the ability to print the menu
func PrintMenu() {
	menu1 := menu{"menu one"}
	fmt.Printf("%v", menu1)
}
