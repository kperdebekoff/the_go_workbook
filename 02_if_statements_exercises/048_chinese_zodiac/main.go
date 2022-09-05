package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	
	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a birth year: ")
	scan.Scan()
	year, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}

	switch year % 12 {
	case 2000 % 12:
		fmt.Println("Dragon")
	case 2001 % 12:
		fmt.Println("Snake")
	case 2002 % 12:
		fmt.Println("Horse")
	case 2003 % 12:
		fmt.Println("Sheep")
	case 2004 % 12:
		fmt.Println("Monkey")
	case 2005 % 12:
		fmt.Println("Rooster")
	case 2006 % 12:
		fmt.Println("Dog")
	case 2007 % 12:
		fmt.Println("Pig")
	case 2008 % 12:
		fmt.Println("Rat")
	case 2009 % 12:
		fmt.Println("Ox")
	case 2010 % 12:
		fmt.Println("Tiger")
	case 2011 % 12:
		fmt.Println("Hare")
	}

}