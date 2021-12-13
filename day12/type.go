package main

const (
	caveStart = Cave(`start`)
	caveEnd   = Cave(`end`)
)

type Cave string

func (c Cave) IsLarge() bool {
	if c == caveStart || c == caveEnd {
		return false
	}

	return c[0] >= 'A' && c[0] <= 'Z'
}

type Path struct {
	traversed map[Cave]bool
	caves     []Cave

	passedSmallTwice bool
}

func NewPath(start Cave) Path {
	return Path{caves: []Cave{start}}
}

func (p Path) Copy() Path {
	path := Path{
		traversed:        make(map[Cave]bool),
		passedSmallTwice: p.passedSmallTwice,
	}

	for k, v := range p.traversed {
		path.traversed[k] = v
	}

	path.caves = append(path.caves, p.caves...)

	return path
}

func (p *Path) Go(cave Cave) bool {
	if cave == caveStart {
		return false
	}

	if !cave.IsLarge() && p.traversed[cave] {
		if p.passedSmallTwice || !part2 {
			return false
		}

		p.passedSmallTwice = true
	}

	p.caves = append(p.caves, cave)
	p.traversed[cave] = true

	return true
}
