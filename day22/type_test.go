package main

import (
	"testing"

	"git.fuyu.moe/Fuyu/assert"
)

func TestSize(t *testing.T) {
	assert := assert.New(t)

	assert.Eq(int64(8), Cuboid{
		X: NewRange(0, 1),
		Y: NewRange(0, 1),
		Z: NewRange(0, 1),
	}.Size())

	assert.Eq(int64(27), Cuboid{
		X: NewRange(-1, 1),
		Y: NewRange(-1, 1),
		Z: NewRange(-1, 1),
	}.Size())

	assert.Eq(int64(0), Cuboid{
		X: NewRange(1, 0),
		Y: NewRange(0, 1),
		Z: NewRange(0, 1),
	}.Size())
}
