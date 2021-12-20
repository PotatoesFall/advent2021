package main

func reduce(pair *Pair) {
	for reduceStep(pair) {
	}
}

func reduceStep(pair *Pair) bool {
	explodable := getExplodable(pair)
	if explodable != nil {
		explode(explodable)
		return true
	}

	return splitOne(pair)
}
