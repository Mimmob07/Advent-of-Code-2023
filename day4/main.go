package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var sum int = 0
var scratchCards map[int]int = map[int]int{} // Key: Card number Value: Amount of cards

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
		cardNum, err := strconv.Atoi(strings.TrimSpace(strings.Split(strings.Split(lineCur, ": ")[0], "Card")[1]))
		if err != nil {
			log.Panic(err)
		}
		scratchCards[cardNum] += 1
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
					score++
				}
			}
		}
		for i := 0; i < score; i++ {
			for j := 0; j < scratchCards[cardNum]; j++ {
				scratchCards[cardNum+i+1] += 1
			}
		}
	}
	for _, value := range scratchCards {
		sum += value
	}
	fmt.Println(scratchCards)
	fmt.Println(sum)
}
