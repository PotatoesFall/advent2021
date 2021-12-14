package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile(`input14`)
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(input, []byte{'\n'})

	polymer := map[[2]byte]int64{}
	polymerLine := lines[0]
	start, end := polymerLine[0], polymerLine[len(polymerLine)-1]
	for i := 0; i < len(polymerLine)-1; i++ {
		key := [2]byte{polymerLine[i], polymerLine[i+1]}
		polymer[key]++
	}

	rulesLines := lines[2:]

	rules := map[[2]byte]byte{}
	for _, rule := range rulesLines {
		var key [2]byte
		copy(key[:], rule[:2])

		rules[key] = rule[6]
	}

	for i := 0; i < 40; i++ {
		polymer = polymerize(polymer, rules)
		fmt.Println(countLetters(polymer, start, end))
	}

	counts := countLetters(polymer, start, end)

	var largestLetter, smallestLetter byte
	var largest, smallest int64 = 0, 1<<63 - 1
	for letter, count := range counts {
		fmt.Println(string(letter), count)
		if count > largest {
			largestLetter = letter
			largest = count
		}

		if count < smallest {
			smallestLetter = letter
			smallest = count
		}
	}

	fmt.Println(largest, smallest, largest-smallest)
	fmt.Println(string(largestLetter), string(smallestLetter))
}

func polymerize(polymer map[[2]byte]int64, rules map[[2]byte]byte) map[[2]byte]int64 {
	added := map[[2]byte]int64{}
	removed := map[[2]byte]int64{}

	for pair, count := range polymer {
		inBetween := rules[pair]

		added[[2]byte{pair[0], inBetween}] += count
		added[[2]byte{inBetween, pair[1]}] += count

		removed[pair] += count
	}

	for addPair, count := range added {
		polymer[addPair] += count
	}

	for removePair, count := range removed {
		polymer[removePair] -= count

		if polymer[removePair] == 0 {
			delete(polymer, removePair)
		}
	}

	return polymer
}

func countLetters(polymer map[[2]byte]int64, start, end byte) map[byte]int64 {
	counts := map[byte]int64{}

	for pair, count := range polymer {
		counts[pair[0]] += count
		counts[pair[1]] += count
	}

	for letter, count := range counts {
		if letter == start || letter == end {
			counts[letter] = (count + 1) / 2
			continue
		}

		counts[letter] = count / 2
	}

	return counts
}
