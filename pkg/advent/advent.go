package advent

import (
	"bufio"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Number interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

type pair struct{ i, j int }
type state struct{ pos, dir pair }

var N = pair{-1, 0}
var W = pair{0, -1}
var E = pair{0, 1}
var S = pair{1, 0}

var NW = pair{-1, -1}
var NE = pair{-1, 1}
var SW = pair{1, -1}
var SE = pair{1, 1}

func directions() []pair { return []pair{NW, N, NE, W, E, SW, S, SE} }
func orthogonal() []pair { return []pair{N, W, E, S} }
func oblique() []pair    { return []pair{NW, NE, SW, SE} }

func parse(s string) int {
	n, _ := strconv.Atoi(s)
	return n
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
	var elements [][]byte
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		elements = append(elements, []byte(sc.Text()))
	}
	return elements
}

func sum[T Number](slice []T) T {
	sum := T(0)
	for _, val := range slice {
		sum += val
	}
	return sum
}

func mul[T Number](slice []T) T {
	if len(slice) == 0 {
		return T(0)
	}

	mul := T(1)
	for _, val := range slice {
		mul *= val
	}
	return mul
}

func transform[From any, To any](slice []From, f func(From) To) []To {
	res := make([]To, 0, len(slice))
	for _, val := range slice {
		res = append(res, f(val))
	}
	return res
}

func accumulate[Elem any, Acc Number](slice []Elem, f func(Acc, Elem) Acc) Acc {
	acc := Acc(0)
	for _, val := range slice {
		acc = f(acc, val)
	}
	return acc
}

func filter[T any](slice []T, f func(T) bool) []T {
	res := make([]T, 0)
	for _, val := range slice {
		if f(val) {
			res = append(res, val)
		}
	}
	return res
}

func keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func inbounds(pos, bounds pair) bool {
	return 0 <= pos.i && pos.i < bounds.i && 0 <= pos.j && pos.j < bounds.j
}

func isdigit[T rune | byte](c T) bool {
	return '0' <= c && c <= '9'
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(integers ...int) int {
	result := integers[0]
	for i := 1; i < len(integers); i++ {
		result = (integers[i] * result) / gcd(integers[i], result)
	}
	return result
}

func all[T any](slice []T, f func(T) bool) bool {
	for _, val := range slice {
		if !f(val) {
			return false
		}
	}
	return true
}

func abs[T Number](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func captures(re *regexp.Regexp, str string) []string {
	captures := make([]string, 0)
	matches := re.FindAllStringSubmatch(str, -1)

	for _, match := range matches {
		for _, group := range match {
			if group != "" {
				captures = append(captures, group)
			}
		}
	}
	return captures
}

func capturesNamed(re *regexp.Regexp, str string) map[string]string {
	captures := make(map[string]string)

	matches := re.FindAllStringSubmatch(str, -1)
	names := re.SubexpNames()

	for _, match := range matches {
		for n, group := range match {
			if names[n] != "" && group != "" {
				captures[names[n]] = group
			}
		}
	}
	return captures
}

func pop[T any](slice *[]T, index int) T {
	elem := (*slice)[index]
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
	return elem
}

func make2d[T any](isize int, jsize int, val T) [][]T {
	slice := make([][]T, isize)
	for i := 0; i < isize; i++ {
		slice[i] = make([]T, jsize)
		for j := 0; j < jsize; j++ {
			slice[i][j] = val
		}
	}
	return slice
}

func sortedInsert[T any](slice []T, value T, f func(T) int) []T {
	i := sort.Search(len(slice), func(i int) bool { return f(slice[i]) > f(value) })
	slice = append(slice, value)
	copy(slice[i+1:], slice[i:])
	slice[i] = value
	return slice
}

func floodFill(src pair, bounds pair, f func(pair) bool) []pair {
	visited := map[pair]bool{}
	queue := []pair{{0, 0}}

	for len(queue) > 0 {
		curr := pop(&queue, 0)
		if visited[curr] {
			continue
		}
		visited[curr] = true

		for _, dir := range orthogonal() {
			next := pair{curr.i + dir.i, curr.j + dir.j}
			if !inbounds(next, bounds) {
				continue
			}
			if !f(next) {
				continue
			}
			queue = append(queue, next)
		}
	}
	return keys(visited)
}
