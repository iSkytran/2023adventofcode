package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/iSkytran/2023adventofcode/utilities"
)

func stringsToInts(s []string) []int {
	// Convert list of strings to a list of integers.
	converted := make([]int, 0)
	for _, str := range s {
		val, _ := strconv.Atoi(str)
		converted = append(converted, val)
	}
	return converted
}

func parseHistory(line string) []int {
	tokens := strings.Fields(line)
	return stringsToInts(tokens)
}

func interpolateNext(history []int) int {
	// Find value of the next number in the history.
	// If all zeros in list, next number is zero.
	if allZeroes(history) {
		return 0
	}

	// Must interpolate deltas first.
	deltas := derivative(history)
	nextDelta := interpolateNext(deltas)

	// Next number is the last number plus the next delta.
	return history[len(history)-1] + nextDelta
}

func interpolatePrev(history []int) int {
	// Find value of the next number in the history.
	// If all zeros in list, next number is zero.
	if allZeroes(history) {
		return 0
	}

	// Must interpolate deltas first.
	deltas := derivative(history)
	prevDelta := interpolatePrev(deltas)

	// Next number is the first number minus the previous delta.
	return history[0] - prevDelta
}

func derivative(history []int) []int {
	// Compute differences between each element.
	deltas := make([]int, 0)
	for i := 1; i < len(history); i++ {
		delta := history[i] - history[i-1]
		deltas = append(deltas, delta)
	}
	return deltas
}

func allZeroes(history []int) bool {
	for _, val := range history {
		if val != 0 {
			return false
		}
	}
	return true
}

func part1(path string) {
	scanner, file := utilities.OpenFile(path)
	defer file.Close()

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		history := parseHistory(line)
		sum += interpolateNext(history)
	}

	fmt.Printf("Total: %d\n", sum)
}

func part2(path string) {
	scanner, file := utilities.OpenFile(path)
	defer file.Close()

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		history := parseHistory(line)
		sum += interpolatePrev(history)
	}

	fmt.Printf("Total: %d\n", sum)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
