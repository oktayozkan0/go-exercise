package main

import (
	"fmt"
	"time"
)

type Game struct {
	isRunning bool
	isPaused  bool
	players   map[string]*Player
	quitCh    chan bool
	pauseCh   chan bool
}

func NewGame() *Game {
	return &Game{
		isRunning: false,
		isPaused:  false,
		players:   make(map[string]*Player),
		quitCh:    make(chan bool),
		pauseCh:   make(chan bool),
	}
}

func (g *Game) Start() {
	g.isRunning = true
	g.gameLoop()
}

func (g *Game) addPlayer(p *Player) {
	g.players[p.Name] = p
	fmt.Printf("Added new player: %s\n", p.Name)
}

func (g *Game) gameLoop() {
	interval := time.Second
	ticker := time.NewTicker(interval)

running:
	for {
		select {
		case <-g.quitCh:
			break running
		case <-g.pauseCh:
			g.isPaused = true
		case <-ticker.C:
			fmt.Println("game is running")
		}
	}
	fmt.Println("game quitted by quitGame")
}

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
	game := NewGame()
	playerA := NewPlayer("DONNIE", 100, 50)
	playerB := NewPlayer("DARCO", 100, 50)
	game.addPlayer(playerA)
	game.addPlayer(playerB)

	go quitGame(game.quitCh)
	game.Start()
}

func quitGame(quitCh chan bool) {
	time.Sleep(time.Second * 5)
	quitCh <- true
}
