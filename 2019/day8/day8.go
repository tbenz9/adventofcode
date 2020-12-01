package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	pixels := 25 * 6
	layer := 100
	decoded := make([]int, pixels)
	numZeros, numOnes, numTwos := 0, 0, 0

	is := loadFile()
	for i := len(is); i > 0; i = i - pixels {
		numZeros, numOnes, numTwos = countDigits(is[i-pixels : i])
		fmt.Println("layer", layer, "contains", numZeros, "zeros, so the answer is", numOnes*numTwos)
		index := 0
		for j := i - pixels; j < i; j++ {
			if is[j] == 0 {
				decoded[index] = 0
			}
			if is[j] == 1 {
				decoded[index] = 1
			}
			index++
		}
		layer--
	}

	// print out the message
	for i := 0; i < len(decoded); i++ {
		if i%25 == 0 {
			fmt.Println("")
		}
		if decoded[i] == 0 {
			fmt.Print(" ")
		} else if decoded[i] == 1 {
			fmt.Print("#")
		}
	}
	fmt.Println("")
}

func countDigits(l []int) (int, int, int) {
	zeros, ones, twos := 0, 0, 0
	for _, v := range l {
		if v == 0 {
			zeros++
		}
		if v == 1 {
			ones++
		}
		if v == 2 {
			twos++
		}
	}
	return zeros, ones, twos
}

func loadFile() []int {
	inputFile, err := ioutil.ReadFile("day8.input")
	if err != nil {
		log.Fatal()
	}
	tmp := strings.Split(strings.TrimSuffix(string(inputFile), "\n"), "")
	input := make([]int, len(tmp))
	for i, v := range tmp {
		input[i], _ = strconv.Atoi(v)
	}
	return input
}
