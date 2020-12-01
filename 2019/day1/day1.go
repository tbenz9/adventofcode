package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var totalModuleFuel = 0
var totalFuel = 0

func main() {
	file, err := os.Open("day1.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		weight, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		calcFuel(weight)
		totalModuleFuel = 0
	}

	if err := scanner.Err(); err != nil {
		log.Fatal()
	}
	fmt.Println("Sum of Fuel: ", totalFuel)
}

func calcFuel(weight int) int {
	fuel := weight/3 - 2
	if fuel <= 0 {
		fmt.Println("exiting with fuel:", totalModuleFuel)
		totalFuel = totalFuel + totalModuleFuel
		return totalModuleFuel
	}
	totalModuleFuel = totalModuleFuel + fuel
	fmt.Println("fuel:", fuel, "\ttotalModuleFuel:", totalModuleFuel)
	calcFuel(fuel)
	return 0
}
