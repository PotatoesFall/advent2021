package main

import "bytes"

func parseInput(input []byte) (Polymer, Rules) {
	lines := bytes.Split(input, []byte{'\n'})

	polymer := parsePolymer(lines[0])
	rules := parseRules(lines[2:])

	return polymer, rules
}

func parseRules(lines [][]byte) Rules {
	rules := make(Rules)

	for _, rule := range lines {
		var key [2]byte
		copy(key[:], rule[:2])

		rules[key] = rule[6]
	}

	return rules
}

func parsePolymer(line []byte) Polymer {
	polymer := make(Polymer)

	for i := 0; i < len(line)-1; i++ {
		pair := [2]byte{line[i], line[i+1]}
		polymer[pair]++
	}

	return polymer
}
