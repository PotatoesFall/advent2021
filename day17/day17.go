package main

import (
	"fmt"
	"sort"
)

const (
	minX, maxX = 57, 116
	minY, maxY = -198, -148
)

type Trajectory struct {
	StartVelX, StartVelY int
	Peak                 int
}

func main() {
	var trajectories []Trajectory

	for velX := 1; velX <= maxX; velX++ {
		var peakY int

	outer:
		for velY := minY; velY <= -minY; velY++ {
			probe := NewProbe(velX, velY)

			peakY = 0
			for probe.Step() {
				if probe.PosY > peakY {
					peakY = probe.PosY
				}

				if inBox(probe) {
					trajectories = append(trajectories, Trajectory{
						StartVelX: velX,
						StartVelY: velY,
						Peak:      peakY,
					})
					continue outer
				}

				if probe.PosY < minY {
					continue outer
				}
			}
		}
	}

	sort.Slice(trajectories, func(i, j int) bool {
		return trajectories[i].Peak < trajectories[j].Peak
	})
	highest := trajectories[len(trajectories)-1]
	fmt.Printf("Part 1 - velocity X %d and velocity Y %d reaches a peak of %d\n",
		highest.StartVelX, highest.StartVelY, highest.Peak)

	fmt.Printf("Part 2 - there are %d trajectories.", len(trajectories))
}

func inBox(p Probe) bool {
	return p.PosX >= minX &&
		p.PosX <= maxX &&
		p.PosY >= minY &&
		p.PosY <= maxY
}

type Probe struct {
	PosX, PosY int
	VelX, VelY int
}

func (p *Probe) Step() bool {
	p.PosX += p.VelX
	p.PosY += p.VelY

	if p.VelX > 0 {
		p.VelX -= 1
	}
	if p.VelX < 0 {
		p.VelX += 1
	}

	p.VelY -= 1

	return true
}

func NewProbe(velX, velY int) Probe {
	return Probe{
		VelX: velX,
		VelY: velY,
	}
}
