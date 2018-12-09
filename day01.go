package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
)

func frequencySum(f io.ReadSeeker) (int, error) {
	s := bufio.NewScanner(f)
	sum := 0
	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		freq, err := strconv.Atoi(s.Text())
		if err != nil {
			return 0, err
		}
		sum += freq
	}

	if s.Err() != nil {
		return 0, s.Err()
	}
	return sum, nil
}

// Returns:
// int - the first frequency to be reached twice
// err - an error if there is one
func calibrate(curFreq int, freqMap map[int]bool, f io.ReadSeeker) (int, error) {
	_, err := f.Seek(0, 0)
	if err != nil {
		return 0, err
	}

	s := bufio.NewScanner(f)
	calibrated := false

	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		freq, err := strconv.Atoi(s.Text())
		if err != nil {
			return 0, err
		}
		curFreq += freq

		if freqMap[curFreq] {
			calibrated = true
			break
		}

		freqMap[curFreq] = true
	}

	if s.Err() != nil {
		return 0, s.Err()
	}

	if !calibrated {
		return calibrate(curFreq, freqMap, f)
	}
	return curFreq, nil
}

func runDay01(f io.ReadSeeker) {
	sum, err := frequencySum(f)
	if err != nil {
		log.Fatalf("day01 frequencySum: Unable to sum inputs: %s", err)
	}

	freqMap := make(map[int]bool)
	freqMap[0] = true
	freq, err := calibrate(0, freqMap, f)
	if err != nil {
		log.Fatalf("day01 calibrate: Unable to calibrate: %s", err)
	}

	fmt.Printf("Q1: %d\n", sum)
	fmt.Printf("Q2: %d\n", freq)
}
