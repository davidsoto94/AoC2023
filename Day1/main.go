package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type numberStruct struct {
	value int
	index int
}

func main() {
	file, error := os.Open("test.txt")

	numberMapping := make(map[string]int)
	numberMapping["1"] = -1
	numberMapping["2"] = -1
	numberMapping["3"] = -1
	numberMapping["4"] = -1
	numberMapping["5"] = -1
	numberMapping["6"] = -1
	numberMapping["7"] = -1
	numberMapping["8"] = -1
	numberMapping["9"] = -1

	if error != nil {
		fmt.Println(error)
	}
	total := int64(0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		first := ""
		last := ""
		numberMapping["1"] = strings.Index(line, "one")
		numberMapping["2"] = strings.Index(line, "two")
		numberMapping["3"] = strings.Index(line, "three")
		numberMapping["4"] = strings.Index(line, "four")
		numberMapping["5"] = strings.Index(line, "five")
		numberMapping["6"] = strings.Index(line, "six")
		numberMapping["7"] = strings.Index(line, "seven")
		numberMapping["8"] = strings.Index(line, "eight")
		numberMapping["9"] = strings.Index(line, "nine")
		firstIndex := -1
		for key, value := range numberMapping {
			if value != -1 {
				if first == "" || value < firstIndex {
					first = key
					last = key
					firstIndex = value
				}
			}
		}

		numberMapping["1"] = strings.Index(line, "1")
		numberMapping["2"] = strings.Index(line, "2")
		numberMapping["3"] = strings.Index(line, "3")
		numberMapping["4"] = strings.Index(line, "4")
		numberMapping["5"] = strings.Index(line, "5")
		numberMapping["6"] = strings.Index(line, "6")
		numberMapping["7"] = strings.Index(line, "7")
		numberMapping["8"] = strings.Index(line, "8")
		numberMapping["9"] = strings.Index(line, "9")
		for key, value := range numberMapping {
			if value != -1 {
				if first == "" || value < firstIndex {
					first = key
					last = key
					firstIndex = value
				}
			}
		}

		numberMapping["1"] = strings.LastIndex(line, "one")
		numberMapping["2"] = strings.LastIndex(line, "two")
		numberMapping["3"] = strings.LastIndex(line, "three")
		numberMapping["4"] = strings.LastIndex(line, "four")
		numberMapping["5"] = strings.LastIndex(line, "five")
		numberMapping["6"] = strings.LastIndex(line, "six")
		numberMapping["7"] = strings.LastIndex(line, "seven")
		numberMapping["8"] = strings.LastIndex(line, "eight")
		numberMapping["9"] = strings.LastIndex(line, "nine")
		lastIndex := -1
		for key, value := range numberMapping {
			if value != -1 {
				if last == "" || value > lastIndex {
					last = key
					lastIndex = value
				}
			}
		}
		numberMapping["1"] = strings.LastIndex(line, "1")
		numberMapping["2"] = strings.LastIndex(line, "2")
		numberMapping["3"] = strings.LastIndex(line, "3")
		numberMapping["4"] = strings.LastIndex(line, "4")
		numberMapping["5"] = strings.LastIndex(line, "5")
		numberMapping["6"] = strings.LastIndex(line, "6")
		numberMapping["7"] = strings.LastIndex(line, "7")
		numberMapping["8"] = strings.LastIndex(line, "8")
		numberMapping["9"] = strings.LastIndex(line, "9")
		for key, value := range numberMapping {
			if value != -1 {
				if last == "" || value > lastIndex {
					last = key
					lastIndex = value
				}
			}
		}

		finalString := "0"
		if first == "" {
			continue
		}
		if last == "" {
			finalString = first + first
		} else {
			finalString = first + last
		}
		fmt.Println(finalString)
		lineNumber, _ := strconv.ParseInt(finalString, 10, 32)
		total += lineNumber

	}
	fmt.Println(total)
}
