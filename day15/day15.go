package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/iSkytran/2023adventofcode/utilities"
)

const numBoxes = 256

type focalLens struct {
	label string
	power int
}

func parseStringFile(path string) []string {
	scanner, file := utilities.OpenFile(path)
	defer file.Close()

	// Only one line to scan.
	scanner.Scan()
	input := scanner.Text()
	sequence := strings.Split(input, ",")

	return sequence
}

func hashString(input string) int {
	value := 0
	for _, char := range input {
		value += int(char)
		value *= 17
		value %= 256
	}

	return value
}

func generateHashMap(sequence []string) [numBoxes][]focalLens {
	// Initialize structure.
	boxes := [numBoxes][]focalLens{}

	for _, step := range sequence {
		dashIndex := strings.Index(step, "-")
		if dashIndex != -1 {
			// Dash operation.
			label := step[:dashIndex]
			hash := hashString(label)

			for i := 0; i < len(boxes[hash]); i++ {
				if label == boxes[hash][i].label {
					// Remove lens.
					boxes[hash] = append(boxes[hash][:i], boxes[hash][i+1:]...)
					break
				}
			}
		} else {
			// Equal operation.
			var lens focalLens
			equalIndex := strings.Index(step, "=")

			lens.label = step[:equalIndex]
			hash := hashString(lens.label)
			lens.power, _ = strconv.Atoi(step[equalIndex+1:])

			found := false
			for i := 0; i < len(boxes[hash]); i++ {
				if lens.label == boxes[hash][i].label {
					// Update power.
					boxes[hash][i].power = lens.power
					found = true
					break
				}
			}

			if !found {
				boxes[hash] = append(boxes[hash], lens)
			}
		}
	}

	return boxes
}

func part1(path string) {
	sequence := parseStringFile(path)

	total := 0
	for _, step := range sequence {
		total += hashString(step)
	}

	fmt.Printf("Total: %d\n", total)
}

func part2(path string) {
	sequence := parseStringFile(path)

	boxes := generateHashMap(sequence)
	total := 0
	for boxNum, box := range boxes {
		for lensNum, lens := range box {
			total += (boxNum + 1) * (lensNum + 1) * lens.power
		}
	}

	fmt.Printf("Total: %d\n", total)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
