package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	left  string
	right string
}

func readData() (pairs []Pair) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		pair := Pair{}
		pair.left = strings.TrimSpace(parts[0])
		pair.right = strings.TrimSpace(parts[1])

		pairs = append(pairs, pair)
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

func part1(pairs []Pair) {
	var mask string
	sum := 0
	memory := make(map[int]int)

	for _, pair := range pairs {
		if pair.left == "mask" {
			mask = pair.right
		} else {
			address, _ := strconv.Atoi(pair.left[4 : len(pair.left)-1])
			value, _ := strconv.Atoi(pair.right)
			memory[address] = applyMask(mask, value)
		}
	}

	for _, value := range memory {
		sum += value
	}

	fmt.Println("Part 1 =", sum)
}

func hasFloatingBit(a string) bool {
	return strings.Index(a, "X") > -1
}

func applyAddressMask(mask string, address int) (addresses []int) {
	addresses = make([]int, 0)

	address2 := strconv.FormatInt(int64(address), 2)
	padding := strings.Repeat("0", 36-len(address2))
	as := padding + address2

	base := ""

	for i, bit := range mask {
		switch bit {
		case '0':
			base += string(as[i])
		case '1', 'X':
			base += string(bit)
		}
	}

	var queue []string
	queue = append(queue, base)

	for len(queue) > 0 {
		e := queue[0]
		queue[0] = ""
		queue = queue[1:]

		if !hasFloatingBit(e) {
			addr, _ := strconv.ParseInt(e, 2, 64)
			addresses = append(addresses, int(addr))
		} else {
			i := strings.Index(e, "X")
			b := e[0:i] + "0" + e[i+1:]
			c := e[0:i] + "1" + e[i+1:]
			queue = append(queue, b)
			queue = append(queue, c)
		}
	}

	return
}

func part2(pairs []Pair) {
	var mask string
	sum := 0
	memory := make(map[int]int)

	for _, pair := range pairs {
		if pair.left == "mask" {
			mask = pair.right
		} else {
			address, _ := strconv.Atoi(pair.left[4 : len(pair.left)-1])
			value, _ := strconv.Atoi(pair.right)
			addresses := applyAddressMask(mask, address)
			for _, address := range addresses {
				memory[address] = value
			}
		}
	}

	for _, value := range memory {
		sum += value
	}

	fmt.Println("Part 2 =", sum)
}

func main() {
	pairs := readData()

	part1(pairs)
	part2(pairs)
}

// Local Variables:
// compile-command: "go build"
// End:
