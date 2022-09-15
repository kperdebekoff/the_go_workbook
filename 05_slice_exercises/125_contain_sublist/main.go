package main

import "fmt"

func main() {

	larger := []int{1, 2, 3, 4}
	smaller1 := []int{3, 4}
	smaller2 := []int{1, 4}
	smaller3 := []int{}

	fmt.Println(isSublist(larger, smaller1))
	fmt.Println(isSublist(larger, smaller2))
	fmt.Println(isSublist(larger, smaller3))

}

func isSublist(larger, smaller []int) bool {

	findFirstInd := func(larger, smaller []int) int {

		firstInd := 0
		loopsEnd:
		for i := 0; i < len(larger); i++ {
			for j := 0; j < len(smaller); j++ {
				if larger[i] == smaller[j] {
					firstInd = i
					break loopsEnd
				}
			}
		}
		return firstInd

	}

	firstIndex := findFirstInd(larger, smaller)
	lastIndex := len(smaller) + firstIndex

	larger = larger[firstIndex:lastIndex]
	count := 0
	for i := 0; i < len(larger); i++ {
		for j := 0; j < len(smaller); j++ {
			if i == j && larger[i] == smaller[j] {
				count += 1
			}
		}
	}

	if count == len(larger) {
		return true
	} else {
		return false
	}
	
}
