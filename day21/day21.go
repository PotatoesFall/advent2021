package main

import "fmt"

const (
	nPlayers      = 2
	rollsPerRound = 3
	boardSize     = 10
	dieSize       = 100
	winScore      = 1_000
)

func main() {
	players := [nPlayers]Player{
		{Position: 8},
		{Position: 3},
	}

	part1(players)

	part2(players)
}

func part1(players [nPlayers]Player) {
	die := NewDeterministicDie(dieSize)

	round := 0
	for round < 1 || players[(round-1)%2].Score < winScore {
		players[round%2].Move(die.RollN(rollsPerRound))
		round++
	}

	fmt.Printf("Part 1 - %d\n", players[(round)%2].Score*round*rollsPerRound)
}

func part2(players [nPlayers]Player) {
	// TODO
}

type Player struct {
	Position int
	Score    int
}

func (p *Player) Move(steps int) {
	p.Position = (p.Position+steps-1)%boardSize + 1
	p.Score += p.Position
}

type DeterministicDie struct {
	next  int
	limit int
}

func NewDeterministicDie(size int) Die {
	return &DeterministicDie{
		next:  1,
		limit: size,
	}
}

func (dd *DeterministicDie) Roll() int {
	roll := dd.next

	dd.next++
	if dd.next > dd.limit {
		dd.next = 1
	}

	return roll
}

func (dd *DeterministicDie) RollN(n int) int {
	sum := 0

	for i := 0; i < n; i++ {
		sum += dd.Roll()
	}

	return sum
}

type Die interface {
	Roll() int
	RollN(n int) int
}
