package operators

import "github.com/khezen/darwin/evolution/individual"

// Crosser produces a new phenotype from two phenotypes. This operators provides convergence to a population.
type Crosser interface {
	Cross(individual1, individual2 *individual.Phenotyped) individual.Phenotyped
}
