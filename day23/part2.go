package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unsafe"
)

/*
PSEUDOCODE

have state + cost, put as only element in list of (state, cost)

start a map of state to cost, starting with the current state to 0

initialize lowest cost as maxint

while len(state-costs) > 0
	for each state
		make [] of possible moves

		for each amphipod
			if in destination sideroom
				continue

			if surrounded by walls and amphipods
				continue

			if in sideroom near door
				for each hallway room
					if path possible
						add to [] of possible moves

			if in hallway room
				if path to destination hallway is possible AND (destination hallway is empty OR destination has only one correct amphipod inside)
					add (to beginning of?) [] of possible moves

		for each possible move
			calculate state after move including cost

			if state is finished AND cost is lower than lowest cost
				lower the lowest cost to the new value
				continue

			append to list of state-costs


	for each possible new state
		if state is already in map and has the same or higher cost
			continue

		if state has higher or equal cost to current lowest cost solution
			continue

		append to outer list of state-costs
}
*/

const print = false

func doPrintln(args ...any) {
	if print {
		fmt.Println(args...)
	}
}

func doPrintf(f string, args ...any) {
	if print {
		fmt.Printf(f, args...)
	}
}

func doPrint(args ...any) {
	if print {
		fmt.Print(args...)
	}
}

func main() {
	state := State{
		AmphipodPosition{D, 12},
		AmphipodPosition{D, 13},
		AmphipodPosition{D, 14},
		AmphipodPosition{B, 15},

		AmphipodPosition{A, 16},
		AmphipodPosition{C, 17},
		AmphipodPosition{B, 18},
		AmphipodPosition{A, 19},

		AmphipodPosition{B, 20},
		AmphipodPosition{B, 21},
		AmphipodPosition{A, 22},
		AmphipodPosition{D, 23},

		AmphipodPosition{C, 24},
		AmphipodPosition{A, 25},
		AmphipodPosition{C, 26},
		AmphipodPosition{C, 27},
	}

	fmt.Println(unsafe.Sizeof(StateCost{}))

	fmt.Printf("PART 2 - %d\n", part2(state))
}

func part2(initialState State) uint { //nolint:funlen,gocyclo
	stateCosts := make([]StateCost, 0, 10_000_000)
	stateCosts = append(stateCosts, StateCost{
		State: initialState,
		Cost:  0,
	})

	seenStates := make(map[State]uint, 10_000_000)
	seenStates[initialState] = 0

	var lowestCost uint = 60_000 // throw away anything with too high of a score

	newStateCosts := make([]StateCost, 0, 10_000_000)

	round := 1
	var start time.Time
	for len(stateCosts) > 0 {
		if !start.IsZero() {
			fmt.Printf("took %d seconds\n", int(time.Since(start).Seconds()))
		}
		start = time.Now()
		fmt.Printf("\n\n### ROUND %02d ###\n", round)
		fmt.Printf("To check %d states...\n", len(stateCosts))
		if print {
			doPrintln(`EARLY EXIT`)
			os.Exit(0)
		}
		round++

		newStateCosts = newStateCosts[:0]

		possibleNewStates := make([]StateCost, 0, 16*11) // number of amphipods times number of possible moves per amphipod
		for _, stateCost := range stateCosts {
			possibleNewStates = possibleNewStates[:0]

		outer:
			for i, amph := range stateCost.State {
				if amph.Position.Sideroom() {
					// check if already in destination
					if amph.InDestinationRoom() {
						correct := true
						for d := amph.Position.SideRoomDepth(); d < 4; d += 1 {
							if stateCost.Get(amph.Position-Position(d)) != amph.Amphipod {
								correct = false
								break
							}
						}
						if correct {
							doPrintln("\n"+amph.Amphipod.String(), "AT", amph.Position, "ALREADY IN CORRECT ROOM")
							continue
						}
					}

					// stuck in deep end
					if depth := amph.Position.SideRoomDepth(); depth > 0 {
						for d := depth; d > 0; d-- {
							if stateCost.Get(amph.Position-Position(d)) != Empty {
								doPrintln("\n"+amph.Amphipod.String(), "AT", amph.Position, "COCKBLOCKED BY", stateCost.Get(amph.Position-1).String(), "AT", amph.Position-1)
								continue outer
							}
						}
					}

					// check all spaces we can move to
					for _, hallwayPosition := range hallwaySpaces {
						move := Move{AmphipodIndex: uint(i), To: hallwayPosition}
						if state, possible := doCheck(stateCost, move); possible {
							possibleNewStates = append(possibleNewStates, state)
						}
					}
				} else { // in hallway
					destinations := amph.Amphipod.Destinations()
					depth := -1
					for d := 0; d < 4; d++ {
						if stateCost.Get(destinations[d]) == Empty {
							depth = d
							continue
						}
						break
					}
					if depth != -1 {
						move := Move{AmphipodIndex: uint(i), To: destinations[depth]}
						if state, possible := doCheck(stateCost, move); possible {
							possibleNewStates = append(possibleNewStates, state)
						}
					}
				}
			}

			for _, possibleNewState := range possibleNewStates {
				if possibleNewState.Done() && possibleNewState.Cost < lowestCost {
					lowestCost = possibleNewState.Cost
					continue
				} else {
					newStateCosts = append(newStateCosts, possibleNewState)
				}
			}
		}

		stateCosts = stateCosts[:0]
		for _, newStateCost := range newStateCosts {
			if newStateCost.Cost > lowestCost {
				continue
			}

			if cost, ok := seenStates[newStateCost.State]; ok && cost <= newStateCost.Cost {
				continue
			}

			seenStates[newStateCost.State] = newStateCost.Cost
			stateCosts = append(stateCosts, newStateCost)
		}
	}

	return lowestCost
}

var (
	hallwaySpaces  = [...]Position{1, 2, 4, 6, 8, 10, 11}
	sideRoomSpaces = [...]Position{12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}
)

/*
Position is a position
Hallway 1-11 // 3, 5, 7 and 9 are forbidden
RoomA 12/13/14/15
RoomB 16/17/18/19
RoomC 20/21/22/23
RoomD 24/25/26/27
*/
type Position uint

func (p Position) Hallway() bool {
	return p < 12
}

func (p Position) Sideroom() bool {
	return p > 11
}

func (p Position) SideRoomDepth() uint {
	return uint((p - 12) % 4)
}

type State [16]AmphipodPosition

func (s State) Get(p Position) Amphipod {
	for _, ap := range s {
		if ap.Position == p {
			return ap.Amphipod
		}
	}

	return Empty
}

func (s State) Done() bool {
	for _, ap := range s {
		if !ap.InDestinationRoom() {
			return false
		}
	}

	doPrintln(`DONE`)
	return true
}

func (s State) String() string {
	var str strings.Builder
	str.WriteString("#############\n#")
	for i := 1; i < 12; i++ {
		str.WriteString(s.Get(Position(i)).String())
	}
	str.WriteString("#\n###")
	str.WriteString(s.Get(12).String())
	str.WriteString(`#`)
	str.WriteString(s.Get(16).String())
	str.WriteString(`#`)
	str.WriteString(s.Get(20).String())
	str.WriteString(`#`)
	str.WriteString(s.Get(24).String())
	str.WriteString("###\n  #")
	str.WriteString(s.Get(13).String())
	str.WriteString(`#`)
	str.WriteString(s.Get(17).String())
	str.WriteString(`#`)
	str.WriteString(s.Get(21).String())
	str.WriteString(`#`)
	str.WriteString(s.Get(25).String())
	str.WriteString("#\n  #")
	str.WriteString(s.Get(14).String())
	str.WriteString(`#`)
	str.WriteString(s.Get(18).String())
	str.WriteString(`#`)
	str.WriteString(s.Get(22).String())
	str.WriteString(`#`)
	str.WriteString(s.Get(26).String())
	str.WriteString("#\n  #")
	str.WriteString(s.Get(15).String())
	str.WriteString(`#`)
	str.WriteString(s.Get(19).String())
	str.WriteString(`#`)
	str.WriteString(s.Get(23).String())
	str.WriteString(`#`)
	str.WriteString(s.Get(27).String())
	str.WriteString("#\n  #########")
	return str.String()
}

type StateCost struct {
	State
	Cost uint
}

type AmphipodPosition struct {
	Amphipod Amphipod
	Position Position
}

func (ap AmphipodPosition) InDestinationRoom() bool {
	switch ap.Position {
	case 12, 13, 14, 15:
		return ap.Amphipod == A
	case 16, 17, 18, 19:
		return ap.Amphipod == B
	case 20, 21, 22, 23:
		return ap.Amphipod == C
	case 24, 25, 26, 27:
		return ap.Amphipod == D
	default:
		return false
	}
}

type Amphipod uint

const (
	Empty Amphipod = iota
	A
	B
	C
	D
)

func (a Amphipod) Cost() uint {
	switch a { //nolint:exhaustive
	case A:
		return 1
	case B:
		return 10
	case C:
		return 100
	case D:
		return 1000
	}

	panic(a)
}

func (a Amphipod) String() string {
	switch a {
	case Empty:
		return `.`
	case A:
		return `A`
	case B:
		return `B`
	case C:
		return `C`
	case D:
		return `D`
	}

	panic(a)
}

func (a Amphipod) Destinations() [4]Position {
	switch a { //nolint:exhaustive
	case A:
		return [4]Position{12, 13, 14, 15}
	case B:
		return [4]Position{16, 17, 18, 19}
	case C:
		return [4]Position{20, 21, 22, 23}
	case D:
		return [4]Position{24, 25, 26, 27}
	}

	panic(a)
}

type Move struct {
	AmphipodIndex uint
	To            Position
}

func distance(p1, p2 Position) uint {
	if p1 < 12 {
		p2, p1 = p1, p2 // p1 is now sideroom, p2 is hallway
	}

	cost := p1.SideRoomDepth() + 1 // 1 step plus depth to get out

	var start Position
	switch p1 {
	case 12, 13, 14, 15:
		start = 3
	case 16, 17, 18, 19:
		start = 5
	case 20, 21, 22, 23:
		start = 7
	case 24, 25, 26, 27:
		start = 9
	}

	if start > p2 {
		cost += uint(start - p2)
	} else {
		cost += uint(p2 - start)
	}

	return cost
}

func doCheck(stateCost StateCost, move Move) (StateCost, bool) {
	doPrintln("\n\nCHECK")
	state, ok := check(stateCost, move)
	doPrintf("%s from %d to %d\n", stateCost.State[move.AmphipodIndex].Amphipod, stateCost.State[move.AmphipodIndex].Position, move.To)
	if ok {
		doPrintln(`possible:`)
		doPrintf("New Cost %d:\n%s\n", state.Cost, state.State)
	} else {
		doPrintln(`not possible:`)
		doPrintf("Cost %d:\n%s\n", stateCost.Cost, stateCost.State)
	}

	return state, ok
}

func check(stateCost StateCost, move Move) (StateCost, bool) {
	start := stateCost.State[move.AmphipodIndex].Position
	if !possible(stateCost.State, start, move.To) {
		return StateCost{}, false
	}

	stateCost.State[move.AmphipodIndex].Position = move.To
	stateCost.Cost += stateCost.State[move.AmphipodIndex].Amphipod.Cost() * distance(start, move.To)

	return stateCost, true
}

func possible(state State, from, to Position) bool {
	if state.Get(to) != Empty {
		doPrint(`destination busy: `)
		return false
	}

	if from < 12 {
		from, to = to, from // from sideroom to hallway for simplicity
	}

	for i := from.SideRoomDepth(); i > 0; i-- {
		if state.Get(from-Position(i)) != Empty {
			doPrint(`cannot leave deep sideroom: `)
			return false
		}
	}

	var start Position
	switch from {
	case 12, 13, 14, 15:
		start = 3
	case 16, 17, 18, 19:
		start = 5
	case 20, 21, 22, 23:
		start = 7
	case 24, 25, 26, 27:
		start = 9
	}

	increment := 1
	if start > to {
		increment = -1
	}

	for p := start; p != to; p = Position(int(p) + increment) {
		if state.Get(p) != Empty {
			doPrintf(`blocked at position %d:`, p)
			return false
		}
	}

	return true
}
