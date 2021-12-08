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

// AllBoxes returns both samples and output boxes
func (l Line) AllBoxes() [14]Box {
	var all [14]Box
	copy(all[:10], l.samples[:])
	copy(all[10:], l.output[:])
	return all
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

// TrueSegment refers to the actual segment that a
// Segment corresponds to
type TrueSegment int

// SegmentLabel is the name of the segment, such as 'a'
type SegmentLabel rune
