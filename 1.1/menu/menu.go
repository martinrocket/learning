package menu

import (
	"bufio"
	"fmt"
	"os"

	"github.com/martinrocket/learning/1.1/handles"
)

func StartMenu() {
	fmt.Println("1. Create JSON")
	fmt.Println("2. Start Web Server")
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("You selected %v \n", string(input))

	switch input {
	case '1':
		fmt.Println("Creating JSON")

	case '2':
		fmt.Println("Starting Web Server")
		handles.Handles()

	default:
		fmt.Println("No Choice.")
	}
}
