package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	
	expectedPercent, simulatedPercent := simulateOneThousandRollDice()

	fmt.Print("Total ", "Simulated ", "Expected\n")
	fmt.Print("      ", "Percent   ", "Percent\n")
	for i := 2; i <= 12; i++ {
		fmt.Print(i, "\t", simulatedPercent[i], "\t", expectedPercent[i], "\n")
	}
	
}

func simulateRollingTwoDice() int {

	min := 1
	max := 6
	rand.Seed(time.Now().UnixNano())
	rollDice1 := rand.Intn(max - min + 1) + min
	rollDice2 := rand.Intn(max - min + 1) + min
	rollDice := rollDice1 + rollDice2
	return rollDice

}

func simulateOneThousandRollDice() (map[int]float64, map[int]float64) {

	expectedResult := map[int]float64{}
	simulatedResult := map[int]float64{1: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0, 11: 0}

	// s2 = 1/36; s3 = 2/36; s4 = 3/36; s5 = 4/36; s6 = 5/36
	// s7 = 6/36; s8 = 5/36; s9 = 4/36; s10 = 3/36; s11 = 2/36; s12 = 1/36
	j := 3
	for i := 2; i <= 12; i++ {
		if i <= 7 {
			expectedResult[i] = math.Floor(float64(i - 1) / 36 * 10000) / 100
		} else {
			expectedResult[i] = math.Floor(float64(i - j) / 36 * 10000) / 100
			j += 2
		}
	}

	for j := 0; j < 1000; j++ {
		dice := simulateRollingTwoDice()
		simulatedResult[dice] += 1
	}

	for i := 2; i <= 12; i++ {
		simulatedResult[i] /= 10
	}

	return expectedResult, simulatedResult

}