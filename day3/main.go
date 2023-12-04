package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type number struct {
	num  string
	x    []int
	line int
}

func main() {
	var (
		line int = 0
		sum  int = 0
	)

	if len(os.Args) != 2 {
		fmt.Println("Usage: ./main <filename>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numbers := []number{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineCur := scanner.Text()
		for i := 0; i < len(lineCur); i++ {
			if _, err := strconv.Atoi(string(lineCur[i])); err == nil {
				numCur := number{line: line}
				j := 0
				for {
					if i+j == len(lineCur) {
						numbers = append(numbers, numCur)
						break
					}
					if _, err := strconv.Atoi(string(lineCur[i+j])); err != nil {
						numbers = append(numbers, numCur)
						break
					}
					numCur.num += string(lineCur[i+j])
					numCur.x = append(numCur.x, i+j)
					j++
				}
				i += j
			}
		}
		line++
	}

	file.Seek(0, 0)

	scanner2 := bufio.NewScanner(file)
	line = 0
	for scanner2.Scan() {
		lineCur := scanner2.Text()
		// Loop over line
		for i := 0; i < len(lineCur); i++ {
			// Check for symbol
			if lineCur[i] == '*' {
				adjacent := []number{}
				for j := 0; j < len(numbers); j++ {
					if numbers[j].line >= line-1 && numbers[j].line <= line+1 {
						if (numbers[j].x[0] >= i-1 && numbers[j].x[0] <= i+1) || (numbers[j].x[len(numbers[j].x)-1] >= i-1 && numbers[j].x[len(numbers[j].x)-1] <= i+1) {
							adjacent = append(adjacent, numbers[j])
						}
					}
				}
				if len(adjacent) == 2 {
					num1, err := strconv.Atoi(adjacent[0].num)
					if err != nil {
						log.Panic(err)
					}
					num2, err := strconv.Atoi(adjacent[1].num)
					if err != nil {
						log.Panic(err)
					}
					sum += num1 * num2
				}
			}
		}

		line++
	}

	// Part 1 solution below. Part 2 solution above.

	// for scanner2.Scan() {
	// 	lineCur := scanner2.Text()
	// 	// Loop over line
	// 	for i := 0; i < len(lineCur); i++ {
	// 		// Check for symbol
	// 		if _, err := strconv.Atoi(string(lineCur[i])); lineCur[i] != '.' && err != nil {
	// 			for j := 0; j < len(numbers); j++ {
	// 				if numbers[j].line >= line-1 && numbers[j].line <= line+1 {
	// 					if (numbers[j].x[0] >= i-1 && numbers[j].x[0] <= i+1) || (numbers[j].x[len(numbers[j].x)-1] >= i-1 && numbers[j].x[len(numbers[j].x)-1] <= i+1) {
	// 						fmt.Printf("Found match! Number: %v\n", numbers[j].num)
	// 						num, err := strconv.Atoi(numbers[j].num)
	// 						if err != nil {
	// 							log.Panic(err)
	// 						}
	// 						sum += num
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// 	line++
	// }

	fmt.Println(sum)
}
