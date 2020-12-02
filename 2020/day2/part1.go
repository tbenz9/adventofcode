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
	fmt.Println("Day 2")
	count := 0
	for _, line := range readLines("part1.input") {
		splits := strings.Fields(line)
		var low, _ = strconv.Atoi(strings.Split(splits[0], "-")[0])
		var high, _ = strconv.Atoi(strings.Split(splits[0], "-")[1])
		var letter = strings.TrimSuffix(splits[1], ":")
		var password = splits[2]
		var charcount = strings.Count(password, letter)
		if charcount >= low && charcount <= high {
			// fmt.Println("bad password", line)
			count++
		}
	}
	fmt.Println(count)
}
