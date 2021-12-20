package main

func alignAllScanners(scanners Scanners) {
	aligned := map[int]bool{0: true}

	notDone := true
	for len(aligned) != nScanners && notDone {
		notDone = false
	outer:
		for i, unalignedScanner := range scanners {
			if aligned[i] {
				continue
			}

			for j, alignedScanner := range scanners {
				if !aligned[j] {
					continue
				}

				overlaps := countVectorOverlaps(alignedScanner, unalignedScanner)
				if overlaps >= requiredVectorOverlap {
					alignScanners(alignedScanner, unalignedScanner)

					notDone = true
					aligned[i] = true
					continue outer
				}
			}
		}
	}
}

func alignScanners(base, scanner *Scanner) {
outer:
	for _, baseDiff := range base.Diffs {
		for _, diff := range scanner.Diffs {
			if baseDiff.Normalized == diff.Normalized && diff.AllDistinct() {
				if tryAlign(base, scanner, baseDiff, diff) {
					align(base, scanner, baseDiff, diff)
					break outer
				}

				if tryAlign(base, scanner, baseDiff.Invert(), diff) {
					align(base, scanner, baseDiff.Invert(), diff)
					break outer
				}

				if tryAlign(base, scanner, baseDiff, diff.Invert()) {
					align(base, scanner, baseDiff, diff.Invert())
					break outer
				}

				if tryAlign(base, scanner, baseDiff.Invert(), diff.Invert()) {
					align(base, scanner, baseDiff.Invert(), diff.Invert())
					break outer
				}
			}
		}
	}
}

func tryAlign(base, scanner *Scanner, baseDiff, diff DiffVector) bool {
	scanner = scanner.Clone()

	align(base, scanner, baseDiff, diff)

	return countPointOverlaps(base, scanner) >= requiredOverlap
}

func align(base, scanner *Scanner, baseDiff, diff DiffVector) {
	matrix := getMatrix(baseDiff.Vector, diff.Vector)

	for i, beacon := range scanner.Beacons {
		scanner.Beacons[i] = matrix.Transform(beacon)
	}

	computeDiffs(scanner)

	scanner.Position = Position(computeDiff(
		scanner.Beacons[diff.ParentBeaconA],
		base.Beacons[baseDiff.ParentBeaconA],
	))

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
