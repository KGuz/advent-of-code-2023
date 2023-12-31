package main

import (
	"aoc/assets"
	"aoc/internal/puzzle"
	"flag"
	"fmt"
	"strings"
)

func main() {
	day := flag.Int("day", 1, "Advent day number [1-25]")
	example := flag.Bool("example", false, "Solves puzzle using example input")
	flag.Parse()

	puzzle, err := puzzle.Dispatch(*day)
	if err != nil {
		fmt.Println(err)
		return
	}

	input, err := assets.Load(*day, *example)
	if err != nil {
		fmt.Println(err)
		return
	}

	pad := strings.Repeat("*", 14)
	fmt.Printf("%s Advent of Code 2023 - Day %02d %s*\n", pad, *day, pad)
	fmt.Printf("Part one: %49s\n", puzzle.PartOne(input))
	fmt.Printf("Part two: %49s\n", puzzle.PartTwo(input))
}
