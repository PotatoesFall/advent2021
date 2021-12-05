package main

import (
	"strconv"
	"strings"
)

// Board represents a bingo board
type Board struct {
	board               [5][5]int
	checked             [5][5]bool
	winningRound, score int
}

func parseBoard(input [][]byte) *Board {
	var board Board
	for i, l := range input {
		str := string(l)

		for j := 0; j < 5; j++ {
			s := strings.TrimSpace(str[j*3 : j*3+2])
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}

			board.board[i][j] = n
		}
	}

	return &board
}

// Check checks all matching numbers
func (b *Board) Check(n int) {
	for i, line := range b.board {
		for j, v := range line {
			if v == n {
				b.checked[i][j] = true
			}
		}
	}
}

// Won checks if the board has won
func (b Board) Won() bool {
	if b.score != 0 || b.winningRound != 0 {
		return true
	}

	for i := 0; i < 5; i++ {
		// horizontal
		if b.checked[i] == [5]bool{true, true, true, true, true} {
			return true
		}

		// vertical
		wonColumn := true
		for j := 0; j < 5; j++ {
			if !b.checked[j][i] {
				wonColumn = false
			}
		}
		if wonColumn {
			return true
		}
	}

	return false
}

// UnmarkedSum calculates the score of the board
func (b Board) UnmarkedSum() int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.checked[i][j] {
				sum += b.board[i][j]
			}
		}
	}

	return sum
}

// Play plays a game with the given numbers.
// if winningRound == 0, the board didn't win
func (b *Board) Play(numbersDrawn []int) {
	for i, n := range numbersDrawn {
		b.Check(n)

		if b.Won() {
			b.winningRound = i + 1
			b.score = n * b.UnmarkedSum()
			return
		}
	}
}
