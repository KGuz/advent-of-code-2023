package advent

import (
	"aoc/pkg/utl"
	"strconv"
)

type Day13 struct {
	/* --- Day 13: Point of Incidence ---
	With your help, the hot springs team locates an appropriate spring which
	launches you neatly and precisely up to the edge of Lava Island.

	There's just one problem: you don't see any lava.

	You do see a lot of ash and igneous rock; there are even what look like
	gray mountains scattered around. After a while, you make your way to a
	nearby cluster of mountains only to discover that the valley between them
	is completely full of large mirrors.  Most of the mirrors seem to be
	aligned in a consistent way; perhaps you should head in that direction?

	As you move through the valley of mirrors, you find that several of them
	have fallen from the large metal frames keeping them in place. The mirrors
	are extremely flat and shiny, and many of the fallen mirrors have lodged
	into the ash at strange angles. Because the terrain is all one color, it's
	hard to tell where it's safe to walk or where you're about to run into a
	mirror.

	You note down the patterns of ash (.) and rocks (#) that you see as you
	walk (your puzzle input); perhaps by carefully analyzing these patterns,
	you can figure out where the mirrors are!

	For example:

	#.##..##.
	..#.##.#.
	##......#
	##......#
	..#.##.#.
	..##..##.
	#.#.##.#.

	#...##..#
	#....#..#
	..##..###
	#####.##.
	#####.##.
	..##..###
	#....#..#

	To find the reflection in each pattern, you need to find a perfect
	reflection across either a horizontal line between two rows or across a
	vertical line between two columns.

	In the first pattern, the reflection is across a vertical line between two
	columns; arrows on each of the two columns point at the line between the
	columns:

	123456789
	    ><
	#.##..##.
	..#.##.#.
	##......#
	##......#
	..#.##.#.
	..##..##.
	#.#.##.#.
	    ><
	123456789

	In this pattern, the line of reflection is the vertical line between
	columns 5 and 6. Because the vertical line is not perfectly in the middle
	of the pattern, part of the pattern (column 1) has nowhere to reflect onto
	and can be ignored; every other column has a reflected column within the
	pattern and must match exactly: column 2 matches column 9, column 3 matches
	8, 4 matches 7, and 5 matches 6.

	The second pattern reflects across a horizontal line instead:

	1 #...##..# 1
	2 #....#..# 2
	3 ..##..### 3
	4v#####.##.v4
	5^#####.##.^5
	6 ..##..### 6
	7 #....#..# 7

	This pattern reflects across the horizontal line between rows 4 and 5. Row
	1 would reflect with a hypothetical row 8, but since that's not in the
	pattern, row 1 doesn't need to match anything. The remaining rows match:
	row 2 matches row 7, row 3 matches row 6, and row 4 matches row 5.

	To summarize your pattern notes, add up the number of columns to the left
	of each vertical line of reflection; to that, also add 100 multiplied by
	the number of rows above each horizontal line of reflection. In the above
	example, the first pattern's vertical line has 5 columns to its left and
	the second pattern's horizontal line has 4 rows above it, a total of 405.

	Find the line of reflection in each of the patterns in your notes. What
	number do you get after summarizing all of your notes? */

	/* --- Part Two ---
	You resume walking through the valley of mirrors and - SMACK! - run
	directly into one. Hopefully nobody was watching, because that must have
	been pretty embarrassing.

	Upon closer inspection, you discover that every mirror has exactly one
	smudge: exactly one . or # should be the opposite type.

	In each pattern, you'll need to locate and fix the smudge that causes a
	different reflection line to be valid. (The old reflection line won't
	necessarily continue being valid after the smudge is fixed.)

	Here's the above example again:

	#.##..##.
	..#.##.#.
	##......#
	##......#
	..#.##.#.
	..##..##.
	#.#.##.#.

	#...##..#
	#....#..#
	..##..###
	#####.##.
	#####.##.
	..##..###
	#....#..#

	The first pattern's smudge is in the top-left corner. If the top-left #
	were instead ., it would have a different, horizontal line of reflection:

	1 ..##..##. 1
	2 ..#.##.#. 2
	3v##......#v3
	4^##......#^4
	5 ..#.##.#. 5
	6 ..##..##. 6
	7 #.#.##.#. 7

	With the smudge in the top-left corner repaired, a new horizontal line of
	reflection between rows 3 and 4 now exists. Row 7 has no corresponding
	reflected row and can be ignored, but every other row matches exactly: row
	1 matches row 6, row 2 matches row 5, and row 3 matches row 4.

	In the second pattern, the smudge can be fixed by changing the fifth symbol
	on row 2 from . to #:

	1v#...##..#v1
	2^#...##..#^2
	3 ..##..### 3
	4 #####.##. 4
	5 #####.##. 5
	6 ..##..### 6
	7 #....#..# 7

	Now, the pattern has a different horizontal line of reflection between rows
	1 and 2.

	Summarize your notes as before, but instead use the new different
	reflection lines. In this example, the first pattern's new horizontal line
	has 3 rows above it and the second pattern's new horizontal line has 1 row
	above it, summarizing to the value 400.

	In each pattern, fix the smudge and find the different line of reflection.
	What number do you get after summarizing the new reflection line in each
	pattern in your notes? */
}

func (d Day13) PartOne(input string) string {
	sum := 0
	maps, bounds := d.parse(input)
	for n, mirrors := range maps {
		nrows, ncols := d.pointOfIncidence(mirrors, bounds[n], false)
		sum += nrows*100 + ncols
	}
	return strconv.Itoa(sum)
}

func (d Day13) PartTwo(input string) string {
	sum := 0
	maps, bounds := d.parse(input)
	for n, mirrors := range maps {
		nrows, ncols := d.pointOfIncidence(mirrors, bounds[n], true)
		sum += nrows*100 + ncols
	}
	return strconv.Itoa(sum)
}

func (d Day13) parse(input string) ([]map[pair]bool, []pair) {
	lines := utl.Lines(input)

	sep := []int{-1}
	for n, line := range lines {
		if len(line) == 0 {
			sep = append(sep, n)
		}
	}
	sep = append(sep, len(lines))

	maps := make([]map[pair]bool, 0, len(sep))
	bounds := make([]pair, 0, len(sep))
	for n := 0; n < len(sep)-1; n++ {
		local := lines[sep[n]+1 : sep[n+1]]
		mirrors := make(map[pair]bool)

		for i := 0; i < len(local); i++ {
			for j := 0; j < len(local[i]); j++ {
				if local[i][j] == '#' {
					mirrors[pair{i, j}] = true
				}
			}
		}
		maps = append(maps, mirrors)
		bounds = append(bounds, pair{len(local), len(local[0])})
	}
	return maps, bounds
}

func symmetry(a, b pair) (float64, float64) {
	if a.i == b.i { // possible vertical symmetry
		aj, bj := min(a.j, b.j), max(a.j, b.j)
		dj := bj - aj
		if dj%2 == 0 {
			return -1, -1 // reflection line should be between columns
		}
		return 0, float64(aj) + float64(dj)/2
	}
	if a.j == b.j { // posible horizontal symmetry
		ai, bi := min(a.i, b.i), max(a.i, b.i)
		di := bi - ai
		if di%2 == 0 {
			return -1, -1 // reflection line should be between rows
		}
		return float64(ai) + float64(di)/2, 0
	}
	return -1, -1 // reflection should be symmetrical about only one axis
}

func reflect(p pair, ir, jr float64) pair {
	if ir == 0 {
		dj := jr - float64(p.j)
		return pair{p.i, p.j + int(dj*2)}
	} else {
		di := ir - float64(p.i)
		return pair{p.i + int(di*2), p.j}
	}
}

func (Day13) pointOfIncidence(mirrors map[pair]bool, bounds pair, smudge bool) (int, int) {
	type fpair struct{ i, j float64 }
	points := utl.Keys(mirrors)

	for n := 0; n < len(points); n++ {
		for m := n + 1; m < len(points); m++ {
			ir, jr := symmetry(points[n], points[m])
			if ir == -1 || jr == -1 {
				continue
			}

			found := true
			smudge := smudge
			for _, p := range points {
				pp := reflect(p, ir, jr)
				if !(0 <= pp.i && pp.i < bounds.i) || !(0 <= pp.j && pp.j < bounds.j) {
					continue // p' would be out of bounds
				}

				if _, ok := mirrors[pp]; !ok {
					if smudge {
						smudge = false // wipe off the smudge
						continue
					}
					found = false
					break
				}
			}

			if found && !smudge {
				return int(ir + 0.5), int(jr + 0.5)
			}
		}
	}
	return 0, 0
}
