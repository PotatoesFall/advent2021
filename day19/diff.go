package main

func computeAllDiffs(scanners [nScanners]*Scanner) {
	for _, scanner := range scanners {
		computeDiffs(scanner)
	}
}

func computeDiffs(scanner *Scanner) {
	scanner.Diffs = []DiffVector{}

	for a, beacon1 := range scanner.Beacons {
		for b, beacon2 := range scanner.Beacons {
			if a == b {
				continue
			}

			diff := DiffVector{
				Parent:        scanner,
				ParentBeaconA: a,
				ParentBeaconB: b,
			}

			diff.Vector = computeDiff(beacon1, beacon2)
			diff.Normalized = diff.Vector.Normalize()

			scanner.Diffs = append(scanner.Diffs, diff)
		}
	}
}

func computeDiff(beacon1, beacon2 Position) Vector {
	return Vector{
		beacon2.X - beacon1.X,
		beacon2.Y - beacon1.Y,
		beacon2.Z - beacon1.Z,
	}
}
