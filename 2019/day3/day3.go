package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type dir struct {
	direction string
	distance  int
}

type coordinate struct {
	x int
	y int
}

func main() {
	fmt.Println("Advent of Code Day 3")
	wi := loadFile()

	cs := make([][]coordinate, 0)
	for _, x := range wi { // get coordinates for each wire
		cs = append(cs, createCoordinates(x))
	}

	shortest := 1000000000000
	intersection := 1000000000000

	for i := 0; i < len(cs[0]); i++ { // loop through first array
		for j := 0; j < len(cs[1]); j++ { // loop through 2nd array
			if cs[0][i] == cs[1][j] {
				md := Abs(cs[0][i].x) + Abs(cs[0][i].y)
				if md != 0 && md < shortest {
					shortest = md
				}
				inter := i + j
				if inter != 0 && inter < intersection {
					intersection = inter
				}
				fmt.Println("intersection found at", cs[0][i], "with a manhattan distance of ", md, "and intersection distance of ", inter)
			}
		}
	}

	fmt.Println("Shortest manhattan distance is", shortest)
	fmt.Println("Shortest intersection distance is", intersection)

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func createCoordinates(x []dir) []coordinate {
	coordinates := []coordinate{coordinate{}}
	for _, d := range x {
		switch op := d.direction; op {
		case "U":
			for u := 0; u < d.distance; u++ {
				last := coordinates[len(coordinates)-1]
				coordinates = append(coordinates, coordinate{last.x, last.y + 1})
			}
		case "D":
			for u := 0; u < d.distance; u++ {
				last := coordinates[len(coordinates)-1]
				coordinates = append(coordinates, coordinate{last.x, last.y - 1})
			}
		case "R":
			for u := 0; u < d.distance; u++ {
				last := coordinates[len(coordinates)-1]
				coordinates = append(coordinates, coordinate{last.x + 1, last.y})
			}
		case "L":
			for u := 0; u < d.distance; u++ {
				last := coordinates[len(coordinates)-1]
				coordinates = append(coordinates, coordinate{last.x - 1, last.y})
			}
		}
	}
	return coordinates
}

func loadFile() [][]dir {

	inputFile, err := ioutil.ReadFile("day3.input")
	if err != nil {
		log.Fatal()
	}

	output := make([][]dir, 2)
	inputString := strings.Split(strings.TrimSuffix(string(inputFile), "\n"), "\n")
	for index, wire := range inputString { // loop through the 2 "wires"
		operations := strings.Split(wire, ",")
		for _, op := range operations { // loop through each direction in each wire
			length, _ := strconv.Atoi(op[1:])
			tmp := dir{op[0:1], length}
			output[index] = append(output[index], tmp)
		}
	}
	return output
}
