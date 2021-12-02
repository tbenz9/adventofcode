package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	fmt.Println("Day 2 part 1")
	var depth, horizontal int = 0, 0
	for _, line := range readLines("day2.input") {
		splits := strings.Fields(line)
		distance, _ := strconv.Atoi(splits[1])
		switch splits[0] {
		case "forward":
			horizontal = horizontal + distance
		case "up":
			depth = depth - distance
		case "down":
			depth = depth + distance

		}
	}
	fmt.Println("horizontal = ", horizontal)
	fmt.Println("depth = ", depth)
	fmt.Println("Answer is: ", horizontal*depth)
}
