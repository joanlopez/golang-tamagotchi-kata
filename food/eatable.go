package food

type Eatable interface {
	Eat() int
}

type Food struct {
	name string
	fulfil int
}

func (f Food) Eat() int {
	// We also could define that it was eaten.
	return f.fulfil
}

func NewTangerine() Eatable {
	return Food{"Tangerine", 10}
}

func NewChocolate() Eatable {
	return Food{"Chocolate", 20}
}

func NewHam() Eatable {
	return Food{"Ham", 50}
}

func NewLettuce() Eatable {
	return Food{"Lettuce", 5}
}