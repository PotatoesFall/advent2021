package main

import (
	"testing"

	"git.fuyu.moe/Fuyu/assert"
)

func TestDay4(t *testing.T) {
	assert := assert.New(t)

	numbersDrawn := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	boards := []*Board{
		parseBoard([][]byte{
			[]byte(`22 13 17 11  0`),
			[]byte(` 8  2 23  4 24`),
			[]byte(`21  9 14 16  7`),
			[]byte(` 6 10  3 18  5`),
			[]byte(` 1 12 20 15 19`),
		}),

		parseBoard([][]byte{
			[]byte(` 3 15  0  2 22`),
			[]byte(` 9 18 13 17  5`),
			[]byte(`19  8  7 25 23`),
			[]byte(`20 11 10 24  4`),
			[]byte(`14 21 16 12  6`),
		}),

		parseBoard([][]byte{
			[]byte(`14 21 17 24  4`),
			[]byte(`10 16 15  9 19`),
			[]byte(`18  8 23 26 20`),
			[]byte(`22 11 13  6  5`),
			[]byte(` 2  0 12  3  7`),
		}),
	}

	winner, _ := getWinnerAndLoser(boards, numbersDrawn)

	assert.Eq(12, winner.winningRound)
	assert.Eq(188, winner.UnmarkedSum())
}
