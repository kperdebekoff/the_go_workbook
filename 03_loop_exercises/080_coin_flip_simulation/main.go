package main

import (
	"fmt"
	"math/rand"
	// "time"
	"unicode/utf8"
	"strings"
)

func main() {
	
	for i := 0; i < 10; i++ {
		flips := coinFlip()
		fmt.Println(flips, utf8.RuneCountInString(flips) , "flips")
	}
	
}

func generateRandomNum() string {
	
	flip := rand.Intn(100)
	if flip % 2 != 0 {
		return "H"
	} else {
		return "T"
	}

}

func coinFlip() string {

	t := ""

	for {
		r := generateRandomNum()
		t += r
		
		if strings.Count(t, "TTT") == 1 || strings.Count(t, "HHH") == 1 {
			break
		}
	}
	return t

}