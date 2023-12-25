package main

import (
	"fmt"
	"math"
	"os"

	"github.com/iSkytran/2023adventofcode/utilities"
)

var (
	up    = utilities.Coordinates{Row: -1, Col: 0}
	down  = utilities.Coordinates{Row: 1, Col: 0}
	left  = utilities.Coordinates{Row: 0, Col: -1}
	right = utilities.Coordinates{Row: 0, Col: 1}
)

func optimalCoverage(grid *utilities.Grid[rune]) int {
	max := 0.0

	// Traverse left to right.
	for i := 0; i < grid.ColSize(); i++ {
		// Check top.
		start := utilities.Vector{Origin: utilities.Coordinates{Row: 0, Col: i}, Direction: down}
		energized := float64(coverage(grid, start))
		max = math.Max(max, energized)

		// Check bottom.
		start = utilities.Vector{Origin: utilities.Coordinates{Row: grid.RowSize() - 1, Col: i}, Direction: up}
		energized = float64(coverage(grid, start))
		max = math.Max(max, energized)
	}

	// Traverse top to bottom.
	for i := 0; i < grid.RowSize(); i++ {
		// Check left
		start := utilities.Vector{Origin: utilities.Coordinates{Row: i, Col: 0}, Direction: right}
		energized := float64(coverage(grid, start))
		max = math.Max(max, energized)

		// Check right.
		start = utilities.Vector{Origin: utilities.Coordinates{Row: i, Col: grid.ColSize() - 1}, Direction: left}
		energized = float64(coverage(grid, start))
		max = math.Max(max, energized)
	}

	return int(max)
}

func coverage(grid *utilities.Grid[rune], start utilities.Vector) int {
	visited := utilities.NewSet[utilities.Vector]()
	visitCoordinate(grid, visited, start)

	// Extract locations from vectors.
	energized := utilities.NewSet[utilities.Coordinates]()
	for _, vector := range visited.ToSlice() {
		energized.Add(vector.Origin)
	}

	return energized.Size()
}

func visitCoordinate(grid *utilities.Grid[rune], set *utilities.Set[utilities.Vector], current utilities.Vector) {
	// Check coordinates are in bounds.
	if !grid.CoordInGrid(current.Origin) {
		return
	}

	// Visit current coordinate.
	if set.Contains(current) {
		// Beam exists, ignore.
		return
	}
	set.Add(current)

	// Compute next coordinate(s).
	currentRune, _ := grid.GetByCoord(current.Origin)
	switch currentRune {
	case '.':
		// Continue in current direction.
		current.Origin = current.Origin.Add(current.Direction)
		visitCoordinate(grid, set, current)
	case '/':
		switch current.Direction {
		case up:
			// Bounce right.
			current.Origin = current.Origin.Add(right)
			current.Direction = right
			visitCoordinate(grid, set, current)
		case down:
			// Bounce left.
			current.Origin = current.Origin.Add(left)
			current.Direction = left
			visitCoordinate(grid, set, current)
		case left:
			// Bounce down.
			current.Origin = current.Origin.Add(down)
			current.Direction = down
			visitCoordinate(grid, set, current)
		case right:
			// Bounce up.
			current.Origin = current.Origin.Add(up)
			current.Direction = up
			visitCoordinate(grid, set, current)
		}
	case '\\':
		switch current.Direction {
		case up:
			// Bounce left.
			current.Origin = current.Origin.Add(left)
			current.Direction = left
			visitCoordinate(grid, set, current)
		case down:
			// Bounce right.
			current.Origin = current.Origin.Add(right)
			current.Direction = right
			visitCoordinate(grid, set, current)
		case left:
			// Bounce up.
			current.Origin = current.Origin.Add(up)
			current.Direction = up
			visitCoordinate(grid, set, current)
		case right:
			// Bounce down.
			current.Origin = current.Origin.Add(down)
			current.Direction = down
			visitCoordinate(grid, set, current)
		}
	case '|':
		switch current.Direction {
		case up, down:
			// Continue in current direction.
			current.Origin = current.Origin.Add(current.Direction)
			visitCoordinate(grid, set, current)
		case left, right:
			// Split into up and down beams.
			next := utilities.Vector{Origin: current.Origin.Add(up), Direction: up}
			visitCoordinate(grid, set, next)
			next = utilities.Vector{Origin: current.Origin.Add(down), Direction: down}
			visitCoordinate(grid, set, next)
		}
	case '-':
		switch current.Direction {
		case up, down:
			// Split into left and right beams.
			next := utilities.Vector{Origin: current.Origin.Add(left), Direction: left}
			visitCoordinate(grid, set, next)
			next = utilities.Vector{Origin: current.Origin.Add(right), Direction: right}
			visitCoordinate(grid, set, next)
		case left, right:
			// Continue in current direction.
			current.Origin = current.Origin.Add(current.Direction)
			visitCoordinate(grid, set, current)
		}
	}
}

func part1(path string) {
	grid := utilities.GridFromFile(path)
	start := utilities.Vector{Origin: utilities.Coordinates{Row: 0, Col: 0}, Direction: right}
	energized := coverage(grid, start)
	fmt.Printf("Energized: %d\n", energized)
}

func part2(path string) {
	grid := utilities.GridFromFile(path)
	energized := optimalCoverage(grid)
	fmt.Printf("Energized: %d\n", energized)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
