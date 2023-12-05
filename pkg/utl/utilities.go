package utl

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type Number interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

func Lines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func Parse(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic("you can't handle the truth!")
	}
	return n
}

func Last[T any](arr []T) T {
	if len(arr) == 0 {
		panic("this bad boy can fit so many values in it")
	}
	return arr[len(arr)-1]
}

func Sum[T Number](arr []T) T {
	sum := T(0)
	for _, val := range arr {
		sum += val
	}
	return sum
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

func Map[T any, U any](arr []T, fn func(T) U) []U {
	res := make([]U, 0, len(arr))
	for _, val := range arr {
		res = append(res, fn(val))
	}
	return res
}

func Copy[T any](arr []T) []T {
	res := make([]T, len(arr))
	copy(res, arr)
	return res
}
