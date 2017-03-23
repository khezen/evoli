package darwin

// Individual is an interface for an individual
type Individual interface {
	Fitness() float32
	SetFitness(float32)
}

// individual https://en.wikipedia.org/wiki/individual
type individual float32

// NewIndividual is the constructor for individuals
func NewIndividual(Fitness float32) Individual {
	indiv := individual(Fitness)
	return &indiv
}

// Fitness returns the strength of a individual regarding to its environement. Higher is stronger.
func (indiv *individual) Fitness() float32 {
	return float32(*indiv)
}

// SetFitness set the strength of a individual regarding to its environement. Higher is stronger.
func (indiv *individual) SetFitness(Fitness float32) {
	*indiv = individual(Fitness)
}
