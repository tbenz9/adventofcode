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
		//		fmt.Println(line)
		splits := strings.Fields(line)
		var low, _ = strconv.Atoi(strings.Split(splits[0], "-")[0])
		var high, _ = strconv.Atoi(strings.Split(splits[0], "-")[1])
		var letter = strings.TrimSuffix(splits[1], ":")
		var password = strings.Split(splits[2], "")

		if (password[low-1] == letter) != (password[high-1] == letter) {
			fmt.Println("bad password", line)
			count++
		}
	}
	fmt.Println("There are:", count, "valid passwords")
}
