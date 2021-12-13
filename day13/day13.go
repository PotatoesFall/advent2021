package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile(`input13`)
	// input, err := os.ReadFile(`input13_test`)
	if err != nil {
		panic(err)
	}

	page, lines := parse(input)

	for i, line := range lines {
		// fmt.Print(page.String())
		page.Fold(line)

		if i == 0 {
			fmt.Printf("Part 1 - there are %d points on the page\n", page.Count())
		}
	}

	fmt.Print(page.String())
}

// Fold is a line to fold along
type Fold struct {
	number     int
	isVertical bool // false means horizontal
}

// Point is a point marked on the page
// x is column, y is row
type Point struct {
	x, y int
}

// Page is a foldable page with points
type Page map[Point]struct{}

// Count returns the number of points on the page
func (p Page) Count() int {
	return len(p)
}

// Fold folds the page along the line
func (p Page) Fold(l Fold) {
	if l.isVertical {
		p.foldVertical(l.number)
		return
	}

	p.foldHorizontal(l.number)
}

func (p Page) foldHorizontal(y int) { //nolint:dupl
	for point := range p {
		if point.y >= y {
			delete(p, point)
		}

		if point.y > y {
			point.y = y - (point.y - y)
			p[point] = struct{}{}
		}
	}
}

func (p Page) foldVertical(x int) { //nolint:dupl
	for point := range p {
		if point.x >= x {
			delete(p, point)
		}

		if point.x > x {
			point.x = x - (point.x - x)
			p[point] = struct{}{}
		}
	}
}

func (p Page) String() string {
	maxX, maxY := 0, 0
	for p := range p {
		if p.x > maxX {
			maxX = p.x
		}

		if p.y > maxY {
			maxY = p.y
		}
	}

	str := strings.Builder{}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			_, found := p[Point{x, y}]
			if found {
				str.WriteRune('#')
			} else {
				str.WriteRune('.')
			}
		}
		str.WriteRune('\n')
	}

	return str.String()
}

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
