package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/iSkytran/2023adventofcode/utilities"
)

const (
	// Card rankings from lowest to highest.
	highCard = iota
	onePair
	twoPair
	threeKind
	fullHouse
	fourKind
	fiveKind
)

var cardValues = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

type camelHand struct {
	cards    string
	bid      int
	handType int
}

func compareCamelHandsNoWildcard(a, b *camelHand) int {
	return compareCamelHands(a, b, false)
}

func compareCamelHandsWildcard(a, b *camelHand) int {
	return compareCamelHands(a, b, true)
}

func compareCamelHands(a, b *camelHand, wildcard bool) int {
	// Sort by hand type.
	if a.handType != b.handType {
		return a.handType - b.handType
	}

	// Sort by card value if handTypes are the same.
	for i := 0; i < len(a.cards); i++ {
		if a.cards[i] != b.cards[i] {
			// Compare using this card.
			valA := cardValues[rune(a.cards[i])]
			valB := cardValues[rune(b.cards[i])]

			if wildcard {
				// Make jokers the lowest value.
				if valA == cardValues['J'] {
					valA = -1
				}

				if valB == cardValues['J'] {
					valB = -1
				}
			}

			return valA - valB
		}
	}

	// Same hands.
	return 0
}

func computeWinnings(hands []*camelHand) int {
	winnings := 0
	for idx, hand := range hands {
		// Rank is one more than the index.
		winnings += (idx + 1) * hand.bid
	}
	return winnings
}

func computeRank(cards string, wildcard bool) int {
	// Compute type of rank from the cards in a hand.
	cardSet := map[rune]int{}

	// Add each card to the set.
	for _, card := range cards {
		_, ok := cardSet[card]
		if ok {
			// Card in set already.
			cardSet[card]++
		} else {
			cardSet[card] = 1
		}
	}

	_, ok := cardSet['J']
	if !ok {
		// Don't do wildcard processing.
		wildcard = false
	}

	// Hand type based on set's length.
	switch len(cardSet) {
	case 1:
		// Five of a kind.
		// No special case for all joker wildcards.
		return fiveKind
	case 2:
		// Promote to five of a kind if wildcard.
		if wildcard {
			return fiveKind
		}
		// Four of a kind or full house.
		for _, count := range cardSet {
			if count == 4 {
				// Four of a kind.
				return fourKind
			}
		}
		// Full house.
		return fullHouse
	case 3:
		// Promote to three of a kind or full house if wildcard.
		if wildcard {
			// Only full house if one joker and two of each of the other cards.
			if cardSet['J'] == 1 {
				for key, value := range cardSet {
					if key != 'J' && value == 2 {
						return fullHouse
					}
				}
			}
			return fourKind
		}
		// Two pair or three of a kind.
		for _, count := range cardSet {
			if count == 3 {
				// Three of a kind.
				return threeKind
			}
		}
		// Two pair.
		return twoPair
	case 4:
		// Promote to three of a kind if wildcard.
		if wildcard {
			return threeKind
		}
		// One pair.
		return onePair
	case 5:
		// Promote to pair if wildcard.
		if wildcard {
			return onePair
		}
		// High card.
		return highCard
	}

	// Invalid hand.
	return -1
}

func parseHands(path string, wildcard bool) []*camelHand {
	scanner, file := utilities.OpenFile(path)
	defer file.Close()

	// Get race times.
	var hands []*camelHand
	for scanner.Scan() {
		line := scanner.Text()

		hand := new(camelHand)
		hands = append(hands, hand)
		fmt.Sscanf(line, "%s %d", &hand.cards, &hand.bid)
		hand.handType = computeRank(hand.cards, wildcard)
	}

	return hands
}

func part1(path string) {
	hands := parseHands(path, false)
	slices.SortFunc(hands, compareCamelHandsNoWildcard)
	winnings := computeWinnings(hands)
	fmt.Printf("Winnings: %d\n", winnings)
}

func part2(path string) {
	hands := parseHands(path, true)
	slices.SortFunc(hands, compareCamelHandsWildcard)
	winnings := computeWinnings(hands)
	fmt.Printf("Winnings: %d\n", winnings)
}

func main() {
	// Input file.
	path := os.Args[1]
	part1(path)
	part2(path)
}
