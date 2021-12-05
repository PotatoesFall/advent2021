package main

import (
	"fmt"
	"strconv"
)

// Lines contains a list of Line
// it represents the problem input
type Lines []Line

func makeLines(lines [][]byte) Lines {
	out := make(Lines, 0, len(lines))
	for _, line := range lines {
		out = append(out, newLine(line))
	}

	return out
}

// Modes finds the mode for each column
func (l Lines) Modes() Line {
	var counts [12]int
	for _, line := range l {
		for i := range counts {
			if line[i] == '1' {
				counts[i]++
			}
		}
	}

	var modes [12]byte
	for i := range modes {
		fmt.Println(counts[i], len(l))
		if counts[i] == len(l)/2 && len(l)%2 == 0 {
			panic(`no mode detected`)
		}

		if counts[i] > len(l)/2 {
			modes[i] = '1'
			continue
		}

		modes[i] = '0'
	}

	return modes
}

// FindLastMatch tries to find a line using the method described in part 2
func (l Lines) FindLastMatch(useMode bool) Line {
	out := make([]bool, len(l))
	var result Line
	for col := 0; col < 12; col++ {
		mode := findModeWithFilter(l, out, col)

		for row, line := range l {
			if (line[col] != mode) == useMode {
				out[row] = true
				continue
			}

			if !out[row] {
				result = line
			}
		}
	}

	if result == (Line{}) {
		panic(`no line found`)
	}

	return result
}

func findModeWithFilter(l Lines, out []bool, col int) byte {
	count, total := 0, 0
	for i, line := range l {
		if out[i] {
			continue
		}

		if line[col] == '1' {
			count++
		}

		total++
	}

	if count > total/2 || (count == total/2 && total%2 == 0) {
		return '1'
	}

	return '0'
}

// Line is a line of input with twelve digits
type Line [12]byte

func newLine(line []byte) Line {
	if len(line) != 12 {
		panic(`line length mismatch`)
	}

	var out Line
	copy(out[:], line)

	return out
}

// Invert turns all 1 into 0 and vice versa
func (l Line) Invert() Line {
	for i, v := range l {
		if v == '1' {
			l[i] = '0'
			continue
		}

		if v == '0' {
			l[i] = '1'
			continue
		}

		panic(`invalid value encountered`)
	}

	return l
}

// Int parses the value as an unsigned binary integer
func (l Line) Int() int {
	var str string
	for _, b := range l {
		str += string(b)
	}

	v, err := strconv.ParseUint(str, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(v)
}

// String returns the original representation as a string
func (l Line) String() string {
	return string(l[:])
}
