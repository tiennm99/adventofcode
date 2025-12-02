package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInvalid(num int) bool {
	s := strconv.Itoa(num)
	n := len(s)
	if n%2 != 0 || n < 2 {
		return false
	}
	half := n / 2
	return s[:half] == s[half:]
}

func part1() {
	data, _ := os.ReadFile("input.txt")
	content := strings.TrimSpace(string(data))
	ranges := strings.Split(content, ",")
	sum := 0
	for _, v := range ranges {
		arr := strings.Split(v, "-")
		start, _ := strconv.Atoi(arr[0])
		end, _ := strconv.Atoi(arr[1])
		for i := start; i <= end; i++ {
			if isInvalid(i) {
				// fmt.Printf("%d\n", i)
				sum += i
			}
		}
	}
	fmt.Printf("%d\n", sum)
}

func isRepeating(num int) bool {
	s := strconv.Itoa(num)
	n := len(s)
	if n <= 1 {
		return false
	}

	for k := 1; k <= n/2; k++ {
		if n%k != 0 {
			continue
		}

		isRepeat := true
		for i := k; i < n; i++ {
			if s[i] != s[i%k] {
				isRepeat = false
				break
			}
		}
		if isRepeat {
			return true
		}
	}
	return false
}

func part2() {
	data, _ := os.ReadFile("input.txt")
	content := strings.TrimSpace(string(data))
	ranges := strings.Split(content, ",")
	sum := 0
	for _, v := range ranges {
		arr := strings.Split(v, "-")
		start, _ := strconv.Atoi(arr[0])
		end, _ := strconv.Atoi(arr[1])
		for i := start; i <= end; i++ {
			if isRepeating(i) {
				// fmt.Printf("%d\n", i)
				sum += i
			}
		}
	}
	fmt.Printf("%d\n", sum)
}

func main() {
	part1()
	part2()
}
