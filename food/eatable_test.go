package food

import "testing"

func TestFood_Eat_ReturnsFulfil(t *testing.T) {
	fulfil := 10
	eatable := &Food{"Fake food", fulfil}
	if fulfil != eatable.Eat() {
		t.Errorf("The returned value when eating food is not the fulfil")
	}
}