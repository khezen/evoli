package evoli

// Evolution - problem solving algorithm by exploring the range of solutions
type Evolution interface {
	Population() Population
	SetPopulation(Population)
	Next() error
	Evaluater() Evaluater
	Alpha() Individual
}

type evolution struct {
	pop       Population
	evaluater Evaluater
}

func newEvolution(pop Population, eval Evaluater) evolution {
	return evolution{pop, eval}
}
func (e *evolution) Population() Population {
	return e.pop
}

func (e *evolution) SetPopulation(pop Population) {
	e.pop = pop
}

func (e *evolution) Alpha() Individual {
	return e.pop.Max()
}

func (e *evolution) Evaluater() Evaluater {
	return e.evaluater
}
