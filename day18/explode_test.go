package main

import (
	"testing"

	"git.fuyu.moe/Fuyu/assert"
)

func TestExplode(t *testing.T) {
	data := [][]byte{
		[]byte("[[[[[9,8],1],2],3],4]"),
		[]byte("[7,[6,[5,[4,[3,2]]]]]"),
		[]byte("[[6,[5,[4,[3,2]]]],1]"),
		[]byte("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"),
		[]byte("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"),
	}

	expected := [][]byte{
		[]byte("[[[[0,9],2],3],4]"),
		[]byte("[7,[6,[5,[7,0]]]]"),
		[]byte("[[6,[5,[7,0]]],3]"),
		[]byte("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"),
		[]byte("[[3,[2,[8,0]]],[9,[5,[7,0]]]]"),
	}

	for i, line := range data {
		t.Run(string(line), func(t *testing.T) {
			assert := assert.New(t)

			_, pair := parsePair(line)

			explode(getExplodable(pair))

			assert.Eq(string(expected[i]), pair.String())
		})
	}
}
