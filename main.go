package main

import (
	"fmt"
	"go-aoc-2018/days"
	"time"
)

var parts = [][]func() string{
	{days.Day1Part1, days.Day1Part2},
	{days.Day2Part1, days.Day2Part2},
	{days.Day3Part1, days.Day3Part2},
}

func main() {
	fmt.Println("--------------------------------------------------------------------")
	for day, part := range parts {
		for n, f := range part {
			fmt.Println("Running day", day+1, "part", n+1)
			start := time.Now()
			fmt.Println(f())
			fmt.Printf("Day %d part %d took %s to execute.\n", day+1, n+1, time.Since(start))
			fmt.Println("--------------------------------------------------------------------")
		}
	}
}
