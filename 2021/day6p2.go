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

func printFish(fish []int, day int) {
	fmt.Println("Day", day, "Fish", fish)
}

func main() {
	fmt.Println("Advent of Code 2021 Day 6 Part 1")
	fish, err := readInput("day6.input")
	check(err)

	days := 256
	answer := 0
	var counts [9]int
	var counts2 [9]int
	for _, f := range fish {
		counts[f] += 1
	}

	for d := 1; d <= days; d++ {
		counts2[0] = counts[1]
		counts2[1] = counts[2]
		counts2[2] = counts[3]
		counts2[3] = counts[4]
		counts2[4] = counts[5]
		counts2[5] = counts[6]
		counts2[6] = counts[7] + counts[0]
		counts2[7] = counts[8]
		counts2[8] = counts[0]
		counts = counts2
		answer = 0
		for _, c := range counts {
			answer = answer + c
		}
		fmt.Println("Day:", d, "Answer:", answer)
	}
}
