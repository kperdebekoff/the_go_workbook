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
	fmt.Print("Enter a grade points: ")
	scan.Scan()
	gradePoints, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	if gradePoints >= 4.0 {
		fmt.Println(gradePoints, "equal an A+ letter point")
	} else if gradePoints >= 3.7 && gradePoints < 4.0 {
		fmt.Println(gradePoints, "equal an A- letter point")
	} else if gradePoints > 3.0 &&  gradePoints < 3.7 {
		fmt.Println(gradePoints, "equal an B+ letter point")
	} else if gradePoints > 2.7 && gradePoints <= 3.0 {
		fmt.Println(gradePoints, "equal an B letter point")
	} else if gradePoints > 2.3 && gradePoints <= 2.7 {
		fmt.Println(gradePoints, "equal an B- letter point")
	} else if gradePoints > 2.0 && gradePoints <= 2.3 {
		fmt.Println(gradePoints, "equal an C+ letter point")
	} else if gradePoints > 1.7 && gradePoints <= 2.0 {
		fmt.Println(gradePoints, "equal an C letter point")
	} else if gradePoints > 1.3 && gradePoints <=1.7 {
		fmt.Println(gradePoints, "equal an C- letter point")
	} else if gradePoints > 1.0 && gradePoints <= 1.3 {
		fmt.Println(gradePoints, "equal an D+ letter point")
	} else if gradePoints > 0 && gradePoints <= 1.0 {
		fmt.Println(gradePoints, "equal an D letter point")
	} else if gradePoints == 0 {
		fmt.Println(gradePoints, "equal an F letter point")
	} else {
		fmt.Println("Enter a correct grade points!")
	}

}