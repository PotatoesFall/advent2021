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

	aligned := map[int]bool{
		0: true,
	}

	notDone := true
	for len(aligned) != nScanners && notDone {
		notDone = false
	outer:
		for i, unalignedScanner := range scanners {
			if aligned[i] {
				continue
			}

			// for j, alignedScanner := range scanners {
			// 	if !aligned[j] {
			// 		continue
			// 	}
			// fmt.Printf("Trying to get scanner %d from %d:\n", i, j)

			alignedScanner := &Scanner{}
			for j, scanner := range scanners {
				if !aligned[j] {
					continue
				}
				alignedScanner.Beacons = append(alignedScanner.Beacons, scanner.Beacons...)
				computeDiffs(alignedScanner)
			}

			overlaps := countVectorOverlaps(alignedScanner, unalignedScanner)
			if overlaps >= requiredVectorOverlap {
				notDone = true

				// if i != j {
				// 	fmt.Printf("Getting scanner %d from %d: %d\n", i, j, overlaps)
				// }

				align(alignedScanner, unalignedScanner)
				if countPointOverlaps(alignedScanner, unalignedScanner) >= requiredOverlap {
					aligned[i] = true
					continue outer
				}
			}
			// }
		}
	}

	points := map[Position]bool{}
	for i, scanner := range scanners {
		if !aligned[i] {
			continue
		}

		for _, beacon := range scanner.Beacons {
			points[beacon] = true
		}
	}

	fmt.Println(len(points))
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

func align(base, scanner *Scanner) {
outer:
	for _, baseDiff := range base.Diffs {
		for _, diff := range scanner.Diffs {
			if baseDiff.Normalized == diff.Normalized && diff.AllDistinct() {
				if tryAlign(base, scanner, baseDiff, diff) {
					doAlign(base, scanner, baseDiff, diff)
					break outer
				}

				if tryAlign(base, scanner, baseDiff.Invert(), diff) {
					doAlign(base, scanner, baseDiff.Invert(), diff)
					break outer
				}

				if tryAlign(base, scanner, baseDiff, diff.Invert()) {
					doAlign(base, scanner, baseDiff, diff.Invert())
					break outer
				}

				if tryAlign(base, scanner, baseDiff.Invert(), diff.Invert()) {
					doAlign(base, scanner, baseDiff.Invert(), diff.Invert())
					break outer
				}
			}
		}
	}
}

func tryAlign(base, scanner *Scanner, baseDiff, diff DiffVector) bool {
	scanner = scanner.Clone()

	doAlign(base, scanner, baseDiff, diff)

	// fmt.Println(scanner.Position)
	// sort.Slice(base.Beacons, func(i, j int) bool {
	// 	return base.Beacons[i].X < base.Beacons[j].X
	// })
	// sort.Slice(scanner.Beacons, func(i, j int) bool {
	// 	return scanner.Beacons[i].X < scanner.Beacons[j].X
	// })
	// fmt.Println(base.Beacons)
	// fmt.Println(scanner.Beacons)
	// panic(nil)

	if countPointOverlaps(base, scanner) >= requiredOverlap {
		// fmt.Println("worked, count:", countPointOverlaps(base, scanner))
		return true
	}

	// fmt.Println("didn't work, count:", countPointOverlaps(base, scanner))
	return false
}

func doAlign(base, scanner *Scanner, baseDiff, diff DiffVector) {
	matrix := getMatrix(baseDiff.Vector, diff.Vector)

	for i, beacon := range scanner.Beacons {
		scanner.Beacons[i] = matrix.Transform(beacon)
	}

	computeDiffs(scanner)

	// fmt.Println("basePos:", base.Beacons[baseDiff.ParentBeaconA])
	// fmt.Println("newPos:", scanner.Beacons[diff.ParentBeaconA])

	scanner.Position = Position(computeDiff(
		scanner.Beacons[diff.ParentBeaconA],
		base.Beacons[baseDiff.ParentBeaconA],
	))

	// fmt.Println(`POSITION:`, scanner.Position)

	for i, beacon := range scanner.Beacons {
		scanner.Beacons[i] = Vector(scanner.Position).Translate(beacon)
	}
}

func countPointOverlaps(a, b *Scanner) int {
	count := 0

	for _, aBeacon := range a.Beacons {
		for _, bBeacon := range b.Beacons {
			if aBeacon == bBeacon {
				count++
			}
		}
	}

	return count
}
