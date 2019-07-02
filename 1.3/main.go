package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter an integer: ")
	s.Scan()
	i, err := strconv.Atoi(s.Text())
	if err != nil {
		fmt.Println("Not an Integer, will use integer = 1")
		i = 1
	}
	i = i * i
	fmt.Printf("%v \n", i)

}
