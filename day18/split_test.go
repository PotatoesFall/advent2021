package main

import (
	"testing"

	"git.fuyu.moe/Fuyu/assert"
)

func TestSplit(t *testing.T) {
	assert := assert.New(t)

	_, pair := parsePair([]byte("[[1,2],[10,8]]"))
	assert.True(splitOne(pair))
	assert.Eq(pair.String(), "[[1,2],[[5,5],8]]")

	_, pair = parsePair([]byte("[[1,2],[11,8]]"))
	assert.True(splitOne(pair))
	assert.Eq(pair.String(), "[[1,2],[[5,6],8]]")

	_, pair = parsePair([]byte("[[1,2],[9,8]]"))
	assert.False(splitOne(pair))
	assert.Eq(pair.String(), "[[1,2],[9,8]]")
}
