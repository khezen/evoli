package evoli

// Crosser produces a new individual from two individuals. This operators provides convergence to a population.
type Crosser interface {
	Cross(parent1, parent2 Individual) (child1, child2 Individual, err error)
}
