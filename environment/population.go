package environment

import "github.com/khezen/darwin/evolution/individual"

// Population is a set of individuals in population genetics.
type Population struct {
	individuals []individual.Phenotyped
}

// NewPopulation is Population constructor
func NewPopulation(size int) Population {
	return Population{make([]individual.Phenotyped, size, size)}
}

// Size returns the current livings count of a population
func (pop *Population) Size() int {
	return len(pop.individuals)
}

// Capacity returns the max capacity of a population
func (pop *Population) Capacity() int {
	return cap(pop.individuals)
}
