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

// 1, 2, 4, 6, 8, 10, 11 hallway
// 12-27 siderooms
type Location uint8

var hallway = [...]Location{1, 2, 4, 6, 8, 10, 11}

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

func main() {
	n := aStarThisShit(startingState, 0)
	fmt.Printf("Part 2 - %d\n", n)
}

func aStarThisShit(state State, cost int) int {
	openSet := minheap.New[CompressedState]()

	cameFrom := map[CompressedState]CompressedState{}

	startCompressed := state.Compress()
	startFScore := h(state)

	gScore := map[CompressedState]int{
		startCompressed: cost,
	}

	fScore := map[CompressedState]int{
		startCompressed: startFScore,
	}

	openSet.Insert(startFScore, startCompressed)

	count := 0
	farthest := 0
	for len(openSet) > 0 {
		count++
		if count%1_000_000 == 0 {
			fmt.Println(`current size of openSet:`, len(openSet))
		}
		// fmt.Println(`lowestF:`, openSet[0].Score)
		currentCompressed, ok := openSet.Extract()
		if !ok {
			break
		}

		if currentCompressed == goalStateCompressed {
			return gScore[currentCompressed]
		}

		current := currentCompressed.Decompress()

		// if count%100_000 == 0 {
		// 	fmt.Println("\n\n")
		// 	bla := currentCompressed
		// 	parent, ok := cameFrom[bla]
		// 	for ok {
		// 		fmt.Println("\ncame from:")
		// 		fmt.Println(bla.Decompress())
		// 		fmt.Println(bla.Decompress().NotMoved)
		// 		fmt.Println(bla.Decompress().MovedOnce)
		// 		bla = parent
		// 		parent, ok = cameFrom[bla]
		// 	}
		// }

		farnessScore := 32
		for i := 0; i < 16; i++ {
			if current.MovedOnce.Get(i) {
				farnessScore -= 1
			}
			if current.NotMoved.Get(i) {
				farnessScore -= 2
			}
		}
		if farnessScore > farthest {
			farthest = farnessScore
			fmt.Println(`NEW FARNESS:`, farthest)
		}

		// fmt.Printf("\n\nneighbors of\n%s\n", current)
		for _, neighbor := range neighbors(current) {
			// fmt.Printf("\n%s\n", neighbor.State)

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

	var bInSpot, bBlocks bool
	for _, amph := range state.Amphipods {
		if amph.Location == 18 && amph.Type == B {
			bInSpot = true
		}
		if amph.Location == 19 && amph.Type != B {
			bBlocks = true
		}

		acc += distance(amph.Location, amph.Type.DeepestGoalLocation()) * amph.Type.Cost()
	}

	if bInSpot && bBlocks {
		acc += B.Cost() * 8
	}

	// fmt.Println(state, acc)

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
	aRoomNum, bRoomNum := getRoomNum(a), getRoomNum(b) // 3, 5, 7, 9
	aRoomDepth, bRoomDepth := getDepth(a), getDepth(b) // 1 through 4
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

func getDepth(l Location) int { // 1 through 4
	return (int(l)-12)%4 + 1
}

func getRoomNum(l Location) int { // 3, 5, 7, 9
	return 2*((int(l)-12)/4) + 3
}

func neighbors(state State) (neighbors []Neighbor) {
	neighbors = make([]Neighbor, 0, 4*7) // at best we have 4 amphipods that can go to 1 of 7 spaces. the more amphipods in the hallway, the fewer options.

outer:
	for i := 0; i < 16; i++ {
		if state.NotMoved.Get(i) {
			// fmt.Println(`not moved:`, i)
			// make first move
			amph := state.Amphipods[i]
			roomNum := getRoomNum(amph.Location)
			depth := getDepth(amph.Location)

			// check if we can reach the hallway
			for j := 1; j < depth; j++ {
				if !locationFree(state, amph.Location-Location(j)) {
					// fmt.Println(`cant reach hallway`, amph.Location, roomNum, amph.Location-Location(i))
					continue outer
				}
			}

			destinations := make([]Location, 0, 7) // max 7 destinations in hallway

			for _, destination := range hallway {
				free := locationFree(state, destination)
				if int(destination) < roomNum { // before room
					if !free {
						destinations = destinations[:0] // clear all previous fields - inaccessible
					} else {
						destinations = append(destinations, destination)
					}
				} else {
					if !free { // no more accessible fields on this side
						break
					} else {
						destinations = append(destinations, destination)
					}
				}
			}
			// fmt.Printf("destinations calculated for %s at %d: %v\n", amph.Type, amph.Location, destinations)

			for _, destination := range destinations {
				neighborState := state

				fmt.Printf("first motion for amphipod i%d %s from %d to %d\n", i, amph.Type, amph.Location, destination)
				neighborState.Amphipods[i].Location = destination
				neighborState.NotMoved.Set(i, false)
				neighborState.MovedOnce.Set(i, true)

				neighbors = append(neighbors, Neighbor{
					State:    neighborState,
					Distance: state.Amphipods[i].Type.Cost() * distance(state.Amphipods[i].Location, destination),
				})
			}
		}

		if state.MovedOnce.Get(i) {
			// bring it home
			amph := state.Amphipods[i]

			destination := amph.Type.DeepestGoalLocation()
			for {
				amphType, ok := getFromLocation(state, destination)
				if !ok { // goal is correct
					break
				}
				if amphType == amph.Type { // goal might be one higher
					destination--
					if (destination-12)%4 == 3 { // already tried all, give up
						panic(`should never happen`)
						// continue outer
					}
					continue
				}

				continue outer // wrong amphipod type -> not possible
			}

			roomNum := getRoomNum(destination)
			step := 1
			if roomNum < int(amph.Location) {
				step = -1
			}

			for l := int(amph.Location) + step; l != roomNum; l += step {
				if !locationFree(state, Location(l)) {
					a, _ := getFromLocation(state, Location(l))
					fmt.Printf("going home: %s at %d blocked by %s at %d\n", amph.Type, amph.Location, a, l)
					continue outer
				}
			}

			neighborState := state

			fmt.Println("DEST:", destination)
			neighborState.Amphipods[i].Location = destination
			neighborState.MovedOnce.Set(i, false)

			neighbors = append(neighbors, Neighbor{
				State:    neighborState,
				Distance: state.Amphipods[i].Type.Cost() * distance(state.Amphipods[i].Location, destination),
			})
		}
	}

	fmt.Printf("\n\nneighbors of:\n%s\n%s\n%s\n", state, state.NotMoved, state.MovedOnce)
	for _, neighbor := range neighbors {
		fmt.Printf("\n%s\n", neighbor)
		fmt.Println(neighbor.NotMoved)
		fmt.Println(neighbor.MovedOnce)
	}

	return neighbors
}

func locationFree(state State, l Location) bool {
	for i := range state.Amphipods {
		if state.Amphipods[i].Location == l {
			return false
		}
	}

	return true
}

func getFromLocation(state State, l Location) (AmphipodType, bool) {
	for i := range state.Amphipods {
		if state.Amphipods[i].Location == l {
			return state.Amphipods[i].Type, true
		}
	}

	return 0, false
}
