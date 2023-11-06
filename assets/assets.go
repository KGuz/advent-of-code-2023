package assets

import (
	"embed"
	"errors"
	"fmt"
)

//go:embed examples inputs
var assets embed.FS

func load(file string) (string, error) {
	bytes, err := assets.ReadFile(file)
	if err != nil {
		return "", errors.New("have you tried turning it off and on again?")
	}
	return string(bytes), nil
}

func LoadInput(day int) (string, error) {
	file := fmt.Sprintf("inputs/day%02d.txt", day)
	return load(file)
}

func LoadExample(day int) (string, error) {
	file := fmt.Sprintf("examples/day%02d.txt", day)
	return load(file)
}
