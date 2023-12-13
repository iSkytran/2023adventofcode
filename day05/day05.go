package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/iSkytran/2023adventofcode/utilities"
)

type almanac struct {
	seedRanges            []*seedRange
	seedToSoil            rangeMaps
	soilToFertilizer      rangeMaps
	fertilizerToWater     rangeMaps
	waterToLight          rangeMaps
	lightToTemperature    rangeMaps
	temperatureToHumidity rangeMaps
	humidityToLocation    rangeMaps
}

type seedRange struct {
	start int
	end   int
}

type rangeMaps []*rangeMap

type rangeMap struct {
	sourceStart       int
	sourceEnd         int
	destinationOffset int
}

func parseSeeds(line string) []*seedRange {
	// Parse strings to a list.
	strSlice := strings.Fields(line[7:])
	values := utilities.StringsToInts(strSlice)

	var seeds []*seedRange
	for _, val := range values {
		// Create seed range with end right after start.
		s := new(seedRange)
		s.start = val
		s.end = val + 1
		seeds = append(seeds, s)
	}

	return seeds
}

func parseRangeOfSeeds(line string) []*seedRange {
	// Parse strings to a list.
	s := strings.Fields(line[7:])
	values := utilities.StringsToInts(s)

	// Go through the whole list.
	var seeds []*seedRange
	for i := 0; i < len(values); i += 2 {
		first := values[i]
		last := first + values[i+1]

		// Create seed range.
		s := new(seedRange)
		s.start = first
		s.end = last
		seeds = append(seeds, s)
	}
	return seeds
}

func generateAlmanac(path string, seedParseFunc func(string) []*seedRange) *almanac {
	scanner, file := utilities.OpenFile(path)
	defer file.Close()

	// Parse each line.
	var curLookupTbl *rangeMaps
	a := new(almanac)
	for scanner.Scan() {
		line := scanner.Text()

		// Determine where in file we are.
		switch {
		case line == "":
			// Empty line.
			continue
		case strings.Contains(line, "seeds"):
			// Parse seed values.
			a.seedRanges = seedParseFunc(line)
		case strings.Contains(line, "seed-to-soil"):
			curLookupTbl = &a.seedToSoil
		case strings.Contains(line, "soil-to-fertilizer"):
			curLookupTbl = &a.soilToFertilizer
		case strings.Contains(line, "fertilizer-to-water"):
			curLookupTbl = &a.fertilizerToWater
		case strings.Contains(line, "water-to-light"):
			curLookupTbl = &a.waterToLight
		case strings.Contains(line, "light-to-temperature"):
			curLookupTbl = &a.lightToTemperature
		case strings.Contains(line, "temperature-to-humidity"):
			curLookupTbl = &a.temperatureToHumidity
		case strings.Contains(line, "humidity-to-location"):
			curLookupTbl = &a.humidityToLocation
		default:
			// Add range to lookup table.
			values := strings.Fields(line)
			ints := utilities.StringsToInts(values)
			newMap := newRangeMap(ints)
			*curLookupTbl = append(*curLookupTbl, newMap)
		}
	}

	return a
}

func (a almanac) minLocation() int {
	// Lookup every single seed in the almanac to find the seed with minimum location.
	min := math.MaxInt64
	for _, sRange := range a.seedRanges {
		for i := sRange.start; i < sRange.end; i++ {
			loc := a.lookup(i)
			if loc < min {
				min = loc
			}
		}
	}
	return min
}

func (a almanac) lookup(num int) int {
	// Lookup location based on seed.
	num = a.seedToSoil.lookup(num)
	num = a.soilToFertilizer.lookup(num)
	num = a.fertilizerToWater.lookup(num)
	num = a.waterToLight.lookup(num)
	num = a.lightToTemperature.lookup(num)
	num = a.temperatureToHumidity.lookup(num)
	num = a.humidityToLocation.lookup(num)
	return num
}

func newRangeMap(input []int) *rangeMap {
	// Parse input of [startOfDestination, startOfSource, rangeLength].
	r := new(rangeMap)
	r.sourceStart = input[1]
	r.sourceEnd = input[1] + input[2] - 1     // Source end is one less than range from start.
	r.destinationOffset = input[0] - input[1] // Offset is destination minus source.
	return r
}

func (ms rangeMaps) lookup(num int) int {
	// Check all ranges.
	for _, m := range ms {
		val := m.lookup(num)
		if val != -1 {
			return val
		}
	}
	// Map to itself if not in a range.
	return num
}

func (m *rangeMap) lookup(num int) int {
	// Check in range.
	if num >= m.sourceStart && num <= m.sourceEnd {
		// In range, map using offset.
		return num + m.destinationOffset
	}
	// Not in range.
	return -1
}

func part1(path string) {
	a := generateAlmanac(path, parseSeeds)
	minLoc := a.minLocation()
	fmt.Printf("Minimum Location: %d\n", minLoc)
}

func part2(path string) {
	a := generateAlmanac(path, parseRangeOfSeeds)
	minLoc := a.minLocation()
	fmt.Printf("Minimum Location: %d\n", minLoc)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
