package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "input.txt")

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lines := strings.Fields(string(file))

	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

// Number of times dial ends at 0 after each rotation.
func part1(lines []string) int {
	position := 50
	timesPassedZero := 0

	for _, line := range lines {
		direction := line[:1]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if direction == "R" {
			position = (position + distance) % 100
		} else if direction == "L" {
			position = (position - distance + 100) % 100
		}

		if position == 0 {
			timesPassedZero++
		}
	}

	return timesPassedZero
}

// Number of times the dial passes 0 during any rotation.
func part2(lines []string) int {
	position := 50
	timesPassedZero := 0

	for _, line := range lines {
		direction := line[:1]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		for i := 0; i < distance; i++ {
			if direction == "R" {
				position = (position + 1) % 100
			} else if direction == "L" {
				position = (position - 1 + 100) % 100
			}

			if position == 0 {
				timesPassedZero++
			}
		}
	}

	return timesPassedZero
}