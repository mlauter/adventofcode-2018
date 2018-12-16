package main

import (
	"reflect"
	"testing"
	"time"
)

func TestDay04NewGuardRecord(t *testing.T) {
	tests := []struct {
		Input  string
		Expect guardRecord
	}{
		{
			"[1518-07-08 00:30] wakes up",
			guardRecord{
				time.Date(1518, 7, 8, 0, 30, 0, 0, time.UTC),
				0,
				actionWakesUp,
			},
		},
		{
			"[1518-03-27 00:16] falls asleep",
			guardRecord{
				time.Date(1518, 3, 27, 0, 16, 0, 0, time.UTC),
				0,
				actionFallsAsleep,
			},
		},
		{
			"[1518-10-18 23:57] Guard #2399 begins shift",
			guardRecord{
				time.Date(1518, 10, 18, 23, 57, 0, 0, time.UTC),
				2399,
				actionBeginsShift,
			},
		},
		{
			"[1518-04-23 00:00] Guard #2857 begins shift",
			guardRecord{
				time.Date(1518, 4, 23, 0, 0, 0, 0, time.UTC),
				2857,
				actionBeginsShift,
			},
		},
	}

	for _, tc := range tests {
		gr, err := newGuardRecord(tc.Input)
		if err != nil {
			t.Errorf("day04 newGuardRecord encountered unexpected exception: %v", err)
		}

		if !reflect.DeepEqual(gr, tc.Expect) {
			t.Errorf("day04 newGuardRecord failed. Expected %v got %v", tc.Expect, gr)
		}
	}
}
