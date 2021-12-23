package main

import "fmt"

func part2(steps []RebootStep) {
	cuboidsOn := Set{}

	for i, step := range steps {
		fmt.Printf("(%3d/%3d)\n", i+1, len(steps))

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

func intersect(c1, c2 Cuboid) (Set, Cuboid, Set) {
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

	for dim := range dimensions {
		split(cuboids, splitter[dim].Min, dim)
		split(cuboids, splitter[dim].Max+1, dim)
	}

	return cuboids
}

func split(cuboids Set, v int64, dim int) {
	newCuboids := Set{}

	for cuboid := range cuboids {
		if !(cuboid[dim].Has(v) && cuboid[dim].Has(v-1)) {
			continue
		}
		cuboids.Delete(cuboid)

		newCuboid := cuboid

		newCuboid[dim] = NewRange(cuboid[dim].Min, v-1)
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
