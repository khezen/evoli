package operators

import "github.com/khezen/darwin/evolution/individual"

// Evaluater computes and set individual resilience
type Evaluater interface {
	Evaluate(individual.Interface) (resilience float32)
}
