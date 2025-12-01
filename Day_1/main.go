package main

import (
	"fmt"
	"strconv"
)

func main() {
	
	rotations := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}

	position := 50
	zero_count := 0 

	for _, line := range rotations {
		
		direction := line[0] // 'L' or 'R'
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("cannot parse distance: %v\n", err)
			continue
		}

		// Apply the rotation
		if direction == 'L' {
			position = (position - distance) % 100
			if position < 0 {
				position += 100
			}
		} else if direction == 'R' {
			position = (position + distance) % 100
		}

		fmt.Printf("After %s: position = %d\n", line, position)

		if position == 0 {
			zero_count++
		}
	}

	fmt.Printf("\nPassword (times dial points at 0): %d\n", zero_count)
}
