package main

import (
	"fmt"
	"go-aoc-2018/days"
	"go-aoc-2018/inputs"
	"go-aoc-2018/testin"
	"os"
	"time"
)

var parts = [][]func(string) string{
	{days.Day1Part1, days.Day1Part2},
	{days.Day2Part1, days.Day2Part2},
	{days.Day3Part1, days.Day3Part2},
	{days.Day4Part1, days.Day4Part2},
}

var dayinput = []string { //Input by day.
	inputs.Input1,
	inputs.Input2,
	inputs.Input3,
	inputs.Input4,
}

var tests = [][][]string {
	{testin.Day1Part1Tests, testin.Day1Part2Tests},
	{testin.Day2Part1Tests, testin.Day2Part2Tests},
	{testin.Day3Part1Tests, testin.Day3Part1Tests},
	{testin.Day4Part1Tests, testin.Day4Part1Tests},
}

var expected = [][][]string { //This should match the structure of tests
	{testin.Day1Part1Expects, testin.Day1Part2Expects},
	{testin.Day2Part1Expects, testin.Day2Part2Expects},
	{testin.Day3Part1Expects, testin.Day3Part2Expects},
	{testin.Day4Part1Expects, testin.Day4Part2Expects},
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "tests" {
		if len(os.Args) > 2 {
			day := 0
			part := 0

			_, _ = fmt.Sscanf(os.Args[2], "%d_%d", &day, &part)
			fmt.Println("--------------------------------------------------------------------")
			fmt.Println("Testing day", day, "part", part)
			for k, in := range tests[day-1][part-1] {
				fmt.Println("Running test", k + 1)
				start := time.Now()
				out := parts[day-1][part-1](in)
				expect := expected[day-1][part-1][k]
				if out == expect {
					fmt.Printf("Succeded test %d. Result: %s, Ran for: %s\n", k+1, out, time.Since(start))
				} else  {
					fmt.Printf("FAILED test %d. Result: %s, Expected: %s, Ran for %s\n", k+1, out, expect, time.Since(start))
				}
			}
			fmt.Println("--------------------------------------------------------------------")
		} else {
			fmt.Println("--------------------------------------------------------------------")
			for day, part := range parts {
				for n, f := range part {
					fmt.Println("Testing day", day + 1, "part", n+1)
					for k, in := range tests[day][n] {
						fmt.Println("Running test", k + 1)
						start := time.Now()
						out := f(in)
						expect := expected[day][n][k]
						if out == expect {
							fmt.Printf("Succeded test %d. Result: %s, Ran for: %s\n", k+1, out, time.Since(start))
						} else  {
							fmt.Printf("FAILED test %d. Result: %s, Expected: %s, Ran for %s\n", k+1, out, expect, time.Since(start))
						}
					}
					fmt.Println("--------------------------------------------------------------------")
				}
			}
		}
	} else if len(os.Args) > 2 && os.Args[2] == "run" {
		day := 0
		part := 0

		_, _ = fmt.Sscanf(os.Args[2], "%d_%d", &day, &part)

		fmt.Println("--------------------------------------------------------------------")
		fmt.Println("Running day", day, "part", part)
		start := time.Now()
		parts[day-1][part-1](dayinput[day-1])
		fmt.Printf("Day %d part %d took %s to execute.\n", day, part, time.Since(start))
		fmt.Println("--------------------------------------------------------------------")
	} else {
		fmt.Println("--------------------------------------------------------------------")
		for day, part := range parts {
			for n, f := range part {
				fmt.Println("Running day", day+1, "part", n+1)
				start := time.Now()
				fmt.Println(f(dayinput[day]))
				fmt.Printf("Day %d part %d took %s to execute.\n", day+1, n+1, time.Since(start))
				fmt.Println("--------------------------------------------------------------------")
			}
		}
	}
}
