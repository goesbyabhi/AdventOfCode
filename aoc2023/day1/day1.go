package main

import ("bufio"; "fmt"; "os"; "strconv")

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("Errir opening file: ", err)
    return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  p1 := 0
  p2 := 0

  for scanner.Scan() {
    line := scanner.Text()
    l1, l2 := processLine(line)
    p1 += l1
    p2 += l2
  }

  if err := scanner.Err(); err != nil {
    fmt.Println("Error reading file: ", err)
    return
  }

  fmt.Println("Sum1: ", p1)
  fmt.Println("Sum2: ", p2)
}

func processLine(line string) (int, int) {
  var p1Process, p2Process []rune

  for i, c := range line {
    if c >= '0' && c <= '9' {
      p1Process = append(p1Process, c)
      p2Process = append(p2Process, c)
    } else {
      for d, val := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}{
        if len(line[i:]) >= len(val) && line[i:i+len(val)] == val {
          p2Process = append(p2Process, []rune(strconv.Itoa(d+1))...)
        }
      }
    }
  }
  p1, _ := strconv.Atoi(string(p1Process[0]) + string(p1Process[len(p1Process)-1]))
	p2, _ := strconv.Atoi(string(p2Process[0]) + string(p2Process[len(p2Process)-1]))

  return p1, p2
}
