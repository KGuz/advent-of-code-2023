package advent

import (
	"bufio"
	"strings"
)

type pair struct{ i, j int }
type direction struct{}

var DIR direction

func (direction) N() pair { return pair{-1, 0} }
func (direction) E() pair { return pair{0, 1} }
func (direction) S() pair { return pair{1, 0} }
func (direction) W() pair { return pair{0, -1} }

func (direction) orthogonal() [4]pair {
	return [4]pair{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
}
func (direction) oblique() [4]pair {
	return [4]pair{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
}
func (d direction) all() [8]pair {
	return [8]pair{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
}

func lines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func elements(s string) [][]byte {
	var matrix [][]byte
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		matrix = append(matrix, []byte(sc.Text()))
	}
	return matrix
}

func transform[From any, To any](slice []From, f func(From) To) []To {
	res := make([]To, 0, len(slice))
	for _, val := range slice {
		res = append(res, f(val))
	}
	return res
}

func sum[T any](slice []T, f func(T) int) int {
	sum := 0
	for _, val := range slice {
		sum += f(val)
	}
	return sum
}

func inbounds(i, j, cols, rows int) bool {
	return 0 <= i && i < cols && 0 <= j && j < rows
}
