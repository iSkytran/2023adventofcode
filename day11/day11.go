package main

import (
	"fmt"
	"math"
	"os"

	"github.com/iSkytran/2023adventofcode/utilities"
)

func cosmicExpansion(grid *utilities.Grid[rune], expansionFactor int) []utilities.Coordinates {
	blankRows := utilities.NewSet[int]()
	blankCols := utilities.NewSet[int]()

	// Find blank rows.
	for i := 0; i < grid.RowSize(); i++ {
		if !grid.RowContains(i, '#') {
			blankRows.Add(i)
		}
	}

	// Find blank columns.
	for i := 0; i < grid.ColSize(); i++ {
		if !grid.ColContains(i, '#') {
			blankCols.Add(i)
		}
	}

	// Find actual coordinates of galaxies.
	galaxies := make([]utilities.Coordinates, 0)
	actualRow := 0
	for row := 0; row < grid.RowSize(); row++ {
		if blankRows.Contains(row) {
			// This row is actually of size expansionFactor.
			actualRow += expansionFactor
			continue
		}

		actualCol := 0
		for col := 0; col < grid.ColSize(); col++ {
			if blankCols.Contains(col) {
				// This row is actually of size expansionFactor.
				actualCol += expansionFactor
				continue
			}

			val, _ := grid.Get(row, col)
			if val == '#' {
				coords := utilities.Coordinates{Row: actualRow, Col: actualCol}
				galaxies = append(galaxies, coords)
			}

			actualCol++
		}

		actualRow++
	}

	return galaxies
}

func manhattanDistance(coords []utilities.Coordinates) int {
	sum := 0
	for i := 0; i < len(coords); i++ {
		iCoords := coords[i]
		for j := i; j < len(coords); j++ {
			jCoords := coords[j]
			sum += int(math.Abs(float64(iCoords.Row - jCoords.Row)))
			sum += int(math.Abs(float64(iCoords.Col - jCoords.Col)))
		}
	}
	return sum
}

func part1(path string) {
	grid := utilities.GridFromFile(path)
	coords := cosmicExpansion(grid, 2)
	sum := manhattanDistance(coords)
	fmt.Printf("Total: %d\n", sum)
}

func part2(path string) {
	grid := utilities.GridFromFile(path)
	coords := cosmicExpansion(grid, 1000000)
	sum := manhattanDistance(coords)
	fmt.Printf("Total: %d\n", sum)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
