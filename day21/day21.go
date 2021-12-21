package main

type Player struct {
	Position int
	Score    int
}

func (p *Player) Move(steps int) {
	p.Position = (p.Position+steps-1)%10 + 1
	p.Score += p.Position
}

func main() {
	players := [2]Player{
		{Position: 8},
		{Position: 3},
	}

	part1(players)

	part2(players)
}
