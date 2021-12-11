package main

import (
	"bytes"
	"fmt"
	"strings"
)

// Octopi is a group of octopi
type Octopi [size][size]octopus

// NewOctopi parses the input intos Octopi
func NewOctopi(input []byte) *Octopi {
	var octopi Octopi

	for i, line := range bytes.Split(input, []byte{'\n'}) {
		for j, numChar := range line {
			octopi[i][j] = octopus(numChar - '0')
		}
	}

	return &octopi
}

// Step does one step and returns the number of flashes in that step
func (o *Octopi) Step() int {
	// step 1 - increase all energy levels
	o.increaseAll()

	// step 2 - flash
	flashCount := newFlasher(o).flash()

	// step 3 - reset
	o.resetFlashedTo0()

	return flashCount
}

func (o *Octopi) points() []point {
	points := make([]point, 0, 100)

	for row, line := range o {
		for col := range line {
			points = append(points, point{row, col})
		}
	}

	return points
}

func (o *Octopi) get(p point) octopus {
	return o[p.row][p.col]
}

func (o *Octopi) increase(p point) {
	o[p.row][p.col]++
}

func (o *Octopi) increaseAll() {
	for _, p := range o.points() {
		o.increase(p)
	}
}

func (o *Octopi) resetFlashedTo0() {
	for _, point := range o.points() {
		if o.get(point).energy() > maxEnergy {
			o.reset(point)
		}
	}
}

func (o *Octopi) reset(p point) {
	o[p.row][p.col] = 0
}

func (o *Octopi) String() string {
	str := &strings.Builder{}

	for _, line := range o {
		for _, octopus := range line {
			fmt.Fprint(str, octopus)
		}
		str.WriteByte('\n')
	}

	return str.String()
}

type octopus int

func (o octopus) energy() int {
	return int(o)
}

func (o octopus) flashes() bool {
	return o.energy() > maxEnergy
}

type point struct {
	row, col int
}

func (p point) inBounds() bool {
	return p.row < size && p.row >= 0 && p.col < size && p.col >= 0
}

func (p point) adjacents() []point {
	all := [8]point{
		{p.row - 1, p.col + 0}, // above
		{p.row - 1, p.col - 1}, // top left diagonal
		{p.row + 0, p.col - 1}, // left
		{p.row + 1, p.col - 1}, // bottom left diagonal
		{p.row + 1, p.col + 0}, // bottom
		{p.row + 1, p.col + 1}, // bottom right diagonal
		{p.row + 0, p.col + 1}, // right
		{p.row - 1, p.col + 1}, // top right diagonal
	}

	var points []point
	for _, p := range all {
		if p.inBounds() {
			points = append(points, p)
		}
	}

	return points
}
