package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var regex = regexp.MustCompile("[0-9]+")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func openFile(path string) *bufio.Scanner {
	file, err := os.Open(path)
	check(err)
	return bufio.NewScanner(file)
}

type schematic struct {
	gridNumbers []*gridNumber
	grid        [][]rune
	gears       map[coordinate]*gear
}

type gridNumber struct {
	value    int
	row      int
	colStart int
	colEnd   int
}

type gridBlock struct {
	origin *coordinate
	block  [][]rune
}

type coordinate struct {
	col, row int
}

type gear struct {
	parts []*gridNumber
}

func newSchematic() *schematic {
	s := new(schematic)
	s.gears = make(map[coordinate]*gear)
	return s
}

func newGridBlock() *gridBlock {
	b := new(gridBlock)
	b.origin = new(coordinate)
	return b
}

func (s *schematic) addLine(line string, lineNum int) {
	// Parse numbers and add to list.
	locs := regex.FindAllStringIndex(line, -1)
	for _, loc := range locs {
		gridNum := new(gridNumber)
		gridNum.row = lineNum
		gridNum.value, _ = strconv.Atoi(line[loc[0]:loc[1]])
		gridNum.colStart = loc[0]
		gridNum.colEnd = loc[1]
		s.gridNumbers = append(s.gridNumbers, gridNum)
	}

	// Add line to grid.
	gridLine := []rune(line)
	s.grid = append(s.grid, gridLine)
}

func (s *schematic) partLookup() []*gridNumber {
	// Find all parts in schematic.
	parts := make([]*gridNumber, 0)
	for _, gridNum := range s.gridNumbers {
		if s.isPart(gridNum) {
			parts = append(parts, gridNum)
		}
	}
	return parts
}

func (s *schematic) isPart(gridNum *gridNumber) bool {
	// Check each element in block.
	block := s.getBlock(gridNum)
	foundSymbol := false
	for rowIdx, row := range block.block {
		for colIdx, val := range row {
			// Check for symbol.
			switch val {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
				// Ignore these.
			case '*':
				// Found a gear, which is a symbol.
				coord := coordinate{row: block.origin.row + rowIdx, col: block.origin.col + colIdx}
				if g, ok := s.gears[coord]; ok {
					g.parts = append(g.parts, gridNum)
				} else {
					g = new(gear)
					g.parts = append(g.parts, gridNum)
					s.gears[coord] = g
				}
				return true
			default:
				foundSymbol = true
			}
		}
	}
	return foundSymbol
}

func (s *schematic) getBlock(gridNum *gridNumber) *gridBlock {
	height, width := len(s.grid), len(s.grid[0])
	startRow, endRow := gridNum.row-1, gridNum.row+2
	startCol, endCol := gridNum.colStart-1, gridNum.colEnd+1

	// Clamp to grid bounds.
	startRow, endRow = max(0, startRow), min(height, endRow)
	startCol, endCol = max(0, startCol), min(width, endCol)

	// Slice grid.
	block := newGridBlock()
	block.origin.row, block.origin.col = startRow, startCol
	for rowIdx := startRow; rowIdx < endRow; rowIdx++ {
		row := s.grid[rowIdx][startCol:endCol]
		block.block = append(block.block, row)
	}

	return block
}

func (s *schematic) gearRatio(g *gear) int {
	if len(g.parts) == 2 {
		// Can only compute gear ratio if there are exactly two parts.
		return g.parts[0].value * g.parts[1].value
	}

	return 0
}

func generateSchematic(path string) *schematic {
	scanner := openFile(path)

	// Iterate per line.
	diagram := newSchematic()
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Parse line.
		diagram.addLine(line, lineNum)
		lineNum++
	}

	return diagram
}

func part1(path string) {
	diagram := generateSchematic(path)

	sum := 0
	parts := diagram.partLookup()
	for _, part := range parts {
		sum += part.value
	}

	fmt.Printf("Total: %d\n", sum)
}

func part2(path string) {
	diagram := generateSchematic(path)

	sum := 0
	diagram.partLookup()
	for _, g := range diagram.gears {
		sum += diagram.gearRatio(g)
	}

	fmt.Printf("Total: %d\n", sum)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
