package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	adventofcode2023 "github.com/schubam/advent-of-code-2023"
)

func main() {
	fmt.Println(V1())
	fmt.Println(V2())
}

func V1() int {
	input := strings.Trim(adventofcode2023.ReadFile("./cmd/day02/part1.txt"), "\n")
	result := processInput(input)

	return result
}

func V2() int {
	input := strings.Trim(adventofcode2023.ReadFile("./cmd/day02/part1.txt"), "\n")
	result := processInput2(input)

	return result
}

var gamePattern = `Game (\d+)`
var gameRegexp = regexp.MustCompile(gamePattern)

var bluePattern = `(\d+)\sblue`
var blueRegexp = regexp.MustCompile(bluePattern)

var redPattern = `(\d+)\sred`
var redRegexp = regexp.MustCompile(redPattern)

var greenPattern = `(\d+)\sgreen`
var greenRegexp = regexp.MustCompile(greenPattern)

func processInput2(input string) int {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		var maxRed, maxGreen, maxBlue int

		for _, b := range blueRegexp.FindAllStringSubmatch(line, -1) {
			n, _ := strconv.Atoi(b[1])
			if n > maxBlue {
				maxBlue = n
			}
		}
		for _, b := range redRegexp.FindAllStringSubmatch(line, -1) {
			n, _ := strconv.Atoi(b[1])
			if n > maxRed {
				maxRed = n
			}
		}
		for _, b := range greenRegexp.FindAllStringSubmatch(line, -1) {
			n, _ := strconv.Atoi(b[1])
			if n > maxGreen {
				maxGreen = n
			}
		}

		sum += maxBlue * maxGreen * maxRed
	}
	return sum
}

func processInput(input string) int {
	red := 12
	green := 13
	blue := 14

	var sum int
	for _, line := range strings.Split(input, "\n") {
		var maxRed, maxGreen, maxBlue int

		game := gameRegexp.FindStringSubmatch(line)
		gameNo, _ := strconv.Atoi(game[1])

		for _, b := range blueRegexp.FindAllStringSubmatch(line, -1) {
			n, _ := strconv.Atoi(b[1])
			if n > maxBlue {
				maxBlue = n
			}
		}
		for _, b := range redRegexp.FindAllStringSubmatch(line, -1) {
			n, _ := strconv.Atoi(b[1])
			if n > maxRed {
				maxRed = n
			}
		}
		for _, b := range greenRegexp.FindAllStringSubmatch(line, -1) {
			n, _ := strconv.Atoi(b[1])
			if n > maxGreen {
				maxGreen = n
			}
		}

		if maxRed <= red && maxGreen <= green && maxBlue <= blue {
			sum += gameNo
		}
	}
	return sum
}
