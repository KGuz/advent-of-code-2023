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

func Keys[K interface{ comparable }, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K interface{ comparable }, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
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
