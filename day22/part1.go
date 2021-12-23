package main

import "fmt"

func part1(steps []RebootStep) {
	range50 := NewRange(-50, 50)
	initArea := Cuboid{range50, range50, range50}

	pointsOn := map[Point]bool{}

	for _, step := range steps {
		if !initArea.Contains(step.Cuboid) {
			break
		}

		for _, point := range step.Points() {
			pointsOn[point] = step.On
		}
	}

	count := 0
	for _, on := range pointsOn {
		if on {
			count++
		}
	}

	fmt.Printf("Part 1 - %d are on.\n", count)
}
