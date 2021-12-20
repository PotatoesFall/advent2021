package main

import (
	"fmt"
	"testing"

	"git.fuyu.moe/Fuyu/assert"
)

func TestNormalize(t *testing.T) {
	assert := assert.New(t)

	cases := [][2]Vector{
		{Vector{1, 2, 3}, Vector{3, 2, 1}},
		{Vector{1, -3, 3}, Vector{3, -3, 1}},
		{Vector{-1, -1, 1}, Vector{1, 1, -1}},
		{Vector{-3, -2, 1}, Vector{3, 2, -1}},
		{Vector{-3, -3, -3}, Vector{3, 3, 3}},
		{Vector{3, 3, 3}, Vector{3, 3, 3}},
		{Vector{3, -3, -3}, Vector{3, 3, -3}},
		{Vector{-3, 3, -3}, Vector{3, 3, -3}},
		{Vector{-3, -3, 3}, Vector{3, 3, -3}},
		{Vector{-3, 3, 3}, Vector{3, 3, -3}},
		{Vector{3, -3, 3}, Vector{3, 3, -3}},
		{Vector{3, 3, -3}, Vector{3, 3, -3}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c[1]), func(t *testing.T) {
			assert.Eq(c[1], c[0].Normalize())
		})
	}
}
