package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readData() (adaptors []int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		adaptor, _ := strconv.Atoi(line)
		adaptors = append(adaptors, adaptor)
	}

	return
}

func removeAt(i int, list []int) (newList []int) {
	newList = make([]int, len(list))
	copy(newList, list)
	copy(newList[i:], newList[i+1:])
	newList = newList[:len(newList)-1]

	return
}

func part1(adaptors []int) {
	currentJoltage := 0
	oneJoltDifferences := 0
	threeJoltDifferences := 0

	for true {
		if len(adaptors) == 0 {
			break
		}

		sort.Ints(adaptors)

		joltage := adaptors[0]
		adaptors = removeAt(0, adaptors)

		difference := joltage - currentJoltage
		if difference == 1 {
			oneJoltDifferences++
		} else if difference == 3 {
			threeJoltDifferences++
		} // difference could be 2 but we ignore it

		currentJoltage = joltage
	}

	threeJoltDifferences++ // for your device
	fmt.Println("Part 1 =", oneJoltDifferences*threeJoltDifferences)
}

func part2(adaptors []int) {
	count := 0
	fmt.Println("Part 2=", count)
}

func main() {
	adaptors := readData()

	part1(adaptors)
	part2(adaptors)
}

// Local Variables:
// compile-command: "go build"
// End:
