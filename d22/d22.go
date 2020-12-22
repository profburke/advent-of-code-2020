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

func part1(decks []Deck) {
	var winner int

	for true {
		aLen := len(decks[0])
		bLen := len(decks[1])

		if aLen == 0 {
			winner = 1
			break
		} else if bLen == 0 {
			winner = 0
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

	// determine winner's score
	score := 0
	multiplier := len(decks[winner])

	for _, card := range decks[winner] {
		score += (card * multiplier)
		multiplier--
	}

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
