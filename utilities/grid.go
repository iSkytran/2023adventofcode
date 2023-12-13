package utilities

import (
	"errors"
)

type Grid[T comparable] struct {
	data    [][]T
	rowSize int
	colSize int
}

func NewGrid[T comparable]() *Grid[T] {
	g := new(Grid[T])
	g.data = make([][]T, 0)
	return g
}

func (g *Grid[_]) Shape() (int, int) {
	return g.rowSize, g.colSize
}

func (g *Grid[T]) Get(rowIndex int, columnIndex int) (T, error) {
	if rowIndex < 0 || rowIndex > g.rowSize {
		return *new(T), errors.New("row index out of bounds")
	}

	if columnIndex < 0 || columnIndex > g.colSize {
		return *new(T), errors.New("column index out of bounds")
	}

	return g.data[rowIndex][columnIndex], nil
}

func (g *Grid[T]) Set(rowIndex int, columnIndex int, item T) error {
	if rowIndex < 0 || rowIndex > g.rowSize {
		return errors.New("row index out of bounds")
	}

	if columnIndex < 0 || columnIndex > g.colSize {
		return errors.New("column index out of bounds")
	}

	g.data[rowIndex][columnIndex] = item
	return nil
}

func (g *Grid[T]) AddRow(index int, row []T) error {
	if index < 0 || index > g.rowSize {
		return errors.New("index out of bounds")
	}

	if index == g.rowSize {
		g.data = append(g.data, row)
	} else {
		g.data = append(g.data[:index+1], g.data[index:]...)
		g.data[index] = row
	}

	g.rowSize++
	return nil
}

func (g *Grid[T]) GetRow(index int) ([]T, error) {
	if index < 0 || index > g.rowSize {
		return []T{}, errors.New("row index out of bounds")
	}

	return g.data[index], nil
}

func (g *Grid[T]) AddColumn(index int, column []T) error {
	if index < 0 || index > g.colSize {
		return errors.New("index out of bounds")
	}

	for i := 0; i < g.rowSize; i++ {
		if index == g.colSize {
			g.data[i] = append(g.data[i], column[i])
		} else {
			g.data[i] = append(g.data[i][:index+1], g.data[i][index:]...)
		}
	}

	g.colSize++
	return nil
}

func (g *Grid[T]) GetColumn(index int) ([]T, error) {
	if index < 0 || index > g.colSize {
		return []T{}, errors.New("column index out of bounds")
	}

	column := make([]T, 0)
	for i := 0; i < g.colSize; i++ {
		column = append(column, g.data[i][index])
	}
	return column, nil
}

func (g *Grid[T]) AppendRow(row []T) {
	g.AddRow(g.rowSize, row)
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
