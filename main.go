// Runner for advent of code challenges
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const usage string = "Usage: adventofcode DAY INFILE"

var dayFuncMap = map[int]func(f io.ReadSeeker){
	1: runDay01,
	2: runDay02,
	3: runDay03,
	4: runDay04,
	5: runDay05,
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal(usage)
	}

	day := args[0]
	path := args[1]

	dayInt, err := strconv.Atoi(day)
	if err != nil {
		fmt.Printf("DAY must be an int. Got type %T value %s\n", day, day)
		log.Fatal(usage)
	}

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Cannot open input file %s", path)
	}
	defer f.Close()

	if dayFunc, ok := dayFuncMap[dayInt]; ok {
		dayFunc(f)
	}

}
