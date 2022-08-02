package main

import (
	"fmt"
	"testing"

	"git.ultraware.nl/martin/assert"
)

func TestCompress(t *testing.T) {
	state := State{
		Amphipods: [16]Amphipod{{B, 1}, {A, 2}, {A, 3}, {A, 4}, {A, 5}, {A, 6}, {A, 21}, {A, 12}, {A, 13}, {A, 14}, {A, 15}, {A, 16}, {A, 17}, {A, 18}, {A, 19}, {A, 20}},
	}
	fmt.Println(state)
	compressed := state.Compress()
	assert.New(t)
	expected := CompressedState{0x21, 0x11, 0x11, 0x00, 0x00, 0x01, 0x11, 0x11, 0x11, 0x11, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	fmt.Println(expected)
	fmt.Println(compressed)
	assert.Eq(t, expected, compressed)
}
