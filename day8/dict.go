package main

var digitsByCount = map[int][]Digit{
	2: {1},
	3: {7},
	4: {4},
	5: {2, 3, 5},
	6: {0, 6, 9},
	7: {8},
}

var possibleTrueSegmentsByCount = map[int][]TrueSegment{
	2: {2, 5},
	3: {0, 2, 5},
	4: {1, 2, 3, 5},
	5: {0, 2, 3, 4, 6, 5, 1},
	6: {0, 1, 2, 4, 5, 6, 3},
	7: {0, 1, 2, 3, 4, 5, 6},
}

// // top to bottom, left to right
// var possibleDigitsByTrueSegmentNumber = map[SegmentNumber][]Digit{
// 	0: {0, 2, 3, 5, 6, 7, 8, 9},
// 	1: {0, 4, 5, 6, 8, 9},
// 	2: {0, 1, 2, 3, 4, 7, 8, 9},
// 	3: {2, 3, 4, 5, 6, 8, 9},
// 	4: {0, 2, 6, 8},
// 	5: {0, 1, 3, 4, 5, 6, 7, 8, 9},
// 	6: {0, 2, 3, 5, 6, 8, 9},
// }

var trueSegmentsByDigit = map[Digit][]TrueSegment{
	0: {0, 1, 2, 4, 5, 6},
	1: {2, 5},
	2: {0, 2, 3, 4, 6},
	3: {0, 2, 3, 5, 6},
	4: {1, 2, 3, 5},
	5: {0, 1, 3, 5, 6},
	6: {0, 1, 3, 4, 5, 6},
	7: {0, 2, 5},
	8: {0, 1, 2, 3, 4, 5, 6},
	9: {0, 1, 2, 3, 5, 6},
}

// var segmentCountByNumber = map[Digit]int{
// 	0: 6,
// 	1: 2,
// 	2: 5,
// 	3: 5,
// 	4: 4,
// 	5: 5,
// 	6: 6,
// 	7: 3,
// 	8: 7,
// 	9: 6,
// }

var segmentByLabel = map[SegmentLabel]Segment{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
}

var segmentLabels = map[Segment]SegmentLabel{
	0: 'a',
	1: 'b',
	2: 'c',
	3: 'd',
	4: 'e',
	5: 'f',
	6: 'g',
}

var (
	allSegments     = [7]Segment{0, 1, 2, 3, 4, 5, 6}
	allTrueSegments = [7]TrueSegment{0, 1, 2, 3, 4, 5, 6}
)
