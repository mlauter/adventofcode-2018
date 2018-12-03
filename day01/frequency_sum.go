package main

import(
	"os"
	"log"
	"bufio"
	"strconv"
	"fmt"
)

// Sum the elements in the scanner
func FrequencySum(s *bufio.Scanner) (int, error) {
	sum := 0
	for s.Scan() {
		if s.Text() == "" {
			continue;
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

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Usage: frequency_sum INFILE")
	}

	path := args[0]
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Cannot open input file %s", path)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum, err := FrequencySum(scanner)
	if err != nil {
		log.Fatalf("Unable to sum inputs: %s", err)
	}

	fmt.Printf("%d", sum)
}
