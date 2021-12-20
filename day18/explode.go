package main

func getExplodable(pair *Pair) *Pair {
	return recurseGetExplodable(pair, 0)
}

func recurseGetExplodable(pair *Pair, nesting int) *Pair {
	if nesting == 4 {
		return pair
	}

	if pair.childX != nil {
		v := recurseGetExplodable(pair.childX, nesting+1)
		if v != nil {
			return v
		}
	}

	if pair.childY != nil {
		v := recurseGetExplodable(pair.childY, nesting+1)
		if v != nil {
			return v
		}
	}

	return nil
}

func explode(pair *Pair) {
	explodeUpwards(pair, pair.Value(X), X)
	explodeUpwards(pair, pair.Value(Y), Y)

	pair.parent.SetChild(pair.parentSide, nil)
	pair.parent.SetValue(pair.parentSide, 0)
}

func explodeUpwards(pair *Pair, val int, side Side) {
	if pair.parent == nil {
		return
	}

	if pair.parentSide != side {
		if !pair.parent.HasChild(side) {
			pair.parent.SetValue(side, pair.parent.Value(side)+val)
			return
		}

		explodeDownwards(pair.parent.Child(side), val, side)
		return
	}

	explodeUpwards(pair.parent, val, side)
}

func explodeDownwards(pair *Pair, val int, side Side) {
	opposite := side.Opposite()
	if !pair.HasChild(opposite) {
		pair.SetValue(opposite, pair.Value(opposite)+val)
		return
	}

	explodeDownwards(pair.Child(opposite), val, side)
}
