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

func runDay02(f io.ReadSeeker) {
	checksum, err := checksum(f)
	if err != nil {
		log.Fatalf("day02 checksum: Encountered unexpected error %v", err)
	}

	fmt.Printf("%d\n", checksum)
}
