package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	const acres = 43560

	fmt.Print("Enter lenth of a field in feet: ")
	length := convertstrToFloat(readEnteredText())

	fmt.Print("Enter width of a field in feet: ")
	width := convertstrToFloat(readEnteredText())

	area := (length * width) / acres

	fmt.Println("Area of a field: ", area, " acres")

}

func readEnteredText() string {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()

}

func convertstrToFloat(dim string) float64 {

	flDim, err := strconv.ParseFloat(dim, 64)
	if err != nil {
		log.Fatal(err)
	}
	return flDim
	
}