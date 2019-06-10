package menu

import (
	"bufio"
	"fmt"
	"os"

	"github.com/martinrocket/learning/1.1/utils"

	"github.com/martinrocket/learning/1.1/handles"
)

func StartMenu() {
	input := '0'
	for input != '3' {
		fmt.Println("1. Create JSON")
		fmt.Println("2. Start Web Server")
		fmt.Println("3. Exit")
		reader := bufio.NewReader(os.Stdin)
		input, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("You selected %v \n", string(input))

		switch input {
		case '1':
			fmt.Println("Creating JSON")
			utils.CreateJSON()

		case '2':
			fmt.Println("Starting Web Server")
			handles.Handles()

		case '3':
			fmt.Println("Exit")
			os.Exit(0)

		default:
			fmt.Println("No Choice.")
		}
	}
}

func readString() string {
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	return input
}

func makeMenu(s []string) {
	for i := range s {
		fmt.Println(s[i])
	}
}
