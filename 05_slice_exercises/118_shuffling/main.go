package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	deck := createDeck()
	deckShuffle := shuffle(deck)

	fmt.Println("Before shuffle:")
	fmt.Println(deck)
	fmt.Println("After shuffle:")
	fmt.Println(deckShuffle)

}

func createDeck() []string {

	var deckSlice []string
	deck := func(s string) []string {

		var deckArr []string
		for i := 2; i <= 14; i++ {

			switch i {
			case 11:
				deckArr = append(deckArr, "J" + s)
			case 12:
				deckArr = append(deckArr, "Q" + s)
			case 13:
				deckArr = append(deckArr, "K" + s)
			case 14:
				deckArr = append(deckArr, "A" + s)
			default:
				deckArr = append(deckArr, strconv.Itoa(i) + s)
			}

		}

		return deckArr

	}

	spades := deck("s")
	hearts := deck("h")
	diamonds := deck("d")
	clubs := deck("c")

	deckSlice = append(deckSlice, spades...)
	deckSlice = append(deckSlice, hearts...)
	deckSlice = append(deckSlice, diamonds...)
	deckSlice = append(deckSlice, clubs...)

	return deckSlice

}

func shuffle(deck []string) []string {

	var shuffleDeck []string
	shuffleDeck = append(shuffleDeck, deck...)
	rand.Seed(time.Now().UnixNano())
	

		for i := 0; i < len(shuffleDeck); i++ {
			n := rand.Intn(53)
			for j := 0; j < n; j++ {
				shuffleDeck[i], shuffleDeck[j] = shuffleDeck[j], shuffleDeck[i]
			}
		}

	return shuffleDeck

}
