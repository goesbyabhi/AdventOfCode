package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

var (
  total = 0
  board []string
  numsMap = make(map[[2]int][]int)
  pattern = regexp.MustCompile(`\d+`)
  gearTotal = 0
)

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    board = append(board, scanner.Text())
  }

  for rowNum := 0; rowNum < len(board); rowNum++ {
    matches := pattern.FindAllStringIndex(board[rowNum], -1)
    for _, match := range matches {
      if isValidNeighbor(rowNum-1, match[0]-1, rowNum+1, match[1], parseInt(board[rowNum][match[0]:match[1]])) {
        total += parseInt(board[rowNum][match[0]:match[1]])
      }
    }
  }

  for _, numVal := range numsMap {
    if len(numVal) == 2 {
      gearTotal += numVal[0] * numVal[1];
    }
  }

  println(total)
  println(gearTotal)
}

func isValidNeighbor(startY, startX, endY, endX, num int) bool  {
  for y := startY; y <= endY; y++ {
    for x := startX; x <= endX; x++ {
      if y >= 0 && y < len(board) && x >= 0 && x < len(board[y]) {
        if !isInvalid(board[y][x]) {
          if board[y][x] == '*' {
            numsMap[[2]int{y,x}] = append(numsMap[[2]int{y, x}], num)
          }
          return true
        }
      }
    }
  }
  return false
}

func isInvalid(char byte) bool {
  return char >= '0' && char <= '9' || char == '.'
}

func parseInt(s string) int {
  num, err := strconv.Atoi(s)
  if err != nil {
    panic(err)
  }
  return num
}
