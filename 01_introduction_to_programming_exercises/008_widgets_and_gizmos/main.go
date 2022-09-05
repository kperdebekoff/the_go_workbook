package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	const (
		widgetsMass = 75
		gizmosMass = 112
	)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a number of widgets: ")
	scanner.Scan()
	numOfWidg := convertStrToInt(scanner.Text())

	fmt.Print("Enter a number of gizmos: ")
	scanner.Scan()
	numOfGiz := convertStrToInt(scanner.Text())

	totalWeight := numOfWidg * widgetsMass + numOfGiz * gizmosMass

	fmt.Print("Total weight of the order: ", totalWeight, " grams")
	
}

func convertStrToInt(text string) int {

	amount, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}
	return amount

}