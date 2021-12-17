package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	input := readInput(`input16`)

	_, packet := readPacket(input)

	fmt.Println("Part 1 -", addVersions(packet))
	fmt.Println("Part 2 -", evaluate(packet))
}

func readInput(path string) []bool {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	input, err := io.ReadAll(hex.NewDecoder(file))
	if err != nil {
		panic(err)
	}

	bools := make([]bool, len(input)*8)
	for i, byt := range input {
		for j := 0; j < 8; j++ {
			bools[i*8+j] = byt&(1<<(7-j)) != 0
		}
	}

	return bools
}

func addVersions(packet Packet) int {
	count := packet.Version

	for _, child := range packet.Children {
		count += addVersions(child)
	}

	return count
}

func evaluate(packet Packet) int {
	switch packet.Type {
	case TypeIDLiteral:
		return packet.Literal

	case TypeIDSum:
		return aggregate(packet.Children, func(agg int, v int) int {
			return agg + v
		}, 0)

	case TypeIDProduct:
		return aggregate(packet.Children, func(agg int, v int) int {
			return agg * v
		}, 1)

	case TypeIDMinimum:
		return aggregate(packet.Children, func(agg int, v int) int {
			if v < agg {
				return v
			}
			return agg
		}, 1<<31-1)

	case TypeIDMaximum:
		return aggregate(packet.Children, func(agg int, v int) int {
			if v > agg {
				return v
			}
			return agg
		}, 0)

	case TypeIDGreaterThan:
		return compare(packet.Children, func(v1, v2 int) bool {
			return v1 > v2
		})

	case TypeIDLessThan:
		return compare(packet.Children, func(v1, v2 int) bool {
			return v1 < v2
		})

	case TypeIDEqualTo:
		return compare(packet.Children, func(v1, v2 int) bool {
			return v1 == v2
		})
	}

	panic(packet.Type)
}

func aggregate(packets []Packet, aggFunc func(agg int, v int) int, agg int) int {
	for _, p := range packets {
		agg = aggFunc(agg, evaluate(p))
	}

	return agg
}

func compare(packets []Packet, compareFunc func(int, int) bool) int {
	if compareFunc(evaluate(packets[0]), evaluate(packets[1])) {
		return 1
	}
	return 0
}
