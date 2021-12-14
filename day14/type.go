package main

// Pair represents two adjacent letters
type Pair [2]byte

// Polymer counts the number of times each pair exists
type Polymer map[Pair]int64

// Rules contains the letters to squeeze between each Pair
type Rules map[Pair]byte
