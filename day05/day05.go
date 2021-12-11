package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile(`input5`)
	if err != nil {
		panic(err)
	}

	lines := parseVentLines(input)

	count1 := lines.OnlyHorizontalOrVertical().Field().CountPointsWithMin(2)
	count2 := lines.Field().CountPointsWithMin(2)
	fmt.Printf("Part 1 - there are %d points with 2 or more vents.\n", count1)
	fmt.Printf("Part 2 - there are %d points with 2 or more vents.\n", count2)
}
