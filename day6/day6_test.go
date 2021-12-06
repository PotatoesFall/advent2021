package main

import (
	"testing"

	"git.fuyu.moe/Fuyu/assert"
)

func TestDay6(t *testing.T) {
	assert := assert.New(t)

	school := School([9]int{0, 1, 1, 2, 1, 0, 0, 0, 0})

	school.Tick()
	expected := School([9]int{1, 1, 2, 1, 0, 0, 0, 0, 0})
	assert.Cmp(expected, school)

	school.Tick()
	expected = School([9]int{1, 2, 1, 0, 0, 0, 1, 0, 1})
	assert.Cmp(expected, school)
}
