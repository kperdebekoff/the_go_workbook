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
	fmt.Print("Enter distance in feets: ")
	scan.Scan()
	dist, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	inches := dist * 12
	yards := dist / 3
	miles := dist / 5280

	fmt.Println(dist, "feets contains", inches, "inches")
	fmt.Println(dist, "feets contains", yards, "yards")
	fmt.Println(dist, "feets contains", miles, "miles")
	
}