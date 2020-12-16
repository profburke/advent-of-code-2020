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

func (t Ticket) Valid(value int, fields map[string]Field) bool {
	for _, field := range fields {
		if field.Valid(value) {
			return true
		}
	}

	return false
}

func readData() (fields map[string]Field, myTicket Ticket, tickets []Ticket) {
	scanner := bufio.NewScanner(os.Stdin)
	fields = make(map[string]Field)
	tickets = make([]Ticket, 0)
	re, _ := regexp.Compile("(.+): (\\d+)-(\\d+) or (\\d+)-(\\d+)")

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := re.FindStringSubmatch(line)
		l1, _ := strconv.Atoi(parts[2])
		h1, _ := strconv.Atoi(parts[3])
		l2, _ := strconv.Atoi(parts[4])
		h2, _ := strconv.Atoi(parts[5])

		field := Field{}
		field.First = Range{Low: l1, High: h1}
		field.Second = Range{Low: l2, High: h2}

		fields[parts[1]] = field
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

func part1(fields map[string]Field, tickets []Ticket) (validTickets []Ticket) {
	esr := 0
	validTickets = make([]Ticket, 0)

	for _, ticket := range tickets {
		valid := true

		for _, value := range ticket {
			if !ticket.Valid(value, fields) {
				valid = false
				esr += value
			}
		}

		if valid {
			validTickets = append(validTickets, ticket)
		}
	}

	fmt.Println("Part 1 =", esr)

	return
}

func part2(fields map[string]Field, myTicket Ticket, Tickets []Ticket) {
	// identifiedFieldNames := make([]string, len(fields))

	product := 1
	// for _, name := range identifiedFieldNames {
	// 	if name[0] == 'd' {
	// 		product *= fields[name]
	// 	}
	// }

	fmt.Println("Part 2 =", product)
}

func main() {
	fields, myTicket, tickets := readData()

	validTickets := part1(fields, tickets)
	part2(fields, myTicket, validTickets)
}

// Local Variables:
// compile-command: "go build"
// End:
