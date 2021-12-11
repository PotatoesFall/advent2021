package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.ReadFile(`input4`)
	if err != nil {
		panic(err)
	}

	numbersDrawn, boards := parseInput(input)

	winner, loser := getWinnerAndLoser(boards, numbersDrawn)

	fmt.Printf("Winning board scores %d\n", winner.score)
	fmt.Printf("Losing board scores %d\n", loser.score)
}

func getWinnerAndLoser(boards []*Board, numbersDrawn []int) (*Board, *Board) {
	winner, loser := boards[0], boards[0]
	winCount, loseCount := 0, 0
	for _, b := range boards {
		b.Play(numbersDrawn)
		// fmt.Printf("board %2d won in round %2d\n", i, b.winningRound)

		if b.winningRound == winner.winningRound {
			winCount++
		}
		if b.winningRound < winner.winningRound {
			winCount = 1
			winner = b
		}

		if b.winningRound == loser.winningRound {
			loseCount++
		}
		if b.winningRound > loser.winningRound {
			loseCount = 1
			loser = b
		}
	}

	if winCount != 1 || loseCount != 1 {
		panic(fmt.Sprintf(`cannot determine winner: %d winners`, winCount))
	}

	return winner, loser
}

func parseInput(input []byte) ([]int, []*Board) {
	lines := bytes.Split(input, []byte{'\n'})

	numbersDrawnBytes := bytes.Split(lines[0], []byte{','})
	numbersDrawn := make([]int, 0, len(numbersDrawnBytes))
	for _, numString := range numbersDrawnBytes {
		n, err := strconv.Atoi(string(numString))
		if err != nil {
			panic(err)
		}
		numbersDrawn = append(numbersDrawn, n)
	}

	var boards []*Board
	for startLine := 2; startLine < len(lines)-4; startLine += 6 {
		board := parseBoard(lines[startLine : startLine+5])

		boards = append(boards, board)
	}

	return numbersDrawn, boards
}
