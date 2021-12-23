package main

import (
	"testing"

	"git.fuyu.moe/Fuyu/assert"
)

var intersectCases = []struct {
	name                 string
	left, right          Cuboid
	expectedLeft         Set
	expectIntersection   bool
	expectedIntersection Cuboid
	expectedRight        Set
}{
	{
		name: `no intersection`,
		left: defaultCuboid(nil),
		right: defaultCuboid(func(c *Cuboid) {
			c.X = Range{Min: 11, Max: 20}
		}),
		expectedLeft:         Set{defaultCuboid(nil): struct{}{}},
		expectIntersection:   false,
		expectedIntersection: Cuboid{},
		expectedRight:        Set{defaultCuboid(func(c *Cuboid) { c.X = Range{Min: 11, Max: 20} }): struct{}{}},
	},
	{
		name:                 `equal`,
		left:                 defaultCuboid(nil),
		right:                defaultCuboid(nil),
		expectedLeft:         Set{},
		expectIntersection:   true,
		expectedIntersection: defaultCuboid(nil),
		expectedRight:        Set{},
	},
	{
		name:  `embedded`,
		left:  defaultCuboid(nil),
		right: Cuboid{X: Range{Min: -1, Max: 1}, Y: Range{Min: -1, Max: 1}, Z: Range{Min: -1, Max: 1}},
		expectedLeft: Set{
			// corners
			defaultCuboid(func(c *Cuboid) { c.X.Min = 2; c.Y.Min = 2; c.Z.Min = 2 }):    struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Min = 2; c.Y.Min = 2; c.Z.Max = -2 }):   struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Min = 2; c.Y.Max = -2; c.Z.Min = 2 }):   struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Max = -2; c.Y.Min = 2; c.Z.Min = 2 }):   struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Min = 2; c.Y.Max = -2; c.Z.Max = -2 }):  struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Max = -2; c.Y.Max = -2; c.Z.Min = 2 }):  struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Max = -2; c.Y.Min = 2; c.Z.Max = -2 }):  struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Max = -2; c.Y.Max = -2; c.Z.Max = -2 }): struct{}{},

			// faces
			defaultCuboid(func(c *Cuboid) { c.X.Min = -1; c.X.Max = 1; c.Y.Min = -1; c.Y.Max = 1; c.Z.Max = -2 }): struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Min = -1; c.X.Max = 1; c.Y.Min = -1; c.Y.Max = 1; c.Z.Min = 2 }):  struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Min = -1; c.X.Max = 1; c.Y.Max = -2; c.Z.Min = -1; c.Z.Max = 1 }): struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Min = -1; c.X.Max = 1; c.Y.Min = 2; c.Z.Min = -1; c.Z.Max = 1 }):  struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Min = 2; c.Y.Min = -1; c.Y.Max = 1; c.Z.Min = -1; c.Z.Max = 1 }):  struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Max = -2; c.Y.Min = -1; c.Y.Max = 1; c.Z.Min = -1; c.Z.Max = 1 }): struct{}{},

			// edges
			defaultCuboid(func(c *Cuboid) { c.X.Min = -1; c.X.Max = 1; c.Y.Min = 2; c.Z.Min = 2 }):   struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Min = -1; c.X.Max = 1; c.Y.Min = 2; c.Z.Max = -2 }):  struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Min = -1; c.X.Max = 1; c.Y.Max = -2; c.Z.Min = 2 }):  struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Min = -1; c.X.Max = 1; c.Y.Max = -2; c.Z.Max = -2 }): struct{}{},

			defaultCuboid(func(c *Cuboid) { c.Y.Min = -1; c.Y.Max = 1; c.X.Min = 2; c.Z.Min = 2 }):   struct{}{},
			defaultCuboid(func(c *Cuboid) { c.Y.Min = -1; c.Y.Max = 1; c.X.Min = 2; c.Z.Max = -2 }):  struct{}{},
			defaultCuboid(func(c *Cuboid) { c.Y.Min = -1; c.Y.Max = 1; c.X.Max = -2; c.Z.Min = 2 }):  struct{}{},
			defaultCuboid(func(c *Cuboid) { c.Y.Min = -1; c.Y.Max = 1; c.X.Max = -2; c.Z.Max = -2 }): struct{}{},

			defaultCuboid(func(c *Cuboid) { c.Z.Min = -1; c.Z.Max = 1; c.Y.Min = 2; c.X.Min = 2 }):   struct{}{},
			defaultCuboid(func(c *Cuboid) { c.Z.Min = -1; c.Z.Max = 1; c.Y.Min = 2; c.X.Max = -2 }):  struct{}{},
			defaultCuboid(func(c *Cuboid) { c.Z.Min = -1; c.Z.Max = 1; c.Y.Max = -2; c.X.Min = 2 }):  struct{}{},
			defaultCuboid(func(c *Cuboid) { c.Z.Min = -1; c.Z.Max = 1; c.Y.Max = -2; c.X.Max = -2 }): struct{}{},
		},
		expectIntersection:   true,
		expectedIntersection: Cuboid{X: Range{Min: -1, Max: 1}, Y: Range{Min: -1, Max: 1}, Z: Range{Min: -1, Max: 1}},
		expectedRight:        Set{},
	},
	{
		name:  `corner intersection`,
		left:  defaultCuboid(nil),
		right: Cuboid{X: Range{0, 20, false}, Y: Range{0, 20, false}, Z: Range{0, 20, false}},
		expectedLeft: Set{
			defaultCuboid(func(c *Cuboid) { c.X.Max = -1; c.Y.Max = -1; c.Z.Max = -1 }): struct{}{},

			defaultCuboid(func(c *Cuboid) { c.X.Min = 0; c.Y.Max = -1; c.Z.Max = -1 }): struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Max = -1; c.Y.Min = 0; c.Z.Max = -1 }): struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Max = -1; c.Y.Max = -1; c.Z.Min = 0 }): struct{}{},

			defaultCuboid(func(c *Cuboid) { c.X.Max = -1; c.Y.Min = 0; c.Z.Min = 0 }): struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Min = 0; c.Y.Max = -1; c.Z.Min = 0 }): struct{}{},
			defaultCuboid(func(c *Cuboid) { c.X.Min = 0; c.Y.Min = 0; c.Z.Max = -1 }): struct{}{},
		},
		expectIntersection:   true,
		expectedIntersection: Cuboid{X: Range{0, 10, false}, Y: Range{0, 10, false}, Z: Range{0, 10, false}},
		expectedRight: Set{
			Cuboid{X: Range{11, 20, false}, Y: Range{11, 20, false}, Z: Range{11, 20, false}}: struct{}{},

			Cuboid{X: Range{0, 10, false}, Y: Range{11, 20, false}, Z: Range{11, 20, false}}: struct{}{},
			Cuboid{X: Range{11, 20, false}, Y: Range{0, 10, false}, Z: Range{11, 20, false}}: struct{}{},
			Cuboid{X: Range{11, 20, false}, Y: Range{11, 20, false}, Z: Range{0, 10, false}}: struct{}{},

			Cuboid{X: Range{0, 10, false}, Y: Range{11, 20, false}, Z: Range{0, 10, false}}: struct{}{},
			Cuboid{X: Range{0, 10, false}, Y: Range{0, 10, false}, Z: Range{11, 20, false}}: struct{}{},
			Cuboid{X: Range{11, 20, false}, Y: Range{0, 10, false}, Z: Range{0, 10, false}}: struct{}{},
		},
	},
}

func TestIntersect(t *testing.T) {
	for _, c := range intersectCases {
		t.Run(c.name, func(t *testing.T) {
			assert := assert.New(t)

			actualLeft, actualIntersection, actualRight := intersect(c.left, c.right)

			if c.expectIntersection {
				assert.Eq(c.expectedIntersection, actualIntersection)
			} else {
				assert.Eq(int64(0), actualIntersection.Size())
			}
			assert.Cmp(c.expectedLeft, actualLeft)
			assert.Cmp(c.expectedRight, actualRight)
		})
	}
}

func defaultCuboid(f func(*Cuboid)) Cuboid {
	cuboid := Cuboid{
		X: Range{Min: -10, Max: 10},
		Y: Range{Min: -10, Max: 10},
		Z: Range{Min: -10, Max: 10},
	}

	if f == nil {
		return cuboid
	}

	f(&cuboid)
	return cuboid
}
