package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.ReadFile(`input1`)
	if err != nil {
		panic(err)
	}

	fmt.Printf("count: %d\n", countIncreases(input))
	fmt.Printf("count with triplets: %d\n", countTripletIncreases(input))
}

func countIncreases(input []byte) int {
	numbers := bytes.Split(input, []byte("\n"))

	count := 0
	prev := 1 << 30
	for _, number := range numbers {
		n, err := strconv.Atoi(string(number))
		if err != nil {
			fmt.Printf("skipping %q\n", string(number))
			continue
		}

		if n > prev {
			count++
		}

		prev = n
	}

	return count
}

func countTripletIncreases(input []byte) int {
	numbers := bytes.Split(input, []byte("\n"))

	nums := make([]int, len(input))
	for i, number := range numbers {
		n, err := strconv.Atoi(string(number))
		if err != nil {
			fmt.Printf("skipping %q\n", string(number))
			continue
		}

		nums[i] = n
	}

	count := 0
	for i := 3; i < len(nums); i++ {
		if nums[i] > nums[i-3] {
			count++
		}
	}

	return count
}
