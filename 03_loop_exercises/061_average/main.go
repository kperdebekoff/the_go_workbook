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
	
	var sum float64 = 0
	var i float64 = 0
	for {
		fmt.Print("Enter some integer number: ")
		num := readInput()
		if num == 0 {
			break
		}
		sum += num
		i += 1
	}

	if sum == 0 && i == 0 {
		fmt.Println("First value can't be 0!")
	} else {
		fmt.Println("Average:", math.Floor((sum / i) * 100) / 100)
	}

}

func readInput() float64 {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	number, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}
	return number
}