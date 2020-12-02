package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ValidationRule struct {
	Upper int
	Lower int
	C     rune
}

func (rule ValidationRule) isValid(candidate string) bool {
	occurences := 0

	for _, char := range candidate {
		if char == rule.C {
			occurences++
		}
	}
	return (occurences >= rule.Lower && occurences <= rule.Upper)
}

func (rule ValidationRule) isValid2(candidate string) bool {
	occurences := 0
	len := len(candidate)

	if rule.Lower <= len && rune(candidate[rule.Lower-1]) == rule.C {
		occurences++
	}

	if rule.Upper <= len && rune(candidate[rule.Upper-1]) == rule.C {
		occurences++
	}

	return (occurences == 1)
}

type Entry struct {
	Candidate string
	Rule      ValidationRule
}

func readData() (entries []Entry) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		pieces := strings.Fields(line)

		bounds := strings.FieldsFunc(pieces[0], func(r rune) bool {
			if r == '-' {
				return true
			} else {
				return false
			}
		})

		lower, _ := strconv.Atoi(bounds[0])
		upper, _ := strconv.Atoi(bounds[1])

		c := rune(strings.Trim(pieces[1], ":")[0])

		candidate := strings.Trim(pieces[2], " ")

		rule := ValidationRule{Lower: lower, Upper: upper, C: c}
		entry := Entry{Candidate: candidate, Rule: rule}
		entries = append(entries, entry)
	}

	return entries
}

func part1(entries []Entry) {
	validPasswords := 0
	for _, entry := range entries {
		if entry.Rule.isValid(entry.Candidate) {
			validPasswords++
		}
	}

	fmt.Println("The number of valid passwords is", validPasswords)
}

func part2(entries []Entry) {
	validPasswords := 0
	for _, entry := range entries {
		if entry.Rule.isValid2(entry.Candidate) {
			validPasswords++
		}
	}

	fmt.Println("The number of valid passwords is", validPasswords)
}

func main() {
	entries := readData()

	part1(entries)
	part2(entries)
}

// Local Variables:
// compile-command: "go build"
// End:
