package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

func checksumPartial(s string) (int, int) {
	var letterMap = map[rune]int{}
	exactlyTwice, exactlyThrice := 0, 0
	for _, r := range s {
		letterMap[r]++
	}

	for _, v := range letterMap {
		switch v {
		case 2:
			exactlyTwice = 1
		case 3:
			exactlyThrice = 1
		}
	}

	return exactlyTwice, exactlyThrice
}

func checksum(f io.ReadSeeker) (int, error) {
	s := bufio.NewScanner(f)

	exactlyTwice, exactlyThrice := 0, 0

	for s.Scan() {
		twice, thrice := checksumPartial(s.Text())
		exactlyTwice += twice
		exactlyThrice += thrice
	}

	if s.Err() != nil {
		return 0, s.Err()
	}

	checksum := exactlyTwice * exactlyThrice
	return checksum, nil
}

// return bool: was exactly one letter different
//         int: index at which a letter was different
//       error
func compareIds(a, b []rune) (bool, int, error) {
	exactlyOneDifferent := false
	indexOfDif := 0

	// I think all the strings are the same length
	if len(a) != len(b) {
		return false, 0, fmt.Errorf("day02 compareIds(%s, %s) received unexpected input. Expected a and b to be the same length", string(a), string(b))
	}
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}

		if exactlyOneDifferent {
			return false, 0, nil
		}

		exactlyOneDifferent = true
		indexOfDif = i
	}

	return exactlyOneDifferent, indexOfDif, nil
}

// there's probably a more efficient way to do this
func findCommonLetters(f io.ReadSeeker) ([]rune, error) {
	var ids [][]rune
	var commonLetters []rune
	s := bufio.NewScanner(f)

ScanIDs:
	for s.Scan() {
		currID := []rune(s.Text())

		for _, id := range ids {
			// this is super dumb, must be a better way
			if len(id) == 0 {
				continue
			}

			isPair, indexOfDif, err := compareIds(id, currID)
			if err != nil {
				return commonLetters, err
			}

			if isPair {
				commonLetters = append(
					currID[:indexOfDif],
					currID[indexOfDif+1:]...,
				)

				break ScanIDs
			}
		}

		ids = append(ids, currID)
	}

	if s.Err() != nil {
		return commonLetters, s.Err()
	}

	return commonLetters, nil
}

func runDay02(f io.ReadSeeker) {
	checksum, err := checksum(f)
	if err != nil {
		log.Fatalf("day02 checksum: Encountered unexpected error %v", err)
	}

	f.Seek(0, 0)
	commonLetters, err := findCommonLetters(f)
	if err != nil {
		log.Fatalf("day02 findCommonLetters: Encountered unexpected error %v", err)
	}
	fmt.Printf("part1: %d\n", checksum)
	fmt.Printf("part2: %s\n", string(commonLetters))
}
