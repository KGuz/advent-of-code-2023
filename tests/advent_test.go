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

func TestAllPuzzles(t *testing.T) {
	var tests = []struct {
		day  int
		part int
		want string
	}{
		{1, 1, "24000"},
		{1, 2, "45000"},
	}

	for _, tt := range tests {
		t.Run(name(tt.day, tt.part), func(t *testing.T) {
			if err := test(tt.day, tt.part, tt.want); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestSinglePuzzle(t *testing.T) {
	day, part, want := 1, 1, "24000"

	if err := test(day, part, want); err != nil {
		t.Error(err)
	}
}
