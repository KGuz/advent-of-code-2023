package advent

import (
	"math"
	"strconv"
)

type Day21 struct {
	/* --- Day 21: Step Counter ---
	You manage to catch the airship right as it's dropping someone else off on
	their all-expenses-paid trip to Desert Island! It even helpfully drops you
	off near the gardener and his massive farm.

	"You got the sand flowing again! Great work! Now we just need to wait until
	we have enough sand to filter the water for Snow Island and we'll have snow
	again in no time."

	While you wait, one of the Elves that works with the gardener heard how
	good you are at solving problems and would like your help. He needs to get
	his steps in for the day, and so he'd like to know which garden plots he
	can reach with exactly his remaining 64 steps.

	He gives you an up-to-date map (your puzzle input) of his starting position
	(S), garden plots (.), and rocks (#). For example:

	...........
	.....###.#.
	.###.##..#.
	..#.#...#..
	....#.#....
	.##..S####.
	.##..#...#.
	.......##..
	.##.#.####.
	.##..##.##.
	...........

	The Elf starts at the starting position (S) which also counts as a garden
	plot. Then, he can take one step north, south, east, or west, but only onto
	tiles that are garden plots. This would allow him to reach any of the tiles
	marked O:

	...........
	.....###.#.
	.###.##..#.
	..#.#...#..
	....#O#....
	.##.OS####.
	.##..#...#.
	.......##..
	.##.#.####.
	.##..##.##.
	...........

	Then, he takes a second step. Since at this point he could be at either
	tile marked O, his second step would allow him to reach any garden plot
	that is one step north, south, east, or west of any tile that he could have
	reached after the first step:

	...........
	.....###.#.
	.###.##..#.
	..#.#O..#..
	....#.#....
	.##O.O####.
	.##.O#...#.
	.......##..
	.##.#.####.
	.##..##.##.
	...........

	After two steps, he could be at any of the tiles marked O above, including
	the starting position (either by going north-then-south or by going
	west-then-east).

	A single third step leads to even more possibilities:

	...........
	.....###.#.
	.###.##..#.
	..#.#.O.#..
	...O#O#....
	.##.OS####.
	.##O.#...#.
	....O..##..
	.##.#.####.
	.##..##.##.
	...........

	He will continue like this until his steps for the day have been exhausted.
	After a total of 6 steps, he could reach any of the garden plots marked O:

	...........
	.....###.#.
	.###.##.O#.
	.O#O#O.O#..
	O.O.#.#.O..
	.##O.O####.
	.##.O#O..#.
	.O.O.O.##..
	.##.#.####.
	.##O.##.##.
	...........

	In this example, if the Elf's goal was to get exactly 6 more steps today,
	he could use them to reach any of 16 garden plots.

	However, the Elf actually needs to get 64 steps today, and the map he's
	handed you is much larger than the example map.

	Starting from the garden plot marked S on your map, how many garden plots
	could the Elf reach in exactly 64 steps? */

	/* --- Part Two ---
	The Elf seems confused by your answer until he realizes his mistake: he was
	reading from a list of his favorite numbers that are both perfect squares
	and perfect cubes, not his step counter.

	The actual number of steps he needs to get today is exactly 26501365.

	He also points out that the garden plots and rocks are set up so that the
	map repeats infinitely in every direction.

	So, if you were to look one additional map-width or map-height out from the
	edge of the example map above, you would find that it keeps repeating:

	.................................
	.....###.#......###.#......###.#.
	.###.##..#..###.##..#..###.##..#.
	..#.#...#....#.#...#....#.#...#..
	....#.#........#.#........#.#....
	.##...####..##...####..##...####.
	.##..#...#..##..#...#..##..#...#.
	.......##.........##.........##..
	.##.#.####..##.#.####..##.#.####.
	.##..##.##..##..##.##..##..##.##.
	.................................
	.................................
	.....###.#......###.#......###.#.
	.###.##..#..###.##..#..###.##..#.
	..#.#...#....#.#...#....#.#...#..
	....#.#........#.#........#.#....
	.##...####..##..S####..##...####.
	.##..#...#..##..#...#..##..#...#.
	.......##.........##.........##..
	.##.#.####..##.#.####..##.#.####.
	.##..##.##..##..##.##..##..##.##.
	.................................
	.................................
	.....###.#......###.#......###.#.
	.###.##..#..###.##..#..###.##..#.
	..#.#...#....#.#...#....#.#...#..
	....#.#........#.#........#.#....
	.##...####..##...####..##...####.
	.##..#...#..##..#...#..##..#...#.
	.......##.........##.........##..
	.##.#.####..##.#.####..##.#.####.
	.##..##.##..##..##.##..##..##.##.
	.................................

	This is just a tiny three-map-by-three-map slice of the
	inexplicably-infinite farm layout; garden plots and rocks repeat as far as
	you can see. The Elf still starts on the one middle tile marked S, though -
	every other repeated S is replaced with a normal garden plot (.).

	Here are the number of reachable garden plots in this new infinite version
	of the example map for different numbers of steps:

	- In exactly 6 steps, he can still reach 16 garden plots.
	- In exactly 10 steps, he can reach any of 50 garden plots.
	- In exactly 50 steps, he can reach 1594 garden plots.
	- In exactly 100 steps, he can reach 6536 garden plots.
	- In exactly 500 steps, he can reach 167004 garden plots.
	- In exactly 1000 steps, he can reach 668697 garden plots.
	- In exactly 5000 steps, he can reach 16733044 garden plots.

	However, the step count the Elf needs is much larger! Starting from the
	garden plot marked S on your infinite map, how many garden plots could the
	Elf reach in exactly 26501365 steps? */
}

func (d Day21) PartOne(input string) string {
	garden := elements(input)
	answer := d.walk(garden, 64)
	return strconv.Itoa(answer)
}

func (d Day21) PartTwo(input string) string {
	garden := elements(input)
	length := len(garden)
	start := length / 2

	xs := []int{start, start + length, start + 2*length}
	ys := transform(xs, func(x int) int { return d.walk(garden, x) })

	a, b, c := cramersRule(xs[0], ys[0], xs[1], ys[1], xs[2], ys[2])
	x := float64(26501365)
	y := a*x*x + b*x + c

	return strconv.Itoa(int(math.Round(y)))
}

func (d Day21) walk(garden [][]byte, steps int) int {
	bounds := point{len(garden), len(garden[0])}
	start := bounds.div(2)

	plots := map[point]bool{start: true}
	for s := 0; s < steps; s++ {
		new := map[point]bool{}
		for pos := range plots {
			for _, dir := range orthogonal() {
				next := pos.add(dir)
				alt := d.translate(next, bounds)
				if garden[alt.i][alt.j] != '#' {
					new[next] = true
				}
			}
		}
		plots = new
	}
	return len(plots)
}

func (Day21) translate(pos point, bounds point) point {
	alt := pos
	if alt.i < 0 {
		alt.i = bounds.i + alt.i%bounds.i
	}
	if alt.i > bounds.i-1 {
		alt.i = alt.i % bounds.i
	}
	if alt.j < 0 {
		alt.j = bounds.j + alt.j%bounds.j
	}
	if alt.j > bounds.j-1 {
		alt.j = alt.j % bounds.j
	}
	return alt
}

func cramersRule(x1, y1 int, x2, y2 int, x3, y3 int) (float64, float64, float64) {
	a, b, c, d := x1*x1, x1, 1, y1
	e, f, g, h := x2*x2, x2, 1, y2
	i, j, k, l := x3*x3, x3, 1, y3

	dt := float64((a * f * k) + (b * g * i) + (c * e * j) - (c * f * i) - (a * g * j) - (b * e * k))
	an := float64((d * f * k) + (b * g * l) + (c * h * j) - (c * f * l) - (d * g * j) - (b * h * k))
	bn := float64((a * h * k) + (d * g * i) + (c * e * l) - (c * h * i) - (a * g * l) - (d * e * k))
	cn := float64((a * f * l) + (b * h * i) + (d * e * j) - (d * f * i) - (a * h * j) - (b * e * l))

	return an / dt, bn / dt, cn / dt
}
