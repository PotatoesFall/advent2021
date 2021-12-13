package main

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
