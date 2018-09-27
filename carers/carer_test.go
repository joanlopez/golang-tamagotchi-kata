package carers

import (
	"testing"
)

func TestHumanCarer_TakeCare(t *testing.T) {
	human := HumanCarer{}
	pet := &FakePet{}
	human.TakeCare(pet)
	if pet.carers != 1 || len(human.pets) != 1 {
		t.Errorf("The human is taking care properly")
	}
}

type FakePet struct {
	carers int
}

func (p *FakePet) Care(c Carer) {
	p.carers++
}