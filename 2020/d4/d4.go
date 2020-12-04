package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func hasRequiredFields(document Document) bool {
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
		if hasRequiredFields(document) {
			validPassports++
		}
	}

	fmt.Println("Part 1 = ", validPassports)
}

func validByr(byr string) bool {
	val, err := strconv.Atoi(byr)
	if err != nil {
		return false
	}

	return (val >= 1920 && val <= 2002)
}

func validIyr(iyr string) bool {
	val, err := strconv.Atoi(iyr)
	if err != nil {
		return false
	}

	return (val >= 2010 && val <= 2020)
}

func validEyr(eyr string) bool {
	val, err := strconv.Atoi(eyr)
	if err != nil {
		return false
	}

	return (val >= 2020 && val <= 2030)
}

func validHgt(hgt string) bool {
	re, _ := regexp.Compile("(\\d+)(in|cm)")
	parts := re.FindStringSubmatch(hgt)

	// first item in slice is entire match
	if len(parts) != 3 {
		return false
	}

	val, _ := strconv.Atoi(parts[1])

	if parts[2] == "cm" {
		return (val >= 150 && val <= 193)
	} else if parts[2] == "in" {
		return (val >= 59 && val <= 76)
	} else {
		return false
	}
}

func validHcl(hcl string) bool {
	matched, _ := regexp.Match(`#[0-9a-f]{6}`, []byte(hcl))
	return matched
}

func validEcl(ecl string) bool {
	return ecl == "amb" || ecl == "blu" || ecl == "brn" ||
		ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth"
}

func validPid(pid string) bool {
	// for some reason, this is returning true for strings that are less than 9 digits long

	if len(pid) != 9 {
		return false
	}

	matched, _ := regexp.Match(`\d{9}`, []byte(pid))
	return matched
}

func fieldsAreValid(document Document) bool {
	return validByr(document.Fields["byr"]) && validIyr(document.Fields["iyr"]) &&
		validEyr(document.Fields["eyr"]) && validHgt(document.Fields["hgt"]) &&
		validHcl(document.Fields["hcl"]) && validEcl(document.Fields["ecl"]) &&
		validPid(document.Fields["pid"])
}

func part2(documents []Document) {
	validPassports := 0

	for _, document := range documents {
		if hasRequiredFields(document) && fieldsAreValid(document) {
			validPassports++
		}
	}

	fmt.Println("Part 2 = ", validPassports)
}

func main() {
	documents := readData()

	part1(documents)
	part2(documents)
}

// Local Variables:
// compile-command: "go build"
// End:
