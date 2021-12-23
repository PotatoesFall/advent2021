package main

import (
	"bytes"
	"os"
	"regexp"
	"strconv"
)

func readInput(filename string) []RebootStep {
	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	split := bytes.Split(input, []byte{'\n'})
	steps := make([]RebootStep, len(split))
	for i, step := range split {
		steps[i] = parseLine(step)
	}

	return steps
}

var inputRegex = regexp.MustCompile(`^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`)

func parseLine(input []byte) RebootStep {
	matches := inputRegex.FindSubmatch(input)
	if len(matches) != 8 {
		panic(input)
	}

	return RebootStep{
		On: string(matches[1]) == `on`,
		Cuboid: Cuboid{
			X: NewRange(parseInt(matches[2]), parseInt(matches[3])),
			Y: NewRange(parseInt(matches[4]), parseInt(matches[5])),
			Z: NewRange(parseInt(matches[6]), parseInt(matches[7])),
		},
	}
}

func parseInt(input []byte) int64 {
	i, err := strconv.ParseInt(string(input), 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}
