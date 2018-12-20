package days

import (
	"fmt"
	"go-aoc-2018/utils"
	"sort"
	"strconv"
	"strings"
	"time"
)

type eventType uint8
const(
	WAKE eventType = iota
	SLEEP
	SHIFT
)

type timeEvent struct {
	unixTime int64 //This is not nanoseconds, just seconds.
	minute int
	eType eventType
	guardNum int
}

type teList []timeEvent

func (p teList) Len() int           { return len(p) }
func (p teList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p teList) Less(i, j int) bool { return p[i].unixTime < p[j].unixTime }

func Day4Part1(input string) string {
	events := make(teList, 0)
	loc, _ := time.LoadLocation("UTC")

	for _,v := range strings.Split(input, "\n") {
		month := 0
		day := 0
		hour := 0
		minute := 0
		eType := WAKE
		guard := 0
		_, _ = fmt.Sscanf(v, "[1518-%d-%d %d:%d]", &month, &day, &hour, &minute)
		//Ignoring it so gofmt doesn't scream at me anymore. We know our inputs are safe from advent of code.
		if strings.Contains(v, "wakes") {
			eType = WAKE
		} else if strings.Contains(v, "falls") {
			eType = SLEEP
		} else {
			eType = SHIFT
			_, _ = fmt.Sscanf(v[19:], "Guard #%d begins shift", &guard)
		}

		events = append(events, timeEvent{time.Date(2000, time.Month(month), day, hour, minute, 0, 0, loc).Unix(), minute, eType, guard })
	}

	sort.Sort(events)

	guardTotalSleep := make(map[int]int)
	guardSleep := make(map[int][60]int)
	currentGuard := 0
	sleeping := false
	sleepMinute := 0
	for _,v := range events {

		switch v.eType {
		case WAKE:
			if sleeping { //Sanity checks.
				sleeping = false
				guardTotalSleep[currentGuard] += v.minute - sleepMinute
				min := guardSleep[currentGuard]
				for _,v := range utils.Range(sleepMinute, v.minute) {
					min[v]++
				}
				guardSleep[currentGuard] = min
			}
		case SLEEP:
			if !sleeping {
				sleeping = true
				sleepMinute = v.minute
			}
		case SHIFT:
			currentGuard = v.guardNum
			sleeping = false //Sanity check.
		}
	}

	mostSleep := 0
	msGuardID := 0
	for k,v := range guardTotalSleep {
		if v >= mostSleep {
			mostSleep = v
			msGuardID = k
		}
	}

	msMinute := 0
	mostSleep = 0
	for k,v := range guardSleep[msGuardID] {
		if v >= mostSleep {
			mostSleep = v
			msMinute = k
		}
	}

	return strconv.Itoa(msMinute * msGuardID)
}

func Day4Part2(input string) string {
	events := make(teList, 0)
	loc, _ := time.LoadLocation("UTC")

	for _,v := range strings.Split(input, "\n") {
		month := 0
		day := 0
		hour := 0
		minute := 0
		eType := WAKE
		guard := 0
		_, _ = fmt.Sscanf(v, "[1518-%d-%d %d:%d]", &month, &day, &hour, &minute)
		//Ignoring it so gofmt doesn't scream at me anymore. We know our inputs are safe from advent of code.
		if strings.Contains(v, "wakes") {
			eType = WAKE
		} else if strings.Contains(v, "falls") {
			eType = SLEEP
		} else {
			eType = SHIFT
			_, _ = fmt.Sscanf(v[19:], "Guard #%d begins shift", &guard)
		}

		events = append(events, timeEvent{time.Date(2000, time.Month(month), day, hour, minute, 0, 0, loc).Unix(), minute, eType, guard })
	}

	sort.Sort(events)

	guardSleep := make(map[int][60]int)
	currentGuard := 0
	sleeping := false
	sleepMinute := 0
	for _,v := range events {

		switch v.eType {
		case WAKE:
			if sleeping { //Sanity checks.
				sleeping = false
				min := guardSleep[currentGuard]
				for _,v := range utils.Range(sleepMinute, v.minute) {
					min[v]++
				}
				guardSleep[currentGuard] = min
			}
		case SLEEP:
			if !sleeping {
				sleeping = true
				sleepMinute = v.minute
			}
		case SHIFT:
			currentGuard = v.guardNum
			sleeping = false //Sanity check.
		}
	}

	outMinute := 0
	omSleep := 0
	omGuard := 0
	for k,v := range guardSleep {
		msMinute := 0
		mostSleep := 0
		for k,v := range v {
			if v >= mostSleep {
				mostSleep = v
				msMinute = k
			}
		}

		if mostSleep > omSleep {
			outMinute = msMinute
			omSleep = mostSleep
			omGuard = k
		}
	}

	return strconv.Itoa(outMinute * omGuard)
}