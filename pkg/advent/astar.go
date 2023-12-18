package advent

import (
	"math"
	"slices"
)

type AStar[T any] struct{}

func astar[T any](graph [][]T, src, dst point, g, f func(T) int) []point {
	return AStar[T]{}.run(graph, src, dst, g, f)
}

func (AStar[T]) initialize(src point, size point) ([][]int, [][]int) {
	gscore := make2d(size.i, size.j, math.MaxInt)
	fscore := make2d(size.i, size.j, math.MaxInt)

	gscore[src.i][src.j] = 0
	fscore[src.i][src.j] = 0

	return gscore, fscore
}

func (astar AStar[T]) run(grid [][]T, src, dst point, g, f func(T) int) []point {
	bounds := point{len(grid), len(grid[0])}
	gscore, fscore := astar.initialize(src, bounds)

	history := map[state]state{}
	queue := []state{{src, E}, {src, S}}

	for len(queue) > 0 {
		idx := astar.lowestScore(queue, fscore)
		curr := pop(&queue, idx)

		if curr.pos == dst {
			path := astar.backtrack(history, curr)
			return transform(path, func(s state) point { return s.pos })
		}

		for _, dir := range orthogonal() {
			next := state{curr.pos.add(dir), dir}
			if !inbounds(next.pos, bounds) {
				continue
			}

			score := gscore[curr.pos.i][curr.pos.j] + g(grid[next.pos.i][next.pos.j])
			if score < gscore[next.pos.i][next.pos.j] {
				history[next] = curr
				gscore[next.pos.i][next.pos.j] = score
				fscore[next.pos.i][next.pos.j] = score + f(grid[next.pos.i][next.pos.j])

				if !slices.Contains(queue, next) {
					queue = append([]state{next}, queue...)
				}
			}
		}
	}
	return []point{} // no path was found
}

func (AStar[T]) lowestScore(queue []state, scores [][]int) int {
	first := queue[0]
	idx, val := 0, scores[first.pos.i][first.pos.j]
	for n, v := range queue {
		if scores[v.pos.i][v.pos.j] < val {
			val = scores[v.pos.i][v.pos.j]
			idx = n
		}
	}
	return idx
}

func (AStar[T]) backtrack(history map[state]state, dst state) []state {
	path := []state{dst}
	curr := dst
	for {
		prev, ok := history[curr]
		if !ok {
			break
		}
		path = append(path, prev)
		curr = prev
	}
	slices.Reverse(path)
	return path
}
