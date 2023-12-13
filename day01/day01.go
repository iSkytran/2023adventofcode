package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/iSkytran/2023adventofcode/utilities"
)

var lookupTable = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func reverse(str string) string {
	slice := []rune(str)
	reversed := make([]rune, len(slice))
	for idx, r := range slice {
		reversed[len(slice)-idx-1] = r
	}
	return string(reversed)
}

func part1(path string) {
	scanner, file := utilities.OpenFile(path)
	defer file.Close()

	// Compile regex.
	regex := regexp.MustCompile("[0-9]")

	// Iterate per line.
	sum := 0
	for scanner.Scan() {
		value := scanner.Text()

		// Grab first number.
		first := regex.FindString(value)
		firstInt := lookupTable[first]

		// Grab last number
		last := reverse(regex.FindString(reverse(value)))
		lastInt := lookupTable[last]

		// Add to sum.
		sum += 10*firstInt + lastInt
	}

	fmt.Printf("Total: %d\n", sum)
}

func part2(path string) {
	scanner, file := utilities.OpenFile(path)
	defer file.Close()

	// Compile regex.
	forwardRegex := regexp.MustCompile("zero|one|two|three|four|five|six|seven|eight|nine|[0-9]")
	backwardRegex := regexp.MustCompile("enin|thgie|neves|xis|evif|ruof|eerht|owt|eno|orez|[0-9]")

	// Iterate per line.
	sum := 0
	for scanner.Scan() {
		value := scanner.Text()

		// Grab first number.
		first := forwardRegex.FindString(value)
		firstInt := lookupTable[first]

		// Grab last number
		last := reverse(backwardRegex.FindString(reverse(value)))
		lastInt := lookupTable[last]

		// Add to sum.
		sum += 10*firstInt + lastInt
	}

	fmt.Printf("Total: %d\n", sum)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
