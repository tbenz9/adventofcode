package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	fmt.Println("Day 5 Part 1")
	input, _ := ioutil.ReadFile("input.txt")
	seats := strings.Split(strings.TrimSpace(string(input)), "\n")

	var seatID, answer float64
	for _, seat := range seats {
		row := getRow(seat[:7])
		col := getCol(seat[7:])
		seatID = row*8 + col
		if seatID > answer {
			answer = seatID
		}
	}
	fmt.Println("Answer is:", answer)
}

func getRow(v string) float64 {
	chars := strings.Split(v, "")
	var min, max float64 = 0, 127
	for i, char := range chars {
		if char == "F" {
			max = max - ((max - min) / 2)
		}
		if char == "B" {
			min = min + ((max - min) / 2)
		}
		if i == 6 && char == "F" {
			return math.Floor(min + 1)
		}
	}
	return math.Floor(max)
}

func getCol(v string) float64 {
	chars := strings.Split(v, "")
	var min, max float64 = 0, 7
	for i, char := range chars {
		if char == "R" {
			min = min + ((max - min) / 2)
		}
		if char == "L" {
			max = max - ((max - min) / 2)
		}
		if i == 3 && char == "L" {
			return math.Floor(min + 1)
		}
	}
	return math.Floor(max)
}
