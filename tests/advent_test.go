package advent_test

import (
	"aoc/assets"
	"aoc/internal/puzzle"
	"fmt"
	"testing"
)

type Params struct {
	day  int
	part int
	want string
}

func TestPuzzlesWithExamples(t *testing.T) {
	test(t, true, []Params{
		{1, 1, "231"},
		{1, 2, "281"},
		{2, 1, "8"},
		{2, 2, "2286"},
		{3, 1, "4361"},
		{3, 2, "467835"},
		{4, 1, "13"},
		{4, 2, "30"},
		{5, 1, "35"},
		{5, 2, "46"},
		{6, 1, "288"},
		{6, 2, "71503"},
		{7, 1, "6440"},
		{7, 2, "5905"},
		{8, 1, "2"},
		{8, 2, "6"},
		{9, 1, "114"},
		{9, 2, "2"},
		{10, 1, "80"},
		{10, 2, "10"},
		{11, 1, "374"},
		{11, 2, "82000210"},
		{12, 1, "21"},
		{12, 2, "525152"},
		{13, 1, "405"},
		{13, 2, "400"},
		{14, 1, "136"},
		{14, 2, "64"},
		{15, 1, "1320"},
		{15, 2, "145"},
		{16, 1, "46"},
		{16, 2, "51"},
		{17, 1, "102"},
		{17, 2, "94"},
		{18, 1, "62"},
		{18, 2, "952408144115"},
		{19, 1, "19114"},
		{19, 2, "167409079868000"},
	})
}

func TestPuzzlesWithInput(t *testing.T) {
	test(t, false, []Params{
		{1, 1, "54081"},
		{1, 2, "54649"},
		{2, 1, "2156"},
		{2, 2, "66909"},
		{3, 1, "528819"},
		{3, 2, "80403602"},
		{4, 1, "28750"},
		{4, 2, "10212704"},
		{5, 1, "175622908"},
		{5, 2, "5200543"},
		{6, 1, "588588"},
		{6, 2, "34655848"},
		{7, 1, "250602641"},
		{7, 2, "251037509"},
		{8, 1, "11309"},
		{8, 2, "13740108158591"},
		{9, 1, "1853145119"},
		{9, 2, "923"},
		{10, 1, "6823"},
		{10, 2, "415"},
		{11, 1, "9686930"},
		{11, 2, "630728425490"},
		{12, 1, "6935"},
		{12, 2, "3920437278260"},
		{13, 1, "34100"},
		{13, 2, "33106"},
		{14, 1, "112773"},
		{14, 2, "98894"},
		{15, 1, "503487"},
		{15, 2, "261505"},
		{16, 1, "7067"},
		{16, 2, "7324"},
		{17, 1, "859"},
		{17, 2, "1027"},
		{18, 1, "50465"},
		{18, 2, "82712746433310"},
		{19, 1, "346230"},
		{19, 2, "124693661917133"},
	})
}

func test(t *testing.T, example bool, params []Params) {
	for _, tt := range params {
		t.Run(name(tt.day, tt.part), func(t *testing.T) {
			input, _ := assets.Load(tt.day, example)
			if got := solve(tt.day, tt.part, input); got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}

func name(day int, part int) string {
	if part == 1 {
		return fmt.Sprintf("Day%02dPartOne", day)
	} else {
		return fmt.Sprintf("Day%02dPartTwo", day)
	}
}

func solve(day int, part int, input string) string {
	puzzle, _ := puzzle.Dispatch(day)
	if part == 1 {
		return puzzle.PartOne(input)
	} else {
		return puzzle.PartTwo(input)
	}
}
