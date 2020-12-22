package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deck []int

func readData() (decks []Deck) {
	scanner := bufio.NewScanner(os.Stdin)
	decks = make([]Deck, 0)
	deck := make(Deck, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Player") {
			continue
		}

		if line == "" {
			decks = append(decks, deck)
			deck = make(Deck, 0)
			continue
		}

		card, _ := strconv.Atoi(line)
		deck = append(deck, card)
	}
	decks = append(decks, deck)

	return
}

func gameOver(decks []Deck) (winner int) {
	for i, deck := range decks {
		if len(deck) == 0 {
			return 1 - i
		}
	}

	return -1
}

func score(deck Deck) (result int) {
	multiplier := len(deck)

	for _, card := range deck {
		result += (card * multiplier)
		multiplier--
	}

	return
}

func part1(decks []Deck) {
	var winner int

	for true {
		if winner = gameOver(decks); winner > -1 {
			break
		}

		aCard := decks[0][0]
		decks[0] = decks[0][1:]

		bCard := decks[1][0]
		decks[1] = decks[1][1:]

		if aCard > bCard {
			decks[0] = append(decks[0], aCard, bCard)
		} else {
			decks[1] = append(decks[1], bCard, aCard)
		}
	}

	score := score(decks[winner])
	fmt.Println("Part 1 =", score)
}

func part2(deck []Deck) {
}

func main() {
	decks := readData()

	part1(decks)
	part2(decks)
}

// Local Variables:
// compile-command: "go build"
// End:
