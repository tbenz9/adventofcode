package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 9 Part 1")
	input, _ := ioutil.ReadFile("input.txt")
	inputs := strings.Split(strings.TrimSpace(string(input)), "\n")
	numbers := []int{}

	start := 0

	for _, v := range inputs { // convert the input to a slice of ints
		i, _ := strconv.Atoi(v)
		numbers = append(numbers, i)
	}

	for validEntry(start, numbers) {
		fmt.Println(numbers[start+5])
		start++
	}
	fmt.Println("bad one is", numbers[start+5])

}

func validEntry(start int, numbers []int) bool {
	end := start + 4
	index := end + 1
	for x := start; x < end; x++ {
		for y := start; y < end; y++ {
			if x == y {
				continue
			}
			if numbers[x]+numbers[y] == numbers[index] {
				return true
			}

		}

	}
	return false
}
