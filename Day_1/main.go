package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day_1.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	position := 50
	zero_count := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("cannot parse distance: %v\n", err)
			continue
		}

		if direction == 'L' {
			position = (position - distance) % 100
			if position < 0 {
				position += 100
			}
		} else if direction == 'R' {
			position = (position + distance) % 100
		}

		if position == 0 {
			zero_count++
		}
	}

	fmt.Printf("password: %d\n", zero_count)
}
