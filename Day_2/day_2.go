package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseRanges(data string) [][2]int {
	var ranges [][2]int
	rangePairs := strings.Split(strings.TrimSpace(data), ",")

	for _, rangePair := range rangePairs {
		parts := strings.Split(rangePair, "-")
		lower, _ := strconv.Atoi(parts[0])
		upper, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, [2]int{lower, upper})
	}

	return ranges
}

func isRepeatingPattern(s string) bool {
	n := len(s)
	// Try all possible pattern lengths from 1 to half the string's length
	for patternLength := 1; patternLength <= n/2; patternLength++ {
		// Only check if the pattern length divides evenly into the total length
		if n%patternLength == 0 {
			pattern := s[:patternLength]
			// Calculate how many times the pattern should repeat
			repetitions := n / patternLength
			// Check if repeating the pattern recreates the original string
			repeated := strings.Repeat(pattern, repetitions)
			if repeated == s {
				return true
			}
		}
	}
	return false
}

func part1(ranges [][2]int) int {
	invalidKeys := []int{}

	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			strI := strconv.Itoa(i)

			// If it's not even length, skip
			if len(strI)%2 != 0 {
				continue
			}

			mid := len(strI) / 2
			left := strI[:mid]
			right := strI[mid:]

			if left == right {
				invalidKeys = append(invalidKeys, i)
			}
		}
	}

	sum := 0
	for _, val := range invalidKeys {
		sum += val
	}
	return sum
}

func part2(ranges [][2]int) int {
	invalidKeys := []int{}

	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			strI := strconv.Itoa(i)

			if isRepeatingPattern(strI) {
				invalidKeys = append(invalidKeys, i)
			}
		}
	}

	sum := 0
	for _, val := range invalidKeys {
		sum += val
	}
	return sum
}

func main() {
	data, err := os.ReadFile("day_2.txt")
	if err != nil {
		panic(err)
	}

	ranges := parseRanges(string(data))

	fmt.Printf("Part 1: %d\n", part1(ranges))
	fmt.Printf("Part 2: %d\n", part2(ranges))
}
