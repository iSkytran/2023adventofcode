package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/iSkytran/2023adventofcode/utilities"
)

const (
	vertical = iota
	horizontal
)

func findReflection(grid *utilities.Grid[rune], direction int) int {
	var size int
	if direction == vertical {
		size = grid.ColSize()
	} else {
		size = grid.RowSize()
	}

	// Check if each line is a reflector line.
	for i := 0; i < size-1; i++ {
		// Check for reflected column.
		beforeIdx, afterIdx := i, i+1
		reflection := true
		for beforeIdx >= 0 && afterIdx < size {
			var before, after []rune
			if direction == vertical {
				before, _ = grid.GetColumn(beforeIdx)
				after, _ = grid.GetColumn(afterIdx)
			} else {
				before, _ = grid.GetRow(beforeIdx)
				after, _ = grid.GetRow(afterIdx)
			}

			if !slices.Equal(before, after) {
				// Not a reflection.
				reflection = false
				break
			}

			beforeIdx--
			afterIdx++
		}

		if reflection {
			return i + 1
		}
	}

	// Reflection not found.
	return 0
}

func findSmudgedReflection(grid *utilities.Grid[rune], direction int) int {
	original := findReflection(grid, direction)

	var size int
	if direction == vertical {
		size = grid.ColSize()
	} else {
		size = grid.RowSize()
	}

	// Check if each line is a reflector line.
	for i := 0; i < size-1; i++ {
		if i == original-1 {
			// If original reflection, ignore.
			continue
		}

		// Check for reflected column.
		beforeIdx, afterIdx := i, i+1
		reflection, fixedSmudge := true, false
		for beforeIdx >= 0 && afterIdx < size {
			var before, after []rune
			if direction == vertical {
				before, _ = grid.GetColumn(beforeIdx)
				after, _ = grid.GetColumn(afterIdx)
			} else {
				before, _ = grid.GetRow(beforeIdx)
				after, _ = grid.GetRow(afterIdx)
			}

			distance := utilities.ElementHammingDistance[rune](before, after)
			if distance != 0 {
				if distance == 1 && !fixedSmudge {
					// Allow one smudge fix.
					fixedSmudge = true
				} else {
					// Not a reflection.
					reflection = false
					break
				}
			}

			beforeIdx--
			afterIdx++
		}

		if reflection {
			return i + 1
		}
	}

	// Reflection not found.
	return 0
}

func parseGrids(path string) []*utilities.Grid[rune] {
	scanner, file := utilities.OpenFile(path)
	defer file.Close()

	grid := utilities.NewGrid[rune]()
	grids := make([]*utilities.Grid[rune], 0)
	grids = append(grids, grid)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// New grid.
			grid = utilities.NewGrid[rune]()
			grids = append(grids, grid)
		} else {
			// Add to current grid.
			grid.AppendRow([]rune(line))
		}
	}

	return grids
}

func part1(path string) {
	grids := parseGrids(path)

	total := 0
	for _, grid := range grids {
		total += findReflection(grid, vertical)
		total += findReflection(grid, horizontal) * 100
	}

	fmt.Printf("Total: %d\n", total)
}

func part2(path string) {
	grids := parseGrids(path)

	total := 0
	for _, grid := range grids {
		total += findSmudgedReflection(grid, vertical)
		total += findSmudgedReflection(grid, horizontal) * 100
	}

	fmt.Printf("Total: %d\n", total)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
