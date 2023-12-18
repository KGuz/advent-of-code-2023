package advent

import (
	"math"
	"slices"
)

type AStar struct{}

func astar(graph [][]byte, src, dst point) []point {
	return AStar{}.run(graph, src, dst)
}

func (AStar) initialize(src point, size point) ([][]int, [][]int) {
	gscore := make2d(size.i, size.j, math.MaxInt)
	fscore := make2d(size.i, size.j, math.MaxInt)

	gscore[src.i][src.j] = 0
	fscore[src.i][src.j] = 0

	return gscore, fscore
}

func (astar AStar) run(graph [][]byte, src, dst point) []point {
	bounds := point{len(graph), len(graph[0])}
	gscore, fscore := astar.initialize(src, bounds)

	history := map[state]state{}
	queue := []state{{src, E}, {src, S}}

	for len(queue) > 0 {
		idx := astar.lowestScore(queue, fscore)
		curr := pop(&queue, idx)

		path := astar.backtrack(history, curr)
		if curr.pos == dst {
			slices.Reverse(path)
			positions := transform(path, func(s state) point { return s.pos })
			return positions[1:]
		}

		for _, dir := range orthogonal() {
			next := state{curr.pos.add(dir), dir}
			if !inbounds(next.pos, bounds) {
				continue
			}
			if curr.dir == dir.mul(-1) {
				continue // cant reverse direction
			}
			if len(path) > 3 && all(path[:3], func(s state) bool { return s.dir == dir }) {
				continue // cant move 4 times in the same direction
			}

			cost := int(graph[next.pos.i][next.pos.j] - '0')
			score := gscore[curr.pos.i][curr.pos.j] + cost

			if score < gscore[next.pos.i][next.pos.j] {
				history[next] = curr
				gscore[next.pos.i][next.pos.j] = score
				fscore[next.pos.i][next.pos.j] = score

				if !slices.Contains(queue, next) {
					queue = append([]state{next}, queue...)
				}
			}
		}
	}
	return []point{} // no path was found
}

func (AStar) lowestScore(queue []state, scores [][]int) int {
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

func (AStar) backtrack(history map[state]state, dst state) []state {
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
	return path
}
