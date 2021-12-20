package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile(`input18`)
	if err != nil {
		panic(err)
	}

	part1(input)

	part2(input)
}

func part1(input []byte) {
	pairs := parsePairs(input)

	sum := addAll(pairs)

	fmt.Println(sum.Magnitude())
}

func part2(input []byte) {
	pairs := parsePairs(input)

	highest := 0
	for _, pair1 := range pairs {
		for _, pair2 := range pairs {
			mag := add(pair1.Clone(), pair2.Clone()).Magnitude()
			if mag > highest {
				highest = mag
			}
		}
	}

	fmt.Println(highest)
}
