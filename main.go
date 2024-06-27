package main

import "fmt"

type Player struct {
	Name        string
	Health      uint32
	AttackPower int
}

func NewPlayer(name string, hp uint32, ap int) *Player {
	return &Player{
		Name:        name,
		Health:      hp,
		AttackPower: ap,
	}
}

func (p *Player) die() {
	fmt.Printf("Killing player %s\n", p.Name)
	p.Health = 0
}

func main() {
	player := NewPlayer("DONNIE", 100, 50)
	player.die()
	fmt.Printf("Player: %s, Health: %d\n", player.Name, player.Health)
}
