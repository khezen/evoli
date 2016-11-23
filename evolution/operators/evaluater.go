package operators

import "github.com/khezen/darwin/evolution/individual"

// Evaluater computes phenotype resilience
type Evaluater interface {
	Evaluate(*individual.Phenotyped)
}
