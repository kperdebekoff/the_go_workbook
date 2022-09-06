package main

import "fmt"

func main() {

	for i:= 0; i <= 100; i += 10 {
		fahrenheit := i * (9 / 5) + 32
		fmt.Println(i, "Celcius degree equal", fahrenheit, "Fahrenheit degree")
	}

}