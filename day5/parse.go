package main

import (
	"bytes"
	"regexp"
	"strconv"
)

func parseVentLines(input []byte) Lines {
	unparsedLines := bytes.Split(input, []byte{'\n'})

	field := make(Lines, 0, len(unparsedLines))
	for _, unparsed := range unparsedLines {
		field = append(field, parseLine(unparsed))
	}

	return field
}

var lineParser = regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

func parseLine(input []byte) Line {
	matches := lineParser.FindStringSubmatch(string(input))

	if len(matches) != 5 {
		panic(`couldn't parse line`)
	}

	var nums [4]int
	for i, str := range matches[1:] {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(`unparsable data`)
		}

		nums[i] = n
	}

	return Line{
		start: Point{nums[0], nums[1]},
		end:   Point{nums[2], nums[3]},
	}
}
