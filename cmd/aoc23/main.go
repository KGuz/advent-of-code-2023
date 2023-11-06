package main

import (
	"aoc23/assets"
	"aoc23/internal/puzzle"
	"flag"
	"fmt"
	"strings"
)

func main() {
	day := flag.Int("day", 1, "Advent day number [1-25]")
	flag.Parse()

	puzzle, err := puzzle.Dispatch(*day)
	if err != nil {
		fmt.Println(err)
		return
	}

	input, err := assets.LoadInput(*day)
	if err != nil {
		fmt.Println(err)
		return
	}

	pad := strings.Repeat("*", 14)
	fmt.Printf("%s Advent of Code 2021 - Day %02d %s*\n", pad, *day, pad)
	fmt.Printf("Part one: %49s\n", puzzle.PartOne(input))
	fmt.Printf("Part two: %49s\n", puzzle.PartTwo(input))
}
