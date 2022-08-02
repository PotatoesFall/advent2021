package main

import (
	"fmt"

	"github.com/PotatoesFall/advent2021/minheap"
)

type AmphipodType uint8

const (
	A AmphipodType = iota + 1
	B
	C
	D
)

func (at AmphipodType) String() string {
	return string(at.Char())
}

var amphipodCosts = [...]int{0, 1, 10, 100, 1000}

func (a AmphipodType) Cost() int {
	return amphipodCosts[a]
}

var amphipodChars = [...]byte{0, 'A', 'B', 'C', 'D'}

func (a AmphipodType) Char() byte {
	return amphipodChars[a]
}

var amphipodDeepestGoalLocations = [...]Location{0, 15, 19, 23, 27}

func (a AmphipodType) DeepestGoalLocation() Location {
	return amphipodDeepestGoalLocations[a]
}

// 1, 2, 4, 6, 7, 10, 11 hallway
// 12-27 siderooms
type Location uint8

type Amphipod struct {
	Type     AmphipodType
	Location Location
}

type Neighbor struct {
	State
	Distance int
}

var startingAmphipods = [16]Amphipod{
	{A, 16},
	{A, 19},
	{A, 22},
	{A, 25},

	{B, 15},
	{B, 18},
	{B, 20},
	{B, 21},

	{C, 17},
	{C, 24},
	{C, 26},
	{C, 27},

	{D, 12},
	{D, 13},
	{D, 14},
	{D, 23},
}

var startingState = State{
	Amphipods: startingAmphipods,
	NotMoved:  Bools16{0b11111111, 0b11111111},
}

var goalAmphipods = [16]Amphipod{
	{A, 12},
	{A, 13},
	{A, 14},
	{A, 15},

	{B, 16},
	{B, 17},
	{B, 18},
	{B, 19},

	{C, 20},
	{C, 21},
	{C, 22},
	{C, 23},

	{D, 24},
	{D, 25},
	{D, 26},
	{D, 27},
}

var goalState = State{
	Amphipods: goalAmphipods,
}

var goalStateCompressed = goalState.Compress()

const prealloc = 50_000_000

func main() {
	stateCosts := make(map[CompressedState]int, prealloc)
	stateCosts[startingState.Compress()] = 0

	aStarThisShit(stateCosts, startingState, 0)
}

func aStarThisShit(stateCosts map[CompressedState]int, state State, cost int) int {
	openSet := minheap.New[CompressedState]()

	cameFrom := map[CompressedState]CompressedState{}

	startCompressed := startingState.Compress()
	startFScore := h(startingState)

	gScore := map[CompressedState]int{
		startCompressed: 0,
	}

	fScore := map[CompressedState]int{
		startCompressed: startFScore,
	}

	openSet.Insert(startFScore, startCompressed)

	for len(openSet) > 0 {
		fmt.Println(`lowestF:`, openSet[0].Score)
		currentCompressed, ok := openSet.Extract()
		if !ok {
			break
		}

		if currentCompressed == goalStateCompressed {
			return gScore[currentCompressed]
		}

		current := currentCompressed.Decompress()

		for _, neighbor := range neighbors(current) {
			neighborCompressed := neighbor.Compress()
			tentativeGScore := gScore[currentCompressed] + neighbor.Distance

			if neighborGScore, ok := gScore[neighborCompressed]; !ok || tentativeGScore < neighborGScore {
				newFScore := tentativeGScore + h(neighbor.State)

				cameFrom[neighborCompressed] = current.Compress()
				gScore[neighborCompressed] = tentativeGScore
				fScore[neighborCompressed] = newFScore
				openSet.Insert(newFScore, neighborCompressed)
			}
		}
	}

	panic(`failure`)
}

func h(state State) int {
	acc := -6666 // make heuristic admissible

	for _, amph := range state.Amphipods {
		acc += distance(amph.Location, amph.Type.DeepestGoalLocation()) * amph.Type.Cost()
	}

	// TODO: account for having to get that B out before putting it back in? might make heuristic more accurate.

	if acc < 0 {
		fmt.Println(state)
		panic(acc)
	}
	return acc
}

var distances = map[[2]Location]int{}

func distance(a, b Location) int {
	if d, ok := distances[[2]Location{a, b}]; ok {
		return d
	}
	d := _d(a, b)
	distances[[2]Location{a, b}] = d
	return d
}

func _d(a, b Location) int {
	// ANY TWO POINTS, NOT LIKE EARLIER
	aSideroom, bSideroom := a > 11, b > 11
	aRoomNum, bRoomNum := roomNum(a), roomNum(b) // 3, 5, 7, 9
	aRoomDepth, bRoomDepth := depth(a), depth(b) // 1 through 4
	if aSideroom && bSideroom {
		if aRoomNum != bRoomNum {
			return abs(aRoomNum-bRoomNum) + aRoomDepth + bRoomDepth
		}

		return abs(aRoomDepth - bRoomDepth)
	}

	if aSideroom && !bSideroom {
		return aRoomDepth + abs(aRoomNum-int(b))
	}

	if !aSideroom && bSideroom {
		return bRoomDepth + abs(bRoomNum-int(a))
	}

	if !aSideroom && !bSideroom {
		return abs(int(a - b))
	}

	panic(`impossible`)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func depth(l Location) int { // 1 through 4
	return (int(l)-12)%4 + 1
}

func roomNum(l Location) int { // 3, 5, 7, 9
	return 2*((int(l)-12)/4) + 3
}

func neighbors(state State) []Neighbor {
	neighbors := make([]Neighbor, 0, 4*7) // at best we have 4 amphipods that can go to 1 of 7 spaces. the more amphipods in the hallway, the fewer options.

	for i := 0; i < 16; i++ {

		if state.NotMoved.Get(i) {
			// make first move
			amph := state.Amphipods[i]
			d := depth(amph.Location)
			rn := roomNum(amph.Location)
			path := make([]int, 0, 10) // longest possible path is 10

			panic(`TODO`)
		}

		if state.MovedOnce.Get(i) {
			// bring it home
			amph := state.Amphipods[i]
			sisters := getSisters(state.Amphipods, i)
			d := depth(amph.Location)
			rn := roomNum(amph.Location)

			for _, sister := range sisters {
				if roomNum(sister.Location) == rn {
				}
			}
			panic(`TODO`)
		}
	}

	return neighbors
}

func getSisters(amphs [16]Amphipod, i int) [3]Amphipod {
	var sisters [3]Amphipod
	offset := i % 4
	copy(sisters[:], amphs[i-offset:i])
	copy(sisters[offset:], amphs[i:i-offset+4])
	return sisters
}
