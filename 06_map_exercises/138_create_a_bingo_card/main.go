package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	bingoCard := createRandomBingoCard()

	displayBingoCards(bingoCard)

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

}
