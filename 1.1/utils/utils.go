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
	fmt.Printf("Enter JSON Name( quit to exit): ")
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
