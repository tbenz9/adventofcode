package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v := scanner.Text()
		//check(err)
		lines = append(lines, v)
	}
	return lines, scanner.Err()
}

func main() {
	fmt.Println("Advent of Code 2021 Day 3 Part 1")
	lines, err := readLines("day3.input")
	check(err)
	var zeros, ones [12]int
	var gamma, epsilon string
	for _, line := range lines {
		for k, bit := range strings.Split(line, "") {
			if bit == "0" {
				zeros[k] = zeros[k] + 1
			} else {
				ones[k] = ones[k] + 1
			}

		}
	}

	for k, _ := range zeros {
		if zeros[k] > ones[k] {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		} else {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		}
	}
	gammaint, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonint, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println("gamma:", gammaint)
	fmt.Println("epsilon:", epsilonint)
	fmt.Println("Answer is:", gammaint*epsilonint)
}
