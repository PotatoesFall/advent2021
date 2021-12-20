package main

import (
	"fmt"
	"strings"
	"testing"

	"git.fuyu.moe/Fuyu/assert"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestAlign(t *testing.T) {
	cases := []struct {
		name            string
		requiredOverlap int
		base, scanner   *Scanner
		expected        *Scanner
	}{
		{
			name:            `no translation, simple rotation`,
			requiredOverlap: 2,
			base: &Scanner{
				Position: Position{0, 0, 0},
				Beacons:  []Position{{1, 2, 3}, {6, 5, 4}},
			},
			scanner: &Scanner{
				Beacons: []Position{{2, 3, 1}, {5, 4, 6}},
			},
			expected: &Scanner{
				Position: Position{0, 0, 0},
				Beacons:  []Position{{1, 2, 3}, {6, 5, 4}},
			},
		},
		{
			name:            `no translation, simple rotation - swapped`,
			requiredOverlap: 3,
			base: &Scanner{
				Position: Position{0, 0, 0},
				Beacons:  []Position{{1, 2, 3}, {6, 5, 4}, {9, 7, 5}},
			},
			scanner: &Scanner{
				Beacons: []Position{{5, 4, 6}, {2, 3, 1}, {7, 5, 9}},
			},
			expected: &Scanner{
				Position: Position{0, 0, 0},
				Beacons:  []Position{{1, 2, 3}, {6, 5, 4}, {9, 7, 5}},
			},
		},
		{
			name:            `simple translation, no rotation`,
			requiredOverlap: 3,
			base: &Scanner{
				Position: Position{0, 0, 0},
				Beacons:  []Position{{1, 2, 3}, {6, 5, 4}, {11, 12, 24}},
			},
			scanner: &Scanner{
				Beacons: []Position{{2, 3, 4}, {7, 6, 5}, {12, 13, 25}},
			},
			expected: &Scanner{
				Position: Position{-1, -1, -1},
				Beacons:  []Position{{1, 2, 3}, {6, 5, 4}, {11, 12, 24}},
			},
		},
		{
			name:            `simple translation, no rotation - swapped`,
			requiredOverlap: 3,
			base: &Scanner{
				Position: Position{0, 0, 0},
				Beacons:  []Position{{1, 2, 3}, {6, 5, 4}, {11, 12, 24}},
			},
			scanner: &Scanner{
				Beacons: []Position{{7, 6, 5}, {12, 13, 25}, {2, 3, 4}},
			},
			expected: &Scanner{
				Position: Position{-1, -1, -1},
				Beacons:  []Position{{1, 2, 3}, {6, 5, 4}, {11, 12, 24}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Println()
			fmt.Println(c.name)

			assert := assert.New(t)
			requiredOverlap = c.requiredOverlap

			computeDiffs(c.base)
			computeDiffs(c.scanner)
			original := c.scanner.Clone()
			align(c.base, c.scanner)

			fmt.Println(c.base.Beacons)
			fmt.Println(c.scanner.Beacons)

			assert.Cmp(c.expected, c.scanner,
				cmpopts.IgnoreFields(Scanner{}, `Diffs`),
				cmpopts.SortSlices(func(a interface{}, b interface{}) bool {
					return strings.Compare(fmt.Sprint(a), fmt.Sprint(b)) < 0
				}),
			)

			assert.Cmp(original.Diffs, c.scanner.Diffs,
				cmpopts.IgnoreFields(DiffVector{}, `Vector`, `Parent`),
				cmpopts.SortSlices(func(a interface{}, b interface{}) bool {
					return strings.Compare(fmt.Sprint(a), fmt.Sprint(b)) < 0
				}),
			)
		})
	}
}
