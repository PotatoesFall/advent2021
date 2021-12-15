package main

import (
	"fmt"
	"io"
)

// Point is a point in a Cavern
type Point struct {
	row, col int
}

func (p Point) Neighbors(height, width int) []Point {
	points := [4]Point{
		{p.row + 1, p.col},
		{p.row, p.col + 1},
		{p.row - 1, p.col},
		{p.row, p.col - 1},
	}

	filtered := make([]Point, 0, 4)
	for _, point := range points {
		if point.row >= 0 && point.row < height && point.col >= 0 && point.col < width {
			filtered = append(filtered, point)
		}
	}

	return filtered
}

// Cavern is a grid of danger values from 1 to 9
// its logic requires it to be rectangular
type Cavern [][]uint8

// DangerAt returns the value at a point in the cavern.
func (c Cavern) DangerAt(p Point) int {
	return int(c[p.row][p.col])
}

// Height returns the height of the cavern
func (c Cavern) Height() int {
	return len(c)
}

// Width returns the width of the cavern
func (c Cavern) Width() int {
	return len(c[0])
}

func (c Cavern) Copy() Cavern {
	out := make(Cavern, len(c))

	for i := range c {
		out[i] = make([]uint8, len(c[i]))
		copy(out[i], c[i])
	}

	return out
}

func (c Cavern) Fprint(w io.Writer) {
	for _, line := range c {
		for _, danger := range line {
			fmt.Fprint(w, danger)
		}

		fmt.Fprintln(w)
	}
}

// Path is a list of points that are traversed in order
type Path []Point

func (p Path) Danger(cavern Cavern) int {
	danger := 0

	for _, point := range p[1:] {
		danger += cavern.DangerAt(point)
	}

	return danger
}

// PathNode caches the Danger so it doesn't need to be recalculated
type PathNode struct {
	Path
	Danger int
}

func (pn PathNode) Copy() PathNode {
	pathNode := PathNode{
		Path:   make(Path, len(pn.Path)),
		Danger: pn.Danger,
	}

	for i, point := range pn.Path {
		pathNode.Path[i] = point
	}

	return pathNode
}

func (pn PathNode) Move(point Point, danger int) PathNode {
	pn.Path = append(pn.Path, point)
	pn.Danger += danger

	return pn
}

type Explorer map[Point]PathNode

func NewExplorer(start Point) Explorer {
	return Explorer{start: {Path: Path{start}}}
}

func (e Explorer) Explore(cavern Cavern) int {
	madeProgress := true

	rounds := 0
	deadPoints := map[Point]bool{}

	for madeProgress && rounds < maxRounds {
		cleaned := 0
		madeProgress = false

		for point, pathNode := range e {
			if deadPoints[point] {
				continue
			}

			pointAlive := false
			for _, neighbor := range point.Neighbors(cavern.Height(), cavern.Width()) {
				alreadyFound, foundAlready := e[neighbor]

				if !foundAlready || pathNode.Danger+cavern.DangerAt(neighbor) < alreadyFound.Danger {
					pointAlive, madeProgress = true, true
					e[neighbor] = pathNode.Copy().Move(neighbor, cavern.DangerAt(neighbor))
					deadPoints[neighbor] = false
				}
			}

			if !pointAlive {
				if deadPoints[point] {
					panic(point)
				}
				cleaned++
				deadPoints[point] = true
			}
		}

		fmt.Printf("Round: %d; Filled: %d; Cleaned: %d;\n", rounds, len(e), cleaned)
		rounds++
	}

	return rounds
}
