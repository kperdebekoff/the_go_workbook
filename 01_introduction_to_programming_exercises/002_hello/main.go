package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Hello, ", text)

}