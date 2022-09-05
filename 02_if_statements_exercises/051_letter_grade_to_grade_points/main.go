package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a letter grade to conver to grade points: ")
	scan.Scan()
	letterGrade := scan.Text()

	switch letterGrade {
	case "A+":
		fmt.Println(letterGrade, "equal 4.0 grade points")
	case "A":
		fmt.Println(letterGrade, "equal 4.0 grade points")
	case "A-":
		fmt.Println(letterGrade, "equal 3.7 grade points")
	case "B+":
		fmt.Println(letterGrade, "equal 3.3 grade points")
	case "B":
		fmt.Println(letterGrade, "equal 3.0 grade points")
	case "B-":
		fmt.Println(letterGrade, "equal 2.7 grade points")
	case "C+":
		fmt.Println(letterGrade, "equal 2.3 grade points")
	case "C":
		fmt.Println(letterGrade, "equal 2.0 grade points")
	case "C-":
		fmt.Println(letterGrade, "equal 2.7 grade points")
	case "D+":
		fmt.Println(letterGrade, "equal 1.3 grade points")
	case "D":
		fmt.Println(letterGrade, "equal 1.0 grade points")
	case "F":
		fmt.Println(letterGrade, "equal 0.0 grade points")
	default:
		fmt.Println("Enter a correct letter grade!")
	}

}