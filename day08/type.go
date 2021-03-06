package main

import (
	"strings"
)

// Digit is a number to be displayed on a box
type Digit int

// Line is a line of input with 10 random sample readings and 4 output readings
type Line struct {
	samples [10]Box
	output  [4]Box
}

// NewLine parses a line of input into a line
func NewLine(str string) Line {
	var lin Line

	split := strings.Split(str, `|`)
	split[0], split[1] = strings.TrimSpace(split[0]), strings.TrimSpace(split[1])

	for i, segmentsStr := range strings.Split(split[0], ` `) {
		if i == 10 {
			break
		}

		lin.samples[i] = NewBox(segmentsStr)
	}

	for i, segmentsStr := range strings.Split(split[1], ` `) {
		if i == 4 {
			break
		}

		lin.output[i] = NewBox(segmentsStr)
	}

	return lin
}

// BoxWithLength returns the box from the samples with the specified length
func (l Line) BoxWithLength(length int) Box {
	for _, box := range l.samples {
		if box.Count() == length {
			return box
		}
	}

	panic(`ashdsjkdhfkjsdf`)
}

// Box is a 7-segment display showing a number
type Box [7]bool

// NewBox parses a new box
func NewBox(str string) Box {
	var seg Box
	for _, r := range str {
		if r < 'a' || r > 'g' {
			panic(r)
		}

		seg[segmentByLabel[SegmentLabel(r)]] = true
	}

	return seg
}

// Count returns the number of segments contained in the Box
func (b Box) Count() int {
	count := 0
	for _, seg := range b {
		if seg {
			count++
		}
	}

	return count
}

// String returns a string representation similar to the original
func (b Box) String() string {
	var str string
	for i, on := range b {
		if on {
			str += string(segmentLabels[Segment(i)])
		}
	}

	return str
}

// Segments returns all segments in the box that are on
func (b Box) Segments() []Segment {
	segments := []Segment{}

	for seg, has := range b {
		if has {
			segments = append(segments, Segment(seg))
		}
	}

	return segments
}

// Segment refers to a specific Segment on a Box
type Segment int

// SegmentLabel is the name of the segment, such as 'a'
type SegmentLabel rune

var segmentByLabel = map[SegmentLabel]Segment{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
}

var segmentLabels = map[Segment]SegmentLabel{
	0: 'a',
	1: 'b',
	2: 'c',
	3: 'd',
	4: 'e',
	5: 'f',
	6: 'g',
}
