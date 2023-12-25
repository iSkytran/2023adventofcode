package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/iSkytran/2023adventofcode/utilities"
)

const (
	north = iota
	west
	south
	east
)

func cycle(grid *utilities.Grid[rune], numCycles int) {
	seen := map[string]int{}
	foundRepeat := false
	for i := 0; i < numCycles; i++ {
		directionalShift(grid, north)
		directionalShift(grid, west)
		directionalShift(grid, south)
		directionalShift(grid, east)

		if !foundRepeat {
			serialization := grid.Serialize()

			if prev, found := seen[serialization]; found {
				cyclesLeft := numCycles - i
				repeatLength := i - prev
				i = numCycles - (cyclesLeft % repeatLength)
				foundRepeat = true
			} else {
				seen[serialization] = i
			}
		}
	}
}

func directionalShift(grid *utilities.Grid[rune], direction int) {
	var size int
	switch direction {
	case north, south:
		size = grid.ColSize()
	case east, west:
		size = grid.RowSize()
	}

	for i := 0; i < size; i++ {
		switch direction {
		case north:
			slice, _ := grid.GetColumn(i)
			slice = shift(slice)
			grid.SetColumn(i, slice)
		case west:
			slice, _ := grid.GetRow(i)
			slice = shift(slice)
			grid.SetRow(i, slice)
		case south:
			slice, _ := grid.GetColumn(i)
			slices.Reverse[[]rune](slice)
			slice = shift(slice)
			slices.Reverse[[]rune](slice)
			grid.SetColumn(i, slice)
		case east:
			slice, _ := grid.GetRow(i)
			slices.Reverse[[]rune](slice)
			slice = shift(slice)
			slices.Reverse[[]rune](slice)
			grid.SetRow(i, slice)
		}
	}
}

func shift(slice []rune) []rune {
	length := len(slice)

	// Modified insertion sort.
	nextEmpty := 0
	for i := 0; i < length; i++ {
		switch slice[i] {
		case '#':
			nextEmpty = i + 1
		case 'O':
			slice[i] = '.'
			slice[nextEmpty] = 'O'
			nextEmpty++
		}
	}

	return slice
}

func calcLoad(grid *utilities.Grid[rune]) int {
	load := 0
	numRows, numCols := grid.Shape()

	for i := 0; i < numRows; i++ {
		rockCount := 0
		for j := 0; j < numCols; j++ {
			if item, _ := grid.Get(i, j); item == 'O' {
				rockCount++
			}
		}

		load += rockCount * (numRows - i)
	}

	return load
}

func part1(path string) {
	grid := utilities.GridFromFile(path)
	directionalShift(grid, north)
	load := calcLoad(grid)
	fmt.Printf("Total Load: %d\n", load)
}

func part2(path string) {
	grid := utilities.GridFromFile(path)
	cycle(grid, 1e9)
	load := calcLoad(grid)
	fmt.Printf("Total Load: %d\n", load)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
