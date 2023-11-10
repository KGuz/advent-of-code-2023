package puzzle

import (
	"aoc/pkg/advent"
	"errors"
)

type Puzzle interface {
	PartOne(string) string
	PartTwo(string) string
}

func Dispatch(day int) (Puzzle, error) {
	switch day {
	case 1:
		return advent.Day01{}, nil
	case 2:
		return advent.Day02{}, nil
	case 3:
		return advent.Day03{}, nil
	case 4:
		return advent.Day04{}, nil
	case 5:
		return advent.Day05{}, nil
	case 6:
		return advent.Day06{}, nil
	case 7:
		return advent.Day07{}, nil
	case 8:
		return advent.Day08{}, nil
	case 9:
		return advent.Day09{}, nil
	case 10:
		return advent.Day10{}, nil
	case 11:
		return advent.Day11{}, nil
	case 12:
		return advent.Day12{}, nil
	case 13:
		return advent.Day13{}, nil
	case 14:
		return advent.Day14{}, nil
	case 15:
		return advent.Day15{}, nil
	case 16:
		return advent.Day16{}, nil
	case 17:
		return advent.Day17{}, nil
	case 18:
		return advent.Day18{}, nil
	case 19:
		return advent.Day19{}, nil
	case 20:
		return advent.Day20{}, nil
	case 21:
		return advent.Day21{}, nil
	case 22:
		return advent.Day22{}, nil
	case 23:
		return advent.Day23{}, nil
	case 24:
		return advent.Day24{}, nil
	case 25:
		return advent.Day25{}, nil
	default:
		return nil, errors.New("are you stupid or something?")
	}
}
