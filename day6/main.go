package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func cleanLine(line []string) []string {
	for i := 0; i < len(line); {
		if line[i] == "" {
			line = append(line[:i], line[i+1:]...)
		} else {
			i++
		}
	}
	return line
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./main <filename>")
		os.Exit(1)
	}

	fileB, err := os.ReadFile(os.Args[1])
	file := string(fileB)
	check(err)

	lines := strings.Split(file, "\n")
	times := cleanLine(strings.Split(lines[0], " ")[1:])
	distances := cleanLine(strings.Split(lines[1], " ")[1:])
	timeS := ""
	distanceS := ""
	for _, nums := range times {
		timeS += nums
	}
	for _, dist := range distances {
		distanceS += dist
	}
	time, err := strconv.Atoi(timeS)
	check(err)
	distance, err := strconv.Atoi(distanceS)
	check(err)

	wins := 0
	for v := 0; v < time; v++ {
		d := v * (time - v)
		if d > distance {
			wins++
		}
	}
	fmt.Println(wins)
}
