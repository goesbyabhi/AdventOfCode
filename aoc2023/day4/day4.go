package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	return txtlines
}

func countWins(line string) float64 {
	var cardWins float64

	parts := strings.Split(line, ": ")[1]
	winningNums, numsIHave := strings.Fields(strings.Split(parts, " | ")[0]), strings.Fields(strings.Split(parts, " | ")[1])

	for _, num := range numsIHave {
		for _, wNum := range winningNums {
			if num == wNum {
				cardWins++
			}
		}
	}
	return cardWins
}

// END HELPERS

func part1(fileLines []string) {
	var result float64

	for _, line := range fileLines {
		cardWins := countWins(line)
		if cardWins > 0 {
			result += math.Pow(2, cardWins-1)
		}
	}

	fmt.Println("Part 1 result:", result)
}

func part2(fileLines []string) {
	var result int

	lastLine := fileLines[len(fileLines)-1]
	maxCard, _ := strconv.Atoi(strings.Fields(strings.Split(lastLine, ":")[0])[1])

	cardEarned := make([]int, maxCard)
	for i := range cardEarned {
		cardEarned[i] = 1
	}

	for cardIndex, line := range fileLines {
		countWins := int(countWins(line))
		for i := cardIndex + 1; i <= cardIndex+countWins; i++ {
			cardEarned[i] += cardEarned[cardIndex]
		}
	}

	result = 0
	for _, val := range cardEarned {
		result += val
	}

	fmt.Println("Part 2 result:", result)
}

func main() {
	fileLines := readFile("input.txt")
	part1(fileLines)
	part2(fileLines)
}

