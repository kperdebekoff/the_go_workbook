package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	const ( 
		planMintes = 50.00
		planMessages = 50.00
		planCost = 15.00
		addMinutes = 0.25
		addMessages = 0.15
		support = 0.44
		commonTax = 0.05
	)

	fmt.Print("Enter used number of minutes in a month: ")
	usedMinutes := readInput()
	fmt.Print("Enter used number of text messages in a month: ")
	usedMessages := readInput()

	minutesCost, messageCost, tax := 0.00, 0.00, 0.00
	if usedMinutes > planMintes && usedMessages > planMessages {
		minutesCost = (usedMinutes - planMintes) * addMinutes
		messageCost = (usedMessages - planMessages) * addMessages
		tax = (planCost + minutesCost + messageCost + support) * commonTax

		fmt.Println("Additional minute charge $", math.Floor(minutesCost * 100) / 100)
		fmt.Println("Additional message charge $", math.Floor(messageCost * 100) / 100)
	} else if usedMinutes <= planMintes && usedMessages <= planMessages {
		tax = (planCost + support) * commonTax
	} else if usedMinutes <= planMintes && usedMessages > planMessages {
		messageCost = (usedMessages - planMessages) * addMessages
		tax = (planCost + messageCost + support) * commonTax

		fmt.Println("Additional message charge $", math.Floor(messageCost * 100) / 100)
	} else if usedMinutes > planMintes && usedMessages <= planMessages {
		minutesCost = (usedMinutes - planMintes) * addMinutes
		tax = (planCost + minutesCost + support) * commonTax

		fmt.Println("Additional minute charge $", math.Floor(minutesCost * 100) / 100)
	}
	
	totalCost := planCost + minutesCost + messageCost + support + tax

	fmt.Println("Base charge $", math.Floor(planCost * 100) / 100)
	fmt.Println("The 911 fee $", math.Floor(support * 100) / 100)
	fmt.Println("Tax $", math.Floor(tax * 100) / 100)
	fmt.Println("Total bill amount $", math.Floor(totalCost * 100) / 100)

}

func readInput() float64 {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	data, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}
	return data

}