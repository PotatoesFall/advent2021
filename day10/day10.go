package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	input, err := os.ReadFile(`input10`)
	if err != nil {
		panic(err)
	}

	lines := parse(input)

	part1, part2 := computeScores(lines)

	fmt.Printf("Part 1 - corruption sum is %d\n", part1)

	fmt.Printf("Part 2 - completion middle score is %d\n", part2)
}

func computeScores(lines []Line) (int, int) {
	corruptionSum := 0
	var completionScores []int

	for _, line := range lines {
		scores := NewParser(line).Scores()

		corruptionSum += scores.Corruption

		if scores.Completion != 0 {
			completionScores = append(completionScores, scores.Completion)
		}
	}

	middleScore := median(completionScores)

	return corruptionSum, middleScore
}

func median(ints []int) int {
	sort.Ints(ints)
	return ints[len(ints)/2]
}

func parse(input []byte) []Line {
	split := bytes.Split(input, []byte{'\n'})

	lines := make([]Line, len(split))
	for i := range split {
		lines[i] = split[i]
	}

	return lines
}
