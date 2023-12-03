package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {}

func Solve(fname string) {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	part1Total := 0
	part2Total := 0
	for fileScanner.Scan() {
		encodedLine := fileScanner.Text()
		part1Solution := Part1(encodedLine)
		part2Solution := Part2(encodedLine)

		part1Total += part1Solution
		part2Total += part2Solution
	}
	fmt.Printf("Part 1: %v\nPart 2: %v", part1Total, part2Total)
}

func Part1(encodedLine string) int {
	firstDigit, lastDigit := "", ""

	for _, char := range encodedLine {
		_, err := strconv.Atoi(string(char))
		if err == nil {
			if firstDigit == "" {
				firstDigit = string(char)
			}
			lastDigit = string(char)
		}
	}

	finalDigit, err := strconv.Atoi(fmt.Sprintf("%v%v", firstDigit, lastDigit))
	if err != nil {
		return 0
	}

	return finalDigit
}

func Part2(encodedLine string) int {
	return 0
}
