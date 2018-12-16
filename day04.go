package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"sort"
	"strconv"
	"time"
)

// Layouts must use the reference time Mon Jan 2 15:04:05 MST 2006
// to show the pattern with which to format/parse a given time/string.
const timeFormat = "2006-01-02 15:04"

const (
	actionBeginsShift actionID = iota // 0
	actionFallsAsleep                 // 1
	actionWakesUp                     // 2
)

type actionID int

var actionRegexes = map[actionID](*regexp.Regexp){
	actionBeginsShift: regexp.MustCompile(`\[(?P<timestamp>.*)\] Guard #(?P<guardid>\d+) begins shift`),
	actionFallsAsleep: regexp.MustCompile(`\[(?P<timestamp>.*)\] falls asleep`),
	actionWakesUp:     regexp.MustCompile(`\[(?P<timestamp>.*)\] wakes up`),
}

// Don't really feel like implementing a sorting algo myself rn,
// maybe i'll come back to it
// byTimestamp implements sort.Interface for []GuardRecord based on
// the timestamp field.
type byTimestamp []guardRecord

func (t byTimestamp) Len() int           { return len(t) }
func (t byTimestamp) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t byTimestamp) Less(i, j int) bool { return t[i].timestamp.Before(t[j].timestamp) }

type guardRecord struct {
	timestamp time.Time
	guardID   int
	actionID  actionID
}

func newGuardRecord(s string) (guardRecord, error) {
	r := guardRecord{}

	for i, re := range actionRegexes {
		if matches := re.FindStringSubmatch(s); matches != nil {
			result := make(map[string]string)
			for j, name := range re.SubexpNames()[1:] {
				result[name] = matches[j+1]
			}
			timestamp, ok := result["timestamp"]
			if !ok {
				return r, fmt.Errorf("day04 newGuardRecord no timestamp in input string %s", s)
			}
			t, err := time.Parse(timeFormat, timestamp)
			if err != nil {
				return r, fmt.Errorf("day04 newGuardRecord could not parse timestamp %s", timestamp)
			}
			r.timestamp = t

			if guardID, ok := result["guardid"]; ok {
				gid, err := strconv.Atoi(guardID)
				if err != nil {
					return r, fmt.Errorf("day04 new GuardRecord guard id is not an int: %s", guardID)
				}
				r.guardID = gid
			}
			r.actionID = i

			break
		}
	}

	if (guardRecord{}) == r {
		return r, fmt.Errorf("Unable to parse guard record %s", s)
	}

	return r, nil
}

func getSortedGuardRecords(f io.ReadSeeker) []guardRecord {
	records := []guardRecord{}
	s := bufio.NewScanner(f)

	for s.Scan() {
		r, err := newGuardRecord(s.Text())
		if err != nil {
			log.Fatalf("day04 getSortedGuardRecords encountered unexpected exception %v", err)
		}
		records = append(records, r)
	}

	if s.Err() != nil {
		log.Fatalf("day04 getSortedGuardRecords: Unable to parse input file: %v", s.Err())
	}

	sort.Sort(byTimestamp(records))

	return records
}

// this is terrible
func recordSleepPeriod(start time.Time, end time.Time, id int, minsAsleep map[int]int, sleepiestMins map[int][]int) int {
	for i := start; i.Before(end); i = i.Add(time.Minute) {
		minsAsleep[id]++
		if _, ok := sleepiestMins[id]; !ok {
			sleepiestMins[id] = make([]int, 60)
		}
		sleepiestMins[id][i.Minute()]++
	}
	totalMinsAsleep, ok := minsAsleep[id]
	if !ok {
		return 0
	}

	return totalMinsAsleep
}

func runDay04(f io.ReadSeeker) {
	records := getSortedGuardRecords(f)
	mostMinsAsleep := 0
	sleepiestGuard := 0
	currentGuard := 0
	fellAsleepTime := time.Time{}

	guardMinsAsleep := map[int]int{}
	guardSleepiestMins := map[int][]int{}

	for _, record := range records {
		switch record.actionID {
		case actionBeginsShift:
			// Not sure if this is possible in the data
			// but in case its possible to sleep through the shift change
			if !fellAsleepTime.Equal(time.Time{}) {
				fmt.Println("Slept through a shift change")
				totalMinsAsleep := recordSleepPeriod(fellAsleepTime, record.timestamp, currentGuard, guardMinsAsleep, guardSleepiestMins)
				if totalMinsAsleep > mostMinsAsleep {
					mostMinsAsleep = totalMinsAsleep
					sleepiestGuard = currentGuard
				}
			}
			currentGuard = record.guardID
			fellAsleepTime = time.Time{}
		case actionFallsAsleep:
			fellAsleepTime = record.timestamp
		case actionWakesUp:
			totalMinsAsleep := recordSleepPeriod(fellAsleepTime, record.timestamp, currentGuard, guardMinsAsleep, guardSleepiestMins)
			if totalMinsAsleep > mostMinsAsleep {
				mostMinsAsleep = totalMinsAsleep
				sleepiestGuard = currentGuard
			}
			fellAsleepTime = time.Time{}
		}
	}

	sleepiestMin := 0
	mostMinsAsleep = 0
	for k, v := range guardSleepiestMins[sleepiestGuard] {
		if v > mostMinsAsleep {
			mostMinsAsleep = v
			sleepiestMin = k
		}
	}

	fmt.Printf("Guard #%d, sleepiest min %d -- %d", sleepiestGuard, sleepiestMin, sleepiestGuard*sleepiestMin)
}
