package main

// func TestCompress(t *testing.T) {
// 	state := State{
// 		Amphipods: [16]Amphipod{{B, 1}, {A, 2}, {A, 3}, {A, 4}, {A, 5}, {A, 6}, {A, 21}, {A, 12}, {A, 13}, {A, 14}, {A, 15}, {A, 16}, {A, 17}, {A, 18}, {A, 19}, {A, 20}},
// 		NotMoved:  Bools16{0b11111111, 0b11111110},
// 	}
// 	fmt.Println(state)
// 	compressed := state.Compress()
// 	decompressed := compressed.Decompress()
// 	fmt.Println(decompressed)
// 	assert.Cmp(t, decompressed, state)
// 	// expected := CompressedState{}
// 	// fmt.Println(expected)
// 	// fmt.Println(compressed)
// 	// assert.Eq(t, expected, compressed)
// }

// func TestBools17(t *testing.T) {
// 	a := Bools16{0b11111111, 0b11111111}
// 	for i := 0; i < 16; i++ {
// 		if a.Get(i) == false {
// 			t.Error(`false`, i)
// 		}
// 		a.Set(i, false)
// 		if a.Get(i) == true {
// 			t.Error(`true`, i)
// 		}
// 	}
// }
