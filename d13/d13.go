package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readData() (timestamp int, buses []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	timestamp, _ = strconv.Atoi(line)

	scanner.Scan()
	line = scanner.Text()

	for _, s := range strings.Split(line, ",") {
		if s == "x" {
			continue
		}

		bus, _ := strconv.Atoi(s)
		buses = append(buses, bus)
	}

	return
}

func part1(timestamp int, buses []int) {
	minWait := math.MaxInt64
	busId := 0

	for _, bus := range buses {
		r := timestamp % bus
		w := bus - r

		if w < minWait {
			minWait = w
			busId = bus
		}
	}

	fmt.Println(minWait, busId, minWait*busId)
}

func part2(timestamp int, buses []int) {
}

func main() {
	timestamp, buses := readData()

	part1(timestamp, buses)
	part2(timestamp, buses)
}

// Local Variables:
// compile-command: "go build"
// End:
