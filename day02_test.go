package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestDay02Checksum(t *testing.T) {
	tests := []struct {
		Input  string
		Expect int
	}{
		{`
abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab
`, 12,
		},
	}

	for _, tc := range tests {
		actual, err := checksum(strings.NewReader(tc.Input))
		if err != nil {
			t.Fatalf("day02 checksum(%v): Received unexpected error %v", tc.Input, err)
		}

		if actual != tc.Expect {
			t.Errorf("day02 checksum expected %v, got %v", tc.Expect, actual)
		}
	}
}

func TestDay02ChecksumPartial(t *testing.T) {
	tests := []struct {
		Input  string
		Expect []int
	}{
		{"abcdef", []int{0, 0}},
		{"bababc", []int{1, 1}},
		{"abbcde", []int{1, 0}},
		{"abcccd", []int{0, 1}},
		{"aabcdd", []int{1, 0}},
		{"abcdee", []int{1, 0}},
		{"ababab", []int{0, 1}},
	}

	for _, tc := range tests {
		twice, thrice := checksumPartial(tc.Input)
		if twice != tc.Expect[0] || thrice != tc.Expect[1] {
			t.Errorf("day02 checksumPartial expected %v, %v got %v, %v", tc.Expect[0], tc.Expect[1], twice, thrice)
		}
	}
}

func TestDay02CompareIds(t *testing.T) {
	tests := []struct {
		i1 []rune
		i2 []rune
		e1 bool
		e2 int
	}{
		{[]rune("abcde"), []rune("axcye"), false, 0},
		{[]rune("fghij"), []rune("fguij"), true, 2},
	}

	for _, tc := range tests {
		isPair, i, err := compareIds(tc.i1, tc.i2)
		if err != nil {
			t.Fatalf("day02 compareIds received unexpected error: %v", err)
		}

		if isPair != tc.e1 || i != tc.e2 {
			t.Errorf("day02 compareIds expected %v, %v got %v, %v", tc.e1, tc.e2, isPair, i)
		}
	}
}

func TestDay02FindCommonLetters(t *testing.T) {
	tests := []struct {
		Input  string
		Expect []rune
	}{
		{`
abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz
`, []rune("fgij"),
		},
	}

	for _, tc := range tests {

		commonLetters, err := findCommonLetters(strings.NewReader(tc.Input))
		if err != nil {
			t.Fatalf("day02 findCommonLetters received unexpected error: %v", err)
		}

		if !reflect.DeepEqual(tc.Expect, commonLetters) {
			t.Errorf("day02 findCommonLetters expected %s, got %s", string(tc.Expect), string(commonLetters))
		}
	}
}
