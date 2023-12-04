package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	Solve("input.txt")
}

func Solve(fname string) {
	readFile, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

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
	fmt.Printf("\nPart 1: %v\nPart 2: %v\n\n", part1Total, part2Total)
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
	firstDigit, lastDigit := "", ""

	for i, char := range encodedLine {
		_, err := strconv.Atoi(string(char))
		if err == nil {
			if firstDigit == "" {
				firstDigit = string(char)
			}
			lastDigit = string(char)
		} else if char == 'z' || char == 'o' || char == 't' || char == 'f' || char == 's' || char == 'e' || char == 'n' {
			writtenDigit := parseWrittenDigit(encodedLine, i)
			if writtenDigit != "" {
				if firstDigit == "" {
					firstDigit = writtenDigit
				}
				lastDigit = writtenDigit
			}
		}
	}

	finalDigit, err := strconv.Atoi(fmt.Sprintf("%v%v", firstDigit, lastDigit))
	if err != nil {
		return 0
	}

	return finalDigit
}

func parseWrittenDigit(fullText string, startingIndex int) string {
	spelledOutDigits := []string{
		"zero",
		"one", "two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	subLine := fullText[startingIndex:]

	for digit, spelledDigit := range spelledOutDigits {
		if i := strings.Index(subLine, spelledDigit); i == 0 {
			return fmt.Sprint(digit)
		}
	}
	return ""
}
