package main

import (
	"testing"

	"git.fuyu.moe/Fuyu/assert"
)

func TestExpand(t *testing.T) {
	assert := assert.New(t)

	cavern := Cavern{{1}}
	cavern = expandCaveByFactor(cavern, 5)
	expected := Cavern{
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 6},
		{3, 4, 5, 6, 7},
		{4, 5, 6, 7, 8},
		{5, 6, 7, 8, 9},
	}
	assert.Cmp(expected, cavern)

	cavern = readInput(`input15_test`)
	cavern = expandCaveByFactor(cavern, 5)
	expected = readInput(`expected15`)
	assert.Cmp(expected, cavern)
}

func TestExplore(t *testing.T) {
	assert := assert.New(t)

	cavern := readInput(`input15_test`)

	path := findSafestPath(cavern, Point{0, 0}, Point{cavern.Height() - 1, cavern.Width() - 1})
	assert.Eq(40, path.Danger(cavern))

	cavern = expandCaveByFactor(cavern, 5)
	path = findSafestPath(cavern, Point{0, 0}, Point{cavern.Height() - 1, cavern.Width() - 1})
	assert.Eq(315, path.Danger(cavern))
}
