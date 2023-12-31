package utilities

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"math"
)

// ******************************************* //
// Coordinate structure and related functions. //
// ******************************************* //
type Coordinates struct {
	Row int
	Col int
}

func (coord Coordinates) Add(otherCoord Coordinates) Coordinates {
	row := coord.Row + otherCoord.Row
	col := coord.Col + otherCoord.Col
	return Coordinates{Row: row, Col: col}
}

func (coord Coordinates) Subtract(otherCoord Coordinates) Coordinates {
	row := coord.Row - otherCoord.Row
	col := coord.Col - otherCoord.Col
	return Coordinates{Row: row, Col: col}
}

func (coord Coordinates) Scale(scalar int) Coordinates {
	row := scalar * coord.Row
	col := scalar * coord.Col
	return Coordinates{Row: row, Col: col}
}

func (coord Coordinates) Abs() Coordinates {
	coord.Row = int(math.Abs(float64(coord.Row)))
	coord.Col = int(math.Abs(float64(coord.Col)))
	return coord
}

func (g *Grid[T]) GetByCoord(coord Coordinates) (T, error) {
	return g.Get(coord.Row, coord.Col)
}

func (g *Grid[T]) SetByCoord(coord Coordinates, val T) error {
	return g.Set(coord.Row, coord.Col, val)
}

func (g *Grid[T]) CoordInGrid(coord Coordinates) bool {
	_, err := g.GetByCoord(coord)
	return err == nil
}

// **************************************** //
// Vector structure (origin and direction). //
// **************************************** //
type Vector struct {
	Origin    Coordinates
	Direction Coordinates
}

// ************************************* //
// Grid structure and related functions. //
// ************************************* //
type Grid[T comparable] struct {
	Data [][]T
}

func NewGrid[T comparable]() *Grid[T] {
	g := new(Grid[T])
	g.Data = make([][]T, 0)
	return g
}

func GridFromFile(path string) *Grid[rune] {
	scanner, file := OpenFile(path)
	defer file.Close()

	g := NewGrid[rune]()
	for scanner.Scan() {
		line := scanner.Text()
		g.AppendRow([]rune(line))
	}
	return g
}

// Size functions.
func (g *Grid[_]) RowSize() int {
	return len(g.Data)
}

func (g *Grid[_]) ColSize() int {
	if g.RowSize() == 0 {
		return 0
	}
	return len(g.Data[0])
}

func (g *Grid[_]) Shape() (int, int) {
	return g.RowSize(), g.ColSize()
}

// General functions.
func (g *Grid[T]) Search(item T) []Coordinates {
	coords := make([]Coordinates, 0)
	for i := 0; i < g.RowSize(); i++ {
		for j := 0; j < g.ColSize(); j++ {
			if g.Data[i][j] == item {
				coord := Coordinates{i, j}
				coords = append(coords, coord)
			}
		}
	}
	return coords
}

func (g *Grid[T]) Get(rowIndex int, columnIndex int) (T, error) {
	if rowIndex < 0 || rowIndex >= g.RowSize() {
		return *new(T), errors.New("row index out of bounds")
	}

	if columnIndex < 0 || columnIndex >= g.ColSize() {
		return *new(T), errors.New("column index out of bounds")
	}

	return g.Data[rowIndex][columnIndex], nil
}

func (g *Grid[T]) Set(rowIndex int, columnIndex int, item T) error {
	if rowIndex < 0 || rowIndex >= g.RowSize() {
		return errors.New("row index out of bounds")
	}

	if columnIndex < 0 || columnIndex >= g.ColSize() {
		return errors.New("column index out of bounds")
	}

	g.Data[rowIndex][columnIndex] = item
	return nil
}

func (g *Grid[T]) Contains(item T) bool {
	for i := 0; i < g.RowSize(); i++ {
		if g.RowContains(i, item) {
			return true
		}
	}
	return false
}

func (g *Grid[_]) PrintGrid() {
	for i := 0; i < g.RowSize(); i++ {
		if row, ok := any(g.Data[i]).([]rune); ok {
			fmt.Println(string(row))
		} else {
			for j := 0; j < g.ColSize(); j++ {
				fmt.Printf("%v", g.Data[i][j])
			}
			fmt.Printf("\n")
		}
	}
}

func (g *Grid[_]) Serialize() string {
	buffer := bytes.Buffer{}
	gob.NewEncoder(&buffer).Encode(g)
	return buffer.String()
}

// Row specific functions.
func (g *Grid[T]) AddRow(index int, row []T) error {
	if index < 0 || index > g.RowSize() {
		return errors.New("index out of bounds")
	}

	if index == g.RowSize() {
		g.Data = append(g.Data, row)
	} else {
		g.Data = append(g.Data[:index+1], g.Data[index:]...)
		g.Data[index] = row
	}

	return nil
}

func (g *Grid[T]) GetRow(index int) ([]T, error) {
	if index < 0 || index >= g.RowSize() {
		return []T{}, errors.New("row index out of bounds")
	}

	return g.Data[index], nil
}

func (g *Grid[T]) SetRow(index int, row []T) error {
	if index < 0 || index >= g.RowSize() {
		return errors.New("row index out of bounds")
	}

	for i := 0; i < g.ColSize(); i++ {
		g.Set(index, i, row[i])
	}

	return nil
}

func (g *Grid[T]) AppendRow(row []T) {
	g.AddRow(g.RowSize(), row)
}

// Column specific functions.
func (g *Grid[T]) RowContains(index int, item T) bool {
	row, err := g.GetRow(index)
	if err != nil {
		return false
	}
	for _, rowItem := range row {
		if rowItem == item {
			return true
		}
	}
	return false
}

func (g *Grid[T]) AddColumn(index int, column []T) error {
	if index < 0 || index > g.ColSize() {
		return errors.New("index out of bounds")
	}

	for i := 0; i < g.RowSize(); i++ {
		if index == g.ColSize() {
			g.Data[i] = append(g.Data[i], column[i])
		} else {
			g.Data[i] = append(g.Data[i][:index+1], g.Data[i][index:]...)
		}
	}

	return nil
}

func (g *Grid[T]) GetColumn(index int) ([]T, error) {
	if index < 0 || index >= g.ColSize() {
		return []T{}, errors.New("column index out of bounds")
	}

	column := make([]T, 0)
	for i := 0; i < g.RowSize(); i++ {
		column = append(column, g.Data[i][index])
	}
	return column, nil
}

func (g *Grid[T]) SetColumn(index int, column []T) error {
	if index < 0 || index >= g.ColSize() {
		return errors.New("column index out of bounds")
	}

	for i := 0; i < g.RowSize(); i++ {
		g.Set(i, index, column[i])
	}

	return nil
}

func (g *Grid[T]) AppendColumn(column []T) {
	g.AddColumn(g.ColSize(), column)
}

func (g *Grid[T]) ColContains(index int, item T) bool {
	col, err := g.GetColumn(index)
	if err != nil {
		return false
	}
	for _, colItem := range col {
		if colItem == item {
			return true
		}
	}
	return false
}
