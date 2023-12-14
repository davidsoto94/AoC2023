package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	readFile, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	symbolAdjasent := false
	sum := 0
	for i, line := range lines {
		currNumber := "0"
		for j := 0; j < len(line); j++ {
			digit, err := strconv.Atoi(string(line[j]))
			if err != nil {
				if currNumber == "0" {
					symbolAdjasent = false
					continue
				}
				if !symbolAdjasent {
					currNumber = "0"
					continue
				}
				val, _ := strconv.Atoi(currNumber)
				sum += val
				symbolAdjasent = false
				currNumber = "0"
				continue
			}

			if !symbolAdjasent {
				symbolAdjasent = symbolAdjacent(i, j, lines)
			}
			currNumber += fmt.Sprintf("%v", digit)
			if j == len(line)-1 {
				if !symbolAdjasent {
					currNumber = "0"
					continue
				}
				val, _ := strconv.Atoi(currNumber)
				sum += val
				symbolAdjasent = false
				currNumber = "0"
			}
		}
	}
	fmt.Println(sum)
}

func symbolAdjacent(i, j int, lines []string) bool {
	if i > 0 && j > 0 {
		if !regexp.MustCompile(`\d`).MatchString(string(lines[i-1][j-1])) && string(lines[i-1][j-1]) != "." {
			return true
		}
	}
	if i > 0 {
		if !regexp.MustCompile(`\d`).MatchString(string(lines[i-1][j])) && string(lines[i-1][j]) != "." {
			return true
		}
	}

	if i > 0 && j < len(lines[i])-1 {
		if !regexp.MustCompile(`\d`).MatchString(string(lines[i-1][j+1])) && string(lines[i-1][j+1]) != "." {
			return true
		}
	}

	if j > 0 {
		if !regexp.MustCompile(`\d`).MatchString(string(lines[i][j-1])) && string(lines[i][j-1]) != "." {
			return true
		}
	}
	if j < len(lines[i])-1 {
		if !regexp.MustCompile(`\d`).MatchString(string(lines[i][j+1])) && string(lines[i][j+1]) != "." {
			return true
		}
	}

	if i < len(lines)-1 && j > 0 {
		if !regexp.MustCompile(`\d`).MatchString(string(lines[i+1][j-1])) && string(lines[i+1][j-1]) != "." {
			return true
		}
	}
	if i < len(lines)-1 {
		if !regexp.MustCompile(`\d`).MatchString(string(lines[i+1][j])) && string(lines[i+1][j]) != "." {
			return true
		}
	}
	if i < len(lines)-1 && j < len(lines[i])-1 {
		if !regexp.MustCompile(`\d`).MatchString(string(lines[i+1][j+1])) && string(lines[i+1][j+1]) != "." {
			return true
		}
	}
	return false
}
