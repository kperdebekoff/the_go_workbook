package main

import "fmt"

func main() {

	for i := 1; i <= 12; i++ {
		fmt.Println(convertIntegerToOrdinal(i))
	}
	
}

func convertIntegerToOrdinal(num int) string {
	
	ordNum := ""

	switch num {
	case 1:
		ordNum = "First"
	case 2:
		ordNum = "Second"
	case 3:
		ordNum = "Third"
	case 4:
		ordNum = "Fourth"
	case 5:
		ordNum = "Fifth"
	case 6:
		ordNum = "Sixth"
	case 7:
		ordNum = "Seventh"
	case 8:
		ordNum = "Eighth"
	case 9:
		ordNum = "Ninth"
	case 10:
		ordNum = "Tenth"
	case 11:
		ordNum = "Eleventh"
	case 12:
		ordNum = "Twelfth"
	}
 return ordNum

}