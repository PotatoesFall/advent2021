package main

import (
	"fmt"
)

const (
	nScanners = 26
)

var (
	requiredOverlap       = 12
	requiredVectorOverlap = requiredOverlap * (requiredOverlap - 1) * 2
)

func main() {
	scanners := readInput()

	computeAllDiffs(scanners)

	alignAllScanners(scanners)

	fmt.Printf("Part 1 - %d\n", countPoints(scanners))

	fmt.Printf("Part 2 - %d\n", maxManhattanDistance(scanners))
}

func countPoints(scanners Scanners) int {
	points := map[Position]bool{}

	for _, scanner := range scanners {
		for _, beacon := range scanner.Beacons {
			points[beacon] = true
		}
	}

	return len(points)
}

func countVectorOverlaps(a, b *Scanner) int {
	count := 0

	for _, aDiff := range a.Diffs {
		for _, bDiff := range b.Diffs {
			if aDiff.Normalized == bDiff.Normalized {
				count++
			}
		}
	}

	return count
}

func maxManhattanDistance(scanners Scanners) int {
	maxDistance := 0

	for _, scanner1 := range scanners {
		for _, scanner2 := range scanners {
			distance := computeDiff(scanner1.Position, scanner2.Position).Manhattan()

			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}

	return maxDistance
}
