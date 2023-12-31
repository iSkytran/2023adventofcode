package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/iSkytran/2023adventofcode/utilities"
)

// Directions that can be added to coordinates.
var (
	up    = utilities.Coordinates{Row: -1, Col: 0}
	down  = utilities.Coordinates{Row: 1, Col: 0}
	left  = utilities.Coordinates{Row: 0, Col: -1}
	right = utilities.Coordinates{Row: 0, Col: 1}
)

// Regex to parse input.
var regex = regexp.MustCompile(`([UDLR]) ([0-9]+) \(#(.....)(.)\)`)

// A dig instruction.
type digInstruction struct {
	direction utilities.Coordinates
	steps     int
}

// Parse file of dig instructions.
func parseInput(path string) []*digInstruction {
	scanner, file := utilities.OpenFile(path)
	defer file.Close()

	instructions := make([]*digInstruction, 0)
	for scanner.Scan() {
		line := scanner.Text()
		fields := regex.FindAllStringSubmatch(line, -1)
		instruction := new(digInstruction)

		direction := fields[0][1]
		switch direction {
		case "U":
			instruction.direction = up
		case "D":
			instruction.direction = down
		case "L":
			instruction.direction = left
		case "R":
			instruction.direction = right
		}

		instruction.steps, _ = strconv.Atoi(fields[0][2])
		instructions = append(instructions, instruction)
	}
	return instructions
}

// Alternative parsing file of dig instructions using the color encoding.
func parseInputHex(path string) []*digInstruction {
	scanner, file := utilities.OpenFile(path)
	defer file.Close()

	instructions := make([]*digInstruction, 0)
	for scanner.Scan() {
		line := scanner.Text()
		fields := regex.FindAllStringSubmatch(line, -1)
		instruction := new(digInstruction)

		hexDistance := fields[0][3]
		distance, _ := strconv.ParseInt(hexDistance, 16, 0)
		instruction.steps = int(distance)

		direction := fields[0][4]
		switch direction {
		case "3":
			instruction.direction = up
		case "1":
			instruction.direction = down
		case "2":
			instruction.direction = left
		case "0":
			instruction.direction = right
		}

		instructions = append(instructions, instruction)
	}
	return instructions
}

// Compute the determinant. Input must be a 2 by 2 slice.
func det(arr [][]int) int {
	return arr[0][0]*arr[1][1] - arr[0][1]*arr[1][0]
}

// Get the area of an enclosed polygon.
func shoelace(instructions []*digInstruction) int {
	coords := []utilities.Coordinates{}

	current := utilities.Coordinates{}
	coords = append(coords, current)

	// Determine vertices.
	edgeLength := 0
	for _, instruction := range instructions {
		edgeLength += instruction.steps
		vector := instruction.direction.Scale(instruction.steps)
		current = current.Add(vector)
		coords = append(coords, current)
	}

	// Build determinant matrix.
	mat := [][]int{}
	for _, coord := range coords {
		row := []int{coord.Row, coord.Col}
		mat = append(mat, row)
	}

	// Compute determinant of the whole matrix.
	area := 0
	for i := 0; i < len(mat)-1; i++ {
		area += det(mat[i : i+2])
	}
	area = int(math.Abs(float64(area) / 2))

	// Pick's theorem accounting for corners.
	return area + edgeLength/2 + 1
}

func part1(path string) {
	instructions := parseInput(path)
	count := shoelace(instructions)
	fmt.Printf("Area: %d\n", count)
}

func part2(path string) {
	instructions := parseInputHex(path)
	count := shoelace(instructions)
	fmt.Printf("Area: %d\n", count)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
