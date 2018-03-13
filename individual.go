package evoli

import "sync"

// Individual is an interface for an individual
type Individual interface {
	Fitness() float64
	SetFitness(float64)
	Equal(Individual) bool
}

// individual https://en.wikipedia.org/wiki/individual
type individual struct {
	fitness float64
	best    Individual
}

// NewIndividual is the constructor for individuals
func NewIndividual(fitness float64) Individual {
	indiv := individual{
		fitness,
		nil,
	}
	return &indiv
}

// Fitness returns the strength of a individual regarding to its environement. Higher is stronger.
func (indiv *individual) Fitness() float64 {
	return indiv.fitness
}

// SetFitness set the strength of a individual regarding to its environement. Higher is stronger.
func (indiv *individual) SetFitness(fitness float64) {
	indiv.fitness = fitness
}

// Equal return true if indiv is equal to toBeCompared
func (indiv *individual) Equal(toBeCompared Individual) bool {
	return indiv == toBeCompared
}

func (indiv *individual) Best() Individual {
	return indiv.best
}

func (indiv *individual) SetBest(best Individual) {
	indiv.best = best
}

// individual https://en.wikipedia.org/wiki/individual
type individualSync struct {
	individual
	fitMut  sync.RWMutex
	bestMut sync.RWMutex
}

// NewIndividualSync is the constructor for threadsafe individuals
func NewIndividualSync(fitness float64) Individual {
	indiv := individual{
		fitness,
		nil,
	}
	return &individualSync{
		indiv,
		sync.RWMutex{},
		sync.RWMutex{},
	}
}

// Fitness returns the strength of a individual regarding to its environement. Higher is stronger.
func (indiv *individualSync) Fitness() float64 {
	indiv.fitMut.RLock()
	defer indiv.fitMut.RUnlock()
	return indiv.individual.Fitness()
}

// SetFitness set the strength of a individual regarding to its environement. Higher is stronger.
func (indiv *individualSync) SetFitness(Fitness float64) {
	indiv.fitMut.Lock()
	defer indiv.fitMut.Unlock()
	indiv.individual.SetFitness(Fitness)
}
