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
	fmt.Print("Enter a sond level in decibels: ")
	scan.Scan()
	soundLevel, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}

	switch soundLevel {

	case 130:
		fmt.Println(soundLevel, "decibel level - Jackhammer noise")
	case 106:
		fmt.Println(soundLevel, "decibel level - Gas lawnmower noise")
	case 70:
		fmt.Println(soundLevel, "decibel level - Alarm clock noise")
	case 40:
		fmt.Println(soundLevel, "decibel level - Quiet room noise")

	}

	if soundLevel > 130 {
		fmt.Println(soundLevel, "dB the larger than maximal 130dB")
	} else if soundLevel > 106 && soundLevel < 130 {
		fmt.Println(soundLevel, "decibel level - between Gas lawnmower and Jackhammer noise")
	} else if soundLevel > 70 && soundLevel < 106 {
		fmt.Println(soundLevel, "decibel level - between Alarm clock and Gas lawnmower noise")
	} else if soundLevel > 40 && soundLevel < 70 {
		fmt.Println(soundLevel, "decibel level - between Quiet room and Alarm clock noise")
	} else if soundLevel < 40 {
		fmt.Println(soundLevel, "dB the loudest than minimal 40dB")
	}

}