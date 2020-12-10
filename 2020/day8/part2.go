package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Day 8 Part 2")
	input, _ := ioutil.ReadFile("input.txt")
	commands := strings.Split(strings.TrimSpace(string(input)), "\n")

	index, accumulator := 0, 0
	jmps := []int{}

	// Make a slice containing the index values that are jmp
	for k, v := range commands {
		instruction := strings.Split(v, " ")
		op := instruction[0]
		if op == "jmp" {
			jmps = append(jmps, k)
		}
	}

	for _, v := range jmps {
		commands = strings.Split(strings.TrimSpace(string(input)), "\n")
		stringSet := make(map[int]bool)
		commands[v] = "nop +0"
		go processCommand(commands[index], index, accumulator, commands, stringSet)
	}
	time.Sleep(10 * time.Second)
	//processCommand(commands[index], index, accumulator, commands, stringSet)
}

func processCommand(ins string, index int, accumulator int, commands []string, stringSet map[int]bool) {
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
	//fmt.Println("ins:", ins, "index:", index, "commands", len(commands))
	if index == len(commands) {
		fmt.Println("Answer is:", accumulator)
		os.Exit(0)
	}
	processCommand(commands[index], index, accumulator, commands, stringSet)
}
