package main

import (
	"fmt"
	"os"
	"strings"
)

func runeToInt(r rune) int {
	return int(r - '0')
}

func getLargestJoltage(raw string) int {
	runes := []rune(raw)
	i1 := 0
	n1 := runes[i1]
	for i := i1; i < len(runes)-1; i++ {
		if runes[i] > n1 {
			i1 = i
			n1 = runes[i1]
			// fmt.Printf("i: %v, n1: %v\n", i, n1)
		}
	}
	i2 := i1 + 1
	n2 := runes[i2]
	for j := i2 + 1; j < len(runes); j++ {
		if runes[j] > n2 {
			i2 = j
			n2 = runes[i2]
			// fmt.Printf("j: %v, n2: %v\n", j, n2)
		}
	}
	return runeToInt(n1)*10 + runeToInt(n2)
}

func part1() {
	data, _ := os.ReadFile("input.txt")
	content := strings.TrimSpace(string(data))
	lines := strings.Split(content, "\n")
	sum := 0
	for _, line := range lines {
		maxJoltage := getLargestJoltage(line)
		// fmt.Println(maxJoltage)
		sum += maxJoltage
	}
	fmt.Printf("%d\n", sum)
}

func getLargestJoltage2(raw string) int {
	runes := []rune(raw)
	num := 12
	r := [12]rune{}
	i := -1
	for idx := range num {
		i++
		r[idx] = runes[i]
		for j := i + 1; j < len(runes)-(num-idx-1); j++ {
			if runes[j] > r[idx] {
				i = j
				r[idx] = runes[i]
			}
		}
	}
	result := 0
	for i := range num {
		result *= 10
		result += runeToInt(r[i])
	}
	return result
}

func part2() {
	data, _ := os.ReadFile("input.txt")
	content := strings.TrimSpace(string(data))
	lines := strings.Split(content, "\n")
	sum := 0
	for _, line := range lines {
		maxJoltage := getLargestJoltage2(line)
		// fmt.Println(maxJoltage)
		sum += maxJoltage
	}
	fmt.Printf("%d\n", sum)
}

func main() {
	part1()
	part2()
}
