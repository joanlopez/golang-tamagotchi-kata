package tamagotchi

type Status int

const (
	DEAD Status = iota
	SICK
	ALIVE
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