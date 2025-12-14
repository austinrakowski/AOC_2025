package main

import (
	"bufio"
	"fmt"
	"os"
)

func readGrid(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func countNeighbors(grid [][]rune, r, c int) int {
	rows := len(grid)
	cols := len(grid[0])
	neighborCount := 0

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		newR := r + dir[0]
		newC := c + dir[1]

		if newR >= 0 && newR < rows && newC >= 0 && newC < cols {
			if grid[newR][newC] == '@' {
				neighborCount++
			}
		}
	}

	return neighborCount
}

func countAccessibleRolls(grid [][]rune) int {
	accessible := 0
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '@' && countNeighbors(grid, r, c) < 4 {
				accessible++
			}
		}
	}

	return accessible
}

func part1(grid [][]rune) int {
	return countAccessibleRolls(grid)
}

func part2(grid [][]rune) int {
	// Create a copy of the grid to work with
	gridCopy := make([][]rune, len(grid))
	for i := range grid {
		gridCopy[i] = make([]rune, len(grid[i]))
		copy(gridCopy[i], grid[i])
	}

	rows := len(gridCopy)
	cols := len(gridCopy[0])
	totalRemoved := 0

	// Track which positions need to be checked
	toCheck := make(map[[2]int]bool)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if gridCopy[r][c] == '@' {
				toCheck[[2]int{r, c}] = true
			}
		}
	}

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for len(toCheck) > 0 {
		toRemove := make([][2]int, 0)

		for pos := range toCheck {
			r, c := pos[0], pos[1]
			if gridCopy[r][c] == '@' && countNeighbors(gridCopy, r, c) < 4 {
				toRemove = append(toRemove, pos)
			}
		}

		if len(toRemove) == 0 {
			break
		}

		// Clear toCheck for the next iteration
		toCheck = make(map[[2]int]bool)

		for _, pos := range toRemove {
			r, c := pos[0], pos[1]
			gridCopy[r][c] = '.'
			totalRemoved++
k
			for _, dir := range directions {
				newR := r + dir[0]
				newC := c + dir[1]

				if newR >= 0 && newR < rows && newC >= 0 && newC < cols {
					if gridCopy[newR][newC] == '@' {
						toCheck[[2]int{newR, newC}] = true
					}
				}
			}
		}
	}

	return totalRemoved
}

func main() {
	grid, err := readGrid("day_4.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(grid) == 0 {
		fmt.Println("Empty grid")
		return
	}

	part1Result := part1(grid)
	fmt.Printf("Part 1: %d\n", part1Result)

	part2Result := part2(grid)
	fmt.Printf("Part 2: %d\n", part2Result)
}
