package main

import (
	"bytes"
	"strconv"
)

func parsePairs(input []byte) []*Pair {
	lines := bytes.Split(input, []byte{'\n'})
	pairs := make([]*Pair, len(lines))

	for i, line := range lines {
		_, pairs[i] = parsePair(line)
	}

	return pairs
}

func parsePair(input []byte) ([]byte, *Pair) {
	var pair Pair

	input = stripFirstChar(input, '[')

	input = parseValue(input, &pair, X)
	input = stripFirstChar(input, ',')

	input = parseValue(input, &pair, Y)
	input = stripFirstChar(input, ']')

	return input, &pair
}

func parseValue(input []byte, dst *Pair, side Side) []byte {
	if input[0] != '[' {
		return parseIntInto(input, dst, side)
	}

	return parsePairInto(input, dst, side)
}

func parsePairInto(input []byte, dst *Pair, side Side) []byte {
	var pair *Pair
	input, pair = parsePair(input)

	dst.SetChild(side, pair)

	return input
}

func parseIntInto(input []byte, dst *Pair, side Side) []byte {
	var v int
	input, v = parseInt(input)

	dst.SetValue(side, v)

	return input
}

func parseInt(input []byte) ([]byte, int) {
	var end int

	for i := 0; i < len(input); i++ {
		if input[i] < '0' || input[i] > '9' {
			end = i
			break
		}
	}

	v, err := strconv.Atoi(string(input[:end]))
	if err != nil {
		panic(string(input))
	}

	return input[end:], v
}

func stripFirstChar(input []byte, char byte) []byte {
	if input[0] != char {
		panic(string(input))
	}

	return input[1:]
}
