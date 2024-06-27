package main

import "fmt"

type Player struct {
	Health      int
	Name        string
	AttackPower int
}

func main() {
	player := Player{
		Health:      100,
		Name:        "DONNIE",
		AttackPower: 50,
	}
	killPlayer(&player)
	fmt.Printf("Player: %s, Health: %d\n", player.Name, player.Health)
}

func killPlayer(p *Player) {
	p.Health = 0
	fmt.Printf("Player %s is dead\n", p.Name)
}
