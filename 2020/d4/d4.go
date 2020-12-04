package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// For part 1, we don't really care what's in the fields...
type Document struct {
	Fields map[string]string
}

func NewDocument() Document {
	d := Document{}
	d.Fields = make(map[string]string)

	return d
}

func readData() (documents []Document) {
	scanner := bufio.NewScanner(os.Stdin)

	current := NewDocument()

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			documents = append(documents, current)
			current = NewDocument()
			continue
		}

		fields := strings.Split(line, " ")

		for _, field := range fields {
			parts := strings.Split(field, ":")
			current.Fields[parts[0]] = parts[1]
		}
	}
	documents = append(documents, current)

	return
}

func isValid(document Document) bool {
	_, byr := document.Fields["byr"]
	_, iyr := document.Fields["iyr"]
	_, eyr := document.Fields["eyr"]
	_, hgt := document.Fields["hgt"]
	_, hcl := document.Fields["hcl"]
	_, ecl := document.Fields["ecl"]
	_, pid := document.Fields["pid"]

	return byr && iyr && eyr && hgt && hcl && ecl && pid
}

func part1(documents []Document) {
	validPassports := 0

	for _, document := range documents {
		if isValid(document) {
			validPassports++
		}
	}

	fmt.Println("Part 1 = ", validPassports)
}

func part2(documents []Document) {
}

func main() {
	documents := readData()

	part1(documents)
	part2(documents)
}

// Local Variables:
// compile-command: "go build"
// End:
