
package main

import (
	"fmt"
	"os"
   	"os/exec"
)

func main(){
	c := exec.Command("clear")
   	c.Stdout = os.Stdout
   	c.Run()
	fmt.Printf("Lesson 1.\nStarting....\n...started\n")
}

