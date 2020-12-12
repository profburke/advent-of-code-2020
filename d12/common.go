package main

import (
	"bufio"
	"os"
	"strconv"
)

type Command struct {
	Direction string
	Distance  int
}

type Coordinates struct {
	X int
	Y int
}

type Heading int

const (
	N Heading = iota
	E
	S
	W
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
