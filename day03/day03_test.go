package main

import (
	"testing"

	"git.fuyu.moe/Fuyu/assert"
)

func TestFindLastMatch(t *testing.T) {
	assert := assert.New(t)

	data := Lines{
		newLine([]byte(`001000010000`)),
		newLine([]byte(`111101111000`)),
		newLine([]byte(`101101011000`)),
		newLine([]byte(`101111011100`)),
		newLine([]byte(`101011010100`)),
		newLine([]byte(`011110111100`)),
		newLine([]byte(`001110011100`)),
		newLine([]byte(`111001110000`)),
		newLine([]byte(`100001000000`)),
		newLine([]byte(`110011100100`)),
		newLine([]byte(`000100001000`)),
		newLine([]byte(`010100101000`)),
	}

	actual := data.FindLastMatch(true)
	expected := newLine([]byte(`101111011100`))
	assert.Eq(expected.String(), actual.String())

	actual = data.FindLastMatch(false)
	expected = newLine([]byte(`010100101000`))
	assert.Eq(expected.String(), actual.String())
}
