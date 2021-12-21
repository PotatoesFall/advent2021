package main

import "fmt"

func part1(players [2]Player) {
	die := DeterministicDie{next: 1, limit: 100}

	round := 0
	for round < 1 || players[(round-1)%2].Score < 1_000 {
		players[round%2].Move(die.RollN(3))
		round++
	}

	fmt.Printf("Part 1 - %d\n", players[(round)%2].Score*round*3)
}

type DeterministicDie struct {
	next  int
	limit int
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
