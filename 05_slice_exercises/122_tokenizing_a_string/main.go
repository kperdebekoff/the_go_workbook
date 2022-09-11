package main

import (
	"fmt"
	"regexp"
)

func main() {

	mathExp1 := "x+(x+10)+4=2*x+14"
	mathExp2 := "(x + 4) ^ 2 - x ^ 2 = x ^ 2 + 16 + 8 * x - x ^ 2 = 16 + 8 * x"

	
	fmt.Println(breakstringToToken(mathExp1))
	fmt.Println(breakstringToToken(mathExp2))

}

func breakstringToToken(mathExp string) []string {

	re, _ := regexp.Compile(`[a-z\(\)\*\/\^\-\+]|\d+`)
	res := re.FindAllString(mathExp, -1)

	return res

}