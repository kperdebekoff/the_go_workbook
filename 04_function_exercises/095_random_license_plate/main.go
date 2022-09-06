package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	fmt.Println(generateLicensePlate())

}

func generateLicensePlate() string {
	
	licPlate := ""
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 3; i++ {
		licPlate += string(rand.Intn(90 - 65 + 1) + 65)
	}

	for i := 0; i < rand.Intn(2) + 3; i++ {
		licPlate += strconv.Itoa(rand.Intn(10))
	}
	return licPlate

}