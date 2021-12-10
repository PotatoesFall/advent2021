package main

import (
	"testing"

	"git.fuyu.moe/Fuyu/assert"
)

var cases1 = []struct {
	line     string
	expected int
}{
	{`{([(<{}[<>[]}>{[]{[(<()>`, 1197},
	{`[[<[([]))<([[{}[[()]]]`, 3},
	{`[{[{({}]{}}([{[{{{}}([]`, 57},
	{`<{([([[(<>()){}]>(<<{{`, 25137},
}

var cases2 = []struct {
	line     string
	expected int
}{
	{`()`, 0},
	{`(`, 1},
	{`<`, 4},
	{`((`, 6},
	{`[({([[{{`, 288957},
	{`(((([{}[]{}()]<{<{{()()()`, 1480781},
}

func Test(t *testing.T) {
	assert := assert.New(t)

	for _, c := range cases1 {
		line := Line(c.line)
		parser := NewParser(line)
		assert.Eq(c.expected, parser.Scores().Corruption)
	}

	for _, c := range cases2 {
		line := Line(c.line)
		parser := NewParser(line)
		assert.Eq(c.expected, parser.Scores().Completion)
	}
}
