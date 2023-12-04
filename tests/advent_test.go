package advent_test

import (
	"aoc/assets"
	"aoc/internal/puzzle"
	"fmt"
	"testing"
)

func name(day int, part int) string {
	if part == 1 {
		return fmt.Sprintf("Day%02dPartOne", day)
	}
	if part == 2 {
		return fmt.Sprintf("Day%02dPartTwo", day)
	}
	panic("yeah, if you could just pick one that would be great")
}

func solve(day int, part int, input string) string {
	puzzle, err := puzzle.Dispatch(day)
	if err != nil {
		panic(err)
	}

	if part == 1 {
		return puzzle.PartOne(input)
	}
	return puzzle.PartTwo(input)
}

func test(day int, part int, want string) error {
	input, err := assets.LoadExample(day)
	if err != nil {
		panic(err)
	}

	if got := solve(day, part, input); got != want {
		return fmt.Errorf("got %s, want %s", got, want)
	}
	return nil
}

func TestPuzzles(t *testing.T) {
	var tests = []struct {
		day  int
		part int
		want string
	}{
		{1, 1, "231"},
		{1, 2, "281"},
		{2, 1, "8"},
		{2, 2, "2286"},
		{3, 1, "4361"},
		{3, 2, "467835"},
		{4, 1, "13"},
		{4, 2, "30"},
	}

	for _, tt := range tests {
		t.Run(name(tt.day, tt.part), func(t *testing.T) {
			if err := test(tt.day, tt.part, tt.want); err != nil {
				t.Error(err)
			}
		})
	}
}
