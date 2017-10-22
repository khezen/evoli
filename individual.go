package darwin

import "sync"

// Individual is an interface for an individual
type Individual interface {
	Fitness() float64
	SetFitness(float64)
	Equal(Individual) bool
}

// individual https://en.wikipedia.org/wiki/individual
type individual float64

// NewIndividual is the constructor for individuals
func NewIndividual(Fitness float64) Individual {
	indiv := individual(Fitness)
	return &indiv
}

// Fitness returns the strength of a individual regarding to its environement. Higher is stronger.
func (indiv *individual) Fitness() float64 {
	return float64(*indiv)
}

// SetFitness set the strength of a individual regarding to its environement. Higher is stronger.
func (indiv *individual) SetFitness(Fitness float64) {
	*indiv = individual(Fitness)
}

// Equal return true if indiv is equal to toBeCompared
func (indiv *individual) Equal(toBeCompared Individual) bool {
	return indiv == toBeCompared
}

// individual https://en.wikipedia.org/wiki/individual
type individualTS struct {
	individual
	mut sync.RWMutex
}

// NewIndividualTS is the constructor for threadsafe individuals
func NewIndividualTS(Fitness float64) Individual {
	indiv := individual(Fitness)
	return &individualTS{
		indiv,
		sync.RWMutex{},
	}
}

// Fitness returns the strength of a individual regarding to its environement. Higher is stronger.
func (indiv *individualTS) Fitness() float64 {
	indiv.mut.RLock()
	defer indiv.mut.RUnlock()
	return indiv.individual.Fitness()
}

// SetFitness set the strength of a individual regarding to its environement. Higher is stronger.
func (indiv *individualTS) SetFitness(Fitness float64) {
	indiv.mut.Lock()
	defer indiv.mut.Unlock()
	indiv.individual.SetFitness(Fitness)
}
