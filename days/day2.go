package days

import (
	"go-aoc-2018/inputs"
	"strconv"
	"strings"
)

func Day2Part1() string {
	input := inputs.Input2
	var twos = 0
	var threes = 0

	for _, v := range strings.Split(input, "\n") {
		var seen = make(map[rune]int)

		for _, r := range v {
			seen[r]++
		}

		var two = false
		var three = false

		for _, i := range seen {
			if i == 2 {
				two = true
			}
			if i == 3 {
				three = true
			}
		}

		if two {
			twos++
		}
		if three {
			threes++
		}
	}

	return strconv.Itoa(twos * threes)
}

func Day2Part2() string {
	input := inputs.Input2
	var ids = strings.Split(input, "\n")

	for _, id1 := range ids {
		for _, id2 := range ids {
			var nonmatch = 0
			var intersection = ""
			for i := range id1 {
				if id1[i] != id2[i] {
					nonmatch++
				} else {
					intersection += string(id1[i])
				}
			}
			if nonmatch == 1 {
				return id1
			}
		}
	}

	return ""
}
