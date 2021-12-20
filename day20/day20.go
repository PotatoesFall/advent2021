package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile(`input20`)
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(input, []byte{'\n'})

	var alg [512]byte
	for i, v := range lines[0] {
		alg[i] = v
	}

	image := lines[2:]

	part1 := countAfterN(image, alg, 2)
	part2 := countAfterN(image, alg, 50)

	fmt.Printf("Part 1 - %d; Part 2 - %d", part1, part2)
}

const (
	Light byte = '#'
	Dark  byte = '.'
)

func countAfterN(image [][]byte, alg [512]byte, rounds int) int {
	// create [rounds] space to expand, and [rounds] space for corruption from the edge
	image = padWithDark(image, rounds*2)

	// apply algorithm
	for i := 0; i < rounds; i++ {
		image = applyAlgorithm(image, alg)
	}

	// trim corruption at edge
	image = trimPadding(image, rounds)
	count := 0
	for _, line := range image {
		for _, pixel := range line {
			if pixel == Light {
				count++
			}
		}
	}

	return count
}

func padWithDark(image [][]byte, thickness int) [][]byte {
	newImage := newDarkImage(len(image) + thickness*2)

	for rowN, row := range image {
		for colN, col := range row {
			newImage[rowN+thickness][colN+thickness] = col
		}
	}

	return newImage
}

func trimPadding(image [][]byte, thickness int) [][]byte {
	newImage := newDarkImage(len(image) - thickness*2)

	for i := thickness; i < len(image)-thickness; i++ {
		for j := thickness; j < len(image)-thickness; j++ {
			newImage[i-thickness][j-thickness] = image[i][j]
		}
	}

	return newImage
}

func newDarkImage(length int) [][]byte {
	newImage := make([][]byte, length)

	for i := range newImage {
		newImage[i] = make([]byte, length) // assumed square
		for j := range newImage[i] {
			newImage[i][j] = Dark
		}
	}

	return newImage
}

func applyAlgorithm(image [][]byte, alg [512]byte) [][]byte {
	newImage := newDarkImage(len(image))

	for i := 1; i < len(image)-1; i++ {
		for j := 1; j < len(image)-1; j++ { // assumed square
			index := computeIndex(image, i, j)

			newImage[i][j] = alg[index]
		}
	}

	return newImage
}

func computeIndex(image [][]byte, i, j int) int {
	index := 0

	for row := i - 1; row < i+2; row++ {
		for col := j - 1; col < j+2; col++ {
			index <<= 1
			if image[row][col] == Light {
				index++
			}
		}
	}

	return index
}
