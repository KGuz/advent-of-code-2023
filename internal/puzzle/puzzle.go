package puzzle

import (
	"aoc23/pkg/days"
	"errors"
)

type Puzzle interface {
	PartOne(string) string
	PartTwo(string) string
}

func Dispatch(day int) (Puzzle, error) {
	switch day {
	case 1:
		return days.Day01{}, nil
	case 2:
		return days.Day02{}, nil
	case 3:
		return days.Day03{}, nil
	case 4:
		return days.Day04{}, nil
	case 5:
		return days.Day05{}, nil
	case 6:
		return days.Day06{}, nil
	case 7:
		return days.Day07{}, nil
	case 8:
		return days.Day08{}, nil
	case 9:
		return days.Day09{}, nil
	case 10:
		return days.Day10{}, nil
	case 11:
		return days.Day11{}, nil
	case 12:
		return days.Day12{}, nil
	case 13:
		return days.Day13{}, nil
	case 14:
		return days.Day14{}, nil
	case 15:
		return days.Day15{}, nil
	case 16:
		return days.Day16{}, nil
	case 17:
		return days.Day17{}, nil
	case 18:
		return days.Day18{}, nil
	case 19:
		return days.Day19{}, nil
	case 20:
		return days.Day20{}, nil
	case 21:
		return days.Day21{}, nil
	case 22:
		return days.Day22{}, nil
	case 23:
		return days.Day23{}, nil
	case 24:
		return days.Day24{}, nil
	case 25:
		return days.Day25{}, nil
	default:
		return nil, errors.New("are you stupid or something?")
	}
}
