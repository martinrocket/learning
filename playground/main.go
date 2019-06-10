package main

import (
	"fmt"
)

func main() {

	makeMenu([]string{"yes", "no"})
}

func makeMenu(s []string) {
	for i := range s {
		fmt.Println(i, s[i])
	}
}
