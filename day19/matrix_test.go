package main

import (
	"testing"

	"git.fuyu.moe/Fuyu/assert"
)

func TestGetMatrix(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		name      string
		base, vec Vector
		expected  Matrix
	}{
		{
			name: `simple`,
			base: Vector{1, 2, 3},
			vec:  Vector{1, 2, 3},
			expected: Matrix{
				[3]int{1, 0, 0},
				[3]int{0, 1, 0},
				[3]int{0, 0, 1},
			},
		},
		{
			name: `rotate`,
			base: Vector{1, 2, 3},
			vec:  Vector{2, 3, 1},
			expected: Matrix{
				[3]int{0, 0, 1},
				[3]int{1, 0, 0},
				[3]int{0, 1, 0},
			},
		},
		{
			name: `flip`,
			base: Vector{1, 2, 3},
			vec:  Vector{-1, -2, -3},
			expected: Matrix{
				[3]int{-1, 0, 0},
				[3]int{0, -1, 0},
				[3]int{0, 0, -1},
			},
		},
		{
			name: `madness`,
			base: Vector{1, -2, -3},
			vec:  Vector{3, 1, -2},
			expected: Matrix{
				[3]int{0, 1, 0},
				[3]int{0, 0, 1},
				[3]int{-1, 0, 0},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Eq(c.expected, getMatrix(c.base, c.vec))
			assert.Eq(Position(c.base), c.expected.Transform(Position(c.vec)))
		})
	}
}
