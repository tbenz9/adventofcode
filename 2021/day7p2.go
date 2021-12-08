package main

import (
	"bufio"
	"fmt"
	"math"
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
	fmt.Println("Advent of Code 2021 Day 7 Part 2")
	positions, err := readInput("day7.input")
	check(err)

	sort.Ints(positions)
	var answer, sum = 0, 0
	var cost [2000]int

	for i, _ := range cost { // make a slice of the fuel cost where index is the distance traveled.
		if i == 0 {
			cost[0] = 0
			continue
		}
		cost[i] = cost[i-1] + i
	}

	for _, m := range positions { // calculate the sum of all crabs location so I can get the mean
		sum = sum + m
	}

	// I don't understand why, but on the sample data this needs to be "math.Ceil"
	// and on the full input data it needs to be "math.Floor".
	mean := math.Floor(float64(sum) / float64(len(positions)))

	fmt.Println("sum", sum)
	fmt.Println("len", len(positions))
	fmt.Println("mean", mean)

	for _, d := range positions {
		distance := d - int(mean)
		if distance < 0 {
			distance = distance * -1
		}
		answer = answer + cost[distance]
		fmt.Println("Move from", d, "to", int(mean), ":", cost[distance], "fuel")
	}
	fmt.Println("Answer:", answer)
}
