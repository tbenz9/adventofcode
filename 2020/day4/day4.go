package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Day 4 Part 1")
	input, _ := ioutil.ReadFile("input.txt")
	rows := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	var count, valid = 0, 0

	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for i := 0; i < len(rows); i++ {
		for _, field := range fields {
			if strings.Contains(rows[i], field) {
				fmt.Println("found a match")
				count++
			}
		}
		if count == 7 {
			valid++
		}
		count = 0
	}
	fmt.Println("Answer: ", valid)
}
