package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	count := 0
	maxSimAmount := 0
	minSimAmount := 0
	for i := 0; i < 1000; i++ {
		card := createRandomBingoCard()
		lotto := simulateLottoGame()
		result := simulateCrossOut(card, lotto)

		if isWin(result) {
			count += 1
			maxSimAmount = i
		}
		if count == 1 {
			minSimAmount = i
		}
	}

	fmt.Println("Minimum numbers of calls", minSimAmount)
	fmt.Println("Maximum numbers of calls", maxSimAmount)
	fmt.Println("Average numbers of calls", (minSimAmount + maxSimAmount) / 2)

}

func simulateCrossOut(card map[string][]int, lotto []int) map[string][]int {

	letter := []string{"B", "I", "N", "G", "O"}
	for i := 0; i < len(lotto); i++ {
		for j := 0; j < len(letter); j++ {
			for k := 0; k < 5; k++ {
				if lotto[i] == card[letter[j]][k] {
					card[letter[j]][k] = 0
				}
			}
		}
	}
	return card

}

func simulateLottoGame() []int {

	const step = 15
	min := 1
	max := 1 + step

	randNum := func(min, max int) int {

		rand.Seed(time.Now().UnixNano())
		return rand.Intn(max - min + 1) + min

	}

	notInArr := func(arr []int, num int) bool {

		not := true
		for i := 0; i < len(arr); i++ {
			if arr[i] == num {
				not = false
				break
			}
		}
		return not

	}

	genUniqNums := func(min, max int) [] int {

		uniqNum := []int{}
		for len(uniqNum) != 5 {
			num := randNum(min, max)
			if notInArr(uniqNum, num) {
				uniqNum = append(uniqNum, num)
			}
		}
		return uniqNum

	}

	lottoNums := []int{}
	for i := 0; i < 5; i++ {
		lottoNums = append(lottoNums, genUniqNums(min, max)...)
		min = 1 + step
		max += step
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(lottoNums), func(i, j int) {
		lottoNums[i], lottoNums[j] = lottoNums[j], lottoNums[i]
	})

	return lottoNums

}

func createRandomBingoCard() map[string][]int {

	const step = 15
	min := 1
	max := 1 + step

	randNum := func(min, max int) int {

		rand.Seed(time.Now().UnixNano())
		return rand.Intn(max - min + 1) + min

	}

	notInArr := func(arr []int, num int) bool {

		not := true
		for i := 0; i < len(arr); i++ {
			if arr[i] == num {
				not = false
				break
			}
		}
		return not

	}

	genUniqNums := func(min, max int) [] int {

		uniqNum := []int{}
		for len(uniqNum) != 5 {
			num := randNum(min, max)
			if notInArr(uniqNum, num) {
				uniqNum = append(uniqNum, num)
			}
		}
		return uniqNum

	}

	letter := []string{ "B", "I", "N", "G", "O"}
	card := map[string][]int{}
	for i := 0; i < len(letter); i++ {
		card[letter[i]] = genUniqNums(min, max)
		min = 1 + step
		max += step
	}

	return card

}

func isWin(card map[string][]int) bool {

	vertSum := map[string]int{}
	keys := []string{"B", "I", "N", "G", "O"}

	for i := 0; i < len(keys); i++ {
		sum := 0
		for j := 0; j < len(card[keys[i]]); j++ {
			sum += card[keys[i]][j]
		}
		vertSum[keys[i]] = sum
		sum = 0
	}

	for _, v := range vertSum {
		if v == 0 {
			return true
		}
	}

	horizSum := map[string]int{}
	keys2 := []string{"row0", "row1", "row2", "row3", "row4"}
	for i := 0; i < len(keys); i++ {
		sum := 0
		for j := 0; j < len(card[keys[i]]); j++ {
			sum += card[keys[j]][i]
		}
		horizSum[keys2[i]] = sum
		sum = 0
	}
	
	for _, v := range horizSum {
		if v == 0 {
			return true
		}
	}

	sumLeftDiag := 0
	for i := 0; i < len(keys); i++ {
		for j := i; j < i + 1; j++ {
			sumLeftDiag += card[keys[i]][j]
		}
	}
	
	if sumLeftDiag == 0 {
		return true
	}

	sumRightDiag := 0
	for i := 0; i < len(keys); i++ {
		for j := 5 - i - 1; j != 5 - i - 2; j-- {
			sumRightDiag += card[keys[i]][j]
		}
	}
	
	if sumRightDiag == 0 {
		return true
	}

	return false

}