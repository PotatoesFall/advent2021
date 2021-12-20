package main

import (
	"bytes"
	"os"
	"regexp"
	"strconv"
)

func readInput() [nScanners]*Scanner {
	input, err := os.ReadFile(`input19`)
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(input, []byte{'\n'})

	scanners := [nScanners]*Scanner{}
	for i := 0; i < nScanners; i++ {
		lines, scanners[i] = scanScanner(lines[1:])
	}

	return scanners
}

func scanScanner(lines [][]byte) ([][]byte, *Scanner) {
	scanner := new(Scanner)

	for len(lines) != 0 && len(lines[0]) != 0 {
		scanner.Beacons = append(scanner.Beacons, scanPosition(lines[0]))
		lines = lines[1:]
	}

	if len(lines) == 0 {
		return nil, scanner
	}

	return lines[1:], scanner
}

var lineRegex = regexp.MustCompile(`^(\-?\d+),(\-?\d+),(\-?\d+)$`)

func scanPosition(line []byte) Position {
	matches := lineRegex.FindSubmatch(line)
	if len(matches) != 4 {
		panic(string(line))
	}

	return Position{
		X: mustAtoi(matches[1]),
		Y: mustAtoi(matches[2]),
		Z: mustAtoi(matches[3]),
	}
}

func mustAtoi(str []byte) int {
	i, err := strconv.Atoi(string(str))
	if err != nil {
		panic(err)
	}

	return i
}
