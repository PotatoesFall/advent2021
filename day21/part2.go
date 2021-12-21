package main

import (
	"fmt"
	"sort"
)

type Scenario struct {
	Players  [2]Player
	LastTurn int
}

func (s Scenario) Roll(roll int) Scenario {
	s.Players[(s.LastTurn)%2].Move(roll)
	s.LastTurn++
	return s
}

func (s Scenario) Winner(winScore int) (int, bool) {
	if s.Players[0].Score >= winScore {
		return 0, true
	}
	if s.Players[1].Score >= winScore {
		return 1, true
	}

	return -1, false
}

func part2(players [2]Player) {
	wins := playScenarios(Scenario{
		Players:  players,
		LastTurn: 0,
	})

	sort.Slice(wins[:], func(i, j int) bool {
		return wins[i] > wins[j]
	})

	fmt.Printf("Part 2 - Winner wins in %d universes, Loser in %d universes\n", wins[0], wins[1])
}

func playScenarios(s Scenario) [2]int64 {
	scenarios := map[Scenario]int64{s: 1}

	wins := [2]int64{0, 0}

	makingProgress := true
	for makingProgress {
		makingProgress = false

		for scenario, universes := range scenarios {
			if universes == 0 {
				continue
			}

			winner, won := scenario.Winner(21)
			if won {
				wins[winner] += universes
				scenarios[scenario] = 0
				continue
			}

			makingProgress = true

			for roll, factor := range universesForRoll {
				scenarios[scenario.Roll(roll)] += int64(factor) * universes
			}

			scenarios[scenario] = 0
		}
	}

	return wins
}

var universesForRoll = map[int]int{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}
