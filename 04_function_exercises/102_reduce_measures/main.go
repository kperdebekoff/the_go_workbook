package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Enter integer number of units: ")
	num := convertStrToInt(readInput())
	fmt.Print("Enter unit of measure in lowercase (cup, tablespoon or teaspoon ): ")
	measure := readInput()

	cups, tablespoons, teaspoons := reduceMeasures(num, measure)
	displayMessage := displayReduceMeasures(cups, tablespoons, teaspoons)

	fmt.Println(displayMessage)
	
}

func reduceMeasures(n int, s string) (int, int, int) {

	const (
		cup = 48  // 1 cup = 48 teaspoon
		tablespoon = 3  // 1 tablespoon = 3 teaspoon
		teaspoon = 1  // 1 teaspoon
	)

	switch s {
	case "cup":
		n *= cup
	case "tablespoon":
		n *= tablespoon
	case "teaspoon":
		n *= teaspoon
	default:
		n = -1
	}

	var cp, tbs, ts int = 0, 0, 0

	if n / cup >= 1 && n > 0 {
		cp = n / cup
	}

	if n % cup >= tablespoon && n > 0  {
		tbs = (n % cup) / tablespoon
	}

	if (n % cup) % tablespoon < tablespoon && n > 0  {
		ts = (n % cup) % tablespoon
	}

	return cp, tbs, ts

}

func displayReduceMeasures(cp, tbs, ts int) string {
	
	message := ""
	if cp == 0 && tbs == 0 && ts == 0 {
		message = "One of the entered parameters wrong!"
	} else if cp > 0 && tbs > 0 && ts > 0 {
		message = strconv.Itoa(cp) + " cups " + strconv.Itoa(tbs) + " tablespoons " + strconv.Itoa(ts) + " teaspoons"
	} else if cp > 0 && tbs > 0 && ts <= 0 {
		message = strconv.Itoa(cp) + " cups " + strconv.Itoa(tbs) + " tablespoons"
	} else if cp > 0 && tbs <= 0 && ts <= 0 {
		message = strconv.Itoa(cp) + " cups"
	} else if cp <= 0 && tbs > 0 && ts > 0 {
		message = strconv.Itoa(tbs) + " tablespoons " + strconv.Itoa(ts) + " teaspoons"
	} else if cp <= 0 && tbs <= 0 && ts > 0 {
		message = strconv.Itoa(ts) + " teaspoons"
	} else if cp <= 0 && tbs > 0 && ts <= 0 {
		message = strconv.Itoa(tbs) + " tablespoons"
	}
	return message

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func convertStrToInt(s string) int {

	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num

}