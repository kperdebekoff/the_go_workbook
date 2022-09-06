package main

import "fmt"

func main() {

	for j := 1; j <= 10; j++ {
		fmt.Print("\t", j, " ")
	}
	fmt.Println("")

	for i := 1; i <= 10; i++ {
		fmt.Print(i, "\t")
		for j := 1; j <= 10; j++ {
			fmt.Print(i * j, "\t")
		}
		fmt.Println("")
	}
	
}