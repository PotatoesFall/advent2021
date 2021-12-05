package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile(`input`)
	if err != nil {
		panic(err)
	}
	lineList := bytes.Split(input, []byte("\n"))

	lines := makeLines(lineList)

	powerConsumption := getPowerConsumption(lines)
	lifeSupportRating := getLifeSupportRating(lines)

	fmt.Printf("power consumption is: %d\n", powerConsumption)
	fmt.Printf("life support rating is: %d\n", lifeSupportRating)
}

func getPowerConsumption(lines Lines) int {
	gamma := lines.Modes()
	epsilon := gamma.Invert()

	return gamma.Int() * epsilon.Int()
}

func getLifeSupportRating(lines Lines) int {
	oxygenGeneratorRating := lines.FindLastMatch(true)
	CO2ScrubberRating := lines.FindLastMatch(false)

	return oxygenGeneratorRating.Int() * CO2ScrubberRating.Int()
}
