package puzzle_test

import (
	"aoc23/assets"
	"aoc23/internal/puzzle"
	"fmt"
	"testing"
)

func solve(day int, part int) string {
	p, _ := puzzle.Dispatch(day)
	input, _ := assets.LoadExample(day)

	if part == 1 {
		return p.PartOne(input)
	}
	return p.PartTwo(input)
}

func TestPuzzles(t *testing.T) {
	var tests = []struct {
		day    int
		part   int
		answer string
	}{
		{1, 1, "24000"},
		{1, 2, "45000"},
	}

	for _, tt := range tests {
		test_name := fmt.Sprintf("Day%02d-Part%d", tt.day, tt.part)

		t.Run(test_name, func(t *testing.T) {
			result := solve(tt.day, tt.part)
			if result != tt.answer {
				t.Errorf("got %s, expected %s", result, tt.answer)
			}
		})
	}
}
