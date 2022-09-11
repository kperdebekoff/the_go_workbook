package main

import "fmt"

func main() {

	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	list2 := []int{11, 27, 31, 49, 52, 67, 73, 88, 91}
	list3 := []int{121, 27, 341, 49, 52, 657, 73, 828, 91}

	fmt.Println(rangeCount(list1, 5, 9))
	fmt.Println(rangeCount(list2, 50, 99))
	fmt.Println(rangeCount(list3, 850, 999))

}

func rangeCount(list []int, min, max int) int {

	count := 0
	for _, v := range list {
		if v >= min && v < max {
			count += 1
		}
	}
	return count

}