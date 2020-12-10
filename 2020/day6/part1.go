package main

import (
	"fmt"
	"io/ioutil"
	//"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 6 Part 1")
	input, _ := ioutil.ReadFile("input.txt")
	groups := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	totalSum := 0

	for _, g := range groups { // loop through each group
		m := make(map[string]int)
		persons := strings.Split(strings.TrimSpace(string(g)), "\n")
		for _, p := range persons { // loop through each person in the group
			questions := strings.Split(p, "")
			for _, q := range questions { // loop through each question (letter) for this person
				m[q] = m[q] + 1
			}
		}
		//fmt.Println(m)
		totalSum = totalSum + len(m)
	}
	fmt.Println("Answer is:", totalSum)
}
