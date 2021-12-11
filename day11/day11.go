package main

import (
	"fmt"
	"os"
)

const (
	size      = 10
	maxEnergy = 9
	rounds    = 100
)

func main() {
	input, err := os.ReadFile(`input11`)
	if err != nil {
		panic(err)
	}

	part1(input)

	part2(input)
}

func part1(input []byte) {
	octopi := NewOctopi(input)
	flashCount := 0

	for i := 0; i < rounds; i++ {
		flashCount += octopi.Step()
	}

	fmt.Printf("Part 1 - In 100 rounds there are %d flashes\n", flashCount)
}

func part2(input []byte) {
	octopi := NewOctopi(input)

	round := 0
	for {
		round++
		if octopi.Step() == 100 {
			break
		}
	}

	fmt.Printf("Part 2 - All of them flash in step %d\n", round)
}
