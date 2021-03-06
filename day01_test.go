package main

import (
	"strings"
	"testing"
)

func TestDay01Sum(t *testing.T) {
	tests := []struct {
		Input  string
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
		actual, err := frequencySum(strings.NewReader(test.Input))
		if err != nil {
			t.Errorf("Got error %s", err)
		}
		if actual != test.Expect {
			t.Errorf("expected %d, got %d", test.Expect, actual)
		}
	}
}

func TestDay01Calibrate(t *testing.T) {
	tests := []struct {
		Input  string
		Expect int
	}{
		{`
0
+1
-1
`, 0,
		},
		{`
+3
+3
+4
-2
-4
`, 10,
		},
		{`
-6
+3
+8
+5
-6
`, 5,
		},
		{`
+7
+7
-2
-7
-4
`, 14,
		},
	}

	for _, test := range tests {
		h := strings.NewReader(test.Input)
		freqMap := make(map[int]bool)
		actual, err := calibrate(0, freqMap, h)
		if err != nil {
			t.Errorf("Got error %s", err)
		}
		if actual != test.Expect {
			t.Errorf("expected %d, got %d", test.Expect, actual)
		}
	}
}
