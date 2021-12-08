package main

// func joinSegNumSlices(a, b []SegmentNumber) []SegmentNumber {
// 	newSlice := []SegmentNumber{}
// 	all := make([]SegmentNumber, len(a)+len(b))
// 	copy(all, a)
// 	copy(all[len(a):], b)

// outer:
// 	for _, val := range all {
// 		for _, v := range newSlice {
// 			if v == val {
// 				continue outer
// 			}
// 		}

// 		newSlice = append(newSlice, val)
// 	}

// 	return newSlice
// }

// func maskSegments(a, b []SegmentNumber) []SegmentNumber {
// 	newSlice := []SegmentNumber{}
// 	for _, v := range a {
// 		for _, vb := range b {
// 			if v == vb {
// 				newSlice = append(newSlice, v)
// 			}
// 		}
// 	}

// 	return newSlice
// }

func containsSegment(segments []Segment, segment Segment) bool {
	for _, seg := range segments {
		if seg == segment {
			return true
		}
	}

	return false
}

func containsTrueSegment(segments []TrueSegment, segment TrueSegment) bool {
	for _, seg := range segments {
		if seg == segment {
			return true
		}
	}

	return false
}

func invertSegments(toInvert []Segment) []Segment {
	inverted := []Segment{}

outer:
	for _, segment := range allSegments {
		for _, segToInvert := range toInvert {
			if segment == segToInvert {
				continue outer
			}
		}

		inverted = append(inverted, segment)
	}

	return inverted
}

func invertTrueSegments(toInvert []TrueSegment) []TrueSegment {
	inverted := []TrueSegment{}

outer:
	for _, segment := range allTrueSegments {
		for _, segToInvert := range toInvert {
			if segment == segToInvert {
				continue outer
			}
		}

		inverted = append(inverted, segment)
	}

	return inverted
}
