package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Range struct {
	Low  int
	High int
}

func (r Range) Contains(value int) bool {
	return value >= r.Low && value <= r.High
}

type Field struct {
	First  Range
	Second Range
}

func (f Field) Valid(value int) bool {
	return f.First.Contains(value) || f.Second.Contains(value)
}

type Ticket []int

func (t Ticket) Valid(value int, fields []Field) bool {
	for _, field := range fields {
		if field.Valid(value) {
			return true
		}
	}

	return false
}

func readData() (fields []Field, myTicket Ticket, tickets []Ticket) {
	scanner := bufio.NewScanner(os.Stdin)
	fields = make([]Field, 0)
	tickets = make([]Ticket, 0)
	re, _ := regexp.Compile(".+: (\\d+)-(\\d+) or (\\d+)-(\\d+)")

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := re.FindStringSubmatch(line)
		l1, _ := strconv.Atoi(parts[1])
		h1, _ := strconv.Atoi(parts[2])
		l2, _ := strconv.Atoi(parts[3])
		h2, _ := strconv.Atoi(parts[4])

		field := Field{}
		field.First = Range{Low: l1, High: h1}
		field.Second = Range{Low: l2, High: h2}

		fields = append(fields, field)
	}

	// deal with my ticket
	myTicket = make([]int, 0)
	scanner.Scan() // skip line "your ticket:"
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(line, ",")
	for _, part := range parts {
		value, _ := strconv.Atoi(part)
		myTicket = append(myTicket, value)
	}

	// deal with nearby tickets
	scanner.Scan() // skip blank line
	scanner.Scan() // skip line "nearby tickets:"
	for scanner.Scan() {
		line = scanner.Text()
		parts := strings.Split(line, ",")
		ticket := make([]int, 0)
		for _, part := range parts {
			value, _ := strconv.Atoi(part)
			ticket = append(ticket, value)
		}
		tickets = append(tickets, ticket)
	}

	return
}

func part1(fields []Field, tickets []Ticket) {
	esr := 0

	for _, ticket := range tickets {
		for _, value := range ticket {
			if !ticket.Valid(value, fields) {
				esr += value
			}
		}
	}

	fmt.Println("Part 1 =", esr)
}

func main() {
	fields, _, tickets := readData()

	part1(fields, tickets)
}

// Local Variables:
// compile-command: "go build"
// End:
