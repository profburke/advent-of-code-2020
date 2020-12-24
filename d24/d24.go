package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinates struct {
	X int
	Y int
}

type Direction int

const (
	West      Direction = iota
	Northwest           = iota
	Northeast           = iota
	East                = iota
	Southeast           = iota
	Southwest           = iota
)

func (d Direction) String() string {
	return [...]string{"West", "Northwest", "Northeast", "East",
		"Southeast", "Southwest"}[d]
}

func (d Direction) From(location Coordinates) (newLocation Coordinates) {
	deltaE, deltaW := 0, -1
	if location.Y%2 != 0 {
		deltaE, deltaW = 1, 0
	}

	switch d {
	case East:
		newLocation = Coordinates{X: location.X + 1, Y: location.Y}
	case West:
		newLocation = Coordinates{X: location.X - 1, Y: location.Y}
	case Northeast:
		newLocation = Coordinates{X: location.X + deltaE, Y: location.Y - 1}
	case Northwest:
		newLocation = Coordinates{X: location.X + deltaW, Y: location.Y - 1}
	case Southeast:
		newLocation = Coordinates{X: location.X + deltaE, Y: location.Y + 1}
	case Southwest:
		newLocation = Coordinates{X: location.X + deltaW, Y: location.Y + 1}
	}

	return
}

func determineSteps(line string) (steps []Direction) {
	steps = make([]Direction, 0)

	complete := false
	var dir Direction
	var prev rune
	for _, d := range line {
		switch d {
		case 'n':
			prev = 'n'
		case 's':
			prev = 's'
		case 'e':
			if prev == 0 {
				dir = East
				complete = true
				prev = 0
			} else {
				if prev == 'n' {
					dir = Northeast
				} else {
					dir = Southeast
				}
				complete = true
				prev = 0
			}
		case 'w':
			if prev == 0 {
				dir = West
				complete = true
				prev = 0
			} else {
				if prev == 'n' {
					dir = Northwest
				} else {
					dir = Southwest
				}
				complete = true
				prev = 0
			}
		}

		if complete {
			steps = append(steps, dir)
			complete = false
		}
	}

	return
}

func readData() (data [][]Direction) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		steps := determineSteps(line)
		data = append(data, steps)
	}

	return
}

func walk(steps []Direction) (location Coordinates) {
	location = Coordinates{X: 0, Y: 0}
	for _, step := range steps {
		location = step.From(location)
	}

	return
}

func part1(data [][]Direction) {
	blackTiles := make(map[Coordinates]bool)

	for _, steps := range data {
		location := walk(steps)

		if _, found := blackTiles[location]; found {
			delete(blackTiles, location)
		} else {
			blackTiles[location] = true
		}
	}

	fmt.Println("Part 1 =", len(blackTiles))
}

func part2(data [][]Direction) {
}

func main() {
	data := readData()

	part1(data)
	part2(data)
}

// Local Variables:
// compile-command: "go build"
// End:
