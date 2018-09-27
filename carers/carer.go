package carers

type Pet interface {
	Care(c Carer)
}

type Carer interface {
	Notify(message string)
	TakeCare(pet Pet)
}

type HumanCarer struct {
	pets []Pet
}

func (human *HumanCarer) Notify(message string) {
	// Do whatever with the message
}

func (human *HumanCarer) TakeCare(pet Pet) {
	pet.Care(human)
	human.pets = append(human.pets, pet)
}