package advent

import "math"

type Dijkstra struct{}

func dijkstra(graph [][]byte, src, dst point) []point {
	return Dijkstra{}.run(graph, src, dst)
}

func (Dijkstra) initialize(src point, size point) ([]state, [][]int) {
	cost := make([][]int, size.i)
	queue := make([]state, 0, 4*size.i*size.j)

	for i := 0; i < size.i; i++ {
		cost[i] = make([]int, size.j)
		for j := 0; j < size.j; j++ {
			cost[i][j] = math.MaxInt
			queue = append(queue, state{point{i, j}, S})
			queue = append(queue, state{point{i, j}, W})
			queue = append(queue, state{point{i, j}, E})
			queue = append(queue, state{point{i, j}, N})
		}
	}
	cost[src.i][src.j] = 0
	return queue, cost
}

func (d Dijkstra) run(graph [][]byte, src, dst point) []point {
	bounds := point{len(graph), len(graph[0])}
	queue, cost := d.initialize(src, bounds)

	history := make(map[state]state)

	for len(queue) > 0 {
		idx := d.lowestScore(queue, cost)
		curr := pop(&queue, idx)

		path := d.backtrack(history, curr)
		for _, dir := range orthogonal() {
			next := state{curr.pos.add(dir), dir}
			if !inbounds(next.pos, bounds) {
				continue
			}
			if len(path) > 1 && (path[0].dir == dir.mul(-1)) {
				continue // cant reverse direction
			}
			if len(path) > 3 && all(path[:3], func(s state) bool { return s.dir == dir }) {
				continue // cant move 4 times in the same direction
			}

			alt := cost[curr.pos.i][curr.pos.j] + int(graph[next.pos.i][next.pos.j]-'0')
			if alt < cost[next.pos.i][next.pos.j] {
				cost[next.pos.i][next.pos.j] = alt
				history[next] = curr
			}
		}
	}
	return transform(d.backtrack(history, state{dst, W}), func(s state) point { return s.pos })
}

func (Dijkstra) lowestScore(queue []state, scores [][]int) int {
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

func (Dijkstra) backtrack(history map[state]state, dst state) []state {
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

func (Dijkstra) last3(history map[point]point, curr point) []point {
	path := []point{curr}
	n := 0
	for n < 3 {
		prev, ok := history[curr]
		if !ok {
			break
		}
		path = append([]point{prev}, path...)
		curr = prev
		n++
	}
	return path
}
