package operators

import "github.com/khezen/darwin/evolution/individual"

// Mutater randomly modify a individual. This operator maintain diversity in a population.
type Mutater interface {
	Mutate(individual.Interface) individual.Interface
}
