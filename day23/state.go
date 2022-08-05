package main

import (
	"strings"
)

type State struct {
	Amphipods [16]Amphipod
	NotMoved  Bools16
	MovedOnce Bools16
}

func (s State) Done() bool {
	for _, a := range s.Amphipods {
		if !a.Home() {
			return false
		}
	}

	return true
}

func (s State) String() string {
	var str strings.Builder

	str.WriteString("#############\n#")
	chars := [28]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'}
	for _, amph := range s.Amphipods {
		chars[amph.Location] = amph.Type.Char()
	}

	for i := 1; i <= 11; i++ {
		str.WriteByte(chars[i])
	}
	str.WriteString("#\n###")
	str.WriteByte(chars[12])
	str.WriteByte('#')
	str.WriteByte(chars[16])
	str.WriteByte('#')
	str.WriteByte(chars[20])
	str.WriteByte('#')
	str.WriteByte(chars[24])
	str.WriteString("###\n  #")
	str.WriteByte(chars[13])
	str.WriteByte('#')
	str.WriteByte(chars[17])
	str.WriteByte('#')
	str.WriteByte(chars[21])
	str.WriteByte('#')
	str.WriteByte(chars[25])
	str.WriteString("#\n  #")
	str.WriteByte(chars[14])
	str.WriteByte('#')
	str.WriteByte(chars[18])
	str.WriteByte('#')
	str.WriteByte(chars[22])
	str.WriteByte('#')
	str.WriteByte(chars[26])
	str.WriteString("#\n  #")
	str.WriteByte(chars[15])
	str.WriteByte('#')
	str.WriteByte(chars[19])
	str.WriteByte('#')
	str.WriteByte(chars[23])
	str.WriteByte('#')
	str.WriteByte(chars[27])
	str.WriteString("#\n  #########")

	return str.String()
}

// func (s State) Compress() CompressedState {
// 	var compressed CompressedState

// 	for i, amph := range s.Amphipods {
// 		b := byte(amph.Location) // bit representation
// 		if b > 11 {
// 			b -= 12
// 		}
// 		index := i / 2 // index in CompressedState array
// 		if i%2 == 0 {
// 			b <<= 4
// 		}

// 		compressed[index] |= b
// 	}

// 	compressed[8] = s.NotMoved[0]
// 	compressed[9] = s.NotMoved[1]
// 	compressed[10] = s.MovedOnce[0]
// 	compressed[11] = s.MovedOnce[1]

// 	return compressed
// }

// type CompressedState [12]byte

// var amphipodTypesByI = [8]AmphipodType{A, A, B, B, C, C, D, D}

// func (cs CompressedState) Decompress() State {
// 	state := State{
// 		NotMoved:  Bools16{cs[8], cs[9]},
// 		MovedOnce: Bools16{cs[10], cs[11]},
// 	}

// 	for i := 0; i < 8; i++ {
// 		if cs[i] == 0 {
// 			continue
// 		}

// 		one, two := cs[i]&0b11110000, cs[i]&0b00001111

// 		if state.NotMoved.Get(i*2) || !(state.MovedOnce.Get(i * 2)) {
// 			one += 12
// 		}
// 		if state.NotMoved.Get(i*2+1) || !(state.MovedOnce.Get(i*2 + 1)) {
// 			two += 12
// 		}

// 		state.Amphipods[i*2] = Amphipod{
// 			Type:     amphipodTypesByI[i],
// 			Location: Location(one),
// 		}
// 		state.Amphipods[i*2+1] = Amphipod{
// 			Type:     amphipodTypesByI[i],
// 			Location: Location(two),
// 		}
// 	}

// 	return state
// }

type Bools16 [2]byte

func (b Bools16) Get(i int) bool {
	return b[i/8]&((0b10000000)>>(i%8)) != 0
}

func (b *Bools16) Set(i int, v bool) {
	if v {
		b[i/8] |= 0b10000000 >> (i % 8)
	} else {
		b[i/8] &^= 0b10000000 >> (i % 8)
	}
}

func (b Bools16) String() string {
	var str strings.Builder
	str.WriteString(`Bools16(0b`)
	for i := 0; i < 16; i++ {
		if b.Get(i) {
			str.WriteRune('1')
		} else {
			str.WriteRune('0')
		}
	}
	str.WriteString(`)`)
	return str.String()
}
