package main

import (
	"bytes"
)

// Point is a specific point in a field
// the first int is a line number, the second the column
type Point [2]int

// Field is a representation of the input data
type Field [100][100]int

// NewField parses a field from input
func NewField(data []byte) Field {
	var field Field

	lines := bytes.Split(data, []byte{'\n'})
	for i, line := range lines {
		for j, b := range line {
			field[i][j] = int(b - '0')
		}
	}

	return field
}

// GetLocalMinima returns a slice of all local minima
// it considers directly adjacent fields only, no diagonals
// in the description, they are called "low points"
func (field Field) GetLocalMinima() []Point {
	minima := []Point{}

	for i, line := range field {
		for j, val := range line {
			if (i == 0 || field[i-1][j] > val) &&
				(i == 99 || field[i+1][j] > val) &&
				(j == 0 || field[i][j-1] > val) &&
				(j == 99 || field[i][j+1] > val) {
				minima = append(minima, Point{i, j})
			}
		}
	}

	return minima
}

// RiskValue of a local minimum is 1 + its value
func (field Field) RiskValue(minimum Point) int {
	return field[minimum[0]][minimum[1]] + 1
}

// GetBasinSize computes the size of a basin around a local minimum recursively
func (field Field) GetBasinSize(minimum Point) int {
	checked := map[Point]bool{}

	return field.recurseBasin(checked, minimum)
}

func (field Field) recurseBasin(checked map[Point]bool, p Point) int {
	if checked[p] {
		return 0
	}

	checked[p] = true

	i, j := p[0], p[1]
	val := field[i][j]

	count := 1 // to count the current point

	if i != 0 && field[i-1][j] >= val && field[i-1][j] != 9 {
		count += field.recurseBasin(checked, Point{i - 1, j})
	}

	if i != 99 && field[i+1][j] >= val && field[i+1][j] != 9 {
		count += field.recurseBasin(checked, Point{i + 1, j})
	}

	if j != 0 && field[i][j-1] >= val && field[i][j-1] != 9 {
		count += field.recurseBasin(checked, Point{i, j - 1})
	}

	if j != 99 && field[i][j+1] >= val && field[i][j+1] != 9 {
		count += field.recurseBasin(checked, Point{i, j + 1})
	}

	return count
}
