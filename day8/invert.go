package main

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

// func invertSegmentMap(toInvert map[Segment]bool) map[Segment]bool {
// 	inverted := map[Segment]bool{}

// 	for _, segment := range allSegments {
// 		if toInvert[segment] {
// 			inverted[segment] = false
// 		} else {
// 			inverted[segment] = true
// 		}
// 	}

// 	return inverted
// }

// func invertTrueSegmentMap(toInvert map[TrueSegment]bool) map[TrueSegment]bool {
// 	inverted := map[TrueSegment]bool{}

// 	for _, segment := range allTrueSegments {
// 		if toInvert[segment] {
// 			inverted[segment] = false
// 		} else {
// 			inverted[segment] = true
// 		}
// 	}

// 	return inverted
// }
