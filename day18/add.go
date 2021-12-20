package main

func addAll(pairs []*Pair) *Pair {
	lastPair := pairs[0]

	for _, pair := range pairs[1:] {
		lastPair = add(lastPair, pair)
	}

	return lastPair
}

func add(x, y *Pair) *Pair {
	pair := &Pair{}

	x.parent = pair
	x.parentSide = X

	y.parent = pair
	y.parentSide = Y

	pair.SetChild(X, x)
	pair.SetChild(Y, y)

	reduce(pair)

	return pair
}
