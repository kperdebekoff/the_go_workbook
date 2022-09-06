package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter some string: ")
	s := readInput()

	intN := isInteger(s)

	fmt.Println("Reprsent an Intger:", intN)
	
}

func isInteger(s string) bool {

	s = strings.Replace(s, " ", "", -1)
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}