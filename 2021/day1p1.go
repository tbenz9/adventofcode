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
	fmt.Println("Advent of Code 2021 Day 1 Part 1")
	lines, err := readLines("day1.input")
	check(err)
	count := 0
	for i, line := range lines {
		if (i + 1) == len(lines) {
			fmt.Println("answer:", count)
			break
		}
		if lines[i+1] > line {
			fmt.Println("Increment", lines[i+1], "->", line)
			count += 1
		}
	}
}
