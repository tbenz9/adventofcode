package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Day 6 Part 2")
	input, _ := ioutil.ReadFile("input.txt")
	groups := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	totalSum, validform := 0, 0

	for _, g := range groups { // loop through each group
		m := make(map[string]int)
		persons := strings.Split(strings.TrimSpace(string(g)), "\n")
		for _, p := range persons { // loop through each person in the group
			questions := strings.Split(p, "")
			for _, q := range questions { // loop through each question (letter) for this person
				m[q] = m[q] + 1
			}
		}
		for _, v := range m { // loop through the map we made and check if everyone answered the question
			if v == len(persons) {
				validform++
			}
		}
		totalSum = totalSum + validform
		validform = 0
	}
	fmt.Println("Answer is:", totalSum)
}
