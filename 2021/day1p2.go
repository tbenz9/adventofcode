package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		check(err)
		lines = append(lines, v)
	}
	return lines, scanner.Err()
}

func main() {
	fmt.Println("Advent of Code 2021 Day 1 Part 2")
	lines, err := readLines("day1.input")
	check(err)
	count := 0
	for i, _ := range lines {
		if (i + 3) == len(lines) {
			fmt.Println("answer:", count)
			break
		}
		current := lines[i] + lines[i+1] + lines[i+2]
		next := lines[i+1] + lines[i+2] + lines[i+3]

		if next > current {
			fmt.Println("Increment", next, "->", current)
			count += 1
		}
	}
}
