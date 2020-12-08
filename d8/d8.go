package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Opcode  string
	Operand int
}

func readData() (program []Instruction) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		i := Instruction{}
		i.Opcode = parts[0]
		operand, _ := strconv.Atoi(parts[1])
		i.Operand = operand

		program = append(program, i)
	}

	return
}

func part1(program []Instruction) {
	pc := 0
	acc := 0
	seen := make(map[int]bool)

	for true {
		_, found := seen[pc]
		if found {
			break
		}

		seen[pc] = true
		i := program[pc]
		fmt.Println(pc, "|", i)
		switch i.Opcode {
		case "acc":
			acc += i.Operand
		case "jmp":
			pc += (i.Operand - 1)
		case "nop":
		}

		pc++

	}

	fmt.Println("part 1 =", acc)
}

func part2(program []Instruction) {
	pc := 0
	acc := 0
	program[600].Opcode = "nop"

	for true {
		if pc == len(program) {
			break
		}

		i := program[pc]

		switch i.Opcode {
		case "acc":
			acc += i.Operand
		case "jmp":
			pc += (i.Operand - 1)
		case "nop":
		}

		pc++

	}

	fmt.Println("part 2 =", acc)
}

func main() {
	program := readData()

	// part1(program)
	part2(program)
}

// Local Variables:
// compile-command: "go build"
// End:
