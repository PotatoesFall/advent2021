package main

import (
	"bytes"
	"fmt"
	"os"
)

const (
	maxRounds = 1000 // should be enough assuming no upwards or leftwards movement
)

func main() {
	cavern := readInput(`input15`)
	path := findSafestPath(cavern, Point{0, 0}, Point{cavern.Height() - 1, cavern.Width() - 1})
	fmt.Println("Part 1 -", path.Danger(cavern))

	cavern = expandCaveByFactor(cavern, 5)
	path = findSafestPath(cavern, Point{0, 0}, Point{cavern.Height() - 1, cavern.Width() - 1})
	fmt.Println("Part 2 -", path.Danger(cavern))
}

func readInput(path string) Cavern {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	cavern := Cavern{}
	for row, line := range bytes.Split(input, []byte{'\n'}) {
		cavern = append(cavern, []uint8{})
		for _, v := range line {
			cavern[row] = append(cavern[row], v-'0')
		}
	}

	return cavern
}

func findSafestPath(cavern Cavern, from, to Point) Path {
	explorer := NewExplorer(from)

	explorer.Explore(cavern)

	return explorer[to].Path
}

func expandCaveByFactor(cavern Cavern, factor int) Cavern { //nolint:unparam
	original := cavern.Copy()

	for i := 0; i < factor-1; i++ {
		for rowN, line := range original {
			for _, danger := range line {
				cavern[rowN] = append(cavern[rowN], increment(danger, i+1))
			}
		}
	}

	topRow := cavern.Copy()

	for i := 0; i < factor-1; i++ {
		for _, line := range topRow {
			newRow := make([]uint8, len(line))
			for j, danger := range line {
				newRow[j] = increment(danger, i+1)
			}

			cavern = append(cavern, newRow)
		}
	}

	return cavern
}

func increment(danger uint8, times int) uint8 {
	for i := 0; i < times; i++ {
		danger += 1

		if danger == 10 {
			danger = 1
		}
	}

	return danger
}

// func diagonal(round int, height, width int) []Point {
// 	points := []Point{}

// 	for row := round; row >= 0; row-- {
// 		col := round - row
// 		point := Point{row, col}

// 		if point.row < height && point.col < width {
// 			points = append(points, point)
// 		}
// 	}

// 	return points
// }
