package main

// Solver collects information, and narrows down possible connections
type Solver map[Segment]map[TrueSegment]bool

// NewSolver populates a new solver with all possible combinations
func NewSolver() Solver {
	s := Solver{}

	for _, segNum := range allSegments {
		segmentsPossible := make(map[TrueSegment]bool)

		for _, segNum := range allTrueSegments {
			segmentsPossible[segNum] = true
		}

		s[segNum] = segmentsPossible
	}

	return s
}

func (s Solver) RestrictTo(segmentToRestrict Segment, restrictTo map[TrueSegment]bool) {
	for segNum, _ := range s[segmentToRestrict] {
		if !restrictTo[segNum] {
			s[segmentToRestrict][segNum] = false
		}
	}
}

func (s Solver) RestrictAllTo(segmentsToRestrict []Segment, restrictTo map[TrueSegment]bool) {
	for _, segment := range segmentsToRestrict {
		s.RestrictTo(segment, restrictTo)
	}
}

func (s Solver) RestrictNegative(segmentsToRestrict Segment, cannotBe map[TrueSegment]bool) {
	for segNum, _ := range s[segmentsToRestrict] {
		if cannotBe[segNum] {
			s[segmentsToRestrict][segNum] = false
		}
	}
}

func (s Solver) RestrictAllNegative(segmentsToRestrict []Segment, cannotBe map[TrueSegment]bool) {
	for _, segment := range segmentsToRestrict {
		s.RestrictNegative(segment, cannotBe)
	}
}

func (s Solver) Inverted() {
}

func (s Solver) GetDigit(box Box) Digit {
	possibleDigits := digitsByCount[box.Count()]

	if len(possibleDigits) == 1 {
		return possibleDigits[0]
	}

	// shit I think i have to invert
	// todo
}
