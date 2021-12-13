package main

import "strings"

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
