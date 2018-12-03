package main

import (
	"testing"
	"bufio"
	"strings"
)

func TestOne(t *testing.T) {
	tests := []struct {
		Input string
		Expect int
	}{
		{`
0
+1
`, 1,
		},
		{`
+1
+1
+1
`, 3,
		},
		{`
+1
+1
-2
`, 0,
		},
		{`
-1
-2
-3
`, -6,
		},
	}

	for _, test := range tests {
		scanner := bufio.NewScanner(strings.NewReader(test.Input))
		actual, err := FrequencySum(scanner)
		if err != nil {
			t.Errorf("Got error %s", err)
		}
		if actual != test.Expect {
			t.Errorf("expected %d, got %d", test.Expect, actual)
		}
	}
}
