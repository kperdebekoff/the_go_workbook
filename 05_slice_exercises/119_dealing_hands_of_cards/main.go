package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	hands := 4
	cards := 5
	deck := createDeck()

	dealing, remCards := deal(hands, cards, deck)

	fmt.Println("All of the hands of cards:")
	fmt.Println(dealing)
	fmt.Println("The cards remaining in the deck:")
	fmt.Println(remCards)

}

func deal(hands int, cards int, deck []string) ([][]string, []string) {

	deckShuffle := shuffle(deck)

	var handsCards = make([][]string, hands)
	for i := 0; i < hands; i++ {
		for j := i; j < cards * hands; j += cards - 1 {
			handsCards[i] = append(handsCards[i], deckShuffle[j])
		}
	}

	deckShuffle = deckShuffle[hands * cards:]

	return handsCards, deckShuffle

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
