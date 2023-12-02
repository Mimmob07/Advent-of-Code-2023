package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// var rgbMax map[string]int = map[string]int{"red": 12, "green": 13, "blue": 14}
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
			lineCur string         = scanner.Text()
			most    map[string]int = map[string]int{"red": 0, "green": 0, "blue": 0}
			power   int            = 0
			// possible bool   = true
		)

		header := strings.Split(lineCur, ": ")
		// game, err := strconv.Atoi(strings.Split(header[0], " ")[1])
		rounds := strings.Split(header[1], "; ")

		if err != nil {
			log.Panic(err)
		}

		for i := 0; i < len(rounds); i++ {
			colors := strings.Split(rounds[i], ", ")
			// if !possible {
			// 	break
			// }
			for j := 0; j < len(colors); j++ {
				colorAndValue := strings.Split(colors[j], " ")
				val, err := strconv.Atoi(colorAndValue[0])
				if err != nil {
					log.Panic(err)
				}
				if val > most[colorAndValue[1]] {
					most[colorAndValue[1]] = val
				}
				// if val > rgbMax[colorAndValue[1]] {
				// 	possible = false
				// 	break
				// }
			}
		}
		for _, value := range most {
			if value == 0 {
				continue
			}
			if power == 0 {
				power = value
			} else {
				power *= value
			}
		}
		sum += power
		// if possible {
		// 	sum += game
		// }
	}
	fmt.Println(sum)
}
