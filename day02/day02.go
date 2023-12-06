package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const redMax = 12
const greenMax = 13
const blueMax = 14

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type game struct {
	id     int
	rounds []*round
}

type round struct {
	red   int
	green int
	blue  int
}

func newGame(str string) *game {
	g := new(game)

	// Parse id.
	idStr, str, _ := strings.Cut(str, ":")
	fmt.Sscanf(idStr[5:], "%d", &g.id)

	// Parse rounds.
	rounds := strings.Split(str, ";")
	for _, rStr := range rounds {
		r := newRound(rStr)
		g.rounds = append(g.rounds, r)
	}

	return g
}

func newRound(str string) *round {
	r := new(round)

	// Parse colors.
	colors := strings.Split(str, ",")
	for _, colorStr := range colors {
		var count int
		var color string
		fmt.Sscanf(colorStr, "%d %s", &count, &color)
		r.addColor(color, count)
	}

	return r
}

func (g *game) valid() bool {
	// Check if all the rounds are valid.
	for _, r := range g.rounds {
		if !r.valid() {
			return false
		}
	}
	return true
}

func (g *game) max() *round {
	// Compute fewest of each color required.
	m := new(round)
	for _, r := range g.rounds {
		if r.red > m.red {
			m.red = r.red
		}
		if r.green > m.green {
			m.green = r.green
		}
		if r.blue > m.blue {
			m.blue = r.blue
		}
	}
	return m
}

func (r *round) power() int {
	return r.red * r.green * r.blue
}

func (r *round) valid() bool {
	// Check if a round is valid.
	return r.red <= redMax && r.green <= greenMax && r.blue <= blueMax
}

func (g *round) addColor(color string, count int) {
	switch color {
	case "red":
		g.red += count
	case "green":
		g.green += count
	case "blue":
		g.blue += count
	}
}

func part1(path string) {
	// Open file.
	file, err := os.Open(path)
	check(err)
	scanner := bufio.NewScanner(file)

	// Parse input.
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		g := newGame(line)

		// Add valid game to sum.
		if g.valid() {
			sum += g.id
		}
	}

	fmt.Printf("Total: %d\n", sum)
}

func part2(path string) {
	// Open file.
	file, err := os.Open(path)
	check(err)
	scanner := bufio.NewScanner(file)

	// Parse input.
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		power := newGame(line).max().power()

		// Add powers together.
		sum += power
	}

	fmt.Printf("Total: %d\n", sum)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
