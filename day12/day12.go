package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/iSkytran/2023adventofcode/utilities"
)

// Used to parse lines.
var regex = regexp.MustCompile(`([\?\.#]*) ([0-9,]*)`)

// Used for memoization.
var cache = make(map[string]int)

type conditionRecord struct {
	conditions string
	numDamaged []int
}

func newConditionRecord() *conditionRecord {
	record := new(conditionRecord)
	record.numDamaged = make([]int, 0)
	return record
}

func parseLine(line string, multiplier int) *conditionRecord {
	record := newConditionRecord()
	matches := regex.FindAllStringSubmatch(line, -1)
	values := utilities.StringsToInts(strings.Split(matches[0][2], ","))

	// Apply multiplier.
	var buffer bytes.Buffer
	for i := 0; i < multiplier; i++ {
		buffer.WriteString(matches[0][1])
		record.numDamaged = append(record.numDamaged, values...)

		// Add in ? delimiter if not the end.
		if i != multiplier-1 {
			buffer.WriteString("?")
		}
	}

	// Add working spring delimiter to the end.
	buffer.WriteString(".")

	record.conditions = buffer.String()
	return record
}

func numArrangements(conditions string, numDamaged []int) int {
	// Check cache first.
	lookup := conditions + strings.Join(utilities.IntsToStrings(numDamaged), ",")
	_, found := cache[lookup]
	if found {
		return cache[lookup]
	}

	arrangements := 0

	// Base cases.
	if len(numDamaged) == 0 {
		// No more groups.
		if !strings.Contains(conditions, "#") {
			// No more damaged springs. A valid arrangement.
			arrangements = 1
		} else {
			// Damaged springs found. An invalid arrangement.
			return 0
		}
	} else if len(conditions) == 0 {
		// No more records to process.
		arrangements = 0
	} else {
		// Check next rune.
		current := rune(conditions[0])

		switch current {
		case '.':
			// Working spring, ignore.
			arrangements = numArrangements(conditions[1:], numDamaged)
		case '#':
			// Damaged spring, need to lookahead.
			groupSize := numDamaged[0]

			if groupSize > len(conditions) {
				// Group does not fit.
				return 0
			}

			group := conditions[:groupSize]
			group = strings.ReplaceAll(group, "?", "#")

			// Check if there are the correct number of damaged.
			for i := 0; i < numDamaged[0]; i++ {
				if group[i] != '#' {
					// Group does not fit all damaged springs.
					return 0
				}
			}

			if conditions[groupSize] == '?' || conditions[groupSize] == '.' {
				// Following separator. Skip it and search for next group.
				reducedDamaged := make([]int, len(numDamaged)-1)
				copy(reducedDamaged, numDamaged[1:])
				arrangements = numArrangements(conditions[groupSize+1:], reducedDamaged)
			}
		case '?':
			// Could be damaged or not.
			conditions := conditions[1:]

			numDamagedCopy := make([]int, len(numDamaged))
			copy(numDamagedCopy, numDamaged)
			arrangements += numArrangements("."+conditions, numDamagedCopy)

			numDamagedCopy = make([]int, len(numDamaged))
			copy(numDamagedCopy, numDamaged)
			arrangements += numArrangements("#"+conditions, numDamagedCopy)
		}
	}

	// Store arrangements in cache.
	cache[lookup] = arrangements
	return arrangements
}

func part1(path string) {
	scanner, file := utilities.OpenFile(path)
	defer file.Close()

	records := make([]*conditionRecord, 0)
	for scanner.Scan() {
		line := scanner.Text()
		records = append(records, parseLine(line, 1))
	}

	total := 0
	for _, record := range records {
		total += numArrangements(record.conditions, record.numDamaged)
	}

	fmt.Printf("Total: %d\n", total)
}

func part2(path string) {
	scanner, file := utilities.OpenFile(path)
	defer file.Close()

	records := make([]*conditionRecord, 0)
	for scanner.Scan() {
		line := scanner.Text()
		records = append(records, parseLine(line, 5))
	}

	total := 0
	for _, record := range records {
		total += numArrangements(record.conditions, record.numDamaged)
	}

	fmt.Printf("Total: %d\n", total)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
