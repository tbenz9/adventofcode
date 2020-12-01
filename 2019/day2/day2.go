package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := ioutil.ReadFile("day2.input")
	if err != nil {
		log.Fatal()
	}
	inputString := strings.Split(strings.TrimSuffix(string(inputFile), "\n"), ",")
	input := make([]int, len(inputString))
	staticInput := make([]int, len(inputString))
	for i, v := range inputString {
		input[i], _ = strconv.Atoi(v)
	}

	copy(staticInput, input)

	inputRange := make([]int, 100)

	for noun, _ := range inputRange {
		for verb, _ := range inputRange {

			input[1] = noun
			input[2] = verb

			result := intCode(input)
			//fmt.Println(input)
			if result == 19690720 {
				fmt.Println("Found it! \nNoun:", noun, "\nVerb", verb)
				break
			}
			copy(input, staticInput)
		}
	}
}

func intCode(input []int) int {
	var pos = 0 //position
	for true {

		if input[pos] == 1 { //add
			val1 := input[pos+1]
			val2 := input[pos+2]
			set := input[pos+3]
			input[set] = input[val1] + input[val2]
			pos = pos + 4
			continue
		}

		if input[pos] == 2 { //multiply
			val1 := input[pos+1]
			val2 := input[pos+2]
			set := input[pos+3]
			input[set] = input[val1] * input[val2]
			pos = pos + 4
			continue
		}

		if input[pos] == 99 { //exit
			pos = 0
			return input[0]
		}

		// Should never get here, but this prevents an infinite loop
		if input[pos] != 1 || input[pos] != 2 || input[pos] != 99 {
			fmt.Println("input[", pos, "] is ", input[pos])
			break
		}

	}
	return 0
}
