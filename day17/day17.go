package main

import (
	"container/heap"
	"fmt"
	"os"

	"github.com/iSkytran/2023adventofcode/utilities"
)

// List of directions that can be added to coordinates.
var directions = []utilities.Coordinates{
	{Row: -1, Col: 0},
	{Row: 1, Col: 0},
	{Row: 0, Col: -1},
	{Row: 0, Col: 1},
}

// A state while exploring the grid taking into account location, direction, and
// number of straight steps.
type pathState struct {
	loc       utilities.Coordinates
	direction utilities.Coordinates
	steps     int
}

// Convert a grid of runes to a grid of integers.
func intGrid(runeGrid *utilities.Grid[rune]) *utilities.Grid[int] {
	grid := utilities.NewGrid[int]()
	for _, row := range runeGrid.Data {
		newRow := make([]int, 0)
		for _, col := range row {
			value := int(col - '0')
			newRow = append(newRow, value)
		}
		grid.AppendRow(newRow)
	}
	return grid
}

// Compute the shortest path with constraints on how far in a straight line one can travel in the grid.
func dijkstra(grid *utilities.Grid[int], start utilities.Coordinates, end utilities.Coordinates, minLine int, maxLine int) int {
	distance := make(map[pathState]int)
	previous := make(map[pathState]pathState)
	pq := new(utilities.MinPriorityQueue[pathState])
	heap.Init(pq)

	// Add start to queue.
	startState := pathState{loc: start, direction: utilities.Coordinates{Row: 0, Col: 0}, steps: 0}
	startElement := utilities.PriorityElement[pathState]{Value: startState, Priority: 0}
	heap.Push(pq, startElement)
	distance[startState] = 0

	for pq.Len() != 0 {
		// Visit next closest coordinate.
		nextElement := heap.Pop(pq).(utilities.PriorityElement[pathState])
		u := nextElement.Value

		if u.steps+1 >= minLine && u.loc == end {
			return nextElement.Priority
		}

		for _, direction := range directions {
			// Disallow backtracking.
			zeroCoord := utilities.Coordinates{Row: 0, Col: 0}
			if direction.Add(u.direction) == zeroCoord {
				continue
			}

			vOrigin := u.loc.Add(direction)
			v := pathState{loc: vOrigin, direction: direction, steps: u.steps + 1}

			if direction == u.direction {
				// Cannot move more than maxLine times in the same direction.
				if v.steps >= maxLine {
					continue
				}
			} else {
				if v.steps < minLine && u.direction != zeroCoord {
					// Disallow turning before minLine. If not starting.
					continue
				}

				// Made a turn. Reset step counter.
				v.steps = 0
			}

			// Check v is in the grid.
			if !grid.CoordInGrid(vOrigin) {
				continue
			}

			// Total distance if v visited.
			edgeWeight, _ := grid.GetByCoord(vOrigin)
			altDist := distance[u] + edgeWeight

			if currDist, found := distance[v]; !found || altDist < currDist {
				// Found shorter path.
				distance[v] = altDist
				previous[v] = u

				// Add v to queue.
				vElement := utilities.PriorityElement[pathState]{Value: v, Priority: altDist}
				heap.Push(pq, vElement)
			}
		}
	}

	// Couldn't find a path.
	return -1
}

func part1(path string) {
	grid := intGrid(utilities.GridFromFile(path))

	start := utilities.Coordinates{Row: 0, Col: 0}
	end := utilities.Coordinates{Row: grid.RowSize() - 1, Col: grid.ColSize() - 1}
	distance := dijkstra(grid, start, end, 0, 3)

	fmt.Printf("Heat Loss: %d\n", distance)
}

func part2(path string) {
	grid := intGrid(utilities.GridFromFile(path))

	start := utilities.Coordinates{Row: 0, Col: 0}
	end := utilities.Coordinates{Row: grid.RowSize() - 1, Col: grid.ColSize() - 1}
	distance := dijkstra(grid, start, end, 4, 10)

	fmt.Printf("Heat Loss: %d\n", distance)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
