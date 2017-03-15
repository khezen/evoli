package individual

// Interface is an interface for an individual
type Interface interface {
	Fitness() float32
	SetFitness(float32)
}

// Individual https://en.wikipedia.org/wiki/individual
type Individual float32

// New is the constructor for individuals
func New(Fitness float32) *Individual {
	indiv := Individual(Fitness)
	return &indiv
}

// Fitness returns the strength of a individual regarding to its environement. Higher is stronger.
func (indiv *Individual) Fitness() float32 {
	return float32(*indiv)
}

// SetFitness set the strength of a individual regarding to its environement. Higher is stronger.
func (indiv *Individual) SetFitness(Fitness float32) {
	*indiv = Individual(Fitness)
}
