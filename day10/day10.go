package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/iSkytran/2023adventofcode/utilities"
)

const (
	north = iota
	south
	west
	east
)

type coordinates struct {
	row int
	col int
}

type pipeMaze struct {
	startCoord *coordinates
	diagram    [][]rune
	loop       *utilities.Set[coordinates]
}

func newMaze() *pipeMaze {
	maze := new(pipeMaze)
	maze.diagram = make([][]rune, 0)
	return maze
}

func (maze *pipeMaze) adjacentToStart() ([]coordinates, *utilities.Set[int]) {
	// Find direction to proceed in from start.
	start := maze.startCoord
	valid := map[int]coordinates{
		north: {start.row - 1, start.col},
		south: {start.row + 1, start.col},
		west:  {start.row, start.col - 1},
		east:  {start.row, start.col + 1},
	}

	if start.row == 0 {
		// Cannot proceed North.
		delete(valid, north)
	}

	if start.row == len(maze.diagram)-1 {
		// Cannot proceed South.
		delete(valid, south)
	}

	if start.col == 0 {
		// Cannot proceed West.
		delete(valid, west)
	}

	if start.col == len(maze.diagram[0])-1 {
		// Cannot proceed East.
		delete(valid, west)
	}

	adjacentCoords := make([]coordinates, 0)
	adjacentDirections := utilities.NewSet[int]()
	for direction, coords := range valid {
		// Check if next pipe is connected.
		pipe := string(maze.diagram[coords.row][coords.col])
		switch {
		case direction == north && strings.Contains("|7F", pipe):
			adjacentCoords = append(adjacentCoords, coords)
			adjacentDirections.Add(direction)
		case direction == south && strings.Contains("|JL", pipe):
			adjacentCoords = append(adjacentCoords, coords)
			adjacentDirections.Add(direction)
		case direction == west && strings.Contains("-FL", pipe):
			adjacentCoords = append(adjacentCoords, coords)
			adjacentDirections.Add(direction)
		case direction == east && strings.Contains("-7J", pipe):
			adjacentCoords = append(adjacentCoords, coords)
			adjacentDirections.Add(direction)
		}
	}

	// No valid pipe found.
	return adjacentCoords, adjacentDirections
}

func (maze *pipeMaze) computeLoop() {
	// Keep track of loop coordinates.
	maze.loop = utilities.NewSet[coordinates]()

	// Find the coordinates of the direction to go in.
	adjacentCoords, _ := maze.adjacentToStart()
	current := adjacentCoords[0]

	leavingStart := true
	for {
		// Determine next coordinate.
		currentPipe := maze.diagram[current.row][current.col]
		valid := make(map[int]coordinates)
		valid[north] = coordinates{current.row - 1, current.col}
		valid[south] = coordinates{current.row + 1, current.col}
		valid[west] = coordinates{current.row, current.col - 1}
		valid[east] = coordinates{current.row, current.col + 1}

		// Add current coordinates to visited.
		maze.loop.Add(current)

		// Remove invalid directions.
		switch currentPipe {
		case '|':
			delete(valid, west)
			delete(valid, east)
		case '-':
			delete(valid, north)
			delete(valid, south)
		case 'L':
			delete(valid, west)
			delete(valid, south)
		case 'J':
			delete(valid, east)
			delete(valid, south)
		case '7':
			delete(valid, north)
			delete(valid, east)
		case 'F':
			delete(valid, north)
			delete(valid, west)
		case 'S':
			// Completed the loop.
			return
		}

		// Remove locations visited.
		for _, coord := range valid {
			// If leaving start, don't visit start.
			if leavingStart && coord == *maze.startCoord {
				leavingStart = false
				continue
			}

			if !maze.loop.Contains(coord) {
				// Location not visited yet.
				current = coord
				break
			}
		}
	}
}

func (maze *pipeMaze) computeEnclosed() [][]rune {
	// Make a copy of the diagram with just the loop.
	diagram := make([][]rune, 0)
	for rowNum, row := range maze.diagram {
		for colNum := range row {
			current := coordinates{rowNum, colNum}
			if !maze.loop.Contains(current) {
				row[colNum] = '.'
			}
		}
		diagram = append(diagram, row)
	}

	// Replace start with pipe.
	_, adjacentDirections := maze.adjacentToStart()
	switch {
	case adjacentDirections.Contains(north) && adjacentDirections.Contains(south):
		diagram[maze.startCoord.row][maze.startCoord.col] = '|'
	case adjacentDirections.Contains(west) && adjacentDirections.Contains(east):
		diagram[maze.startCoord.row][maze.startCoord.col] = '-'
	case adjacentDirections.Contains(north) && adjacentDirections.Contains(west):
		diagram[maze.startCoord.row][maze.startCoord.col] = 'J'
	case adjacentDirections.Contains(north) && adjacentDirections.Contains(east):
		diagram[maze.startCoord.row][maze.startCoord.col] = 'L'
	case adjacentDirections.Contains(west) && adjacentDirections.Contains(south):
		diagram[maze.startCoord.row][maze.startCoord.col] = '7'
	case adjacentDirections.Contains(east) && adjacentDirections.Contains(south):
		diagram[maze.startCoord.row][maze.startCoord.col] = 'F'
	}

	for rowNum := range diagram {
		// Find number of times loop crossed.
		// Even counts are outside the loop, odd counts are in the loop.
		inLoop := false
		for colNum := range diagram[0] {
			currentPipe := diagram[rowNum][colNum]

			// Simple vertical case.
			if currentPipe == '|' {
				inLoop = !inLoop
				continue
			}

			// Lookahead vertical cases.
			if currentPipe == 'F' {
				for i := colNum; i < len(diagram[0]); i++ {
					futurePipe := diagram[rowNum][i]
					if futurePipe == 'J' {
						// Zig-zag detected.
						inLoop = !inLoop
						break
					}
					if futurePipe == '7' {
						// No zig-zag.
						break
					}
				}
			}

			if currentPipe == 'L' {
				for i := colNum; i < len(diagram[0]); i++ {
					futurePipe := diagram[rowNum][i]
					if futurePipe == '7' {
						// Zig-zag detected.
						inLoop = !inLoop
						break
					}
					if futurePipe == 'J' {
						// No zig-zag.
						break
					}
				}
			}

			// Only modify if not a pipe.
			if currentPipe == '.' {
				if inLoop {
					diagram[rowNum][colNum] = 'I'
				} else {
					diagram[rowNum][colNum] = 'O'
				}
			}
		}
	}

	return diagram
}

func parseMaze(path string) *pipeMaze {
	scanner, file := utilities.OpenFile(path)
	defer file.Close()

	maze := newMaze()
	rowNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		for colNum, pipeChar := range line {
			if pipeChar == 'S' {
				// Found start coordinates.
				maze.startCoord = &coordinates{rowNum, colNum}
			}
		}

		maze.diagram = append(maze.diagram, []rune(line))
		rowNum++
	}

	// Figure out loop coordinates.
	maze.computeLoop()
	return maze
}

func (maze *pipeMaze) stepsToEnd() int {
	return maze.loop.Size() / 2
}

func part1(path string) {
	maze := parseMaze(path)
	steps := maze.stepsToEnd()
	fmt.Printf("Steps to Furthest: %d\n", steps)
}

func part2(path string) {
	maze := parseMaze(path)
	diagram := maze.computeEnclosed()
	count := 0
	for _, row := range diagram {
		for _, col := range row {
			if col == 'I' {
				count++
			}
		}
	}
	fmt.Printf("Number of Inner Tiles: %d\n", count)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
