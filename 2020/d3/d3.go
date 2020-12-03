package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	X, Y int
}

func (p Position) addMod(other Position, bound int) (sum Position) {
	sum.Y = (p.Y + other.Y) % bound
	sum.X = p.X + other.X

	return
}

func readData() (trees [][]string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		trees = append(trees, strings.Split(line, ""))
	}

	return
}

func countCollisions(trees [][]string, slope Position) int {
	rows := len(trees)
	columns := len(trees[0])
	treeCount := 0
	position := Position{X: 0, Y: 0}

	for position.X < rows {
		if trees[position.X][position.Y] == "#" {
			treeCount++
		}

		position = position.addMod(slope, columns)
	}

	return treeCount
}

func part1(trees [][]string) {
	slope := Position{X: 1, Y: 3}

	treeCount := countCollisions(trees, slope)
	fmt.Println("Trees passed =", treeCount)
}

func part2(trees [][]string) {
	product := 1
	slopes := []Position{
		Position{X: 1, Y: 1},
		Position{X: 1, Y: 3},
		Position{X: 1, Y: 5},
		Position{X: 1, Y: 7},
		Position{X: 2, Y: 1},
	}

	for _, slope := range slopes {
		product *= countCollisions(trees, slope)
	}
	fmt.Println("Product of the collisions =", product)
}

func main() {
	trees := readData()

	part1(trees)
	part2(trees)
}

// Local Variables:
// compile-command: "go build"
// End:
