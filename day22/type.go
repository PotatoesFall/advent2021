package main

const (
	X = iota
	Y
	Z
)

var dimensions = [3]int{X, Y, Z}

type Range struct {
	Min, Max int64
	Invalid  bool
}

func NewRange(min, max int64) Range {
	if max > min {
		return Range{Min: min, Max: max}
	}

	return Range{Invalid: true}
}

func (r Range) All() []int64 {
	if r.Invalid {
		return nil
	}

	n := make([]int64, 0, r.Max-r.Min+1)

	for i := r.Min; i < r.Max; i++ {
		n = append(n, i)
	}

	return n
}

func (r Range) Has(v int64) bool {
	if r.Invalid {
		return false
	}

	return v >= r.Min && v < r.Max
}

func (r Range) Contains(rang Range) bool {
	if r.Invalid || rang.Invalid {
		return false
	}

	return rang.Min >= r.Min && rang.Max <= r.Max
}

func (r Range) Intersects(rang Range) bool {
	if r.Invalid || rang.Invalid {
		return false
	}

	return rang.Max > r.Min && rang.Min < r.Max
}

func (r Range) Size() int64 {
	if r.Invalid {
		return 0
	}

	return r.Max - r.Min
}

type Cuboid [3]Range

func (c Cuboid) Contains(cuboid Cuboid) bool {
	return c[X].Contains(cuboid[X]) && c[Y].Contains(cuboid[Y]) && c[Z].Contains(cuboid[Z])
}

func (c Cuboid) Intersects(cuboid Cuboid) bool {
	return c[X].Intersects(cuboid[X]) && c[Y].Intersects(cuboid[Y]) && c[Z].Intersects(cuboid[Z])
}

func (c Cuboid) Points() []Point {
	var points []Point

	for _, x := range c[X].All() {
		for _, y := range c[Y].All() {
			for _, z := range c[Z].All() {
				points = append(points, Point{x, y, z})
			}
		}
	}

	return points
}

func (c Cuboid) Size() int64 {
	var size int64 = 1

	for i := range c {
		size *= c[i].Size()
	}

	return size
}

type RebootStep struct {
	Cuboid

	On bool
}

type Set map[Cuboid]struct{}

func (s Set) Put(c Cuboid) {
	if c.Size() == 0 {
		return
	}

	s[c] = struct{}{}
}

func (s Set) Has(c Cuboid) bool {
	_, ok := s[c]
	return ok
}

func (s Set) Delete(c Cuboid) {
	delete(s, c)
}

type Point [3]int64
