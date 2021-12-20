package main

func splitOne(pair *Pair) bool {
	for _, side := range bothSides {
		if !pair.HasChild(side) && pair.Value(side) > 9 {
			pair.SetChild(side, split(pair.Value(side)))
			pair.SetValue(side, 0)
			return true
		}

		if pair.HasChild(side) && splitOne(pair.Child(side)) {
			return true
		}
	}

	return false
}

func split(val int) *Pair {
	var pair Pair

	pair.x = val / 2
	pair.y = (val + 1) / 2

	return &pair
}
