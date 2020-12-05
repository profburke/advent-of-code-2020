package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readData() (passes []string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		pass := scanner.Text()
		passes = append(passes, pass)
	}

	return
}

func rowEncodingToInt(encoding string) int {
	upper := 128
	lower := 0

	for _, c := range encoding {
		if c == 'B' {
			lower += (upper - lower) / 2
		} else {
			upper = lower + (upper-lower)/2
		}
	}

	return lower
}

func seatEncodingToInt(encoding string) int {
	upper := 8
	lower := 0

	for _, c := range encoding {
		if c == 'R' {
			lower += (upper - lower) / 2
		} else {
			upper = lower + (upper-lower)/2
		}
	}

	return lower
}

func findSeatId(pass string) int {
	row := 0
	seat := 0

	rowEncoding := pass[0:7]
	row = rowEncodingToInt(rowEncoding)

	seatEncoding := pass[7:10]
	seat = seatEncodingToInt(seatEncoding)

	return row*8 + seat
}

func part1(passes []string) {
	var seatIds []int

	for _, pass := range passes {
		seatId := findSeatId(pass)
		seatIds = append(seatIds, seatId)
	}

	sort.Ints(seatIds)
	fmt.Println("Part 1 =", seatIds[len(seatIds)-1])
}

func part2(passes []string) {
}

func main() {
	passes := readData()

	part1(passes)
	part2(passes)
}

// Local Variables:
// compile-command: "go build"
// End:
