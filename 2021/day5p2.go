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
	var coordinates [][]int
	var c2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v := scanner.Text()
		lines = append(lines, v)
	}
	for _, line := range lines {
		s := strings.Split(line, "->") // Split at "->"
		for _, coor := range s {
			c := strings.Split(coor, ",") // Split at ","
			for _, ints := range c {
				i, err := strconv.Atoi(strings.TrimSpace(ints)) // Chomp whitepspace and convert to an Int
				check(err)
				c2 = append(c2, i)
			}
		}
		coordinates = append(coordinates, c2)
		c2 = nil
	}
	return coordinates, scanner.Err()
}

// Print the Diagram
func printDiagram(diagram [boardsize][boardsize]int) {
	for _, diag := range diagram {
		for _, point := range diag {
			if point == 0 {
				fmt.Print(".")
				continue
			}
			fmt.Print(point)
		}
		fmt.Println("")
	}
}

func main() {
	fmt.Println("Advent of Code 2021 Day 5 Part 2")
	coordinates, err := readInput("day5.input")
	check(err)
	var x1, y1, x2, y2, answer int
	var diagram [boardsize][boardsize]int
	for _, coordinate := range coordinates {
		x1 = coordinate[0]
		y1 = coordinate[1]
		x2 = coordinate[2]
		y2 = coordinate[3]
		//fmt.Println(coordinate)
		if x1 == x2 || y1 == y2 { // Vertical and horizontal lines
			if x1 > x2 { //Draw horizontal line right to left
				for i := x2; i <= x1; i++ {
					diagram[y1][i] += 1
				}
				continue
			} else if x2 > x1 { //Draw horizontal line left to right
				for i := x1; i <= x2; i++ {
					diagram[y1][i] += 1
				}
				continue
			} else if y1 > y2 { //Draw vertical line bottom to top
				for i := y2; i <= y1; i++ {
					diagram[i][x1] += 1
				}
				continue
			} else if y2 > y1 { //Draw vertical line top to bottom
				for i := y1; i <= y2; i++ {
					diagram[i][x1] += 1
				}
				continue
			} else if x1 == x2 && y1 == y2 { // A single point (not actually in the puzzle input)
				diagram[y1][x1] += 1
			} else {
				fmt.Println("This should never get printed")
			}
		} else { // Diagonal lines
			if y2 > y1 {
				if x2 > x1 { //top left -> bottom right
					count := 0
					for i := x1; i <= x2; i++ {
						diagram[y1+count][i] += 1
						count += 1
					}
					continue
				} else { // bottom left -> top right
					count := 0
					for i := x2; i <= x1; i++ {
						diagram[y2-count][i] += 1
						count += 1
					}
					continue
				}
			} else {
				if x2 > x1 { // bottom left -> top right
					count := 0
					for i := x1; i <= x2; i++ {
						diagram[y1-count][i] += 1
						count += 1
					}
					continue
				} else { // bottom right -> top left
					count := 0
					for i := x2; i <= x1; i++ {
						diagram[y2+count][i] += 1
						count += 1
					}
					continue
				}
				fmt.Println("You should also never get here")
			}
			fmt.Println("And you should definitely never get here")
		}
	}
	//printDiagram(diagram)

	for _, row := range diagram {
		for _, column := range row {
			if column > 1 {
				answer += 1
			}
		}
	}
	fmt.Println("Answer:", answer)
}
