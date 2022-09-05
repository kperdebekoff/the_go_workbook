package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Enter note letter: ")
	letter := readInput()

	fmt.Print("Enter octave number: ")
	octNumber, err := strconv.ParseFloat(readInput(), 64)
	if err != nil {
		log.Fatal(err)
	}

	freq := 0.00
	switch letter {
	case "C":
		freq = 261.63
	case "D":
		freq = 293.66
	case "E":
		freq = 329.63
	case "F":
		freq = 349.23
	case "G":
		freq = 392.00
	case "A":
		freq = 440.00
	case "B":
		freq = 493.88
	}

	noteFreq := freq / math.Pow(2, 4 - octNumber)

	fmt.Println(letter + fmt.Sprintf("%v", octNumber), "note's frequency is", noteFreq, "Hz")

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}