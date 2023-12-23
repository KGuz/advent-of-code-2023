package advent

import (
	// . "aoc/internal/heap"

	"fmt"
	"slices"
	"strconv"
)

func (d Day23) PartTwo2(input string) string {
	hike, bounds := d.parse(input)
	start, end := point{0, 1}, point{bounds.i - 1, bounds.j - 2}

	segments := d.segments(hike, start, end)
	trails, src, dst := d.trails(segments, start, end)

	for n, t := range trails {
		if n == src {
			fmt.Printf("SRC: %v: %v\n", n, t)
		} else if n == dst {
			fmt.Printf("DST: %v: %v\n", n, t)
		} else {
			fmt.Printf("%v: %v\n", n, t)
		}
	}

	d.print(hike, segments, bounds)
	d.hike(trails, src, dst)

	return strconv.Itoa(0) // 2840
}

func (d Day23) hike(trails []Trail, src, dst int) {
	queue := [][]int{{src}}
	paths := [][]int{}

	for len(queue) > 0 {
		visited := pop(&queue, len(queue)-1)
		curr := visited[0]

		if curr == dst {
			paths = append(paths, visited)
			fmt.Printf("found path: %v\n", visited)
			continue
		}

		for _, t := range trails[curr].adj {
			if !slices.Contains(visited, t) {
				queue = append(queue, append([]int{t}, visited...))
			}
		}
	}

	longestIdx, longestVal := 0, 0
	for i, path := range paths {
		length := accumulate(path, func(acc, n int) int {
			return acc + len(trails[n].path)
		})
		if length > longestVal {
			longestVal = length
			longestIdx = i
		}
	}

	points := d.backtrack(trails, paths[longestIdx])

	fmt.Printf("longest: %v\n", points)
}

func (Day23) backtrack(trails []Trail, path []int) []point {

	points := map[point]bool{}
	// for _, n := range path {

	// }
	return keys(points)
}

type Trail struct {
	adj  []int
	path []point
}

func (t Trail) String() string {
	return fmt.Sprintf("Trail(adj: %v, path: (%v %v))", t.adj, t.path[0], t.path[len(t.path)-1])
}

func (d Day23) segments(hike map[point]byte, start point, end point) map[point][]point {
	segments := map[point][]point{}
	queue := []point{start}
	visited := map[point]bool{}
	nodes := []point{start, end}

	for len(queue) > 0 {
		curr := pop(&queue, 0)
		if _, ok := visited[curr]; ok {
			continue
		}

		// traverse hike paths until intersection
		path := []point{curr}
		prev := curr
		for {
			neighbours := transform(orthogonal(), func(dir point) point { return curr.add(dir) })
			neighbours = filter(neighbours, func(p point) bool {
				_, exists := hike[p]
				_, seen := visited[p]
				return exists && prev != p && !seen
			})

			if len(neighbours) != 1 {
				nodes = append(nodes, curr)
				queue = append(queue, neighbours...)
				break
			}

			path = append(path, neighbours[0])
			prev, curr = curr, neighbours[0]
		}

		// mark whole path as visited
		for _, point := range path {
			visited[point] = true
		}

		// if path[len(path)-1] != end {
		// 	path = path[:len(path)-1]
		// }
		segments[path[0]] = path
	}
	return segments
}

func (d Day23) trails(segments map[point][]point, start, end point) ([]Trail, int, int) {
	keys := keys(segments)
	trails := transform(keys, func(p point) Trail {
		return Trail{path: segments[p], adj: []int{}}
	})

	for i := 0; i < len(keys); i++ {
		s1 := segments[keys[i]]
		for j := 0; j < len(keys); j++ {
			if i != j {
				s2 := segments[keys[j]]
				if d.pathsAreAdjecent(s1, s2) {
					trails[i].adj = append(trails[i].adj, j)
				}
			}
		}
	}

	src, dst := 0, 0
	for n, t := range trails {
		if t.path[0] == start {
			src = n
		}
		if t.path[len(t.path)-1] == end {
			dst = n
		}
	}

	return trails, src, dst
}

func (Day23) pathsAreAdjecent(a, b []point) bool {
	s1, e1 := a[0], a[len(a)-1]
	s2, e2 := b[0], b[len(b)-1]

	return e1.sub(s2).l1dist(point{}) <= 2 ||
		s1.sub(e2).l1dist(point{}) <= 2 ||
		s1.sub(s2).l1dist(point{}) <= 2 ||
		e1.sub(e2).l1dist(point{}) <= 2
}
