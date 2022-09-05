package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter number of sides in shapes (from 3 to 10): ")
	scan.Scan()
	sides, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}

	switch sides {
		
	case 3:
		fmt.Println(sides, "sides in Triangle")
	case 4:
		fmt.Println(sides, "sides in Quadrilateral")
	case 5:
		fmt.Println(sides, "sides in Pentagon")
	case 6:
		fmt.Println(sides, "sides in Hexagon")
	case 7:
		fmt.Println(sides, "sides in Heptagon")
	case 8:
		fmt.Println(sides, "sides in Octagon")
	case 9:
		fmt.Println(sides, "sides in Nonagon")
	case 10:
		fmt.Println(sides, "sides in Decagon")
	default:
		fmt.Println(sides, "sides out of range")
	}

}