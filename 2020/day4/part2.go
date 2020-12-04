package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 4 Part 1")
	input, _ := ioutil.ReadFile("input.txt")
	rows := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	finalcount := 0

	for _, p := range rows { // loop through each passport
		validcount := 0
		p = strings.Replace(p, "\n", " ", -1) // Remove newlines to make it more readable, not strictly required.
		passport := make(map[string]string)
		fields := strings.Split(p, " ") // split into fields by whitespace i.e. "iyr:2011" is one field
		for _, field := range fields {
			pair := strings.Split(field, ":") // split field into a key:value pair
			passport[string(pair[0])] = string(pair[1])
		}
		fmt.Println("Passport: ", passport)

		if validate_byr(passport["byr"]) {
			validcount++
		}

		if validate_iyr(passport["iyr"]) {
			validcount++
		}

		if validate_eyr(passport["eyr"]) {
			validcount++
		}

		if validate_hgt(passport["hgt"]) {
			validcount++
		}

		if validate_hcl(passport["hcl"]) {
			validcount++
		}

		if validate_ecl(passport["ecl"]) {
			validcount++
		}

		if validate_pid(passport["pid"]) {
			validcount++
		}
		if validcount == 7 {
			finalcount++
		}
	}
	fmt.Println("Answer is: ", finalcount)
}

func validate_byr(v string) bool {
	if v == "" {
		return false
	}
	year, _ := strconv.Atoi(v)
	if 1920 <= year && year <= 2002 {
		return true
	}
	return false
}

func validate_iyr(v string) bool {
	if v == "" {
		return false
	}
	year, _ := strconv.Atoi(v)
	if 2010 <= year && year <= 2020 {
		return true
	}
	return false
}

func validate_eyr(v string) bool {
	if v == "" {
		return false
	}
	year, _ := strconv.Atoi(v)
	if 2020 <= year && year <= 2030 {
		return true
	}
	return false
}

func validate_hgt(v string) bool {
	if v == "" {
		return false
	}
	if v[len(v)-2:] == "in" {
		h, _ := strconv.Atoi(v[:2])
		if 59 <= h && h <= 76 {
			return true
		} else {
			return false
		}
	}
	if v[len(v)-2:] == "cm" {
		h, _ := strconv.Atoi(v[:3])
		if 150 <= h && h <= 193 {
			return true
		} else {
			return false
		}
	}
	return false
}

func validate_hcl(v string) bool {
	if v == "" {
		return false
	}
	if v[0:1] == "#" && len(v[1:]) == 6 {
		return true
	}
	return false
}

func validate_ecl(v string) bool {
	if v == "" {
		return false
	}
	colors := map[string]int{
		"amb": 1,
		"blu": 1,
		"brn": 1,
		"gry": 1,
		"grn": 1,
		"hzl": 1,
		"oth": 1,
	}
	if colors[v] == 1 {
		return true
	}
	return false
}

func validate_pid(v string) bool {
	if v == "" {
		return false
	}
	if _, err := strconv.Atoi(v); err == nil {
		if len(v) == 9 {
			return true
		}
	}
	return false
}
