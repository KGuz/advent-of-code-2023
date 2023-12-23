package advent

import (
	// . "aoc/internal/heap"

	"aoc/pkg/dbg"
	"fmt"
	"slices"
	"strconv"
	"time"
)

type Day23 struct {
	/* --- Day 23: A Long Walk ---
	The Elves resume water filtering operations! Clean water starts flowing
	over the edge of Island Island.

	They offer to help you go over the edge of Island Island, too! Just hold on
	tight to one end of this impossibly long rope and they'll lower you down a
	safe distance from the massive waterfall you just created.

	As you finally reach Snow Island, you see that the water isn't really
	reaching the ground: it's being absorbed by the air itself. It looks like
	you'll finally have a little downtime while the moisture builds up to
	snow-producing levels. Snow Island is pretty scenic, even without any snow;
	why not take a walk?

	There's a map of nearby hiking trails (your puzzle input) that indicates
	paths (.), forest (#), and steep slopes (^, >, v, and <).

	For example:

	#.#####################
	#.......#########...###
	#######.#########.#.###
	###.....#.>.>.###.#.###
	###v#####.#v#.###.#.###
	###.>...#.#.#.....#...#
	###v###.#.#.#########.#
	###...#.#.#.......#...#
	#####.#.#.#######.#.###
	#.....#.#.#.......#...#
	#.#####.#.#.#########v#
	#.#...#...#...###...>.#
	#.#.#v#######v###.###v#
	#...#.>.#...>.>.#.###.#
	#####v#.#.###v#.#.###.#
	#.....#...#...#.#.#...#
	#.#########.###.#.#.###
	#...###...#...#...#.###
	###.###.#.###v#####v###
	#...#...#.#.>.>.#.>.###
	#.###.###.#.###.#.#v###
	#.....###...###...#...#
	#####################.#

	You're currently on the single path tile in the top row; your goal is to
	reach the single path tile in the bottom row. Because of all the mist from
	the waterfall, the slopes are probably quite icy; if you step onto a slope
	tile, your next step must be downhill (in the direction the arrow is
	pointing). To make sure you have the most scenic hike possible, never step
	onto the same tile twice. What is the longest hike you can take?

	In the example above, the longest hike you can take is marked with O, and
	your starting position is marked S:

	#S#####################
	#OOOOOOO#########...###
	#######O#########.#.###
	###OOOOO#OOO>.###.#.###
	###O#####O#O#.###.#.###
	###OOOOO#O#O#.....#...#
	###v###O#O#O#########.#
	###...#O#O#OOOOOOO#...#
	#####.#O#O#######O#.###
	#.....#O#O#OOOOOOO#...#
	#.#####O#O#O#########v#
	#.#...#OOO#OOO###OOOOO#
	#.#.#v#######O###O###O#
	#...#.>.#...>OOO#O###O#
	#####v#.#.###v#O#O###O#
	#.....#...#...#O#O#OOO#
	#.#########.###O#O#O###
	#...###...#...#OOO#O###
	###.###.#.###v#####O###
	#...#...#.#.>.>.#.>O###
	#.###.###.#.###.#.#O###
	#.....###...###...#OOO#
	#####################O#

	This hike contains 94 steps. (The other possible hikes you could have taken
	were 90, 86, 82, 82, and 74 steps long.)

	Find the longest hike you can take through the hiking trails listed on your
	map. How many steps long is the longest hike? */

	/* --- Part Two ---
	As you reach the trailhead, you realize that the ground isn't as slippery
	as you expected; you'll have no problem climbing up the steep slopes.

	Now, treat all slopes as if they were normal paths (.). You still want to
	make sure you have the most scenic hike possible, so continue to ensure
	that you never step onto the same tile twice. What is the longest hike you
	can take?

	In the example above, this increases the longest hike to 154 steps:

	#S#####################
	#OOOOOOO#########OOO###
	#######O#########O#O###
	###OOOOO#.>OOO###O#O###
	###O#####.#O#O###O#O###
	###O>...#.#O#OOOOO#OOO#
	###O###.#.#O#########O#
	###OOO#.#.#OOOOOOO#OOO#
	#####O#.#.#######O#O###
	#OOOOO#.#.#OOOOOOO#OOO#
	#O#####.#.#O#########O#
	#O#OOO#...#OOO###...>O#
	#O#O#O#######O###.###O#
	#OOO#O>.#...>O>.#.###O#
	#####O#.#.###O#.#.###O#
	#OOOOO#...#OOO#.#.#OOO#
	#O#########O###.#.#O###
	#OOO###OOO#OOO#...#O###
	###O###O#O###O#####O###
	#OOO#OOO#O#OOO>.#.>O###
	#O###O###O#O###.#.#O###
	#OOOOO###OOO###...#OOO#
	#####################O#

	Find the longest hike you can take through the surprisingly dry hiking
	trails listed on your map. How many steps long is the longest hike? */
}

func (d Day23) PartOne(input string) string {
	return ""
	hike := elements(input)
	src, dst := point{0, 1}, point{len(hike) - 1, len(hike[0]) - 2}

	longest := d.scenicHike(hike, src, dst, func(p point) []point {
		switch hike[p.i][p.j] {
		case '^':
			return []point{p.add(N)}
		case '>':
			return []point{p.add(E)}
		case 'v':
			return []point{p.add(S)}
		case '<':
			return []point{p.add(W)}
		case '.':
			return []point{p.add(N), p.add(E), p.add(S), p.add(W)}
		default:
			return []point{}
		}
	})
	return strconv.Itoa(longest - 1)
}

func (d Day23) PartTwo(input string) string {
	d.PartTwo2(input)
	return ""
	hike, bounds := d.parse(input)
	src, dst := point{0, 1}, point{bounds.i - 1, bounds.j - 2}

	segments, intersections := d.makeSegments(hike, src, dst)
	fmt.Printf("intersections: %v\n", intersections)
	d.print(hike, segments, bounds)

	paths := d.traverse(hike, segments, intersections, bounds)
	lenghts := transform(paths, func(path []point) int {
		return accumulate(path, func(acc int, pos point) int { return acc + len(segments[pos]) })
	})

	fmt.Printf("lenghts: %v\n", lenghts)
	return strconv.Itoa(0) // 2840
	// return ""
}

func (Day23) parse(input string) (map[point]byte, point) {
	hike := map[point]byte{}
	elements := elements(input)

	for i, row := range elements {
		for j, elem := range row {
			if elem != '#' {
				hike[point{i, j}] = elem
			}
		}
	}
	return hike, point{len(elements), len(elements[0])}
}

func (Day23) makeSegments(hike map[point]byte, src, dst point) (map[point][]point, []point) {
	segments := map[point][]point{}
	queue := []point{src, dst}
	visited := map[point]bool{}
	intersections := []point{}

	for len(queue) > 0 {
		curr := pop(&queue, 0)
		if _, ok := visited[curr]; ok {
			continue
		}

		// traverse hike paths until intersection
		segment := []point{curr}
		prev := curr
		for {
			neighbours := transform(orthogonal(), func(dir point) point { return curr.add(dir) })
			neighbours = filter(neighbours, func(p point) bool {
				_, exists := hike[p]
				_, seen := visited[p]
				return exists && prev != p && !seen
			})

			if len(neighbours) != 1 {
				// start new segments
				queue = append(queue, neighbours...)
				break
			}
			segment = append(segment, neighbours[0])
			prev, curr = curr, neighbours[0]
		}
		// mark whole path as visited
		for _, point := range segment {
			visited[point] = true
		}

		first, last, inter := segment[0], segment[len(segment)-2], segment[len(segment)-1]
		intersections = append(intersections, inter)
		segment = segment[:len(segment)-1] // truncate segment to not include intersection

		reversed := slices.Clone(segment)
		slices.Reverse(reversed)

		segments[first] = segment
		segments[last] = reversed

	}
	return segments, intersections
}

func (Day23) leadsTo(segments map[point][]point, inter point) []point {
	leads := []point{}
	for _, dir := range orthogonal() { // searching for intersection
		next := inter.add(dir)
		if _, ok := segments[next]; ok {
			leads = append(leads, next)
		}
	}
	return leads
}

func (Day23) print(hike map[point]byte, segments map[point][]point, bounds point) {
	buffer := make2d(bounds.i, bounds.j, ".")

	colors := []string{dbg.BLUE, dbg.YELLOW, dbg.PURPLE, dbg.GREEN, dbg.CYAN, dbg.RED}
	for n, path := range values(segments) {
		c := colors[n%len(colors)]
		for _, p := range path {
			if hike[p] == '.' {
				buffer[p.i][p.j] = fmt.Sprintf("%s%c%s", c, '#', dbg.END)
			} else {
				buffer[p.i][p.j] = fmt.Sprintf("%s%c%s", c, hike[p], dbg.END)
			}
		}
	}

	for i := 0; i < bounds.i; i++ {
		for j := 0; j < bounds.j; j++ {
			fmt.Print(buffer[i][j])
		}
		println()
	}
}

func (d Day23) traverse(hike map[point]byte, segments map[point][]point, intersections []point, bounds point) [][]point {
	src, dst := point{0, 1}, point{bounds.i - 1, bounds.j - 2}

	queue := [][]point{{src, segments[src][len(segments[src])-1]}}
	finished := [][]point{}

	for len(queue) > 0 {
		path := pop(&queue, 0) // path - visited segment starts and ends
		end := path[len(path)-1]

		if end == dst { // if we finished on the dst point
			finished = append(finished, path) // add traversed segment starts as one of possibilities

			fmt.Print("finished path\n")
			continue
		}

		fmt.Printf("path %v\n", path)
		d.DebugSegments(hike, segments, path, bounds)
		time.Sleep(time.Second)

		inter := d.goToIntersection(intersections, end)
		fmt.Printf("inter: %v\n", inter)

		for _, next := range d.leadsTo(segments, inter) {
			fmt.Printf("considerig next: %v\n", next)
			if !slices.Contains(path, next) { // if we havent visited this segment yet
				start, end := next, segments[next][len(segments[next])-1] // first and last point of segment
				new := append(slices.Clone(path), start, end)             // copy hike history and include this segment
				queue = append([][]point{new}, queue...)
				fmt.Printf("we will visit %v\n", start)
			}
		}
	}
	return finished
}

func (Day23) goToIntersection(intersections []point, curr point) point {
	for _, dir := range orthogonal() { // searching for intersection
		next := curr.add(dir)
		if slices.Contains(intersections, next) {
			return next
		}
	}
	return curr
}

func (d Day23) DebugSegments(hike map[point]byte, segments map[point][]point, path []point, bounds point) {
	dbgsegments := map[point][]point{}
	for k, v := range segments {
		if slices.Contains(path, k) {
			dbgsegments[k] = v
		}
	}
	d.print(hike, dbgsegments, bounds)
}

type HikeState struct {
	pos point
	// path map[point]bool
	path []point
}

func (self HikeState) Less(other any) bool {
	// return self.pos.l2dist(point{0, 0}) < self.pos.l2dist(point{0, 0})
	return len(self.path) > len(other.(HikeState).path)
}

func (Day23) scenicHike(hike [][]byte, src, dst point, neighbours func(point) []point) int {
	// heap := MakeHeap(HikeState{src, map[point]bool{src: true}})
	bounds := point{len(hike), len(hike[0])}
	longest := 0
	queue := []HikeState{{src, []point{src}}}

	for len(queue) > 0 {
		curr := pop(&queue, 0)
		if curr.pos == dst {
			fmt.Printf("longest: %v\n", longest)
			longest = max(longest, len(curr.path))
			continue
		}

		// traverseUntilIntersection(hike, &curr.path, neighbours)
		curr.pos = curr.path[0]

		for _, next := range neighbours(curr.pos) {
			// _, ok := curr.path[next]
			ok := slices.Contains(curr.path, next)

			if !ok && inbounds(next, bounds) {
				// path := CloneMap(curr.path)
				// path[next] = true
				// heap.Push(HikeState{next, path})

				path := append([]point{next}, curr.path...)
				queue = append(queue, HikeState{next, path})
			}
		}
	}
	return longest
}

func (Day23) scenicHike2(hike map[point]byte, src, dst point, neighbours func(point) []point) int {
	// heap := MakeHeap(HikeState{src, map[point]bool{src: true}})
	longest := 0
	queue := [][]point{{src}}

	for len(queue) > 0 {
		curr := pop(&queue, 0)
		if curr[len(curr)-1] == dst {
			fmt.Printf("longest: %v\n", longest)
			longest = max(longest, len(curr))
			continue
		}

		// fmt.Println(curr)
		curr = traverseUntilIntersection(hike, curr, neighbours)
		last := curr[len(curr)-1]
		// fmt.Println(curr)

		for _, neighbour := range neighbours(last) {
			visited := slices.Contains(curr, neighbour)
			if !visited {
				next := slices.Clone(curr)
				next = append(next, neighbour)
				queue = append(queue, next)
			}
		}
	}
	return longest
}

func traverseUntilIntersection(hike map[point]byte, path []point, neighbours func(point) []point) []point {
	last := path[len(path)-1]
	previous := last
	for {
		neighbours := neighbours(last)
		neighbours = slices.DeleteFunc(neighbours, func(p point) bool { return p == previous })
		if len(neighbours) != 1 {
			break
		}
		path = append(path, neighbours[0])
		previous = last
		last = neighbours[0]
	}
	return path
}
