package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/iSkytran/2023adventofcode/utilities"
)

func parseSliceToInt(strSlice []string) []int {
	intSlice := make([]int, 0)
	for _, str := range strSlice {
		val, _ := strconv.Atoi(str)
		intSlice = append(intSlice, val)
	}
	return intSlice
}

func copySliceOfGames(games []*scratchoffGame) []*scratchoffGame {
	copy := make([]*scratchoffGame, 0)
	return append(copy, games...)
}

type scratchoffGame struct {
	id          int
	winningNums []int
	actualNums  []int
	numMatches  int
}

func newScratchoffGame(input string) *scratchoffGame {
	// Parse line into a game struct.
	game := new(scratchoffGame)
	gameStr, input, _ := strings.Cut(input, ":")
	fmt.Sscanf(gameStr[5:], "%d", &game.id)

	// Split around vertical bar and parse into two slice.
	arrays := strings.SplitN(input, "|", -1)
	winningNums := strings.Fields(arrays[0])
	actualNums := strings.Fields(arrays[1])
	game.winningNums = parseSliceToInt(winningNums)
	game.actualNums = parseSliceToInt(actualNums)
	game.computeMatches()

	return game
}

func (game *scratchoffGame) computeMatches() {
	// Store number of matches between winning and actual numbers.
	count := 0
	for _, actualNum := range game.actualNums {
		for _, winningNum := range game.winningNums {
			if actualNum == winningNum {
				count++
			}
		}
	}
	game.numMatches = count
}

func (game *scratchoffGame) computePoints() int {
	if game.numMatches == 0 {
		return 0
	}
	return int(math.Pow(2, float64(game.numMatches-1)))
}

func part1(path string) {
	scanner, file := utilities.OpenFile(path)
	defer file.Close()

	games := make([]*scratchoffGame, 0)
	for scanner.Scan() {
		line := scanner.Text()
		game := newScratchoffGame(line)
		games = append(games, game)
	}

	sum := 0
	for _, game := range games {
		sum += game.computePoints()
	}

	fmt.Printf("Total: %d\n", sum)
}

func part2(path string) {
	scanner, file := utilities.OpenFile(path)
	defer file.Close()

	games := make([]*scratchoffGame, 0)
	for scanner.Scan() {
		line := scanner.Text()
		game := newScratchoffGame(line)
		games = append(games, game)
	}

	// Continuously get game copies until there aren't any more.
	count := 0
	toProcess := copySliceOfGames(games)
	for len(toProcess) != 0 {
		// Dequeue toProcess.
		game := toProcess[0]
		toProcess = toProcess[1:]

		// Add following games based on number of matches.
		for i := 0; i < game.numMatches; i++ {
			// Game id is one more than index, but next one after current is desired.
			anotherGameIdx := game.id + i
			toProcess = append(toProcess, games[anotherGameIdx])
		}

		count++
	}

	fmt.Printf("Count: %d\n", count)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
