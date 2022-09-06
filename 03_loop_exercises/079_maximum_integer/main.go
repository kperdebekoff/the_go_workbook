package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	min := 1
	max := 100

	maxRand := 0
	count := 0
	i := 0
	for i != 100 {
		rand.Seed(time.Now().UnixNano())
		randomNum := rand.Intn(max - min) + min
		if randomNum > maxRand {
			maxRand = randomNum
			count += 1
			fmt.Print("Update ==> ")
		}
		fmt.Println(randomNum)
		i += 1
	}

	fmt.Println("The maximum value:", maxRand)
	fmt.Println("The maximum was updated", count, "times")
	
}