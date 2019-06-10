package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
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

type MyJSON struct {
	Field1 interface{} `json:"Name"`
	Field2 interface{} `json:"Value"`
}

//Write JSON
func CreateJSON() {
	q := "quit"
	s := MyJSON{}
	for s.Field1 != q {

		s.Field1 = rName()
		if s.Field1 == q {
			break
		}
		s.Field2 = rValue()
		j, _ := json.MarshalIndent(s, "", " ")
		os.Stdout.Write(j)
		fmt.Println()
	}
}

func rName() string {
	scanName := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter JSON Name(quit to exit): ")
	scanName.Scan()
	fmt.Println()
	return scanName.Text()
}
func rValue() string {
	scanValue := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter JSON Value: ")
	scanValue.Scan()
	fmt.Println()
	return scanValue.Text()
}

func createJSONMenu() {
	input := '0'
	for input != '3' {
		fmt.Println("1. Add Key Value Pair")
		fmt.Println("2. Add Child to Previous")
		fmt.Println("3. Done")
		reader := bufio.NewReader(os.Stdin)
		input, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("You selected %v \n", string(input))

		switch input {
		case '1':
			fmt.Println("Adding new key value pair.")
			addParent()

		case '2':
			fmt.Println("Adding child to previous")
			addChild()

		case '3':
			fmt.Println("Done")
			return

		default:
			fmt.Println("No Choice.")
		}
	}
}

func addParent() {
	return
}

func addChild() {
	return
}
