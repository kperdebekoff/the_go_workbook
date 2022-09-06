package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	fmt.Print("Enter an math operator: ")
	oprtr := readInput()

	prcdnc := precendence(oprtr)

	fmt.Println(prcdnc)

}

func precendence(operator string) int {

	level := -1
	switch operator {
	case "+", "-":
		level = 1
	case "*", "/":
		level = 2
	case "^":
		level = 3
	}
	return level

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}