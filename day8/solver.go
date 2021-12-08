package main

import "fmt"

type Link struct {
	From Segment
	To   TrueSegment
}

// Solver collects information, and narrows down possible connections
type Solver map[Link]bool

// NewSolver populates a new solver with all possible combinations
func NewSolver() Solver {
	s := Solver{}

	for _, trueSegment := range allTrueSegments {
		for _, segment := range allSegments {
			s[Link{segment, trueSegment}] = true
		}
	}

	return s
}

func (s Solver) Invalidate(segment Segment, trueSegment TrueSegment) {
	s[Link{segment, trueSegment}] = false
}

func (s Solver) InvalidateAll(segment Segment, trueSegments []TrueSegment) {
	for _, trueSegment := range trueSegments {
		s.Invalidate(segment, trueSegment)
	}
}

func (s Solver) InvalidateAllBackwards(trueSegment TrueSegment, segments []Segment) {
	for _, segment := range segments {
		s.Invalidate(segment, trueSegment)
	}
}

func (s Solver) InvalidateRoundRobin(segments []Segment, trueSegments []TrueSegment) {
	for _, segment := range segments {
		for _, trueSegment := range trueSegments {
			s.Invalidate(segment, trueSegment)
		}
	}
}

func (s Solver) GetDigit(box Box) Digit {
	possibleDigits := digitsByCount[box.Count()]

	// if len(possibleDigits) == 1 {
	// 	return possibleDigits[0]
	// }

	var foundDigit Digit
	count := 0
	for _, digit := range possibleDigits {
		if s.CheckDigit(box, digit) {
			foundDigit = digit
			count++
		}
	}

	// if count != 1 {
	// 	panic(count)
	// }

	fmt.Println(foundDigit)

	return foundDigit
}

func (s Solver) CheckDigit(box Box, digit Digit) bool {
	trueSegments := trueSegmentsByDigit[digit]

outer:
	for _, trueSegment := range trueSegments {
		for _, segment := range box.Segments() {
			if s[Link{segment, trueSegment}] {
				continue outer
			}
		}

		return false
	}

	return true
}

func (s Solver) PossibleTrueSegments(segment Segment) []TrueSegment {
	possibleTrueSegments := []TrueSegment{}

	for link, possible := range s {
		if possible && link.From == segment {
			possibleTrueSegments = append(possibleTrueSegments, link.To)
		}
	}

	return possibleTrueSegments
}

func (s Solver) PossibleSegments(trueSegment TrueSegment) []Segment {
	possibleSegments := []Segment{}

	for link, possible := range s {
		if possible && link.To == trueSegment {
			possibleSegments = append(possibleSegments, link.From)
		}
	}

	return possibleSegments
}

// func (s Solver) BackwardsRestrictTo(toRestrict TrueSegment, restrictTo []Segment) {
// 	for trueSegment, _ := range s.backwards[toRestrict] {
// 		if !containsSegment(restrictTo, trueSegment) {
// 			s.backwards[toRestrict][trueSegment] = false
// 		}
// 	}
// }

// func (s Solver) BackwardsRestrictAllTo(toRestrict []TrueSegment, restrictTo []Segment) {
// 	for _, segment := range toRestrict {
// 		s.BackwardsRestrictTo(segment, restrictTo)
// 	}
// }

// func (s Solver) RestrictNegative(segmentsToRestrict Segment, cannotBe map[TrueSegment]bool) {
// 	for segNum, _ := range s.forwards[segmentsToRestrict] {
// 		if cannotBe[segNum] {
// 			s.forwards[segmentsToRestrict][segNum] = false
// 		}
// 	}
// }

// func (s Solver) RestrictAllNegative(segmentsToRestrict []Segment, cannotBe map[TrueSegment]bool) {
// 	for _, segment := range segmentsToRestrict {
// 		s.RestrictNegative(segment, cannotBe)
// 	}
// }

// func (s Solver) Inverted() {
// }
