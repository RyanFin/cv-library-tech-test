package model

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
			BeeType:        "QueenBee",
		},
	}
}

func NewWorkerBee() WorkerBee {
	return WorkerBee{
		Bee: Bee{
			Health:         75,
			DamageToPlayer: 5,
			DamageByPlayer: 25,
			BeeType:        "WorkerBee",
		},
	}
}

func NewDroneBee() DroneBee {
	return DroneBee{
		Bee: Bee{
			Health:         60,
			DamageToPlayer: 1,
			DamageByPlayer: 30,
			BeeType:        "DroneBee",
		},
	}
}

func NewPlayer() Player {
	return Player{
		Health: 100,
	}
}

func (b *Bee) HitByPlayer() {
	b.Health -= b.DamageByPlayer
}
