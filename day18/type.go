package main

import "fmt"

type Side int

const (
	X Side = iota
	Y
)

var bothSides = [2]Side{X, Y}

func (s Side) Opposite() Side {
	switch s {
	case X:
		return Y
	case Y:
		return X
	}

	panic(s)
}

type Pair struct {
	childX, childY *Pair
	x, y           int

	parent     *Pair
	parentSide Side
}

func (p Pair) HasChild(side Side) bool {
	switch side {
	case X:
		return p.childX != nil
	case Y:
		return p.childY != nil
	}

	return false
}

func (p Pair) HasValue(side Side) bool {
	return !p.HasChild(side)
}

func (p Pair) Child(side Side) *Pair {
	switch side {
	case X:
		return p.childX
	case Y:
		return p.childY
	}

	return nil
}

func (p Pair) Value(side Side) int {
	switch side {
	case X:
		return p.x
	case Y:
		return p.y
	}

	return 0
}

func (p Pair) String() string {
	xStr, yStr := fmt.Sprint(p.x), fmt.Sprint(p.y)

	if p.childX != nil {
		xStr = p.childX.String()
	}
	if p.childY != nil {
		yStr = p.childY.String()
	}

	return fmt.Sprintf(`[%s,%s]`, xStr, yStr)
}

func (p *Pair) SetChild(side Side, child *Pair) {
	if child != nil {
		child.parent = p
		child.parentSide = side
	}

	switch side {
	case X:
		p.childX = child
	case Y:
		p.childY = child
	}
}

func (p *Pair) SetValue(side Side, val int) {
	switch side {
	case X:
		p.x = val
	case Y:
		p.y = val
	}
}

func (p Pair) Magnitude() int {
	return 3*p.magnitude(X) + 2*p.magnitude(Y)
}

func (p Pair) magnitude(side Side) int {
	if p.HasChild(side) {
		return p.Child(side).Magnitude()
	}

	return p.Value(side)
}

func (p *Pair) Clone() *Pair {
	pair := &Pair{}
	*pair = *p

	for _, side := range bothSides {
		if p.HasChild(side) {
			pair.SetChild(side, p.Child(side).Clone())
		}
	}

	return pair
}
