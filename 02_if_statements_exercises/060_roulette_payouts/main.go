package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	rand.Seed(time.Now().Unix())
	result := rand.Intn(37)

	win1 := result
	win2 := "black"
	win3 := "even"
	win4 := "19 to 36"


	switch result {
	case 1, 3, 5, 7, 9, 12, 14, 16, 18, 19, 21, 23, 25, 27, 30, 32, 34, 36:
		win2 = "red"
	}

	if result % 2 != 0 {
		win3 = "odd"
	}

	if result <= 18 {
		win4 = "1 to 18"
	}

	fmt.Println("The spin resulted in", result, "...")
	if result == 0 {
		fmt.Println("Pay", win1)
	} else {
		fmt.Println("Pay", win1)
		fmt.Println("Pay", win2)
		fmt.Println("Pay", win3)
		fmt.Println("Pay", win4)
	}
	
}