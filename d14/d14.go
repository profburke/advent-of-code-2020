package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readData() (memory map[int]int) {
	var mask string
	var parts []string
	memory = make(map[int]int)
	re, _ := regexp.Compile("mem\\[(\\d+)\\] = (\\d+)")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if line[0:4] == "mask" {
			parts = strings.Split(line, "=")
			mask = strings.TrimSpace(parts[1])
		} else {
			parts = re.FindStringSubmatch(line)
			address, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			memory[address] = applyMask(mask, value)
		}
	}

	return
}

func applyMask(mask string, value int) int {
	high := len(mask) - 1

	for index, bit := range mask {
		if bit == 'X' {
			continue
		}

		value2 := 1 << (high - index)
		if bit == '1' {
			value = value | value2
		} else {
			if value&value2 > 0 {
				value -= value2
			}
		}
	}

	return value
}

func part1(memory map[int]int) {
	sum := 0

	for _, value := range memory {
		sum += value
	}

	fmt.Println("Part 1 =", sum)
}

func part2(memory map[int]int) {
}

func main() {
	memory := readData()

	part1(memory)
	part2(memory)
}

// Local Variables:
// compile-command: "go build"
// End:
