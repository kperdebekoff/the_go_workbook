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

	fmt.Print("Enter the latitude in degree for first point: ")
	t1 := convertDegToRadians(readInput())
	fmt.Print("Enter the longtitude in degree for first point: ")
	g1 := convertDegToRadians(readInput())
	fmt.Print("Enter the latitude in degree for second point: ")
	t2 := convertDegToRadians(readInput())
	fmt.Print("Enter the longtitude in degree for second point: ")
	g2 := convertDegToRadians(readInput())

	distance := 6371.01 * math.Acos(math.Sin(t1) * math.Sin(t2) + math.Cos(t1) * math.Cos(t2) * math.Cos(g1 -g2))

	fmt.Println("Distance between the points:", math.Floor(distance * 1000) / 1000, "km.")
}

func readInput() float64 {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	l, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	return l

}

func convertDegToRadians(deg float64) float64 {

	return deg * (math.Pi / 180)

}