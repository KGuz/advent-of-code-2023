package advent

type Dijkstra struct{}

func dijkstra(graph [][]byte, src, dst pair) []pair {
	return Dijkstra{}.run(graph, src, dst)
}

func (Dijkstra) initialize(src pair, size pair) ([]state, [][]int) {
	INFINITY := int(^uint(0) >> 1)

	cost := make([][]int, size.i)
	queue := make([]state, 0, 4*size.i*size.j)

	for i := 0; i < size.i; i++ {
		cost[i] = make([]int, size.j)
		for j := 0; j < size.j; j++ {
			cost[i][j] = INFINITY
			queue = append(queue, state{pair{i, j}, S})
			queue = append(queue, state{pair{i, j}, W})
			queue = append(queue, state{pair{i, j}, E})
			queue = append(queue, state{pair{i, j}, N})
		}
	}
	cost[src.i][src.j] = 0
	return queue, cost
}

func (d Dijkstra) run(graph [][]byte, src, dst pair) []pair {
	bounds := pair{len(graph), len(graph[0])}
	queue, cost := d.initialize(src, bounds)

	history := make(map[state]state)

	for len(queue) > 0 {
		idx := d.lowestScore(queue, cost)
		curr := pop(&queue, idx)

		path := d.backtrack(history, curr)
		for _, dir := range orthogonal() {
			next := state{pair{curr.pos.i + dir.i, curr.pos.j + dir.j}, dir}
			if !inbounds(next.pos, bounds) {
				continue
			}
			if len(path) > 1 && (path[0].dir == pair{-dir.i, -dir.j}) {
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

	path1 := transform(d.backtrack(history, state{dst, S}), func(s state) pair { return s.pos })
	path2 := transform(d.backtrack(history, state{dst, W}), func(s state) pair { return s.pos })
	if len(path1) < len(path2) {
		return path1
	} else {
		return path2
	}
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

func (d Dijkstra) directions(pos pair, bounds pair, history map[pair]pair) []pair {
	last := d.last3(history, pos)
	return filter(orthogonal(), func(dir pair) bool {
		next := pair{pos.i + dir.i, pos.j + dir.j}
		if !inbounds(next, bounds) {
			return false
		}
		if len(last) > 1 && (last[0] == pair{-dir.i, -dir.j}) {
			return false // cant reverse direction
		}
		if len(last) > 3 && dir == last[0] && dir == last[1] && dir == last[2] {
			return false // cant move 4 times in the same direction
		}
		return true
	})
}

func (Dijkstra) last3(history map[pair]pair, curr pair) []pair {
	path := []pair{curr}
	n := 0
	for n < 3 {
		prev, ok := history[curr]
		if !ok {
			break
		}
		path = append([]pair{prev}, path...)
		curr = prev
		n++
	}
	return path
}
