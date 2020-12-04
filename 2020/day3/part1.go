package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("day 3 part 1")
	input, _ := ioutil.ReadFile("day3.input")
	rows := strings.Split(strings.TrimSpace(string(input)), "\n")
	var count, column int
	for i := 0; i < len(rows); i++ {
		if column >= len(rows[i]) {
			column = column - len(rows[i])
		}
		if string(rows[i][column]) == "#" {
			count++
		}
		column = column + 3
	}
	fmt.Println("answer:", count)
}
