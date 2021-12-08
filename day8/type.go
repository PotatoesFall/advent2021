package main

import (
	"strings"
)

type digit int

type line struct {
	samples [10]segments
	output  [4]segments
}

func (l line) AllSegments() [14]segments {
	var all [14]segments
	copy(all[:10], l.samples[:])
	copy(all[10:], l.output[:])
	return all
}

func lineFromStr(str string) line {
	var lin line

	split := strings.Split(str, `|`)
	split[0], split[1] = strings.TrimSpace(split[0]), strings.TrimSpace(split[1])

	for i, segmentsStr := range strings.Split(split[0], ` `) {
		if i == 10 {
			break
		}

		lin.samples[i] = segmentsFromStr(segmentsStr)
	}

	for i, segmentsStr := range strings.Split(split[1], ` `) {
		if i == 4 {
			break
		}

		lin.output[i] = segmentsFromStr(segmentsStr)
	}

	return lin
}

type segments [7]bool

type segNum int

type segName rune

func (s segments) Count() int {
	count := 0
	for _, seg := range s {
		if seg {
			count++
		}
	}

	return count
}

func (s segments) String() string {
	var str string
	for i, on := range s {
		if on {
			str += string(segmentNameByNumer[segNum(i)])
		}
	}

	return str
}

func segmentsFromStr(str string) segments {
	var seg segments
	for _, r := range str {
		if r < 'a' || r > 'g' {
			panic(r)
		}

		seg[segmentNumberByName[segName(r)]] = true
	}

	return seg
}
