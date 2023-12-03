package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("Error opening file: ", err)
    return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  gIDregex := regexp.MustCompile("[^A-Za-z\\s][0-9]*")
  ans := 0
  powerSum := 0

  for scanner.Scan() {
    line := scanner.Text()
    currID := getGameID(line, gIDregex)
    redCount, greenCount, blueCount := getColrs(line)
    if redCount <= 12 && greenCount <= 13 && blueCount <= 14 {
      ans += currID
    }
    calcPower := calculateMinSetPower(line)
		powerSum += calcPower
  }

  if err := scanner.Err(); err != nil {
    fmt.Println("Error reading file's lines: ", err)
  }

  fmt.Println("Sum: ", ans)
  fmt.Println("Power Sum: ", powerSum)
}

func getGameID(line string, gIDregex *regexp.Regexp) int {
  currID := gIDregex.FindString(line[:8])
  numcurrID, _ := strconv.Atoi(currID)
  return numcurrID
}

func getColrs(line string) (int, int, int) {
  redCount := getCount(line, "[0-9]{1,2} red")
  greenCount := getCount(line, "[0-9]{1,2} green")
  blueCount := getCount(line, "[0-9]{1,2} blue")

  return redCount, greenCount, blueCount
}

func getCount(line string, colorRegex string) int {
  count := 0
  regex := regexp.MustCompile(colorRegex)

  semiColon := strings.Split(line, ";")
  for _, substring := range semiColon {
    for _, element := range regex.FindAllString(substring, -1) {
      if number, err := strconv.Atoi(strings.TrimSpace(element[:2])); err == nil && number > count {
        count = number
      }
    }
  }
  return count
}
// p2
func calculateMinSetPower(line string) int {
	redMin := -1
	greenMin := -1
	blueMin := -1

	subsetPattern := regexp.MustCompile(`(\d+) (\w+)`)
	subsetMatches := subsetPattern.FindAllStringSubmatch(line, -1)

	for _, subset := range subsetMatches {
		count, err := strconv.Atoi(subset[1])
		if err != nil {
			panic(err)
		}

		color := subset[2]
		switch color {
		case "red":
			if redMin == -1 || count > redMin {
				redMin = count
			}
		case "green":
			if greenMin == -1 || count > greenMin {
				greenMin = count
			}
		case "blue":
			if blueMin == -1 || count > blueMin {
				blueMin = count
			}
		}
	}
	return redMin * greenMin * blueMin
}
