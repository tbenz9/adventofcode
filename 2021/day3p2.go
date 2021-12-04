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

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v := scanner.Text()
		//check(err)
		lines = append(lines, v)
	}
	return lines, scanner.Err()
}

// returns 0 if there are more zeros than ones
// returns 1 if there are more ones than zeros
// returns 2 if their is an equal amount of ones and zeros
func moarZeros(lines []string, col int) int {
	var zeros, ones int = 0, 0
	for _, line := range lines {
		bits := strings.Split(line, "")
		if bits[col] == "0" {
			zeros += 1
		} else {
			ones += 1
		}
	}
	if zeros == ones {
		return 2
	} else if zeros > ones {
		return 0
	}
	return 1
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	fmt.Println("Advent of Code 2021 Day 3 Part 2")
	generator, err := readLines("day3.input")
	scrubber, err := readLines("day3.input")
	check(err)

	var generatorint, scrubberint int64
	loops := 12

	// Generator
	for i := 0; i < loops; i++ {
		keep := moarZeros(generator, i)
		if keep == 2 {
			keep = 1
		}
		if keep == 0 {
			// keep the zeros
			for j := len(generator) - 1; j >= 0; j-- {
				bits := strings.Split(generator[j], "")
				if bits[i] != "0" {
					//fmt.Println("removing", bits, "because it has a 1 in spot", i)
					generator = remove(generator, j)
				}
			}
		} else {
			// Keep the ones
			for j := len(generator) - 1; j >= 0; j-- {
				bits := strings.Split(generator[j], "")
				if bits[i] != "1" {
					generator = remove(generator, j)
					//fmt.Println("removing", bits, "because it has a 0 in spot", i)
				}
			}
		}
		//fmt.Println("Generator:", generator)
		if len(generator) == 1 {
			generatorint, _ = strconv.ParseInt(generator[0], 2, 64)
			break
		}
	}

	// Scrubbers!
	for i := 0; i < loops; i++ {
		keep := moarZeros(scrubber, i)
		if keep == 2 {
			keep = 1
		}
		if keep == 1 {
			// Keep the zeros
			for j := len(scrubber) - 1; j >= 0; j-- {
				bits := strings.Split(scrubber[j], "")
				if bits[i] != "0" {
					//fmt.Println("removing", bits, "because it has a 1 in spot", i)
					scrubber = remove(scrubber, j)
				}
			}
		} else {
			// Keep the ones
			for j := len(scrubber) - 1; j >= 0; j-- {
				bits := strings.Split(scrubber[j], "")
				if bits[i] != "1" {
					scrubber = remove(scrubber, j)
					//fmt.Println("removing", bits, "because it has a 0 in spot", i)
				}
			}
		}
		if len(scrubber) == 1 {
			scrubberint, _ = strconv.ParseInt(scrubber[0], 2, 64)
			break
		}
	}

	fmt.Println("generator:", generatorint)
	fmt.Println("scrubber:", scrubberint)
	fmt.Println("Answer is:", generatorint*scrubberint)
}
