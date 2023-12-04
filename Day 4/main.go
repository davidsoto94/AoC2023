package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	file, error := os.Open("test.txt")
	if error != nil {
		fmt.Println(error)
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		res1 := strings.Split(scanner.Text(), ":")
		winners := strings.Split(res1[1], "|")
		winners = strings.Split(winners[0], " ")
		numbers := strings.Split(res1[1], "|")
		numbers = strings.Split(numbers[1], " ")
		lineSum := 0
		for _, number := range numbers {
			if number != "" && slices.Contains(winners, number) {
				if lineSum == 0 {
					lineSum = 1
					continue
				}
				lineSum += lineSum
			}
		}
		sum += lineSum
	}
	fmt.Println(sum)
}
