package tamagotchi

import "github.com/joanlopez/golang-tamagotchi-kata/food"

const INITIAL_LIFE = 100
const SICK_LIFE = 25
const DEAD_LIFE = 0

const MAX_LIFE = 100

type Tamagotchi struct {
	life int
	status Status
}

func NewTamagotchi() *Tamagotchi {
	return &Tamagotchi{INITIAL_LIFE, newStatus(INITIAL_LIFE)}
}

func (t *Tamagotchi) Eat(eatable food.Eatable) {
	t.life += eatable.Eat()
	if t.life > MAX_LIFE {
		t.getSick()
	} else {
		t.status = newStatus(t.life)
	}
}

func (t *Tamagotchi) Life() int {
	return t.life
}

func (t *Tamagotchi) IsAlive() bool {
	return t.status == ALIVE
}

func (t *Tamagotchi) IsSick() bool {
	return t.status == SICK
}

func (t *Tamagotchi) IsDead() bool {
	return t.status == DEAD
}

func (t *Tamagotchi) setLife(lifePoints int) {
	t.life = lifePoints
	t.status = newStatus(lifePoints)
}

func (t *Tamagotchi) getSick() {
	t.setLife(SICK_LIFE)
}