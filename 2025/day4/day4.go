package main

import (
	"fmt"
	"os"
	"strings"
)

var arounds = []Vec2{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

type Vec2 struct {
	X, Y int
}

func isValid(input []string, pos Vec2) bool {
	if input[pos.X][pos.Y] != '@' {
		return false
	}
	count := 0
	for _, around := range arounds {
		x := pos.X + around.X
		y := pos.Y + around.Y
		if x < 0 || x >= len(input) || y < 0 || y >= len(input[x]) {
			continue
		}
		if input[x][y] == '@' {
			count++
		}
	}
	return count < 4
}

func count(input []string) int {
	count := 0
	for i, l := range input {
		for j := range l {
			if isValid(input, Vec2{i, j}) {
				count++
			}
		}
	}
	return count
}

func part1() {
	data, _ := os.ReadFile("input.txt")
	content := strings.TrimSpace(string(data))
	lines := strings.Split(content, "\n")
	count := count(lines)
	fmt.Printf("%d\n", count)
}

func isValid2(input [][]rune, pos Vec2) bool {
	if input[pos.X][pos.Y] != '@' {
		return false
	}
	count := 0
	for _, around := range arounds {
		x := pos.X + around.X
		y := pos.Y + around.Y
		if x < 0 || x >= len(input) || y < 0 || y >= len(input[x]) {
			continue
		}
		if input[x][y] == '@' {
			count++
		}
	}
	return count < 4
}

func count2(input []string) int {
	runes := make([][]rune, len(input))
	for i, l := range input {
		runes[i] = []rune(l)
	}
	queue := make(map[Vec2]bool)
	for i, l := range runes {
		for j := range l {
			if runes[i][j] == '@' {
				queue[Vec2{i, j}] = true
			}
		}
	}
	accepted := make(map[Vec2]bool)
	for len(queue) > 0 {
		for pos, _ := range queue {
			delete(queue, pos)
			if isValid2(runes, pos) {
				accepted[pos] = true
				runes[pos.X][pos.Y] = '.'
				for _, around := range arounds {
					x := pos.X + around.X
					y := pos.Y + around.Y
					if x < 0 || x >= len(runes) || y < 0 || y >= len(runes[x]) {
						continue
					}
					if accepted[Vec2{x, y}] {
						continue
					}
					if runes[x][y] == '@' {
						queue[Vec2{x, y}] = true
					}
				}
			}
		}
	}
	return len(accepted)
}

func part2() {
	data, _ := os.ReadFile("input.txt")
	content := strings.TrimSpace(string(data))
	lines := strings.Split(content, "\n")
	count := count2(lines)
	fmt.Printf("%d\n", count)
}

func main() {
	part1()
	part2()
}
