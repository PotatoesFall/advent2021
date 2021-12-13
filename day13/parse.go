package main

import (
	"bytes"
	"regexp"
	"strconv"
)

var (
	regexPoint = regexp.MustCompile(`^(\d+),(\d+)$`)
	regexFold  = regexp.MustCompile(`^fold along (x|y)=(\d+)$`)
)

func parse(input []byte) (Page, []Fold) {
	page := Page{}
	folds := []Fold{}

	for _, line := range bytes.Split(input, []byte{'\n'}) {
		if regexPoint.Match(line) {
			page[parsePoint(line)] = struct{}{}
		}

		if regexFold.Match(line) {
			folds = append(folds, parseFold(line))
		}
	}

	return page, folds
}

func parsePoint(src []byte) Point {
	subMatches := regexPoint.FindSubmatch(src)
	if len(subMatches) != 3 {
		panic(subMatches)
	}

	return Point{
		x: parseInt(subMatches[1]),
		y: parseInt(subMatches[2]),
	}
}

func parseFold(src []byte) Fold {
	subMatches := regexFold.FindSubmatch(src)
	if len(subMatches) != 3 {
		panic(subMatches)
	}

	vertical := false
	if bytes.Equal(subMatches[1], []byte("x")) {
		vertical = true
	}

	return Fold{
		number:     parseInt(subMatches[2]),
		isVertical: vertical,
	}
}

func parseInt(src []byte) int {
	n, err := strconv.Atoi(string(src))
	if err != nil {
		panic(err)
	}

	return n
}
