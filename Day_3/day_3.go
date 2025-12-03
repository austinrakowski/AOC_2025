package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(line string) int {
	n := len(line)
	maxAfter := make([]byte, n)

	// Track the largest digit that comes after each position
	maxAfter[n-1] = line[n-1]
	for i := n - 2; i >= 0; i-- {
		// Store the max of current digit vs. max of everything to its right
		if line[i] > maxAfter[i+1] {
			maxAfter[i] = line[i]
		} else {
			maxAfter[i] = maxAfter[i+1]
		}
	}

	// Find best first digit and pair it with best digit after it
	maxJoltage := 0
	for i := 0; i < n-1; i++ {
		digit1 := int(line[i] - '0')
		// Get the best possible second digit (from positions after i)
		digit2 := int(maxAfter[i+1] - '0')
		joltage := digit1*10 + digit2
		if joltage > maxJoltage {
			maxJoltage = joltage
		}
	}

	return maxJoltage
}

func part2(line string) int {

	x := 12
	n := len(line)

	// for each position in our result,
	// pick the largest digit from the remaining window
	result := ""
	starting_pos := 0

	for i := 0; i < x; i++ {
	
		remaining := x - i - 1
		endPos := n - remaining

		// Find the maximum digit in this window
		max_digit := line[starting_pos]
		max_idx := starting_pos
		for j := starting_pos; j < endPos; j++ {
			if line[j] > max_digit {
				max_digit = line[j]
				max_idx = j
			}
		}
		result += string(max_digit)

		starting_pos = max_idx + 1
	}

	num, _ := strconv.Atoi(result)
	return num
}

func main() {
	file, err := os.Open("day_3.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pt1_joltage := 0
	pt2_joltage := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		pt1_joltage += part1(line)
		pt2_joltage += part2(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part 1 - Total joltage: %d\n", pt1_joltage)
	fmt.Printf("Part 2 - Total joltage: %d\n", pt2_joltage)
}
