package main

import (
	"testing"

	"git.fuyu.moe/Fuyu/assert"
)

func TestDay5(t *testing.T) {
	assert := assert.New(t)

	data := []byte(
		"0,9 -> 5,9\n" +
			"8,0 -> 0,8\n" +
			"9,4 -> 3,4\n" +
			"2,2 -> 2,1\n" +
			"7,0 -> 7,4\n" +
			"6,4 -> 2,0\n" +
			"0,9 -> 2,9\n" +
			"3,4 -> 1,4\n" +
			"0,0 -> 8,8\n" +
			"5,5 -> 8,2",
	)

	lines := parseVentLines(data)
	assert.Eq(10, len(lines))
	expectedLine := Line{start: Point{0, 9}, end: Point{5, 9}}
	assert.Eq(expectedLine, lines[0])

	lines1 := lines.OnlyHorizontalOrVertical()
	assert.Eq(6, len(lines1))

	field := lines1.Field()
	expectedField := Field([][]int{
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 1, 1, 2, 1, 1, 1, 2, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
	})
	assert.Cmp(expectedField, field)

	count := field.CountPointsWithMin(2)
	assert.Eq(5, count)

	// part2
	field = lines.Field()
	expectedField = Field([][]int{
		{1, 0, 1, 0, 0, 0, 0, 1, 1, 0},
		{0, 1, 1, 1, 0, 0, 0, 2, 0, 0},
		{0, 0, 2, 0, 1, 0, 1, 1, 1, 0},
		{0, 0, 0, 1, 0, 2, 0, 2, 0, 0},
		{0, 1, 1, 2, 3, 1, 3, 2, 1, 1},
		{0, 0, 0, 1, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
	})
	assert.Cmp(expectedField, field)

	count = field.CountPointsWithMin(2)
	assert.Eq(12, count)
}
