package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type game struct {
	id         int
	redTotal   int
	greenTotal int
	blueTotal  int
}

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	red := 12
	green := 13
	blue := 14
	runningTotal := 0
	runningPower := 0
	for scanner.Scan() {
		gameTotals := getGameTotals(scanner.Text())
		if gameTotals.redTotal <= red && gameTotals.greenTotal <= green && gameTotals.blueTotal <= blue {
			runningTotal += gameTotals.id
		}
		gamePower := gameTotals.redTotal * gameTotals.blueTotal * gameTotals.greenTotal
		runningPower += gamePower
	}
	fmt.Printf("Part 1: %v\nPart 2: %v\n", runningTotal, runningPower)
}

func getGameTotals(gameLine string) game {
	initialSplit := strings.Split(gameLine, ": ")
	gameID, err := strconv.Atoi(strings.Replace(initialSplit[0], "Game ", "", 1))
	if err != nil {
		fmt.Println(err)
	}
	gameResults := strings.Split(initialSplit[1], "; ")
	gameTotals := game{
		id: gameID,
	}
	for _, results := range gameResults {
		colors := strings.Split(results, ", ")
		for _, color := range colors {
			value, err := strconv.Atoi(strings.TrimSpace(color[0:2]))
			if err != nil {
				fmt.Println(err)
			}
			if strings.Contains(color, "red") {
				if gameTotals.redTotal < value {
					gameTotals.redTotal = value
				}
			} else if strings.Contains(color, "blue") {
				if gameTotals.blueTotal < value {
					gameTotals.blueTotal = value
				}
			} else if strings.Contains(color, "green") {
				if gameTotals.greenTotal < value {
					gameTotals.greenTotal = value
				}
			}

		}
	}
	return gameTotals
}
