package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	totalPoint := 0.0
	count := 0.0
	for {
		fmt.Print("Enter a letter grades (blank to quit): ")
		letterGrade := readInput()
		if letterGrade == "" {
			break
		}
		point := letterGradeToPointGrade(letterGrade)
		totalPoint += point
		count += 1
	}

	fmt.Println("Total grade points:", totalPoint)
	fmt.Println("Grade point average:", totalPoint / count)
	
}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func letterGradeToPointGrade(letter string) float64 {
	gp := 0.0
	switch letter {
	case "A+":
		gp = 4.0
	case "A":
		gp =  4.0
	case "A-":
		gp = 3.7
	case "B+":
		gp = 3.3
	case "B":
		gp = 3.0
	case "B-":
		gp = 2.7
	case "C+":
		gp = 2.3
	case "C":
		gp = 2.0
	case "C-":
		gp = 2.7
	case "D+":
		gp = 1.3
	case "D":
		gp = 1.0
	case "F":
		gp = 0.0
	}
	return gp

}