package main

import (
	"fmt"
)

func main() {

	card1 := map[string][]int {
		"B": {0, 0, 0, 0, 0}, "I": {1, 2, 3, 4, 6},
		"N": {1, 2, 3, 4, 7}, "G": {1, 2, 3, 4, 8}, "O": {9, 2, 3, 4, 9},
	}
	displayBingoCards(card1)

	card2 := map[string][]int {
		"B": {0, 0, 0, 0, 1}, "I": {1, 2, 0, 4, 6},
		"N": {1, 2, 0, 4, 7}, "G": {1, 2, 0, 4, 8}, "O": {1, 2, 0, 4, 9},
	}
	displayBingoCards(card2)

	card3 := map[string][]int {
		"B": {0, 0, 0, 0, 1}, "I": {1, 0, 0, 4, 6},
		"N": {1, 2, 0, 4, 7}, "G": {1, 2, 0, 0, 8}, "O": {1, 2, 3, 4, 0},
	}
	displayBingoCards(card3)


	card4 := map[string][]int {
		"B": {1, 2, 3, 4, 1}, "I": {3, 0, 0, 4, 6},
		"N": {1, 2, 0, 4, 7}, "G": {1, 2, 0, 0, 8}, "O": {1, 2, 3, 4, 5},
	}
	displayBingoCards(card4)

	card5 := map[string][]int {
		"B": {5, 2, 3, 4, 0}, "I": {3, 0, 1, 0, 6},
		"N": {1, 2, 0, 4, 7}, "G": {1, 0, 1, 5, 8}, "O": {0, 2, 3, 4, 1},
	}
	displayBingoCards(card5)

	fmt.Println(isWin(card1))  // true - all vertical values is 0
	fmt.Println(isWin(card2))  // true - all horizontal values is 0
	fmt.Println(isWin(card3))  // true - left diagonal values is 0
	fmt.Println(isWin(card4)) // false - no values is 0
	fmt.Println(isWin(card5)) // true - right diagonal values is 0

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

func displayBingoCards(bingoCard map[string][]int) {

	displayCards := func(card map[string][]int, key string, item int) {
		for k, v := range card {
			if k == key {
				for j := item; j < len(v); j++ {
					fmt.Print(v[j], "\t")
					break
				}
			}
		}
	}

	letter := []string{ "B", "I", "N", "G", "O"}
	for _, v := range letter {
		fmt.Print(v, "\t")
	}
	fmt.Println()
	for j := 0; j < 5; j++ {
		for i := 0; i < len(letter); i++ {
			displayCards(bingoCard, letter[i], j)
		}
		fmt.Println()
	}
	fmt.Println("---------------------------------")

}
