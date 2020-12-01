package main

import (
	"fmt"
	"strconv"
	"strings"
)

 lk k,const lower = 124075
const upper = 580769 // input:580769

func main() {
	fmt.Println("Advent of Code Day 4")

	answer := make(map[int]int)

	// Check for never decreasing digits
	for a := lower; a < upper; a++ {
		tmp := verifyDigits(a)
		if tmp != 0 {
			answer[tmp] = 0
		}
	}

	// Check if adjacent digits are the same first
	for a, _ := range answer {
		tmp := adjacentDigits(a)
		if tmp == 0 {
			delete(answer, a)
		}
	}

	// part 2 Check that matching digits are not part of a larger group
	for a, _ := range answer {
		tmp := part2(a)
		if tmp == 0 {
			delete(answer, a)
		}
	}

	fmt.Println("answer length is ", len(answer))
}

// Requirement 3, Two adjacent digits are the same.
func adjacentDigits(a int) int {
	tmp := strings.Split(strconv.Itoa(a), "")
	for i := 0; i < 5; i++ {
		if tmp[i] == tmp[i+1] {
			return a
		}
	}
	return 0
}

// Requirement 4, Going from left to right, the digits never decrease.
func verifyDigits(d int) int {
	var counter = 0
	tmp := strings.Split(strconv.Itoa(d), "")
	for i := 0; i < 5; i++ {
		if tmp[i+1] >= tmp[i] {
			counter += 1
		}
	}
	if counter == 5 {
		return d
	}
	return 0
}

// Part 2, Adjacent Digits cannot be groups larger than 2
func part2(a int) int {
	tmp := strings.Split(strconv.Itoa(a), "")

	//check if there are more than one set of doubles
	count := 0
	f := make(map[string]int)
	for i := 0; i < 5; i++ {
		if tmp[i] == tmp[i+1] {
			count++
			f[tmp[i]] += 1
		}
	}

	for _, v := range f { // 1 or more single digit pairs (i.e. 144789, 577788)
		if v == 1 {
			return a
		}
	}

	if len(f) == 1 { // Only 1 digit pair, but it's repeated more than once (i.e. 366679, 244447)
		for _, v := range f {
			if v > 1 {
				return 0
			}
		}
	}

	return 0
}
