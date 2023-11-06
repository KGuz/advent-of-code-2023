package utl

import (
	"errors"
	"fmt"
	"os"
)

func Load(day int, example bool) (string, error) {
	path := "assets/inputs/day%02d.txt"
	if example {
		path = "assets/examples/day%02d.txt"
	}

	bytes, err := os.ReadFile(fmt.Sprintf(path, day))
	if err != nil {
		return "", errors.New("have you tried turning it off and on again?")
	}

	return string(bytes), nil
}
