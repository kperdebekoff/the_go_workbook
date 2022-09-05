package main

import "fmt"

func main() {
	
	count := 0
	var pi float64 = 3

	for i := 2.0; i <= 30.0; i += 4.0 {

		pi += 4 / (i * (i + 1) * (i + 2))
		count += 1
		fmt.Println(count, "approximate:", pi)

		if i < 30 {
			for j := i + 2; j < i + 3; j += 1 {

				pi -= 4 / (j * (j + 1) * (j + 2))
				count += 1
				fmt.Println(count, "approximate:", pi)
	
			}
		}
	}

}