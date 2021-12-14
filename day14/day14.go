package main

import (
	"fmt"
	"os"
	"sort"
)

const rounds = 40

func main() {
	input, err := os.ReadFile(`input14`)
	if err != nil {
		panic(err)
	}

	polymer, rules := parseInput(input)

	for i := 0; i < rounds; i++ {
		polymer = polymerize(polymer, rules)
	}

	counts := letterCounts(polymer)
	sortInt64(counts)

	fmt.Println(counts[len(counts)-1] - counts[0])
}

func polymerize(polymer Polymer, rules Rules) Polymer {
	diff := make(Polymer)

	for pair, count := range polymer {
		between := rules[pair]

		diff[pair] -= count

		diff[Pair{pair[0], between}] += count
		diff[Pair{between, pair[1]}] += count
	}

	for pair, count := range diff {
		polymer[pair] += count

		if polymer[pair] == 0 {
			delete(polymer, pair)
		}
	}

	return polymer
}

func letterCounts(polymer Polymer) []int64 {
	counts := map[byte]int64{}

	for pair, count := range polymer {
		counts[pair[0]] += count
		counts[pair[1]] += count
	}

	for letter, count := range counts {
		// start and end only counted once rather than double, will be odd
		counts[letter] = (count + 1) / 2
	}

	countSlice := make([]int64, 0, len(counts))
	for _, count := range counts {
		countSlice = append(countSlice, count)
	}

	return countSlice
}

func sortInt64(slice []int64) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}
