package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestDay03ClaimGetCoordinates(t *testing.T) {
	tests := []struct {
		Input  *claim
		Expect []tuple
	}{
		{
			&claim{id: 1, xOff: 3, yOff: 2, width: 5, height: 4},
			[]tuple{tuple{3, 2}, tuple{3, 3}, tuple{3, 4}, tuple{3, 5},
				tuple{4, 2}, tuple{4, 3}, tuple{4, 4}, tuple{4, 5},
				tuple{5, 2}, tuple{5, 3}, tuple{5, 4}, tuple{5, 5},
				tuple{6, 2}, tuple{6, 3}, tuple{6, 4}, tuple{6, 5},
				tuple{7, 2}, tuple{7, 3}, tuple{7, 4}, tuple{7, 5},
			},
		},
		{
			&claim{id: 2, xOff: 1, yOff: 3, width: 4, height: 4},
			[]tuple{tuple{1, 3}, tuple{1, 4}, tuple{1, 5}, tuple{1, 6},
				tuple{2, 3}, tuple{2, 4}, tuple{2, 5}, tuple{2, 6},
				tuple{3, 3}, tuple{3, 4}, tuple{3, 5}, tuple{3, 6},
				tuple{4, 3}, tuple{4, 4}, tuple{4, 5}, tuple{4, 6},
			},
		},
	}

	for _, tc := range tests {
		coords := tc.Input.getCoordinates()

		// would be nice to have a visual representation to print
		if !reflect.DeepEqual(tc.Expect, coords) {
			t.Error("day03 claim.getCoordinates failed")
		}
	}
}

func TestNewClaim(t *testing.T) {
	tests := []struct {
		Input  string
		Expect claim
	}{
		{
			"#1253 @ 683,604: 22x23",
			claim{
				id:     1253,
				xOff:   683,
				yOff:   604,
				width:  22,
				height: 23,
			},
		},
	}

	for _, tc := range tests {
		c, err := newClaim(tc.Input)
		if err != nil {
			t.Fatalf("day03 newClaim encountered unexpected error %v", err)
		}

		if !reflect.DeepEqual(*c, tc.Expect) {
			t.Error("day03 newClaim failed")
		}
	}
}

func TestConstructFabric(t *testing.T) {
	tests := []struct {
		Input  string
		Expect int
	}{
		{
			`#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`, 4,
		},
	}

	for _, tc := range tests {
		fabric := constructFabric(strings.NewReader(tc.Input), 8, 8)
		if fabric.overlappingSquares != tc.Expect {
			t.Errorf("day03 getOverlappingSquareInches expected %d got %d", tc.Expect, fabric.overlappingSquares)
		}
		fmt.Printf("%s\n", fabric.print())
	}
}
