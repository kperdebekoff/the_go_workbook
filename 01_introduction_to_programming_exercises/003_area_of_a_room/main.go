package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	
	fmt.Print("Enter length of a room in meters: ")
	strRoomLength := enterText()

	fmt.Print("Enter width of a room in meters: ")
	strRoomWidth := enterText()

	floatRoomLength := strToFloat(strRoomLength)
	floatRoomWidth := strToFloat(strRoomWidth)

	roomArea := floatRoomLength * floatRoomWidth

	fmt.Print("Area of a room: ", roomArea, " m^2")

}

func enterText() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	length := scanner.Text()

	return length
}

func strToFloat(roomDim string) float64 {
	floatDim, err2 := strconv.ParseFloat(roomDim, 64)
	if err2 != nil {
		log.Fatal(err2)
	}

	return floatDim
}