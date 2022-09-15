package main

import "fmt"

func main() {

	larger := []int{1, 2, 3, 4}

	allPossibleSublists := findAllPossibleSublists(larger)
	for _, sublist := range allPossibleSublists {
		fmt.Println(sublist)
	}

}

func findAllPossibleSublists(list []int) [][]int {

	sublists := [][]int{ {}, }

	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list) + 1; j++ {
			sublists = append(sublists, list[i:j])
		}
	} 

	return sublists
}