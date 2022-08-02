package main

// import (
// 	"fmt"
// 	"math"
// 	"strings"
// )

// /*
// PSEUDOCODE

// have state + cost, put as only element in list of (state, cost)

// start a map of state to cost, starting with the current state to 0

// initialize lowest cost as maxint

// while len(state-costs) > 0
// 	for each state
// 		make [] of possible moves

// 		for each amphipod
// 			if in destination sideroom
// 				continue

// 			if surrounded by walls and amphipods
// 				continue

// 			if in sideroom near door
// 				for each hallway room
// 					if path possible
// 						add to [] of possible moves

// 			if in hallway room
// 				if path to destination hallway is possible AND (destination hallway is empty OR destination has only one correct amphipod inside)
// 					add (to beginning of?) [] of possible moves

// 		for each possible move
// 			calculate state after move including cost

// 			if state is finished AND cost is lower than lowest cost
// 				lower the lowest cost to the new value
// 				continue

// 			append to list of state-costs

// 	for each possible new state
// 		if state is already in map and has the same or higher cost
// 			continue

// 		if state has higher or equal cost to current lowest cost solution
// 			continue

// 		append to outer list of state-costs
// }
// */

// func main() {
// 	state := State{
// 		AmphipodPosition{D, 12},
// 		AmphipodPosition{B, 13},
// 		AmphipodPosition{A, 14},
// 		AmphipodPosition{A, 15},
// 		AmphipodPosition{B, 16},
// 		AmphipodPosition{D, 17},
// 		AmphipodPosition{C, 18},
// 		AmphipodPosition{C, 19},
// 	}

// 	fmt.Printf("PART 1 - %d\n", part1(state))
// }

// func part1(initialState State) uint { //nolint:funlen,gocyclo
// 	stateCosts := []StateCost{{
// 		State: initialState,
// 		Cost:  0,
// 	}}

// 	seenStates := map[State]uint{
// 		initialState: 0,
// 	}

// 	var lowestCost uint = math.MaxUint

// 	round := 1
// 	for len(stateCosts) > 0 {
// 		fmt.Printf("\n\n### ROUND %02d ###", round)
// 		round++

// 		newStateCosts := make([]StateCost, 0)

// 		for _, stateCost := range stateCosts {
// 			var possibleNewStates []StateCost

// 			for i, amph := range stateCost.State {
// 				if amph.Position.Sideroom() {
// 					// check if already in destination
// 					if amph.InDestinationRoom() {
// 						continue
// 					}

// 					// stuck in deep end
// 					if amph.Position.DeepSideroom() && stateCost.Get(amph.Position-1) != Empty {
// 						// fmt.Println(`COCKBLOCKED BY`, stateCost.Get(amph.Position-1).String())
// 						continue
// 					}

// 					// check all spaces we can move to
// 					for _, hallwayPosition := range hallwaySpaces {
// 						move := Move{AmphipodIndex: uint(i), To: hallwayPosition}
// 						if state, possible := doCheck(stateCost, move); possible {
// 							possibleNewStates = append(possibleNewStates, state)
// 						}
// 					}
// 				} else { // in hallway
// 					shallow, deep := amph.Amphipod.Destinations()
// 					// add check that we either move to deep position or other amphipod is already there
// 					if stateCost.State.Get(deep) == Empty {
// 						move := Move{AmphipodIndex: uint(i), To: deep}
// 						if state, possible := doCheck(stateCost, move); possible {
// 							possibleNewStates = append(possibleNewStates, state)
// 						}
// 					}
// 					if stateCost.State.Get(deep) == amph.Amphipod && stateCost.State.Get(shallow) == Empty {
// 						move := Move{AmphipodIndex: uint(i), To: shallow}
// 						if state, possible := doCheck(stateCost, move); possible {
// 							possibleNewStates = append(possibleNewStates, state)
// 						}
// 					}
// 				}
// 			}

// 			for _, possibleNewState := range possibleNewStates {
// 				if possibleNewState.Done() && possibleNewState.Cost < lowestCost {
// 					lowestCost = possibleNewState.Cost
// 					continue
// 				} else {
// 					newStateCosts = append(newStateCosts, possibleNewState)
// 				}
// 			}
// 		}

// 		stateCosts = nil
// 		for _, newStateCost := range newStateCosts {
// 			if cost, ok := seenStates[newStateCost.State]; ok && cost <= newStateCost.Cost {
// 				continue
// 			}

// 			seenStates[newStateCost.State] = newStateCost.Cost
// 			stateCosts = append(stateCosts, newStateCost)
// 		}
// 	}

// 	return lowestCost
// }

// var (
// 	hallwaySpaces  = [...]Position{1, 2, 4, 6, 8, 10, 11}
// 	sideRoomSpaces = [...]Position{12, 13, 14, 15, 16, 17, 18, 19}
// )

// /*
// Position is a position
// Hallway 1-11 // 3, 5, 7 and 9 are forbidden
// RoomA 12/13
// RoomB 14/15
// RoomC 16/17
// RoomD 18/19
// */
// type Position uint

// func (p Position) Hallway() bool {
// 	return p < 12
// }

// func (p Position) Sideroom() bool {
// 	return p > 11
// }

// func (p Position) DeepSideroom() bool {
// 	return p%2 == 1
// }

// type State [8]AmphipodPosition

// func (s State) Get(p Position) Amphipod {
// 	for _, ap := range s {
// 		if ap.Position == p {
// 			return ap.Amphipod
// 		}
// 	}

// 	return Empty
// }

// func (s State) Done() bool {
// 	for _, ap := range s {
// 		if !ap.InDestinationRoom() {
// 			return false
// 		}
// 	}

// 	fmt.Println(`DONE`)
// 	return true
// }

// func (s State) String() string {
// 	var str strings.Builder
// 	str.WriteString("#############\n#")
// 	for i := 1; i < 12; i++ {
// 		str.WriteString(s.Get(Position(i)).String())
// 	}
// 	str.WriteString("#\n###")
// 	str.WriteString(s.Get(12).String())
// 	str.WriteString(`#`)
// 	str.WriteString(s.Get(14).String())
// 	str.WriteString(`#`)
// 	str.WriteString(s.Get(16).String())
// 	str.WriteString(`#`)
// 	str.WriteString(s.Get(18).String())
// 	str.WriteString("###\n  #")
// 	str.WriteString(s.Get(13).String())
// 	str.WriteString(`#`)
// 	str.WriteString(s.Get(15).String())
// 	str.WriteString(`#`)
// 	str.WriteString(s.Get(17).String())
// 	str.WriteString(`#`)
// 	str.WriteString(s.Get(19).String())
// 	str.WriteString("#\n  #########")
// 	return str.String()
// }

// type StateCost struct {
// 	State
// 	Cost uint
// }

// type AmphipodPosition struct {
// 	Amphipod Amphipod
// 	Position Position
// }

// func (ap AmphipodPosition) InDestinationRoom() bool {
// 	switch ap.Position {
// 	case 12, 13:
// 		return ap.Amphipod == A
// 	case 14, 15:
// 		return ap.Amphipod == B
// 	case 16, 17:
// 		return ap.Amphipod == C
// 	case 18, 19:
// 		return ap.Amphipod == D
// 	default:
// 		return false
// 	}
// }

// type Amphipod uint

// const (
// 	Empty Amphipod = iota
// 	A
// 	B
// 	C
// 	D
// )

// func (a Amphipod) Cost() uint {
// 	switch a { //nolint:exhaustive
// 	case A:
// 		return 1
// 	case B:
// 		return 10
// 	case C:
// 		return 100
// 	case D:
// 		return 1000
// 	}

// 	panic(a)
// }

// func (a Amphipod) String() string {
// 	switch a {
// 	case Empty:
// 		return `.`
// 	case A:
// 		return `A`
// 	case B:
// 		return `B`
// 	case C:
// 		return `C`
// 	case D:
// 		return `D`
// 	}

// 	panic(a)
// }

// func (a Amphipod) Destinations() (Position, Position) {
// 	switch a { //nolint:exhaustive
// 	case A:
// 		return 12, 13
// 	case B:
// 		return 14, 15
// 	case C:
// 		return 16, 17
// 	case D:
// 		return 18, 19
// 	}

// 	panic(a)
// }

// type Move struct {
// 	AmphipodIndex uint
// 	To            Position
// }

// func distance(p1, p2 Position) uint {
// 	if p1 < 12 {
// 		p2, p1 = p1, p2 // p1 is now sideroom, p2 is hallway
// 	}

// 	var cost uint
// 	if p1%2 == 1 {
// 		cost = 2 // deep, takes an additional step to get out
// 	} else {
// 		cost = 1 // just one step to get out
// 	}

// 	var start Position
// 	switch p1 {
// 	case 12, 13:
// 		start = 3
// 	case 14, 15:
// 		start = 5
// 	case 16, 17:
// 		start = 7
// 	case 18, 19:
// 		start = 9
// 	}

// 	if start > p2 {
// 		cost += uint(start - p2)
// 	} else {
// 		cost += uint(p2 - start)
// 	}

// 	return cost
// }

// func doCheck(stateCost StateCost, move Move) (StateCost, bool) {
// 	// fmt.Println("\n\nCHECK")
// 	state, ok := check(stateCost, move)
// 	// fmt.Printf("%s from %d to %d\n", stateCost.State[move.AmphipodIndex].Amphipod, stateCost.State[move.AmphipodIndex].Position, move.To)
// 	if ok {
// 		// fmt.Println(`possible:`)
// 		// fmt.Printf("New Cost %d:\n%s\n", state.Cost, state.State)
// 	} else {
// 		// fmt.Println(`not possible:`)
// 		// fmt.Printf("Cost %d:\n%s\n", stateCost.Cost, stateCost.State)
// 	}

// 	return state, ok
// }

// func check(stateCost StateCost, move Move) (StateCost, bool) {
// 	start := stateCost.State[move.AmphipodIndex].Position
// 	if !possible(stateCost.State, start, move.To) {
// 		return StateCost{}, false
// 	}

// 	stateCost.State[move.AmphipodIndex].Position = move.To
// 	stateCost.Cost += stateCost.State[move.AmphipodIndex].Amphipod.Cost() * distance(start, move.To)

// 	return stateCost, true
// }

// func possible(state State, from, to Position) bool {
// 	if state.Get(to) != Empty {
// 		// fmt.Print(`destination busy: `)
// 		return false
// 	}

// 	if from < 12 {
// 		from, to = to, from // from sideroom to hallway for simplicity
// 	}

// 	if from.DeepSideroom() && state.Get(from-1) != Empty {
// 		// fmt.Print(`cannot leave deep sideroom: `)
// 		return false
// 	}

// 	var start Position
// 	switch from {
// 	case 12, 13:
// 		start = 3
// 	case 14, 15:
// 		start = 5
// 	case 16, 17:
// 		start = 7
// 	case 18, 19:
// 		start = 9
// 	}

// 	increment := 1
// 	if start > to {
// 		increment = -1
// 	}

// 	for p := start; p != to; p = Position(int(p) + increment) {
// 		if state.Get(p) != Empty {
// 			// fmt.Printf(`blocked at position %d:`, p)
// 			return false
// 		}
// 	}

// 	return true
// }
