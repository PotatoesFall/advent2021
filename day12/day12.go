package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
)

const part2 = true

var connections = map[Cave][]Cave{}

func main() {
	input, err := os.ReadFile(`input12`)
	if err != nil {
		panic(err)
	}

	parseConnections(input)

	paths := getPaths(NewPath(caveStart))

	fmt.Printf(`Part 2 - there are %d paths`, len(paths))
}

func parseConnections(input []byte) {
	regexConnection := regexp.MustCompile(`^(\w+)\-(\w+)$`)

	for _, line := range bytes.Split(input, []byte{'\n'}) {
		matches := regexConnection.FindStringSubmatch(string(line))
		if len(matches) != 3 {
			panic(string(line))
		}

		connections[Cave(matches[1])] = append(connections[Cave(matches[1])], Cave(matches[2]))
		connections[Cave(matches[2])] = append(connections[Cave(matches[2])], Cave(matches[1]))
	}
}

func getPaths(path Path) []Path {
	if path.caves[len(path.caves)-1] == caveEnd {
		return []Path{path}
	}

	var paths []Path
	for _, cave := range connections[path.caves[len(path.caves)-1]] {
		newPath := path.Copy()
		if newPath.Go(cave) {
			paths = append(paths, getPaths(newPath)...)
		}
	}

	return paths
}
