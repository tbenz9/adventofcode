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
	fish, err := readInput("day6.sample")
	check(err)

	days := 18

	for d := 0; d < days; d++ {
		for i, f := range fish {
			if f == 0 {
				fish[i] = 6
				fish = append(fish, 8)
			} else {
				fish[i] -= 1
			}
		}
		//printFish(fish, d+1)
	}
	fmt.Println("Answer:", len(fish))
}
