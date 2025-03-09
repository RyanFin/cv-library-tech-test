package model

import "fmt"

type Bee struct {
	Health, DamageToPlayer, DamageByPlayer int
	BeeType                                string
}

type QueenBee struct {
	Bee
}

type WorkerBee struct {
	Bee
}

type DroneBee struct {
	Bee
}

type Player struct {
	Health int
	Name   string
}

func NewQueenBee() QueenBee {
	return QueenBee{
		Bee: Bee{
			Health:         100,
			DamageToPlayer: 10,
			DamageByPlayer: 10,
			BeeType:        "Queen",
		},
	}
}

func NewWorkerBee() WorkerBee {
	return WorkerBee{
		Bee: Bee{
			Health:         75,
			DamageToPlayer: 5,
			DamageByPlayer: 25,
			BeeType:        "Worker",
		},
	}
}

func NewDroneBee() DroneBee {
	return DroneBee{
		Bee: Bee{
			Health:         60,
			DamageToPlayer: 1,
			DamageByPlayer: 30,
			BeeType:        "Drone",
		},
	}
}

func NewPlayer() Player {
	return Player{
		Health: 100,
	}
}

func (b *Bee) HitByPlayer() {
	fmt.Printf("Direct Hit! You took %v hit points from a %s Bee\n", b.DamageByPlayer, b.BeeType)
	b.Health -= b.DamageByPlayer
}

func (b *Bee) HitsPlayer(player *Player) {
	fmt.Printf("%s Bee deals %v damage to the player.", b.BeeType, b.DamageToPlayer)
	player.Health -= b.DamageToPlayer
}
