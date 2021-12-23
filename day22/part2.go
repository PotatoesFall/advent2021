package main

import "fmt"

func part2(steps []RebootStep) {
	cuboidsOn := Set{}

	for _, step := range steps {
		// fmt.Printf("(%3d/%3d)\n", i+1, len(steps))

		if step.On {
			applyOnStep(cuboidsOn, step.Cuboid)
		} else {
			applyOffStep(cuboidsOn, step.Cuboid)
		}
	}

	count := countCubes(cuboidsOn)

	fmt.Printf("Part 2 - %d cubes are lit.\n", count)
}

func countCubes(cuboids Set) int64 {
	var count int64 = 0
	for cuboid := range cuboids {
		count += cuboid.Size()
	}

	return count
}

func applyOffStep(cuboidsOn Set, stepCuboid Cuboid) {
	for cuboid := range cuboidsOn {
		if !cuboid.Intersects(stepCuboid) {
			continue
		}

		left, _, _ := intersect(cuboid, stepCuboid)

		// remove old large cuboid
		cuboidsOn.Delete(cuboid)

		// add new smaller cuboids, without intersection
		for c := range left {
			cuboidsOn.Put(c)
		}
	}
}

func applyOnStep(cuboidsOn Set, stepCuboid Cuboid) {
	newCuboidsOn := Set{}
	newCuboidsOn.Put(stepCuboid)

	for cuboid := range cuboidsOn {
		newNewCuboidsOn := Set{}

		for newCuboid := range newCuboidsOn {
			if !cuboid.Intersects(newCuboid) {
				continue
			}

			// get intersection cuboids
			_, _, right := intersect(cuboid, newCuboid)

			// remove former cuboid
			newCuboidsOn.Delete(newCuboid)

			// add new smaller cuboids
			for c := range right {
				newNewCuboidsOn.Put(c)
			}
		}

		for cuboid := range newNewCuboidsOn {
			newCuboidsOn.Put(cuboid)
		}
	}

	for cuboid := range newCuboidsOn {
		cuboidsOn.Put(cuboid)
	}
}

func intersect(c1, c2 Cuboid) (Set, Cuboid, Set) { //nolint:unparam
	left := splitByCuboid(c1, c2)
	right := splitByCuboid(c2, c1)
	intersect := intersection(c1, c2)

	left.Delete(intersect)
	right.Delete(intersect)

	return left, intersect, right
}

func splitByCuboid(target Cuboid, splitter Cuboid) Set {
	cuboids := Set{}
	cuboids.Put(target)

	// split into up to 27 pieces
	for dim := range dimensions {
		split(cuboids, splitter[dim].Min, dim)
		split(cuboids, splitter[dim].Max, dim)
	}

	// improves performance by undoing unnecessary split
	attemptMergeCuboids(cuboids)

	return cuboids
}

func split(cuboids Set, v int64, dim int) {
	newCuboids := Set{}

	for cuboid := range cuboids {
		if !(cuboid[dim].Has(v)) {
			continue
		}
		cuboids.Delete(cuboid)

		newCuboid := cuboid

		newCuboid[dim] = NewRange(cuboid[dim].Min, v)
		newCuboids.Put(newCuboid)

		newCuboid[dim] = NewRange(v, cuboid[dim].Max)
		newCuboids.Put(newCuboid)
	}

	for cuboid := range newCuboids {
		cuboids.Put(cuboid)
	}
}

func intersection(c1, c2 Cuboid) Cuboid {
	cuboid := Cuboid{}

	for dim := range dimensions {
		cuboid[dim] = NewRange(max(c1[dim].Min, c2[dim].Min), min(c1[dim].Max, c2[dim].Max))
	}

	return cuboid
}

func attemptMergeCuboids(cuboids Set) {
outer:
	for {
		for cuboid1 := range cuboids {
			for cuboid2 := range cuboids {
				cuboid, success := attemptMergeTwoCuboids(cuboid1, cuboid2)
				if success {
					cuboids.Delete(cuboid1)
					cuboids.Delete(cuboid2)
					cuboids.Put(cuboid)
					continue outer
				}
			}
		}

		break
	}
}

func attemptMergeTwoCuboids(c1, c2 Cuboid) (Cuboid, bool) {
	for dim1 := range dimensions {
		if c1[dim1] == c2[dim1] {
			for dim2 := range dimensions {
				if dim1 == dim2 {
					continue
				}

				if c1[dim2] == c2[dim2] {
					for dim3 := range dimensions {
						if dim3 == dim2 {
							continue
						}

						var cuboid Cuboid

						cuboid[dim1] = c1[dim1]
						cuboid[dim2] = c1[dim2]
						cuboid[dim3] = NewRange(min(c1[dim3].Min, c2[dim3].Min), max(c1[dim3].Max, c2[dim3].Max))

						return cuboid, true
					}
				}
			}
		}
	}

	return Cuboid{}, false
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}

	return a
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}
