package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var sum int = 0

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./main <filename>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for scanner.Scan() {
		var (
			score          = 0
			lineCur string = scanner.Text()
		)
		winNums := strings.Split(strings.Split(strings.Split(lineCur, ": ")[1], " | ")[0], " ")
		myNums := strings.Split(strings.Split(strings.Split(lineCur, ": ")[1], " | ")[1], " ")
		for i := 0; i < len(winNums); {
			if winNums[i] == "" || winNums[i] == " " {
				winNums = append(winNums[:i], winNums[i+1:]...)
			} else {
				i++
			}
		}
		for i := 0; i < len(myNums); {
			if myNums[i] == "" || myNums[i] == " " {
				myNums = append(myNums[:i], myNums[i+1:]...)
			} else {
				i++
			}
		}
		for i := 0; i < len(winNums); i++ {
			for j := 0; j < len(myNums); j++ {
				if winNums[i] == myNums[j] {
					fmt.Println(winNums[i], myNums[j])
					if score == 0 {
						score = 1
					} else {
						score *= 2
					}
				}
			}
		}
		fmt.Println(score)
		sum += score
	}
	fmt.Println(sum)
}
