package main

import(
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, error := os.Open("test.txt")
	if error != nil {
		fmt.Println(error)
	}
	total := int64(0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		first := ""
		last := ""
		for i := 0; i < len(line); i++{
			ch := string(line[i])
			if s, err := strconv.ParseInt(ch, 10, 32); err == nil {
				if first == ""{
					first = fmt.Sprintf("%v",s)
				} else {
					last = fmt.Sprintf("%v",s)
				}
			}
		}
		finalString := "0"
		if first==""{
			continue
		}
		if last==""{
			finalString = first+first
		} else {
			finalString = first+last
		}
		lineNumber,_ := strconv.ParseInt(finalString, 10, 32)
		total += lineNumber
		
	}
	fmt.Println(total)
}