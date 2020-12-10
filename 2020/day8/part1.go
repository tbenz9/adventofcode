package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 8 Part 1")
	input, _ := ioutil.ReadFile("input.txt")
	commands := strings.Split(strings.TrimSpace(string(input)), "\n")

	index, accumulator := 0, 0
	stringSet := make(map[int]bool)
	processCommand(commands[index], index, accumulator, commands, stringSet)
}

func processCommand(ins string, index int, accumulator int, commands []string, stringSet map[int]bool) {
	if stringSet[index] {
		fmt.Println("Answer is:", accumulator)
		os.Exit(0)
	}
	stringSet[index] = true
	instruction := strings.Split(ins, " ")
	op := instruction[0]
	arg, _ := strconv.Atoi(instruction[1])

	switch op {
	case "nop":
		index++
	case "acc":
		accumulator = accumulator + arg
		index++
	case "jmp":
		index = index + arg
	}
	processCommand(commands[index], index, accumulator, commands, stringSet)
}
