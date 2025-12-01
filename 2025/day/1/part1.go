package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	current := 50
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		multiplier := 1
		if direction == 'L' {
			multiplier = -1
		}
		number, _ := strconv.Atoi(line[1:])
		current += number * multiplier
		current = ((current % 100) + 100) % 100
		if current == 0 {
			counter++
		}
		// fmt.Printf("%s %c %d %d %d %d\n", line, direction, number, multiplier, current, counter)
	}
	fmt.Printf("%d\n", counter)
}
