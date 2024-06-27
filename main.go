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
	fmt.Printf("player name: %s - attack power: %d", player.Name, player.AttackPower)
}
