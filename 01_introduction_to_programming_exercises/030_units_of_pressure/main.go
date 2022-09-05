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
	fmt.Print("Enter pressure in kilo-pascals: ")
	scan.Scan()
	p, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	psi := p * 0.145038
	mmMerc := p * 0.133322387415
	atm := p * 0.00986923

	fmt.Println(p, "kilo-pascals is", psi, "PSI")
	fmt.Println(p, "kilo-pascals is", mmMerc, "millimeters of mercury")
	fmt.Println(p, "kilo-pascals is", atm, "atmospheres")

}