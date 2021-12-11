package main

type flasher struct {
	octopi  *Octopi
	flashed map[point]bool
	count   int
}

func newFlasher(octopi *Octopi) *flasher {
	return &flasher{octopi: octopi, flashed: make(map[point]bool)}
}

func (f *flasher) flash() int {
	for _, point := range f.octopi.points() {
		f.recurseFlash(point)
	}

	return f.count
}

func (f *flasher) recurseFlash(p point) {
	if !f.octopi.get(p).flashes() || f.hasFlashed(p) {
		return
	}

	f.count++
	f.setFlashed(p)

	for _, p := range p.adjacents() {
		f.octopi.increase(p)
		f.recurseFlash(p)
	}
}

func (f *flasher) hasFlashed(p point) bool {
	return f.flashed[p]
}

func (f *flasher) setFlashed(p point) {
	f.flashed[p] = true
}
