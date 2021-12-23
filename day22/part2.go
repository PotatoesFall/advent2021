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

	splitX(cuboids, splitter.X.Min)
	splitX(cuboids, splitter.X.Max+1)

	splitY(cuboids, splitter.Y.Min)
	splitY(cuboids, splitter.Y.Max+1)

	splitZ(cuboids, splitter.Z.Min)
	splitZ(cuboids, splitter.Z.Max+1)

	return cuboids
}

func splitX(cuboids Set, x int64) {
	newCuboids := Set{}

	for cuboid := range cuboids {
		if !(cuboid.X.Has(x) && cuboid.X.Has(x-1)) {
			continue
		}
		cuboids.Delete(cuboid)

		newCuboids.Put(Cuboid{
			X: NewRange(cuboid.X.Min, x-1),
			Y: cuboid.Y,
			Z: cuboid.Z,
		})

		newCuboids.Put(Cuboid{
			X: NewRange(x, cuboid.X.Max),
			Y: cuboid.Y,
			Z: cuboid.Z,
		})
	}

	for cuboid := range newCuboids {
		cuboids.Put(cuboid)
	}
}

func splitY(cuboids Set, y int64) {
	newCuboids := Set{}

	for cuboid := range cuboids {
		if !(cuboid.Y.Has(y) && cuboid.Y.Has(y-1)) {
			continue
		}
		cuboids.Delete(cuboid)

		newCuboids.Put(Cuboid{
			X: cuboid.X,
			Y: NewRange(cuboid.Y.Min, y-1),
			Z: cuboid.Z,
		})

		newCuboids.Put(Cuboid{
			X: cuboid.X,
			Y: NewRange(y, cuboid.Y.Max),
			Z: cuboid.Z,
		})
	}

	for cuboid := range newCuboids {
		cuboids.Put(cuboid)
	}
}

func splitZ(cuboids Set, z int64) {
	newCuboids := Set{}

	for cuboid := range cuboids {
		if !(cuboid.Z.Has(z) && cuboid.Z.Has(z-1)) {
			continue
		}
		cuboids.Delete(cuboid)

		newCuboids.Put(Cuboid{
			X: cuboid.X,
			Y: cuboid.Y,
			Z: NewRange(cuboid.Z.Min, z-1),
		})

		newCuboids.Put(Cuboid{
			X: cuboid.X,
			Y: cuboid.Y,
			Z: NewRange(z, cuboid.Z.Max),
		})
	}

	for cuboid := range newCuboids {
		cuboids.Put(cuboid)
	}
}

func intersection(c1, c2 Cuboid) Cuboid {
	return Cuboid{
		X: NewRange(max(c1.X.Min, c2.X.Min), min(c1.X.Max, c2.X.Max)),
		Y: NewRange(max(c1.Y.Min, c2.Y.Min), min(c1.Y.Max, c2.Y.Max)),
		Z: NewRange(max(c1.Z.Min, c2.Z.Min), min(c1.Z.Max, c2.Z.Max)),
	}
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
