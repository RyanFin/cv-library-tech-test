package model

type Bee struct {
	Health, DamageToPlayer, DamageByPlayer int
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

func NewQueenBee() QueenBee {
	return QueenBee{
		Bee: Bee{
			Health:         100,
			DamageToPlayer: 10,
			DamageByPlayer: 10,
		},
	}
}

func NewWorkerBee() WorkerBee {
	return WorkerBee{
		Bee: Bee{
			Health:         75,
			DamageToPlayer: 5,
			DamageByPlayer: 25,
		},
	}
}

func NewDroneBee() DroneBee {
	return DroneBee{
		Bee: Bee{
			Health:         60,
			DamageToPlayer: 1,
			DamageByPlayer: 30,
		},
	}
}

func (b *Bee) HitByPlayer() {
	b.Health -= b.DamageByPlayer
}
