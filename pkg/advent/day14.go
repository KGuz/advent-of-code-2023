package advent

import (
	"aoc/pkg/utl"
	"bytes"
	"fmt"
	"strconv"
)

type Day14 struct {
	/* --- Day 14: Parabolic Reflector Dish ---
	You reach the place where all of the mirrors were pointing: a massive
	parabolic reflector dish attached to the side of another large mountain.

	The dish is made up of many small mirrors, but while the mirrors themselves
	are roughly in the shape of a parabolic reflector dish, each individual
	mirror seems to be pointing in slightly the wrong direction. If the dish is
	meant to focus light, all it's doing right now is sending it in a vague
	direction.

	This system must be what provides the energy for the lava! If you focus the
	reflector dish, maybe you can go where it's pointing and use the light to
	fix the lava production.

	Upon closer inspection, the individual mirrors each appear to be connected
	via an elaborate system of ropes and pulleys to a large metal platform
	below the dish. The platform is covered in large rocks of various shapes.
	Depending on their position, the weight of the rocks deforms the platform,
	and the shape of the platform controls which ropes move and ultimately the
	focus of the dish.

	In short: if you move the rocks, you can focus the dish. The platform even
	has a control panel on the side that lets you tilt it in one of four
	directions! The rounded rocks (O) will roll when the platform is tilted,
	while the cube-shaped rocks (#) will stay in place. You note the positions
	of all of the empty spaces (.) and rocks (your puzzle input). For example:

	O....#....
	O.OO#....#
	.....##...
	OO.#O....O
	.O.....O#.
	O.#..O.#.#
	..O..#O..O
	.......O..
	#....###..
	#OO..#....

	Start by tilting the lever so all of the rocks will slide north as far as
	they will go:

	OOOO.#.O..
	OO..#....#
	OO..O##..O
	O..#.OO...
	........#.
	..#....#.#
	..O..#.O.O
	..O.......
	#....###..
	#....#....

	You notice that the support beams along the north side of the platform are
	damaged; to ensure the platform doesn't collapse, you should calculate the
	total load on the north support beams.

	The amount of load caused by a single rounded rock (O) is equal to the
	number of rows from the rock to the south edge of the platform, including
	the row the rock is on. (Cube-shaped rocks (#) don't contribute to load.)
	So, the amount of load caused by each rock in each row is as follows:

	OOOO.#.O.. 10
	OO..#....#  9
	OO..O##..O  8
	O..#.OO...  7
	........#.  6
	..#....#.#  5
	..O..#.O.O  4
	..O.......  3
	#....###..  2
	#....#....  1

	The total load is the sum of the load caused by all of the rounded rocks.
	In this example, the total load is 136.

	Tilt the platform so that the rounded rocks all roll north. Afterward, what
	is the total load on the north support beams? */

	/* --- Part Two ---
	The parabolic reflector dish deforms, but not in a way that focuses the
	beam. To do that, you'll need to move the rocks to the edges of the
	platform. Fortunately, a button on the side of the control panel labeled
	"spin cycle" attempts to do just that!

	Each cycle tilts the platform four times so that the rounded rocks roll
	north, then west, then south, then east. After each tilt, the rounded rocks
	roll as far as they can before the platform tilts in the next direction.
	After one cycle, the platform will have finished rolling the rounded rocks
	in those four directions in that order.

	Here's what happens in the example above after each of the first few cycles:

	After 1 cycle:

	.....#....
	....#...O#
	...OO##...
	.OO#......
	.....OOO#.
	.O#...O#.#
	....O#....
	......OOOO
	#...O###..
	#..OO#....

	After 2 cycles:

	.....#....
	....#...O#
	.....##...
	..O#......
	.....OOO#.
	.O#...O#.#
	....O#...O
	.......OOO
	#..OO###..
	#.OOO#...O

	After 3 cycles:

	.....#....
	....#...O#
	.....##...
	..O#......
	.....OOO#.
	.O#...O#.#
	....O#...O
	.......OOO
	#...O###.O
	#.OOO#...O

	This process should work if you leave it running long enough, but you're
	still worried about the north support beams. To make sure they'll survive
	for a while, you need to calculate the total load on the north support
	beams after 1000000000 cycles.

	In the above example, after 1000000000 cycles, the total load on the north
	support beams is 64.

	Run the spin cycle for 1000000000 cycles. Afterward, what is the total load
	on the north support beams? */
}

func (d Day14) PartOne(input string) string {
	state := d.parse(input)
	d.north(state)
	return strconv.Itoa(d.load(state))
}

func (d Day14) PartTwo(input string) string {
	state := d.parse(input)
	past := make(map[string]int)

	for cycle := 1; ; cycle++ {
		d.north(state)
		d.west(state)
		d.south(state)
		d.east(state)

		if final := d.check(state, past, cycle); final != nil {
			d.show(final)
			return strconv.Itoa(d.load(final))
		}
	}
}

func (Day14) load(state [][]byte) int {
	sum := 0
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state[0]); j++ {
			if state[i][j] == 'O' {
				sum += len(state) - i
			}
		}
	}
	return sum
}

func (Day14) parse(input string) [][]byte {
	return utl.Map(utl.Lines(input), utl.ToBytes)
}

func (Day14) show(state [][]byte) {
	for y := 0; y < len(state); y++ {
		for x := 0; x < len(state[0]); x++ {
			switch state[y][x] {
			case '#':
				fmt.Print("●")
			case '.':
				fmt.Print(" ")
			case 'O':
				fmt.Print(utl.YELLOW, "●", utl.END)
			}
		}
		fmt.Print("\n")
	}
}

func (Day14) hash(state [][]byte) string {
	return string(bytes.Join(state, []byte{'\n'}))
}

func (Day14) north(state [][]byte) {
	for j := 0; j < len(state[0]); j++ {
		y := 0
		for i := 0; i < len(state); i++ {
			if state[i][j] == 'O' {
				state[i][j], state[y][j] = '.', 'O'
				y++
			} else if state[i][j] == '#' {
				y = i + 1
			}
		}
	}
}

func (Day14) west(state [][]byte) {
	for i := 0; i < len(state); i++ {
		x := 0
		for j := 0; j < len(state[0]); j++ {
			if state[i][j] == 'O' {
				state[i][j], state[i][x] = '.', 'O'
				x++
			} else if state[i][j] == '#' {
				x = j + 1
			}
		}
	}
}

func (Day14) south(state [][]byte) {
	for j := 0; j < len(state[0]); j++ {
		y := len(state) - 1
		for i := len(state) - 1; i >= 0; i-- {
			if state[i][j] == 'O' {
				state[i][j], state[y][j] = '.', 'O'
				y--
			} else if state[i][j] == '#' {
				y = i - 1
			}
		}
	}
}

func (Day14) east(state [][]byte) {
	for i := 0; i < len(state); i++ {
		x := len(state[0]) - 1
		for j := len(state[0]) - 1; j >= 0; j-- {
			if state[i][j] == 'O' {
				state[i][j], state[i][x] = '.', 'O'
				x--
			} else if state[i][j] == '#' {
				x = j - 1
			}
		}
	}
}

func (d Day14) check(state [][]byte, past map[string]int, curr int) [][]byte {
	hash := d.hash(state)
	prev, ok := past[hash]

	if !ok {
		past[hash] = curr
		return nil
	}

	period := curr - prev
	last := ((1000000000 - curr) % period) + prev

	for cycle, idx := range past {
		if idx == last {
			return d.parse(cycle)
		}
	}
	panic("unreachable")
}
