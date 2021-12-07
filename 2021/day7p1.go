package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// readInput reads a whole file into memory
// and returns a slice of its lines.
func readInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	var s2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v := scanner.Text()
		//check(err)
		lines = append(lines, v)
	}
	for _, line := range lines {
		s := strings.Split(line, ",") // Split at ","
		for _, in := range s {
			i, err := strconv.Atoi(strings.TrimSpace(in)) // Chomp whitepspace and convert to an Int
			check(err)
			s2 = append(s2, i)
		}
	}
	return s2, scanner.Err()
}

func main() {
	fmt.Println("Advent of Code 2021 Day 7 Part 1")
	positions, err := readInput("day7.input")
	check(err)

	sort.Ints(positions)
	median := positions[len(positions)/2]
	answer := 0

	for _, d := range positions {
		distance := d - median
		if distance < 0 {
			distance = distance * -1
		}
		answer = answer + distance
	}
	fmt.Println("Answer:", answer)
}
