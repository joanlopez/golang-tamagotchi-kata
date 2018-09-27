package tamagotchi

type Status string

const (
	DEAD Status = "Dead"
	SICK Status = "Sick"
	ALIVE Status = "Alive"
)

func newStatus(lifePoints int) Status {
	if lifePoints == DEAD_LIFE {
		return DEAD
	}

	if lifePoints <= SICK_LIFE {
		return SICK
	}

	return ALIVE
}