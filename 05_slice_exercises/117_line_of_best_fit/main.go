package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Enter the x-part a collection of points: ")
	x := readInput()
	fmt.Print("Enter the y-part a collection of points: ")
	y := readInput()

	lineOfBestFit := calcLineOfBestFit(x, y)

	fmt.Println("The line of best fit:", lineOfBestFit)

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func convertStrToFloat(x string) float64 {

	num, err := strconv.ParseFloat(x, 64)
	if err != nil {
		log.Fatal(err)
	}
	return num

}

func converFloatToStr(x float64) string {

	s := fmt.Sprintf("%.2f", x)
	return s

}

func calcLineOfBestFit(x, y string) string {

	var xArr []float64
	var yArr []float64

	for x != "" {

		xArr = append(xArr, convertStrToFloat(x))
		yArr = append(yArr, convertStrToFloat(y))

		fmt.Print("Enter the x-part a collection of points (blank to quit): ")
		x = readInput()
		if x == "" { break }
		fmt.Print("Enter the y-part a collection of points: ")
		y = readInput()

	}

	xy := calcMultXAndY(xArr, yArr)
	sumXy := sumArrItems(xy)
	sumX := sumArrItems(xArr)
	sumY := sumArrItems(yArr)
	n := float64(len(xArr))
	xx := calcPowX(xArr)
	sumXx := sumArrItems(xx)

	m :=  (sumXy - (sumX * sumY / n)) / (sumXx - sumX * sumX / n)
	b := sumY / n - m * (sumX / n)
	fit := "y = " + converFloatToStr(m) + "x + " + converFloatToStr(b)

	return fit

}

func sumArrItems(n []float64) float64 {
	
	sum := 0.0
	for _, v := range n {
		sum += v
	}
	return sum

}

func calcMultXAndY(x, y []float64) []float64 {

	var multArr []float64
	for i := 0; i < len(x); i++ {
		multArr = append(multArr, x[i] * y[i])
	}
	return multArr

}

func calcPowX(x []float64) []float64 {

	for i := 0; i < len(x); i++ {
		x[i] *= x[i]
	}
	return x

}
