package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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
}

func newMaze() *pipeMaze {
	maze := new(pipeMaze)
	maze.diagram = make([][]rune, 0)
	return maze
}

func (maze *pipeMaze) directionFromStart() (*coordinates, error) {
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

	for direction, coords := range valid {
		// Check if next pipe is connected.
		pipe := string(maze.diagram[coords.row][coords.col])
		switch {
		case direction == north && strings.Contains("|7F", pipe):
			return &coords, nil
		case direction == south && strings.Contains("|JL", pipe):
			return &coords, nil
		case direction == west && strings.Contains("-FL", pipe):
			return &coords, nil
		case direction == east && strings.Contains("-7J", pipe):
			return &coords, nil
		}
	}

	// No valid pipe found.
	return nil, errors.New("no valid next step from start")
}

func (maze *pipeMaze) computeLoopSteps() int {
	// Keep track of visited coordinates.
	visited := make(map[coordinates]struct{})

	// Find the coordinates of the direction to go in.
	current, ok := maze.directionFromStart()
	check(ok)

	steps := 0
	leavingStart := true
	for {
		steps++

		// Determine next coordinate.
		currentPipe := maze.diagram[current.row][current.col]
		valid := make(map[int]*coordinates)
		valid[north] = &coordinates{current.row - 1, current.col}
		valid[south] = &coordinates{current.row + 1, current.col}
		valid[west] = &coordinates{current.row, current.col - 1}
		valid[east] = &coordinates{current.row, current.col + 1}

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
			return steps
		}

		// Add current coordinates to visited.
		visited[*current] = struct{}{}

		// Remove locations visited.
		for _, coord := range valid {
			// If leaving start, don't visit start.
			if leavingStart && *coord == *maze.startCoord {
				leavingStart = false
				continue
			}

			_, ok := visited[*coord]
			if !ok {
				// Location not visited yet.
				current = coord
				break
			}
		}
	}
}

func parseMaze(path string) *pipeMaze {
	scanner, file := openFile(path)
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

	return maze
}

func part1(path string) {
	maze := parseMaze(path)
	steps := maze.computeLoopSteps()
	fmt.Printf("Steps to Furthest: %d\n", steps / 2)
}

func part2(path string) {

}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
