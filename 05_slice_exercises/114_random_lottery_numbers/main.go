package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	ticketNums := generateRandomNumbers()

	for _, v := range ticketNums {
		fmt.Println(v)
	}

}

func generateRandomNumbers() []int {

	var randNums []int
	
	rand.Seed(time.Now().UnixNano())
	const (
		min = 1
		max = 49
	)

	randNums = append(randNums, rand.Intn(max - min + 1) + min)

	sort.Slice(randNums, func(i int, j int) bool {
		return randNums[i] < randNums[j]
	})

	for i := 0; i < 6; i++ {

		randNum := rand.Intn(max - min + 1) + min
		for _, v := range randNums {
			if v != randNum {
				randNums = append(randNums, randNum)
				break
			}
		}

	}

	return randNums 

}