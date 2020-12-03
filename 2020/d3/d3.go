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

func main() {
	trees := readData()
	rows := len(trees)
	columns := len(trees[0])
	treeCount := 0
	position := Position{X: 0, Y: 0}
	slope := Position{X: 1, Y: 3}

	for position.X < rows {
		if trees[position.X][position.Y] == "#" {
			treeCount++
		}

		position = position.addMod(slope, columns)
	}

	fmt.Println("Trees passed = ", treeCount)
}

// Local Variables:
// compile-command: "go build"
// End:
