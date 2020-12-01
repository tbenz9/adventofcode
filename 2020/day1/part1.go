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
	fmt.Println("Day 1 Part 1")
	lines, err := readLines("part1.input")
	check(err)
	for _, line1 := range lines {
		for _, line2 := range lines {
			if line1+line2 == 2020 {
				fmt.Println("value 1 is ", line1, "value 2 is ", line2)
				fmt.Println("Answer is: ", line1*line2)
				os.Exit(0)
			}
		}
	}
}
