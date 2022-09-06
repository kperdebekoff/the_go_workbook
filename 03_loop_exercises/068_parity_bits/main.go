package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	for {
		scan := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter group of 8 bits (blank to quit): ")
		scan.Scan()
		bits := scan.Text()
		if bits == "" {
			break
		} else if utf8.RuneCountInString(bits) > 8 || utf8.RuneCountInString(bits) < 8 {
			fmt.Println("Entered something other than 8 bits!")
		} else {
			parityBits := strings.Count(bits, "1")
			if parityBits % 2 == 0 {
				fmt.Println("The parity bit should be 0")
			} else {
				fmt.Println("The parity bit should be 1")
			}
		}
	}
	
}