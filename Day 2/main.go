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
	for scanner.Scan() {
		games := strings.Split(scanner.Text(), ";")
		gameValue, _ := strconv.ParseUint(strings.Split(strings.Split(games[0], ":")[0], " ")[1], 10, 0)
		fmt.Println(gameValue)
		games[0] = strings.Split(games[0], ":")[1]
		for _, val := range games {
			number := strings.Split(val, " ")[0]
		}
	}
}
