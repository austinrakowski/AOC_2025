package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func parseInput(filename string) ([]Range, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var ranges []Range
	var ids []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			ranges = append(ranges, Range{Start: start, End: end})
		} else if line != "" {
			id, _ := strconv.Atoi(line)
			ids = append(ids, id)
		}
	}

	return ranges, ids, scanner.Err()
}

func mergeRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return []Range{}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	merged := []Range{ranges[0]}

	for _, r := range ranges[1:] {
		lastIdx := len(merged) - 1
		if r.Start <= merged[lastIdx].End+1 {
			if r.End > merged[lastIdx].End {
				merged[lastIdx].End = r.End
			}
		} else {
			merged = append(merged, r)
		}
	}

	return merged
}

func part1(ranges []Range, ids []int) int {
	count := 0
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r.Start && id <= r.End {
				count++
				break
			}
		}
	}
	return count
}

func part2(ranges []Range) int {
	merged := mergeRanges(ranges)
	count := 0
	for _, r := range merged {
		count += (r.End - r.Start) + 1
	}
	return count
}

func main() {
	ranges, ids, err := parseInput("day_5.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Part 1:", part1(ranges, ids))
	fmt.Println("Part 2:", part2(ranges))
}
