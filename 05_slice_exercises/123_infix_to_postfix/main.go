package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	mathExp1 := "3+4"
	mathExp2 := "(3 + 4) ^ 2"
	mathExp3 := "16 + -8 * 2 - +10"

	tokens1 := findUnary(breakstringToToken(mathExp1))
	tokens2 := findUnary(breakstringToToken(mathExp2))
	tokens3 := findUnary(breakstringToToken(mathExp3))
	
	fmt.Println(convertInfixToPostfix(tokens1))
	fmt.Println(convertInfixToPostfix(tokens2))
	fmt.Println(convertInfixToPostfix(tokens3))

}

func breakstringToToken(mathExp string) []string {

	mathExp = strings.Replace(mathExp, " ", "", -1)
	re, _ := regexp.Compile(`[a-z\(\)\*\/\^\-\+]|\d+`)
	res := re.FindAllString(mathExp, -1)

	return res

}

func findUnary(tokens []string) []string {

	for i := 0; i < len(tokens); i++ {
		if i < len(tokens) - 1 {
			for j := i + 1; j < i + 2; j++ {
				if tokens[i] == "+" && tokens[j] == "+" {
					tokens[j] = "u+"
				} else if tokens[i] == "-" && tokens[j] == "-" {
					tokens[j] = "u-"
				} else if tokens[i] == "-" && tokens[j] == "+" {
					tokens[j] = "u+"
				} else if tokens[i] == "+" && tokens[j] == "-" {
					tokens[j] = "u-"
				} else if tokens[i] == "(" && tokens[j] == "-" {
					tokens[j] = "u-"
				} else if tokens[i] == "(" && tokens[j] == "+" {
					tokens[j] = "u+"
				}
			}
		}
	}

	return tokens
}

func convertInfixToPostfix(tokens []string) []string {

	var operators []string
	var postfix []string

	for i := 0; i < len(tokens); i++ {
		if isStrToInt(tokens[i]) {
			postfix = append(postfix, tokens[i])
		}
		if isOperator(tokens[i]) {
			for len(operators) != 0 && operators[len(operators) - 1] != "(" && tokens[i] != "u+" && tokens[i] != "u-" {
				lst := operators[len(operators) - 1:]
				operators = operators[:len(operators) - 1]
				postfix = append(postfix, lst...)
			}
			operators = append(operators, tokens[i])
		}
		if tokens[i] == "(" {
			operators = append(operators, tokens[i])
		}
		if tokens[i] == ")" {
			for operators[(len(operators) - 1)] != "(" {
				lst := operators[len(operators) - 1:]
				operators = operators[:len(operators) - 1]
				postfix = append(postfix, lst...)
			}
			for i := 0; i < len(operators); i++ {
				if operators[i] == ")" {
					operators = operators[:i]
				}
			}
		}
	}

	for len(operators) != 0 {
		lst := operators[len(operators) - 1:]
		operators = operators[:len(operators) - 1]
		postfix = append(postfix, lst...)
	}

	return postfix

}

func isStrToInt(s string) bool {

	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true

}

func isOperator(s string) bool {
	
	operator := []string{"+", "-", "*", "/", "u-", "u+"}
	for i := 0; i < len(operator); i++ {
		if operator[i] == s {
			return true
		}
	}

	return false

}