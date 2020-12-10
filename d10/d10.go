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

type Entry struct {
	Index   int
	Joltage int
}

// func pickBest(adaptors []int, currentJoltage int) (index, joltage int, err error) {
// 	oneJolts := make([]Entry, 0)
// 	twoJolts := make([]Entry, 0)
// 	threeJolts := make([]Entry, 0)

// 	for index, joltage = range adaptors {
// 		difference := adaptors[index] - currentJoltage

// 		if difference == 1 {
// 			oneJolts = append(oneJolts, Entry{Index: index, Joltage: joltage})
// 		} else if difference == 2 {
// 			twoJolts = append(twoJolts, Entry{Index: index, Joltage: joltage})
// 		} else if difference == 3 {
// 			threeJolts = append(threeJolts, Entry{Index: index, Joltage: joltage})
// 		}
// 	}

// 	if len(oneJolts) > 0 {
// 		index = oneJolts[0].Index
// 		joltage = oneJolts[0].Joltage
// 	} else if len(twoJolts) > 0 {
// 		index = twoJolts[0].Index
// 		joltage = twoJolts[0].Joltage
// 	} else if len(threeJolts) > 0 {
// 		index = threeJolts[0].Index
// 		joltage = threeJolts[0].Joltage
// 	} else {
// 		err = errors.New("couldn't find a compatible adaptor")
// 	}

// 	return
// }

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

	// for len(adaptors) > 0 {
	// 	index, joltage, err := pickBest(adaptors, currentJoltage)
	// 	if err != nil {
	// 		// TODO
	// 	}

	// 	difference := joltage - currentJoltage

	// 	if difference == 1 {
	// 		oneJoltDifferences++
	// 	} else if difference == 3 {
	// 		threeJoltDifferences++
	// 	} // difference could be 2 but we ignore it

	// 	currentJoltage += joltage
	// 	adaptors = removeAt(index, adaptors)
	// }

	// get max of adaptors + 3 to get joltage of your device
	// take final currentJoltage and find difference between it and your device

	threeJoltDifferences++ // for your device
	fmt.Println("Part 1 =", oneJoltDifferences*threeJoltDifferences)
}

func part2(adaptors []int) {
}

func main() {
	adaptors := readData()

	part1(adaptors)
	part2(adaptors)
}

// Local Variables:
// compile-command: "go build"
// End:
