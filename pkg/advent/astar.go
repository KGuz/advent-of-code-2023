package advent

import "slices"

type AStar struct{}

func astar(graph [][]byte, src, dst pair) []pair {
	return AStar{}.run(graph, src, dst)
}

func (AStar) initialize(src pair, size pair) ([][]int, [][]int) {
	INFINITY := int(^uint(0) >> 1)

	gscore := make2d(size.i, size.j, INFINITY)
	fscore := make2d(size.i, size.j, INFINITY)

	gscore[src.i][src.j] = 0
	fscore[src.i][src.j] = 0

	return gscore, fscore
}

func (astar AStar) run(graph [][]byte, src, dst pair) []pair {
	bounds := pair{len(graph), len(graph[0])}
	gscore, fscore := astar.initialize(src, bounds)

	history := map[state]state{}
	queue := []state{{src, E}, {src, S}}

	for len(queue) > 0 {
		idx := astar.lowestScore(queue, fscore)
		curr := pop(&queue, idx)

		path := astar.backtrack(history, curr)
		if curr.pos == dst {
			slices.Reverse(path)
			positions := transform(path, func(s state) pair { return s.pos })
			return positions[1:]
		}

		for _, dir := range orthogonal() {
			next := state{pair{curr.pos.i + dir.i, curr.pos.j + dir.j}, dir}
			if !inbounds(next.pos, bounds) {
				continue
			}
			if (curr.dir == pair{-dir.i, -dir.j}) {
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
	return []pair{} // no path was found
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
