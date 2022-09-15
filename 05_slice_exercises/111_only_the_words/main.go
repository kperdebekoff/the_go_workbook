package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	fmt.Print("Enter some sentnces: ")
	sentnces := readInput()

	sliceOfWords := createSliceOfWords(sentnces)

	fmt.Println(sliceOfWords)
}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func createSliceOfWords(text string) []string {

	re, _ := regexp.Compile(`[,.!?:]*`)
	res := re.FindAllString(text, -1)

	for _, v := range res {
		text = strings.ReplaceAll(text, v, "")
	}

	words := strings.Split(text, " ")

	return words

}