package main

import (
	"strings"
)

type State struct {
	Amphipods [16]Amphipod
	NotMoved  Bools16
	MovedOnce Bools16
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
	str.WriteString("#\n  ########")

	return str.String()
}

func (s State) Compress() CompressedState {
	var compressed CompressedState

	for _, amph := range s.Amphipods {
		b := byte(amph.Type)             // bit representation
		index := (amph.Location - 1) / 2 // index in CompressedState array
		panic(`problem is somewhere here, they being swapped...`)
		if amph.Location%2 == 1 {
			b <<= 4
		}

		compressed[index] |= b
	}

	compressed[14] = s.NotMoved[0]
	compressed[15] = s.NotMoved[1]
	compressed[16] = s.MovedOnce[0]
	compressed[17] = s.MovedOnce[1]

	return compressed
}

type CompressedState [18]byte

func (cs CompressedState) Decompress() State {
	var state State

	c := 0
	for i := 0; i < 14; i++ {
		if cs[i] == 0 {
			continue
		}

		one, two := cs[i]&0b11110000, cs[i]&0b00001111

		if one != 0 {
			state.Amphipods[c] = Amphipod{AmphipodType(one >> 4), Location(i*2 + 1)}
			// fmt.Println(cs[i], AmphipodType(one>>4), i*2+1)
			c++
		}
		if two != 0 {
			state.Amphipods[c] = Amphipod{AmphipodType(two), Location(i*2 + 2)}
			// fmt.Println(cs[i], AmphipodType(two), i*2+2)
			c++
		}
	}

	state.NotMoved = Bools16{cs[14], cs[15]}
	state.MovedOnce = Bools16{cs[16], cs[17]}

	return state
}

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
