package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	boardsize = 1000
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// readInput reads a whole file into memory
// and returns a slice of its lines.
func readInput(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	var grid [][]int
	var c2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v := scanner.Text()
		lines = append(lines, v)
	}
	for _, line := range lines {
		s := strings.Split(line, "") // Split each char
		for _, ints := range s {
			i, err := strconv.Atoi(strings.TrimSpace(ints)) // Chomp whitepspace and convert to an Int
			check(err)
			c2 = append(c2, i)
		}
		grid = append(grid, c2)
		c2 = nil
	}
	return grid, scanner.Err()
}

func main() {
	fmt.Println("Advent of Code 2021 Day 9 Part 1")
	grid, err := readInput("day9.input")
	check(err)
	var lowpoints []int
	answer := 0
	for r, row := range grid { // loop through each grid row
		for c, col := range row { // loop through each row column
			lowpointCount := 0
			possibleLowpoints := 0
			// Check if the number above is larger
			if r > 0 {
				possibleLowpoints += 1
				if grid[r-1][c] > col {
					lowpointCount += 1
				}
			}
			// Check if number below is larger
			if r < len(grid)-1 {
				possibleLowpoints += 1
				if grid[r+1][c] > col {
					lowpointCount += 1
				}
			}
			// Check if number to left is larger
			if c > 0 {
				possibleLowpoints += 1
				if grid[r][c-1] > col {
					lowpointCount += 1
				}
			}
			// Check if number to the right is larger
			if c < len(row)-1 {
				possibleLowpoints += 1
				if grid[r][c+1] > col {
					lowpointCount += 1
				}
			}
			if lowpointCount == possibleLowpoints {
				lowpoints = append(lowpoints, col)
			}
		}
	}

	for _, point := range lowpoints {
		answer = answer + point + 1
	}
	fmt.Println("Answer:", answer)
}
