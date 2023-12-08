package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func openFile(path string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(path)
	check(err)
	return bufio.NewScanner(file), file
}

func stringsToInts(s []string) []int {
	// Convert list of strings to a list of integers.
	converted := make([]int, 0)
	for _, str := range s {
		val, _ := strconv.Atoi(str)
		converted = append(converted, val)
	}
	return converted
}

type boatRace struct {
	time           int
	recordDistance int
}

func makeBoatRaces(path string) []*boatRace {
	scanner, file := openFile(path)
	defer file.Close()

	// Get race times.
	scanner.Scan()
	timeLine := scanner.Text()[9:]

	// Get distance records.
	scanner.Scan()
	distanceLine := scanner.Text()[9:]

	// Parse numbers.
	times := stringsToInts(strings.Fields(timeLine))
	distances := stringsToInts(strings.Fields(distanceLine))

	var boatRaces []*boatRace
	for i := 0; i < len(times); i++ {
		race := new(boatRace)
		race.time = times[i]
		race.recordDistance = distances[i]
		boatRaces = append(boatRaces, race)
	}

	return boatRaces
}

func makeBoatRace(path string) *boatRace {
	scanner, file := openFile(path)
	defer file.Close()

	// Get race times.
	scanner.Scan()
	timeLine := scanner.Text()[9:]

	// Get distance records.
	scanner.Scan()
	distanceLine := scanner.Text()[9:]

	// Parse numbers.
	times, _ := strconv.Atoi(strings.ReplaceAll(timeLine, " ", ""))
	distances, _ := strconv.Atoi(strings.ReplaceAll(distanceLine, " ", ""))

	race := new(boatRace)
	race.time = times
	race.recordDistance = distances

	return race
}

func (race *boatRace) waysToWin() int {
	// Compute all the ways to win the race.
	wins := 0
	for velocity := 1; velocity < race.time; velocity++ {
		timeLeft := race.time - velocity
		totalDistance := timeLeft * velocity
		if totalDistance > race.recordDistance {
			wins++
		}
	}
	return wins
}

func part1(path string) {
	boatRaces := makeBoatRaces(path)
	product := 1
	for _, race := range boatRaces {
		product *= race.waysToWin()
	}

	fmt.Printf("Product: %d\n", product)
}

func part2(path string) {
	boatRace := makeBoatRace(path)
	ways := boatRace.waysToWin()

	fmt.Printf("Ways to Win: %d\n", ways)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
