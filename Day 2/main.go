package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, error := os.Open("test.txt")

	if error != nil {
		fmt.Println(error)
	}
	scanner := bufio.NewScanner(file)
	maxValues := make(map[string]int)
	total := 0
	for scanner.Scan() {
		maxValues["red"] = 1
		maxValues["green"] = 1
		maxValues["blue"] = 1
		sets := strings.Split(scanner.Text(), ";")
		sets[0] = strings.Split(sets[0], ":")[1]
		for _, val := range sets {
			colors := strings.Split(val, ",")
			for _, dices := range colors {
				key := strings.Split(dices, " ")[2]
				number, _ := strconv.Atoi(strings.Split(dices, " ")[1])
				if maxValues[key] < number {
					maxValues[key] = number
				}
			}
		}
		gameValue := maxValues["blue"] * maxValues["green"] * maxValues["red"]
		total += gameValue
	}
	fmt.Println(total)
}
