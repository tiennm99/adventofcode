package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start, End int
}

func readInput() ([]Range, []int) {
	data, _ := os.ReadFile("input.txt")
	content := strings.TrimSpace(string(data))
	lines := strings.Split(content, "\n")
	ranges := []Range{}
	for len(lines) > 0 {
		line := lines[0]
		lines = lines[1:]
		parts := strings.Split(line, "-")
		if len(parts) < 2 {
			break
		}

		start, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		end, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		ranges = append(ranges, Range{start, end})
	}

	candidates := []int{}
	for len(lines) > 0 {
		line := lines[0]
		lines = lines[1:]
		candidate, _ := strconv.Atoi(strings.TrimSpace(line))
		candidates = append(candidates, candidate)

	}

	return ranges, candidates
}

func part1() {
	ranges, candidates := readInput()
	count := 0
	for _, candidate := range candidates {
		for _, r := range ranges {
			if candidate >= r.Start && candidate <= r.End {
				count++
				break
			}
		}
	}

	fmt.Printf("%d\n", count)
}

func part2() {
	ranges, _ := readInput()
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	merged := []Range{ranges[0]}
	for _, cur := range ranges[1:] {
		last := &merged[len(merged)-1]

		if cur.Start <= last.End {
			if cur.End > last.End {
				last.End = cur.End
			}
		} else {
			merged = append(merged, cur)
		}
	}

	count := 0
	for _, m := range merged {
		count += (m.End - m.Start + 1)
	}
	fmt.Printf("%d\n", count)
}

func main() {
	part1()
	part2()
}
