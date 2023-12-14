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
	maxValues["red"] = 12
	maxValues["green"] = 13
	maxValues["blue"] = 14
	total := 0
	for scanner.Scan() {
		games := strings.Split(scanner.Text(), ";")
		gameValue, _ := strconv.Atoi(strings.Split(strings.Split(games[0], ":")[0], " ")[1])
		games[0] = strings.Split(games[0], ":")[1]
		impossible := false
		for _, val := range games {
			values := strings.Split(val, ",")
			for _, dices := range values {
				key := strings.Split(dices, " ")[2]
				number, _ := strconv.Atoi(strings.Split(dices, " ")[1])
				if maxValues[key] < number {
					impossible = true
					break
				}
			}

		}
		if !impossible {
			total += gameValue
		}
	}
	fmt.Println(total)
}
