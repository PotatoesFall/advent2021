package main

import (
	"sort"
)

type Scanner struct {
	Position Position
	Beacons  []Position
	Diffs    []DiffVector
}

type Scanners [26]*Scanner

type Position struct {
	X, Y, Z int
}

type Vector struct {
	X, Y, Z int
}

func (v Vector) Manhattan() int {
	return abs(v.X) + abs(v.Y) + abs(v.Z)
}

func (v Vector) Invert() Vector {
	return Vector{-v.X, -v.Y, -v.Z}
}

func (v Vector) Normalize() Vector {
	values := []int{v.X, v.Y, v.Z}

	sort.Slice(values, func(i, j int) bool {
		if abs(values[i]) == abs(values[j]) {
			return values[i] > values[j]
		}

		return abs(values[i]) > abs(values[j])
	})

	return Vector{abs(values[0]), abs(values[1]), abs(values[2])}
}

func (v Vector) AllDistinct() bool {
	return abs(v.X) != abs(v.Y) && abs(v.X) != abs(v.Z) && abs(v.Y) != abs(v.Z) && v.X != 0 && v.Y != 0 && v.Z != 0
}

func (v Vector) Translate(p Position) Position {
	return Position{
		p.X + v.X,
		p.Y + v.Y,
		p.Z + v.Z,
	}
}

func (s *Scanner) Clone() *Scanner {
	scanner := &Scanner{
		Position: s.Position,
		Beacons:  make([]Position, len(s.Beacons)),
		Diffs:    make([]DiffVector, len(s.Diffs)),
	}

	for i, beacon := range s.Beacons {
		scanner.Beacons[i] = beacon
	}

	for i, diff := range s.Diffs {
		scanner.Diffs[i] = diff
		scanner.Diffs[i].Parent = scanner
	}

	return scanner
}

type DiffVector struct {
	Vector
	Normalized                   Vector
	Parent                       *Scanner
	ParentBeaconA, ParentBeaconB int // index in parent's Beacons slice
}

func (d DiffVector) Invert() DiffVector {
	return DiffVector{
		Vector:        d.Vector.Invert(),
		Normalized:    d.Normalized,
		Parent:        d.Parent,
		ParentBeaconA: d.ParentBeaconB,
		ParentBeaconB: d.ParentBeaconA,
	}
}

type Matrix [3][3]int

func (m Matrix) Transform(v Position) Position {
	return Position{
		X: m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z,
		Y: m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z,
		Z: m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z,
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}
