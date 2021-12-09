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
	count := 0

	for _, box := range l.output {
		count *= 10
		count += int(getDigit(l, box))
	}

	return count
}

func getDigit(l Line, box Box) Digit {
	box1 := l.BoxWithLength(2)
	box4 := l.BoxWithLength(4)

	switch box.Count() {
	case 2:
		return Digit(1)

	case 3:
		return Digit(7)

	case 4:
		return Digit(4)

	case 5:
		// 3 has full overlap with 1
		if countBoxOverlap(box, box1) == 2 {
			return 3
		}

		// 5 has overlap 3 with 4
		if countBoxOverlap(box, box4) == 3 {
			return 5
		}

		// 2 has overlap 2 with 4
		if countBoxOverlap(box, box4) == 2 {
			return 2
		}

	case 6:
		// 6 has overlap 1 with 1
		if countBoxOverlap(box, box1) == 1 {
			return 6
		}

		// 9 has overlap 4 with 4
		if countBoxOverlap(box, box4) == 4 {
			return 9
		}

		// 0 has overlap 3 with 4
		if countBoxOverlap(box, box4) == 3 {
			return 0
		}

	case 7:
		return Digit(8)
	}

	panic(box)
}

func countBoxOverlap(a, b Box) int {
	count := 0

	for _, segmentA := range a.Segments() {
		for _, segmentB := range b.Segments() {
			if segmentA == segmentB {
				count++
				break
			}
		}
	}

	return count
}
