package utl

import (
	"strconv"
	"strings"
)

type Number interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

func Lines(s string) []string {
	return strings.Split(strings.TrimRight(s, "\n"), "\n")
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
