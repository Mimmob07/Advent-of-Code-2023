package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var digitsMap map[int]string = map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine"}

func main() {
	// Check arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./main <filename>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	var sum int = 0
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for scanner.Scan() {
		var (
			lineCur   string = scanner.Text()
			firstNumB bool   = false
			firstNum  string
			lastNum   string
			finalNum  int
		)
		for i := 0; i < len(lineCur); i++ {
			tmp := match(lineCur[i:])
			if _, err := strconv.Atoi(string(lineCur[i])); err == nil || tmp != "" {
				if !firstNumB {
					if tmp != "" {
						firstNum = tmp
						lastNum = tmp
						firstNumB = true
					} else {
						firstNum = string(lineCur[i])
						lastNum = string(lineCur[i])
						firstNumB = true
					}
				} else {
					if tmp != "" {
						lastNum = tmp
					} else {
						lastNum = string(lineCur[i])
					}
				}
			}
		}
		finalNum, err := strconv.Atoi(firstNum + lastNum)
		if err != nil {
			log.Fatal(err)
		}
		sum += finalNum
	}

	fmt.Println(sum)
}

func match(line string) string {
	for i := 0; i < 9; i++ {
		if len(line) < len(digitsMap[i+1]) {
			continue
		}
		for j := 0; j < len(digitsMap[i+1]); j++ {
			if !(digitsMap[i+1][j] == line[j]) {
				break
			}
			if j+1 == len(digitsMap[i+1]) {
				return fmt.Sprint(i + 1)
			}
		}
	}
	return ""
}
