package utl

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type Number interface {
	Integer | float32 | float64
}

type Integer interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64
}

func Lines(str string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(str))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func Parse(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic("you can't handle the truth!")
	}
	return n
}

func Captures(re *regexp.Regexp, str string) []string {
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

func CapturesNamed(re *regexp.Regexp, str string) map[string]string {
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

func Last[T any](slice []T) T {
	if len(slice) == 0 {
		panic("this bad boy can fit so many values in it")
	}
	return slice[len(slice)-1]
}

func Sum[T Number](slice []T) T {
	sum := T(0)
	for _, val := range slice {
		sum += val
	}
	return sum
}

func Map[T any, U any](slice []T, fn func(T) U) []U {
	res := make([]U, 0, len(slice))
	for _, val := range slice {
		res = append(res, fn(val))
	}
	return res
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	res := make([]T, 0)
	for _, val := range slice {
		if fn(val) {
			res = append(res, val)
		}
	}
	return res
}

// Keys returns the keys of the map m.
// The keys will be in an indeterminate order.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Values returns the values of the map m.
// The values will be in an indeterminate order.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// Greatest common divisor (GCD) via Euclidean algorithm
func GCD[T Integer](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Least Common Multiple (LCM) via GCD
func LCM[T Integer](a, b T, integers ...T) T {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func All[T any](slice []T, fn func(T) bool) bool {
	for _, val := range slice {
		if !fn(val) {
			return false
		}
	}
	return true
}

func Any[T any](slice []T, fn func(T) bool) bool {
	for _, val := range slice {
		if fn(val) {
			return true
		}
	}
	return false
}

func None[T any](slice []T, fn func(T) bool) bool {
	return !Any(slice, fn)
}

func FloodFill(graph [][]byte, si, sj int, fn func(i, j int) bool) map[[2]int]bool {
	directions := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	visited := make(map[[2]int]bool)
	queue := [][2]int{{si, sj}}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if _, ok := visited[node]; ok {
			continue
		}

		visited[node] = true
		for _, dir := range directions {
			next := [2]int{node[0] + dir[0], node[1] + dir[1]}

			inbounds := 0 <= next[0] && next[0] < len(graph) && 0 <= next[1] && next[1] < len(graph[next[0]])
			if inbounds && fn(next[0], next[1]) {
				queue = append(queue, next)
			}
		}
	}

	return visited
}

const (
	CYAN      = "\033[96m"
	PURPLE    = "\033[95m"
	BLUE      = "\033[94m"
	YELLOW    = "\033[93m"
	GREEN     = "\033[92m"
	RED       = "\033[91m"
	BOLD      = "\033[1m"
	UNDERLINE = "\033[4m"
	END       = "\033[0m"
)

func AbsDiff[T Number](a, b T) T {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}

func ToBytes(s string) []byte { return []byte(s) }

func Clone2d[T any](matrix [][]T) [][]T {
	duplicate := make([][]T, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]T, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}
