package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readData() (data []int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)

		if err != nil {
			fmt.Fprintln(os.Stderr, "converting string to int: ", err)
			os.Exit(1)
		}

		data = append(data, i)
	}

	return data
}

func part1(expenses []int) {
	for _, val := range expenses {
		other := 2020 - val
		index := sort.SearchInts(expenses, other)

		if expenses[index] == other {
			fmt.Fprintln(os.Stdout, val, "*", other, "=", val*other)
			return
		}
	}

	fmt.Fprintln(os.Stdout, "Couldn't find a match")
}

func part2(expenses []int) {
	for i, val1 := range expenses {
		target := 2020 - val1

		var reducedExpenses = make([]int, len(expenses))
		copy(reducedExpenses, expenses)

		copy(reducedExpenses[i:], reducedExpenses[i+1:])
		tail := len(reducedExpenses) - 1
		reducedExpenses[tail] = 0
		reducedExpenses = reducedExpenses[:tail]

		for _, val := range reducedExpenses {
			other := target - val
			index := sort.SearchInts(reducedExpenses, other)

			if reducedExpenses[index] == other {
				fmt.Fprintln(os.Stdout, val1, "*", val, "*", other, "=",
					val1*val*other)
				return
			}
		}
	}

	fmt.Fprintln(os.Stdout, "Couldn't find a match")
}

func main() {
	expenses := readData()
	sort.Ints(expenses)

	part1(expenses)
	part2(expenses)
}

// Local Variables:
// compile-command: "go build"
// End:
