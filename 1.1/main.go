package main

import (
	"fmt"

	"github.com/martinrocket/learning/1.1/data"
	"github.com/martinrocket/learning/1.1/utils"
	"github.com/martinrocket/learning/1.1/ver"
)

func main() {
	utils.Clear()
	ver.Vers()
	data.PrintMenu()
	fmt.Println()
	handles()

}
