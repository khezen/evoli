package operators

import "github.com/khezen/darwin/evolution/individual"

// Mutater randomly modify a phenotype. This operator maintain diversity in a population.
type Mutater interface {
	Mutate(*individual.Phenotyped)
}
