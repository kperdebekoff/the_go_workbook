package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter some math expression: ")
	mathExp := readInput()

	tokens := findUnary(breakstringToToken(mathExp))
	postfix := convertInfixToPostfix(tokens)
	result := evaluateInfix(postfix)
	
	fmt.Println(postfix)
	fmt.Println(result)

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

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
				if tokens[i] == "-" && i == 0 {
					tokens[i] = "u-"
				} else if tokens[i] == "+" && tokens[j] == "+" {
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

func evaluateInfix(infix []string) int {

	var values [] int
	for i := 0; i < len(infix); i++ {
		if isStrToInt(infix[i]) {
			num, _ := strconv.Atoi(infix[i])
			values = append(values, num)
		}
		if infix[i] == "u-" || infix[i] == "u+" {
			if infix[i] == "u-" {
				lst := values[len(values) - 1]
				values = values[:len(values) - 1]
				values = append(values, -1 * lst)
			}
		}
		if infix[i] == "+" || infix[i] == "-" || infix[i] == "*" || infix[i] == "/" {
			right := values[len(values) - 1]
			values = values[:len(values) - 1]
			left := values[len(values) - 1]
			values = values[:len(values) - 1]
			if infix[i] == "+" { 
				values = append(values, right + left)
			} else if infix[i] == "-" {
				values = append(values, right - left)
			} else if infix[i] == "*" {
				values = append(values, right * left)
			} else if infix[i] == "/" {
				values = append(values, right / left)
			}
		}
	}

	return values[0]

}