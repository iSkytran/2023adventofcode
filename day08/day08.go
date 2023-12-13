package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/iSkytran/2023adventofcode/utilities"
)

func lcm(numSlice []int) int {
	// Least common multiple of a list of numbers.
	current := 1
	for _, num := range numSlice {
		current = num * current / gcd(num, current)
	}
	return current
}

func gcd(a, b int) int {
	// Euclidean algorithm.
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

type navigationMap struct {
	instructions []rune
	network      map[string]*networkNode
}

type networkNode struct {
	left  string
	right string
}

func parseMap(path string) *navigationMap {
	scanner, file := utilities.OpenFile(path)
	defer file.Close()

	regex := regexp.MustCompile(`([A-Z0-9]+) = \(([A-Z0-9]+), ([A-Z0-9]+)\)`)
	navMap := new(navigationMap)

	scanner.Scan()
	navMap.instructions = []rune(scanner.Text())
	navMap.network = make(map[string]*networkNode, 0)

	scanner.Scan() // Ignore blank line.
	for scanner.Scan() {
		// Parse line.
		line := scanner.Text()
		tokens := regex.FindStringSubmatch(line)

		// Create map node.
		node := new(networkNode)
		node.left = tokens[2]
		node.right = tokens[3]
		navMap.network[tokens[1]] = node
	}

	return navMap
}

func (navMap *navigationMap) stepsToExit(start string, end string) int {
	// Find starts.
	starts := make([]string, 0)
	for key := range navMap.network {
		if strings.Contains(key, start) {
			starts = append(starts, key)
		}
	}

	// Invalid start string.
	if len(starts) == 0 {
		return 0
	}

	allSteps := make([]int, 0)
	for _, current := range starts {
		steps := 0
		for {
			instruction := navMap.instructions[0]
			if instruction == 'L' {
				// Move left.
				current = navMap.network[current].left
			} else {
				// Move right.
				current = navMap.network[current].right
			}

			// Move instruction to the end.
			navMap.instructions = navMap.instructions[1:]
			navMap.instructions = append(navMap.instructions, instruction)
			steps++

			// Break if end reached.
			if strings.Contains(current, end) {
				allSteps = append(allSteps, steps)
				break
			}
		}
	}

	return lcm(allSteps)
}

func part1(path string) {
	navMap := parseMap(path)
	steps := navMap.stepsToExit("AAA", "ZZZ")
	fmt.Printf("Steps: %d\n", steps)
}

func part2(path string) {
	navMap := parseMap(path)
	steps := navMap.stepsToExit("A", "Z")
	fmt.Printf("Steps: %d\n", steps)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
