package main

import "testing"

func TestValidByr(t *testing.T) {
	tests := []struct {
		doc  Document
		want bool
	}{
		{Document{Fields: map[string]string{"byr": "1920"}}, true},
		{Document{Fields: map[string]string{"byr": "2000"}}, true},
		{Document{Fields: map[string]string{"byr": "2002"}}, true},
		{Document{Fields: map[string]string{"byr": ""}}, false},
		{Document{Fields: map[string]string{"byr": "1919"}}, false},
		{Document{Fields: map[string]string{"byr": "2003"}}, false},
	}

	for _, tt := range tests {
		got := validByr(tt.doc.Fields["byr"])
		if got != tt.want {
			t.Errorf("got %v want %v", got, tt.want)
		}
	}
}

func TestValidIyr(t *testing.T) {
	tests := []struct {
		doc  Document
		want bool
	}{
		{Document{Fields: map[string]string{"iyr": "2010"}}, true},
		{Document{Fields: map[string]string{"iyr": "2020"}}, true},
		{Document{Fields: map[string]string{"iyr": "2012"}}, true},
		{Document{Fields: map[string]string{"iyr": ""}}, false},
		{Document{Fields: map[string]string{"iyr": "1999"}}, false},
		{Document{Fields: map[string]string{"iyr": "2023"}}, false},
	}

	for _, tt := range tests {
		got := validIyr(tt.doc.Fields["iyr"])
		if got != tt.want {
			t.Errorf("got %v want %v", got, tt.want)
		}
	}
}

func TestValidEyr(t *testing.T) {
	tests := []struct {
		doc  Document
		want bool
	}{
		{Document{Fields: map[string]string{"eyr": "2020"}}, true},
		{Document{Fields: map[string]string{"eyr": "2020"}}, true},
		{Document{Fields: map[string]string{"eyr": "2022"}}, true},
		{Document{Fields: map[string]string{"eyr": ""}}, false},
		{Document{Fields: map[string]string{"eyr": "1919"}}, false},
		{Document{Fields: map[string]string{"eyr": "2033"}}, false},
	}

	for _, tt := range tests {
		got := validEyr(tt.doc.Fields["eyr"])
		if got != tt.want {
			t.Errorf("got %v want %v", got, tt.want)
		}
	}
}

func TestValidHgt(t *testing.T) {
	tests := []struct {
		doc  Document
		want bool
	}{
		{Document{Fields: map[string]string{"hgt": "60in"}}, true},
		{Document{Fields: map[string]string{"hgt": "190cm"}}, true},
		{Document{Fields: map[string]string{"hgt": "190in"}}, false},
		{Document{Fields: map[string]string{"hgt": ""}}, false},
		{Document{Fields: map[string]string{"hgt": ""}}, false},
		{Document{Fields: map[string]string{"hgt": "190"}}, false},
		{Document{Fields: map[string]string{"hgt": "60cz"}}, false},
	}

	for _, tt := range tests {
		hgt := tt.doc.Fields["hgt"]
		got := validHgt(hgt)
		if got != tt.want {
			t.Errorf("%s: got %v want %v", hgt, got, tt.want)
		}
	}
}

func TestValidHcl(t *testing.T) {
	tests := []struct {
		doc  Document
		want bool
	}{
		{Document{Fields: map[string]string{"hcl": "#000000"}}, true},
		{Document{Fields: map[string]string{"hcl": "#123abc"}}, true},
		{Document{Fields: map[string]string{"hcl": "#847912"}}, true},
		{Document{Fields: map[string]string{"hcl": "#abcdef"}}, true},
		{Document{Fields: map[string]string{"hcl": ""}}, false},
		{Document{Fields: map[string]string{"hcl": "#0012"}}, false},
		{Document{Fields: map[string]string{"hcl": "123456"}}, false},
		{Document{Fields: map[string]string{"hcl": "#123GHI"}}, false},
	}

	for _, tt := range tests {
		hcl := tt.doc.Fields["hcl"]
		got := validHcl(hcl)
		if got != tt.want {
			t.Errorf("%s: got %v want %v", hcl, got, tt.want)
		}
	}
}

func TestValidEcl(t *testing.T) {
	tests := []struct {
		doc  Document
		want bool
	}{
		{Document{Fields: map[string]string{"ecl": "amb"}}, true},
		{Document{Fields: map[string]string{"ecl": "blu"}}, true},
		{Document{Fields: map[string]string{"ecl": "brn"}}, true},
		{Document{Fields: map[string]string{"ecl": "gry"}}, true},
		{Document{Fields: map[string]string{"ecl": "grn"}}, true},
		{Document{Fields: map[string]string{"ecl": "hzl"}}, true},
		{Document{Fields: map[string]string{"ecl": "oth"}}, true},
		{Document{Fields: map[string]string{"ecl": ""}}, false},
		{Document{Fields: map[string]string{"ecl": "Amb"}}, false},
		{Document{Fields: map[string]string{"ecl": "xyz"}}, false},
		{Document{Fields: map[string]string{"ecl": "am"}}, false},
	}

	for _, tt := range tests {
		ecl := tt.doc.Fields["ecl"]
		got := validEcl(ecl)
		if got != tt.want {
			t.Errorf("%s: got %v want %v", ecl, got, tt.want)
		}
	}
}

func TestValidPid(t *testing.T) {
	tests := []struct {
		doc  Document
		want bool
	}{
		{Document{Fields: map[string]string{"pid": "000000001"}}, true},
		{Document{Fields: map[string]string{"pid": "100000001"}}, true},
		{Document{Fields: map[string]string{"pid": "000239801"}}, true},
		{Document{Fields: map[string]string{"pid": "00000001"}}, false},
		{Document{Fields: map[string]string{"pid": "0000000001"}}, false},
		{Document{Fields: map[string]string{"pid": ""}}, false},
		{Document{Fields: map[string]string{"pid": "0K0000001"}}, false},
	}

	for _, tt := range tests {
		pid := tt.doc.Fields["pid"]
		got := validPid(pid)
		if got != tt.want {
			t.Errorf("%s: got %v want %v", pid, got, tt.want)
		}
	}
}
