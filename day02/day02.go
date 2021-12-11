package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var errEmpty = errors.New(`empty`)

func main() {
	input, err := os.ReadFile(`input2`)
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(input, []byte("\n"))

	s1, s2 := new(state), new(state)
	for _, line := range lines {
		m, err := parse(line)
		if err != nil {
			fmt.Printf("Skipping %q\n", line)
			continue
		}

		s1.mutate(m)
		s2.mutate2(m)
	}

	fmt.Printf("Product Part 1: %d\n", s1.product())
	fmt.Printf("Product Part 2: %d\n", s2.product())
}

func parse(data []byte) (movement, error) {
	s := string(data)

	if strings.HasPrefix(s, `forward `) {
		n, err := strconv.Atoi(s[8:])
		if err != nil {
			return movement{}, err
		}
		return movement{Forward: n}, err
	}

	if strings.HasPrefix(s, `down `) {
		n, err := strconv.Atoi(s[5:])
		if err != nil {
			return movement{}, err
		}
		return movement{Vertical: n}, err
	}

	if strings.HasPrefix(s, `up `) {
		n, err := strconv.Atoi(s[3:])
		if err != nil {
			return movement{}, err
		}
		return movement{Vertical: -n}, err
	}

	return movement{}, errEmpty
}

type state struct {
	Depth, Position, Aim int
}

func (s *state) mutate(m movement) {
	s.Position += m.Forward
	s.Depth += m.Vertical
}

func (s *state) mutate2(m movement) {
	s.Aim += m.Vertical
	s.Depth += s.Aim * m.Forward
	s.Position += m.Forward
}

func (s state) product() int {
	return s.Depth * s.Position
}

type movement struct {
	Forward, Vertical int
}
