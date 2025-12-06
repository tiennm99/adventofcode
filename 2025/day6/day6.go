package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() [][]string {
	data, _ := os.ReadFile("input.txt")
	content := strings.TrimSpace(string(data))
	lines := strings.Split(content, "\n")
	input := [][]string{}
	size := len(strings.Fields(lines[0]))
	for range size {
		input = append(input, []string{})
	}
	for _, line := range lines {
		parts := strings.Fields(line)
		for i, part := range parts {
			input[i] = append(input[i], strings.TrimSpace(part))
		}
	}

	return input
}

func part1() {
	input := readInput()
	result := 0
	for _, operation := range input {
		total, _ := strconv.Atoi(operation[0])
		if operation[len(operation)-1] == "+" {
			for _, value := range operation[1 : len(operation)-1] {
				num, _ := strconv.Atoi(value)
				total += num
			}
		} else {
			for _, value := range operation[1 : len(operation)-1] {
				num, _ := strconv.Atoi(value)
				total *= num
			}
		}
		result += total
	}

	fmt.Printf("%d\n", result)
}

func readInput2() [][]rune {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}
	r, c := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		if len(line) > c {
			c = len(line)
		}
		r++
	}

	// Because operands in + and * in each operation and when calculate sum can be swaped
	// We can do this
	// Flip the matrix to make it easier to read
	r, c = c, r
	input := [][]rune{}
	for i := range r {
		slice := []rune{}
		for j := range c {
			if j < len(lines) && i < len(lines[j]) {
				slice = append(slice, rune(lines[j][i]))
			} else {
				slice = append(slice, ' ')
			}
		}
		input = append(input, slice)
	}

	return input
}

func part2() {
	input := readInput2()
	result := 0
	newOperation := true
	total := 0
	operation := ' '
	for _, line := range input {
		str := strings.TrimSpace(string(line))
		if len(str) == 0 {
			// fmt.Printf("Total: %d\n", total)
			result += total
			total = 0
			newOperation = true
			continue
		}
		if newOperation {
			first := str[:len(str)-1]
			// fmt.Printf("First: %s\n", first)
			number, err := strconv.Atoi(strings.TrimSpace(first))
			if err != nil {
				fmt.Println("Error parsing number:", err)
				return
			}
			total = number
			operation = rune(str[len(str)-1])
			newOperation = false
			// fmt.Printf("First: %d, Operation: %c\n", total, operation)
		} else {
			number, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error parsing number:", err)
				return
			}
			// fmt.Println(number)
			switch operation {
			case '+':
				total += number
			case '*':
				total *= number
			}
		}
	}
	result += total

	fmt.Printf("%d\n", result)
}

func main() {
	part1()
	part2()
}
