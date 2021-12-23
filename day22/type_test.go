package main

import (
	"testing"

	"git.fuyu.moe/Fuyu/assert"
)

func TestSize(t *testing.T) {
	assert := assert.New(t)

	assert.Eq(int64(8), Cuboid{
		X: NewRange(0, 2),
		Y: NewRange(0, 2),
		Z: NewRange(0, 2),
	}.Size())

	assert.Eq(int64(27), Cuboid{
		X: NewRange(-1, 2),
		Y: NewRange(-1, 2),
		Z: NewRange(-1, 2),
	}.Size())

	assert.Eq(int64(0), Cuboid{
		X: NewRange(1, 1),
		Y: NewRange(0, 2),
		Z: NewRange(0, 2),
	}.Size())
}
