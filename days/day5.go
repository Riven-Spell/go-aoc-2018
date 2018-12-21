package days

import (
	"math"
	"strconv"
	"strings"
)

func react(input string) int {
	for {
		found := false

		last := ""
		for _,v := range strings.Split(input, "") {
			if strings.EqualFold(v, last) && v != last {
				found = true
				input = strings.Replace(input, last + v, "", -1)
			}

			last = v
		}

		if !found {
			break
		}
	}

	return len(input)
}

func Day5Part1(input string) string {
	return strconv.Itoa(react(input))
}

func Day5Part2(input string) string {
	lowestCount := math.MaxInt64
	for _,v := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") { // this is slow as fuck. There's gotta be a better way. This is effectively 26x50k characters.
		r := strings.NewReplacer(v, "", strings.ToUpper(v), "")
		count := react(r.Replace(input))
		if count < lowestCount {
			lowestCount = count
		}
	}

	return strconv.Itoa(lowestCount)
}