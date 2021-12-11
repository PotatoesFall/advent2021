package main

// School is a school of lanternfish
type School [9]int

// Tick models the growth during one day
func (s *School) Tick() {
	s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8] = s[1], s[2], s[3], s[4], s[5], s[6], s[7]+s[0], s[8], s[0]
}

// Num returns the number of fish
func (s *School) Num() int {
	sum := 0
	for _, n := range s {
		sum += n
	}

	return sum
}
