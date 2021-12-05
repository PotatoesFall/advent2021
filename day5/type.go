package main

// Field is a vent field
type Field [][]int

// CountPointsWithMin counts all points in the field with more than n vents
func (f Field) CountPointsWithMin(n int) int {
	count := 0

	for _, line := range f {
		for _, v := range line {
			if v >= n {
				count++
			}
		}
	}

	return count
}

// Lines is a set of vent lines
// it can be parsed from the input using parseVentLines()
type Lines []Line

// Field returns the field resulting from the set of lines
func (l Lines) Field() Field {
	var maxX, maxY int
	for _, line := range l {
		localMaxX := max(line.start.x, line.end.x)
		localMaxY := max(line.start.y, line.end.y)

		if localMaxX > maxX {
			maxX = localMaxX
		}
		if localMaxY > maxY {
			maxY = localMaxY
		}
	}

	field := make([][]int, maxY+1)
	for i := range field {
		field[i] = make([]int, maxX+1)
	}

	for _, l := range l {
		for _, p := range l.Points() {
			field[p.y][p.x]++
		}
	}

	return field
}

// OnlyHorizontalOrVertical filters the lines
// it returns only horizontal and vertical lines
func (l Lines) OnlyHorizontalOrVertical() Lines {
	field := Lines{}
	for _, line := range l {
		if line.IsHorizontalOrVertical() {
			field = append(field, line)
		}
	}

	return field
}

// Line represents a line of vents
type Line struct {
	start, end Point
}

// Points returns all points on the line
func (l Line) Points() []Point {
	var points []Point
	minY, maxY := min(l.start.y, l.end.y), max(l.start.y, l.end.y)
	minX, maxX := min(l.start.x, l.end.x), max(l.start.x, l.end.x)

	switch {
	case minX == maxX: // vertical
		for y := minY; y <= maxY; y++ {
			points = append(points, Point{x: l.start.x, y: y})
		}

	case minY == maxY: // horizontal
		for x := minX; x <= maxX; x++ {
			points = append(points, Point{x: x, y: l.start.y})
		}

	default: // diagonal
		x, y := l.start.x, l.start.y
		for {
			points = append(points, Point{x: x, y: y})

			if x == l.end.x {
				break
			}

			if l.start.x > l.end.x {
				x--
			} else {
				x++
			}
			if l.start.y > l.end.y {
				y--
			} else {
				y++
			}
		}
	}

	return points
}

// IsHorizontalOrVertical returns true if the line is horizontal or vertical
func (l Line) IsHorizontalOrVertical() bool {
	if l.start == l.end {
		panic(`same start and end!`)
	}

	return l.start.x == l.end.x || l.start.y == l.end.y
}

// Point is a point in the field
type Point struct {
	x, y int
}
