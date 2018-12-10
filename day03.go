package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"sync"
)

const claimPattern = `#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`

var claimRe = regexp.MustCompile(claimPattern)

type claim struct {
	id     int
	xOff   int
	yOff   int
	width  int
	height int
}

// Return all (i, j) coordinates of square inches within this claim
func (c *claim) getCoordinates() []tuple {
	var coords []tuple
	for i := c.xOff; i < (c.xOff + c.width); i++ {
		for j := c.yOff; j < (c.yOff + c.height); j++ {
			coords = append(coords, tuple{i, j})
		}
	}

	return coords
}

// is this a bad idea?
// maybe a hash would be better
type tuple struct {
	i, j int
}

// dictionary of keys sparse matrix
// of potentially infinite size
// there is no out of bounds in this matrix,
// the matrix simply expands to fit a key set outside its bounds
type fabric struct {
	sync.Mutex         // super overkill since i'm not actually parallelizing this work but maybe i'll try it later
	rows               int
	cols               int
	elements           map[tuple]int
	overlappingSquares int
}

func (f *fabric) get(t tuple) int {
	if v, ok := f.elements[t]; ok {
		return v
	}
	return 0
}

func (f *fabric) unsafeSet(t tuple, v int) {
	f.elements[t] = v
	if t.i >= f.rows {
		f.rows = t.i + 1
	}
	if t.j >= f.cols {
		f.cols = t.j + 1
	}

}

func (f *fabric) increment(t tuple) int {
	f.Lock()
	v := f.get(t)
	v = v + 1
	f.unsafeSet(t, v)
	f.Unlock()

	return v
}

func (f *fabric) processClaim(c *claim) {
	for _, coord := range c.getCoordinates() {
		v := f.increment(coord)

		if v == 2 {
			f.overlappingSquares++
		}
	}
}

func (f *fabric) print() string {
	fabricVis := make([]rune, (f.rows+1)*f.cols) // add 1 to row length for newline
	for j := 0; j < f.cols; j++ {
		for i := 0; i < f.rows; i++ {
			var r rune
			v := f.get(tuple{i, j})
			switch v {
			case 0:
				r = '.'
			case 1:
				r = 'O'
			default:
				r = 'X'
			}

			fabricVis = append(fabricVis, r)
		}
		fabricVis = append(fabricVis, '\n')
	}
	return string(fabricVis)
}

func newFabric(rows, cols int) *fabric {
	if rows < 0 {
		panic("fabric cannot have negative row count")
	}
	if cols < 0 {
		panic("fabric cannot have negative col count")
	}

	return &fabric{
		rows:               rows,
		cols:               cols,
		elements:           make(map[tuple]int),
		overlappingSquares: 0,
	}
}

func newClaim(s string) (*claim, error) {
	// regex is probably pretty slow but let's see
	matches := claimRe.FindStringSubmatch(s)
	if len(matches) != 6 {
		return nil, fmt.Errorf("day03 newClaim encountered unparseable claim %s", s)
	}

	// gross
	dims := make([]int, len(matches)-1)
	for i, m := range matches[1:] {
		dim, err := strconv.Atoi(m)
		if err != nil {
			return nil, fmt.Errorf("day03 newClaim encountered unparseable claim %s", s)
		}
		dims[i] = dim
	}

	return &claim{
		id:     dims[0],
		xOff:   dims[1],
		yOff:   dims[2],
		width:  dims[3],
		height: dims[4],
	}, nil
}

func constructFabric(in io.ReadSeeker, r, c int) *fabric {
	fp := newFabric(r, c)
	s := bufio.NewScanner(in)

	for s.Scan() {
		cp, err := newClaim(s.Text())
		if err != nil {
			log.Fatalf("day03 constructFabric: Unable to parse claims: %v", err)
		}

		fp.processClaim(cp)
	}

	if s.Err() != nil {
		log.Fatalf("day03 constructFabric: Unable to parse claims: %v", s.Err())
	}

	return fp
}

// having some regrets about not just reading the whole input file
// into an array. it's not that big
func runDay03(f io.ReadSeeker) {
	// "The whole piece of fabric they're working on is a very large square -
	// at least 1000 inches on each side."
	fabric := constructFabric(f, 1000, 1000)

	fmt.Printf("How many square inches of fabric are within two or more claims? %d\n", fabric.overlappingSquares)
}
