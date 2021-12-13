package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile(`input13`)
	if err != nil {
		panic(err)
	}

	page, lines := parse(input)

	for i, line := range lines {
		page.Fold(line)

		if i == 0 {
			fmt.Printf("Part 1 - there are %d points on the page\n", page.Count())
		}
	}

	fmt.Print(page.String())
}
