package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("day 3 part 22")
	input, _ := ioutil.ReadFile("day3.input")
	rows := strings.Split(strings.TrimSpace(string(input)), "\n")

	paths := [5][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2}}

	var count, column int
	final := 1

	for _, p := range paths {
		for i := 0; i < len(rows); i = i + p[1] {
			if column >= len(rows[i]) {
				column = column - len(rows[i])
			}
			if string(rows[i][column]) == "#" {
				count++
			}
			column = column + p[0]
		}
		final = final * count
		count, column = 0, 0
	}
	fmt.Println("answer:", final)
}
