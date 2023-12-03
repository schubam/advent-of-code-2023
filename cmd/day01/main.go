package main

import (
	"fmt"
	"strconv"
	"strings"

	adventofcode2023 "github.com/schubam/advent-of-code-2023"
)

var substitutions = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "th3ee",
	"four":  "fo4r",
	"five":  "fi5e",
	"six":   "s6x",
	"seven": "se7en",
	"eight": "ei8ht",
	"nine":  "ni9e",
}

func main() {
	fmt.Println(V1())
	fmt.Println(V2())
}

func V2() int {
	input := strings.Trim(adventofcode2023.ReadFile("./cmd/day01/v1.txt"), "\n")
	lines := strings.Split(input, "\n")

	var sum int
	for i, l := range lines {
		for key, value := range substitutions {
			l = strings.ReplaceAll(l, key, value)
		}
		vals := extractDigits(l)
		fmt.Printf("Line %d (%d, %d) = %s (%s)\n", i, vals[0], vals[1], lines[i], l)
		sum += combineDigits(vals)
	}
	return sum
}

func V1() int {
	input := strings.Trim(adventofcode2023.ReadFile("./cmd/day01/v1.txt"), "\n")
	lines := strings.Split(input, "\n")
	var sum int
	for _, line := range lines {
		vals := extractDigits(line)
		sum += combineDigits(vals)
	}
	return sum
}

func combineDigits(digits [2]int) int {
	combinedNumber, err := strconv.Atoi(fmt.Sprintf("%d%d", digits[0], digits[1]))
	if err != nil {
		fmt.Println("error combining digits into number", err)
		return 0
	}
	return combinedNumber
}

func extractDigits(s string) [2]int {
	cleanStr := ""
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			cleanStr += string(ch)
		}
	}

	if cleanStr == "" {
		return [2]int{0, 0}
	}

	first := string(cleanStr[0])
	last := string(cleanStr[len(cleanStr)-1])
	firstN, err := strconv.Atoi(first)
	if err != nil {
		fmt.Println("error converting string to number", err)
		return [2]int{0, 0}
	}
	lastN, err := strconv.Atoi(last)
	if err != nil {
		fmt.Println("error converting string to number", err)
		return [2]int{0, 0}
	}
	return [2]int{firstN, lastN}
}

