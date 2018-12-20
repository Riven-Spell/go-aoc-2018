package days

import (
	"adventofcode/inputs"
	"strconv"
	"strings"
)

func Day1Part1() string {
	input := inputs.Input1
	var freq int64
	for _, v := range strings.Split(input, "\n") {
		change, _ := strconv.ParseInt(v, 10, 64)
		freq += change
	}

	return strconv.FormatInt(freq, 10)
}

type LinkedInt64 struct {
	next *LinkedInt64
	value int64
}

func Day1Part2() string {
	input := inputs.Input1
	var freq int64
	var seenfreqs = make(map[int64]bool)
	var first *LinkedInt64

	{
		var last *LinkedInt64

		for _, v := range strings.Split(input, "\n") {
			change, _ := strconv.ParseInt(v, 10, 64)

			if last == nil {
				last = &LinkedInt64{
					value: change,
				}
				first = last
			} else {
				last.next = &LinkedInt64{
					value: change,
				}
				last = last.next
			}
		}

		last.next = first
	}

	{
		var last = first

		for {
			if _, ok := seenfreqs[freq]; ok {
				break
			}
			seenfreqs[freq] = true

			freq += last.value
			last = last.next
		}
	}

	return strconv.FormatInt(freq, 10)
}