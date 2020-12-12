package main

import (
	"bufio"
	"os"
	"strconv"
)

func readData() (data []Command) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		distance, _ := strconv.Atoi(line[1:])
		command := Command{Direction: line[0:1], Distance: distance}
		data = append(data, command)
	}

	return
}

func main() {
	commands := readData()

	part1(commands)
	part2(commands)
}

// Local Variables:
// compile-command: "go build"
// End:
