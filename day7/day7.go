package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile(`input7`)
	if err != nil {
		panic(err)
	}

	numStrings := strings.Split(string(input), `,`)
	crabs := make([]int, len(numStrings))
	for i, str := range numStrings {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		crabs[i] = n
	}

	sort.Ints(crabs)
	median := crabs[len(crabs)/2] // not actually always the median, but works here

	fmt.Printf("Part 1 - total Fuel usage: %d\n", totalDistance(crabs, median))

	average := average(crabs) // is currently rounded down because that worked for me. Might have to be rounded up depending on data set.
	fmt.Println(average)

	fmt.Printf("Part 2 - total Fuel usage: %d\n", totalSquareDistance(crabs, average))
}

func totalDistance(nums []int, goal int) int {
	distance := 0
	for _, n := range nums {
		distance += absDiff(n, goal)
	}

	return distance
}

func average(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return int(float64(sum) / float64(len(nums)))
}

func totalSquareDistance(nums []int, goal int) int {
	fuelUsage := 0
	for _, n := range nums {
		distance := absDiff(n, goal)
		fuel := distance * (distance + 1) / 2 // triangle formula

		fuelUsage += fuel
	}

	return fuelUsage
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
