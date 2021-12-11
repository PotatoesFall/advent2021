package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	data, err := os.ReadFile(`input9`)
	if err != nil {
		panic(err)
	}

	field := NewField(data)

	fmt.Printf("Part 1 - Sum of all Risk Values is %d\n", part1(field))

	fmt.Printf("Part 2 - Product of three largest basin sizes is %d\n", part2(field))
}

func part1(field Field) int {
	minima := field.GetLocalMinima()

	riskSum := 0
	for _, minimum := range minima {
		riskSum += field.RiskValue(minimum)
	}

	return riskSum
}

func part2(field Field) int {
	minima := field.GetLocalMinima()

	basinSizes := make([]int, len(minima))
	for i, minimum := range minima {
		basinSizes[i] = field.GetBasinSize(minimum)
	}

	sort.Ints(basinSizes)
	l := len(basinSizes)

	return basinSizes[l-1] * basinSizes[l-2] * basinSizes[l-3]
}
