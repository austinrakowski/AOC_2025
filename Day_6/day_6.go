package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day_6.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) int {
	var columns [][]string

	for _, line := range lines {
		fields := strings.Fields(line)

		for len(columns) < len(fields) {
			columns = append(columns, []string{})
		}

		for i, field := range fields {
			columns[i] = append(columns[i], field)
		}
	}

	grandTotal := 0
	for _, col := range columns {
		if len(col) == 0 {
			continue
		}

		operator := col[len(col)-1]

		var numbers []int
		for j := 0; j < len(col)-1; j++ {
			num, err := strconv.Atoi(col[j])
			if err != nil {
				continue
			}
			numbers = append(numbers, num)
		}

		if len(numbers) == 0 {
			continue
		}

		result := numbers[0]
		for j := 1; j < len(numbers); j++ {
			if operator == "+" {
				result += numbers[j]
			} else if operator == "*" {
				result *= numbers[j]
			}
		}

		grandTotal += result
	}

	return grandTotal
}

func part2(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	grandTotal := 0
	col := maxLen - 1

	for col >= 0 {
		
		isEmpty := true
		for row := 0; row < len(lines); row++ {
			if col < len(lines[row]) && lines[row][col] != ' ' {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			col--
			continue
		}

		startCol := col
		for startCol > 0 {
			hasData := false
			for row := 0; row < len(lines); row++ {
				if startCol-1 < len(lines[row]) && lines[row][startCol-1] != ' ' {
					hasData = true
					break
				}
			}
			if hasData {
				startCol--
			} else {
				break
			}
		}

		var numbers []int
		var operator string

		for c := col; c >= startCol; c-- {
			var digitStr string
			var opStr string

			for row := 0; row < len(lines); row++ {
				if c < len(lines[row]) && lines[row][c] != ' ' {
					char := lines[row][c]
					
					if char == '+' || char == '*' {
						opStr = string(char)
					} else {
						digitStr += string(char)
					}
				}
			}

			
			if opStr != "" {
				operator = opStr
			}

			if digitStr != "" {
				num, err := strconv.Atoi(digitStr)
				if err == nil {
					numbers = append(numbers, num)
				}
			}
		}

		if len(numbers) > 0 && operator != "" {
			result := numbers[0]
			for j := 1; j < len(numbers); j++ {
				if operator == "+" {
					result += numbers[j]
				} else if operator == "*" {
					result *= numbers[j]
				}
			}
			grandTotal += result
		}

		col = startCol - 1
	}

	return grandTotal
}
