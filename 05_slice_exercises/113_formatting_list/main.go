package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Enter list items: ")
	item := readInput()

	strs := storeDataInSlice(item)

	fmt.Println(addListFormatting(strs))

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func storeDataInSlice(s string) []string {

	var sliceOfItem []string
	for s != "" {
		sliceOfItem = append(sliceOfItem, s)
		fmt.Print("Enter list items: ")
		s = readInput()
	}

	return sliceOfItem

}

func addListFormatting(strs []string) string {

	fmtList := ""

	if len(strs) < 2 {

		fmtList += strs[0]

	} else if len(strs) > 2 {

		for i, v := range strs {
			if i < len(strs) - 1 {
				fmtList += v + ","
			} else {
				fmtList += " and " + v
			}

		}
	} else {

		for i, v := range strs {
			if i == 0 {
				fmtList += v
			} else {
				fmtList += " and " + v
			}
		}

	}

	return fmtList

}