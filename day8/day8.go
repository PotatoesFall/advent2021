package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile(`input8`)
	if err != nil {
		panic(err)
	}

	strLines := strings.Split(string(input), "\n")
	lines := make([]Line, len(strLines))
	for i, strLine := range strLines {
		lines[i] = NewLine(strLine)
	}

	totalObviousCount := 0
	for _, l := range lines {
		totalObviousCount += countPart1(l)
	}

	fmt.Printf("Part 1 - obvious count is %d\n", totalObviousCount)

	totalCount := 0
	for _, l := range lines {
		totalCount += countPart2(l)
	}

	fmt.Printf("Part 2 - full count is %d\n", totalCount)
}

func countPart1(l Line) int {
	count := 0

	for _, seg := range l.output {
		segCount := seg.Count()
		if segCount == 2 || segCount == 3 || segCount == 4 || segCount == 7 {
			count++
		}
	}

	return count
}

func countPart2(l Line) int {
	solver := NewSolver()

	for _, box := range l.AllBoxes() {
		actualSegments := box.Segments()
		possibleTrueSegments := possibleTrueSegmentsByCount[box.Count()]

		solver.RestrictAllTo(actualSegments, possibleTrueSegments)

		notActualSegments := invertSegments(actualSegments)
		solver.RestrictAllNegative(notActualSegments, possibleTrueSegments)

	}

	return -1
}

// func count(l Line) int {
// 	// initialize map with all values
// 	possibleTrueSegNums := map[SegmentNumber][]SegmentNumber{}
// 	for _, segN := range allSegNums {
// 		possibleTrueSegNums[segN] = allSegNums[:]
// 	}

// 	// get all counts per segment
// 	countsPerSeg := getSegmentCountsPerSegNum(l)

// 	// remove all trueSegs that are not possible with all of those counts
// 	maskTrueSegNumsByCounts(l, possibleTrueSegNums, countsPerSeg)

// 	// for each output, try and get a result
// 	outputNumber := 0
// 	for _, segs := range l.output {
// 		outputNumber *= 10
// 		outputNumber += int(getDigit(segs, possibleTrueSegNums))
// 	}

// 	return outputNumber
// }

// func getDigit(segs Box, possibleTrueSegNums map[SegmentNumber][]SegmentNumber) Digit {
// 	possibleDigits := digitsBySegmentCount[segs.Count()]

// 	if len(possibleDigits) == 1 {
// 		return possibleDigits[0]
// 	}

// 	foundDigit := Digit(-1)
// 	count := 0

// 	// try to find digit by going through possible ones
// outer:
// 	for _, possibleDigit := range possibleDigits {

// 		// check that all necessary segments are possible
// 		for _, needTrueSegN := range trueSegmentNumbersByDigit[possibleDigit] {

// 			found := false

// 			// go through current segments to see if we have the true segment
// 			for segN, on := range segs {
// 				if !on {
// 					continue
// 				}

// 				possibleTrue := possibleTrueSegNums[SegmentNumber(segN)]

// 				// if we have it, we are good
// 				for _, possibleSeg := range possibleTrue {
// 					if needTrueSegN == possibleSeg {
// 						found = true
// 						break
// 					}
// 				}

// 			}
// 			if !found {
// 				continue outer
// 			}

// 			break
// 		}

// 		count++
// 		foundDigit = possibleDigit
// 		// break
// 	}

// 	if foundDigit == -1 {
// 		panic("oof")
// 	}

// 	if count != 1 {
// 		panic("oh no")
// 	}

// 	return foundDigit
// }

// func maskTrueSegNumsByCounts(l Line, possibleTrueSegNums map[SegmentNumber][]SegmentNumber, countsPerSeg map[SegmentNumber]map[int]bool) {
// 	for segN, counts := range countsPerSeg {
// 		for count, seen := range counts { // assume no false entries
// 			if !seen {
// 				continue
// 			}

// 			possibleTrueSegs, ok := possibleTrueSegmentsByCount[count]
// 			if !ok {
// 				panic(`aaaah`)
// 			}

// 			possibleTrueSegNums[segN] = maskSegments(possibleTrueSegNums[segN], possibleTrueSegs)
// 		}
// 	}
// }

// func getSegmentCountsPerSegNum(l Line) map[SegmentNumber]map[int]bool {
// 	segmentCounts := map[SegmentNumber]map[int]bool{}

// 	for _, segments := range l.AllSegments() {
// 		for i, on := range segments {
// 			if !on {
// 				continue
// 			}
// 			segN := SegmentNumber(i)

// 			_, ok := segmentCounts[segN]
// 			if !ok {
// 				segmentCounts[segN] = map[int]bool{}
// 			}

// 			segmentCounts[segN][segments.Count()] = true

// 		}
// 	}

// 	return segmentCounts
// }
