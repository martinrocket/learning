package data

import (
	"fmt"
	"os"
)

type menu struct {
	a string
}

// PrintMenu exports the ability to print the menu
func PrintMenu() {
	s, err := os.Create("./data/menu.config")
	if err != nil {
		fmt.Printf("error opening file %s\n", s.Name())
		fmt.Println(s.Name())
	}

	defer s.Close()

	menu1 := menu{"menu one"}
	fmt.Println(menu1.a)
}
