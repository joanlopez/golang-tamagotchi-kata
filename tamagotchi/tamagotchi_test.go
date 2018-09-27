package tamagotchi

import (
	"github.com/joanlopez/golang-tamagotchi-kata/food"
	"testing"
	"github.com/joanlopez/golang-tamagotchi-kata/carers"
	"reflect"
)

type eatTC struct {
	name           string
	tamagotchi     *Tamagotchi
	eatable        food.Eatable
	expectedLife   int
	expectedStatus Status
}

var eatTCs = []eatTC{
	{"tired eat tangerine", tiredTamagotchi(), food.NewTangerine(), 50, ALIVE},
	{"tired eat chocolate", tiredTamagotchi(), food.NewChocolate(), 60, ALIVE},
	{"tired eat ham", tiredTamagotchi(), food.NewHam(), 90, ALIVE},
	{"tired eat lettuce", tiredTamagotchi(), food.NewLettuce(), 45, ALIVE},
	{"rested eat something", tamagotchiWith(100), food.NewLettuce(), 25, SICK},
	{"bit tired eat ham", tamagotchiWith(90), food.NewChocolate(), 25, SICK},
}

func TestTamagotchi_Eat(t *testing.T) {
	for _, tc := range eatTCs {
		t.Run(tc.name, func(t *testing.T) {
			tc.tamagotchi.Eat(tc.eatable)
			if tc.expectedLife != tc.tamagotchi.Life() {
				t.Errorf("life got: %v; expected: %v", tc.tamagotchi.Life(), tc.expectedLife)
			}
			if tc.expectedStatus != tc.tamagotchi.status {
				t.Errorf("status got: %v; expected: %v", tc.tamagotchi.status, tc.expectedStatus)
			}
		})
	}
}

func TestTamagotchi_Care(t *testing.T) {
	expectedNotifications := []string {
		"Status changed from: Alive; to: Sick",
	}
	human := &FakeHuman{}
	tamagotchi := tamagotchiWith(MAX_LIFE)
	human.TakeCare(tamagotchi)
	tamagotchi.Eat(food.NewLettuce())
	if !reflect.DeepEqual(human.receivedNotifications, expectedNotifications) {
		t.Errorf("The tamagotchi's carer was not notified properly")
	}
}

// Private types and functions

type FakeHuman struct {
	receivedNotifications []string
}

func (h *FakeHuman) Notify(message string) {
	h.receivedNotifications = append(h.receivedNotifications, message)
}

func (h *FakeHuman) TakeCare(pet carers.Pet) {
	pet.Care(h)
}

func tiredTamagotchi() *Tamagotchi {
	return &Tamagotchi{40, ALIVE, []carers.Carer{}}
}

func tamagotchiWith(lifePoints int) *Tamagotchi {
	return &Tamagotchi{lifePoints, ALIVE, []carers.Carer{}}
}
