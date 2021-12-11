package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile(`input6`)
	if err != nil {
		panic(err)
	}

	numStrings := strings.Split(string(input), `,`)
	var school School
	for _, str := range numStrings {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		school[n]++
	}

	for i := 0; i < 80; i++ {
		school.Tick()
	}
	fmt.Println(school.Num())

	for i := 80; i < 256; i++ {
		school.Tick()
	}
	fmt.Println(school.Num())
}
