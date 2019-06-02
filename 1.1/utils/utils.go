package utils

import (
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
	Field1 string `json:"Name"`
	Field2 string `json:"Function"`
}

//Write JSON
func CreateJSON(s1 string, s2 string) {
	fmt.Println(s1)
	fmt.Println(s2)
	s := MyJSON{s1, s2}
	j, _ := json.MarshalIndent(s, "", " ")
	os.Stdout.Write(j)
	fmt.Println()
}
