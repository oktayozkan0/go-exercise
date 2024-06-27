package main

import "fmt"

type Player struct {
	Health      int
	Name        string
	AttackPower int
}

func (p *Player) die() {
	fmt.Printf("Killing player %s\n", p.Name)
	p.Health = 0
}

func main() {
	player := Player{
		Health:      100,
		Name:        "DONNIE",
		AttackPower: 50,
	}
	player.die()
	fmt.Printf("Player: %s, Health: %d\n", player.Name, player.Health)
}
