package tamagotchi

import (
	"github.com/joanlopez/golang-tamagotchi-kata/food"
	"github.com/joanlopez/golang-tamagotchi-kata/carers"
	"fmt"
)

const INITIAL_LIFE = 100
const SICK_LIFE = 25
const DEAD_LIFE = 0

const MAX_LIFE = 100

type Tamagotchi struct {
	life int
	status Status
	carers []carers.Carer
}

func NewTamagotchi() *Tamagotchi {
	return &Tamagotchi{INITIAL_LIFE, newStatus(INITIAL_LIFE), []carers.Carer{}}
}

func (t *Tamagotchi) Eat(eatable food.Eatable) {
	t.setLife(t.life + eatable.Eat())
	if t.life > MAX_LIFE {
		t.getSick()
	} else {
		t.setStatus(newStatus(t.life))
	}
}

func (t *Tamagotchi) Life() int {
	return t.life
}

func (t *Tamagotchi) Care(carer carers.Carer) {
	t.carers = append(t.carers, carer)
}

func (t *Tamagotchi) getSick() {
	t.setLife(SICK_LIFE)
}

func (t *Tamagotchi) setLife(lifePoints int) {
	t.life = lifePoints
	t.setStatus(newStatus(lifePoints))
}

func (t *Tamagotchi) setStatus(newStatus Status) {
	if t.status == newStatus {
		return
	}
	t.notifyCarers(t.status, newStatus)
	t.status = newStatus
}

func (t *Tamagotchi) notifyCarers(prevStatus, nextStatus Status) {
	message := fmt.Sprintf("Status changed from: %v; to: %v", prevStatus, nextStatus)
	for _, carer := range t.carers {
		carer.Notify(message)
	}
}