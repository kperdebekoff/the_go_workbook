package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Enter a decimal number: ")
	decimal := readInput()
	
	result := ""
	for {
		result += strconv.Itoa(decimal % 2)
		decimal /= 2
		if decimal == 0 {
			break
		}
	}

	fmt.Println(reverseString(result))

}

func readInput() int {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}
	return num

}

func reverseString(text string) string {

	revStr := ""
 for i := len([]rune(text)) - 1; i != -1; i-- {
	revStr += string([]rune(text)[i])
 }
 return revStr

}