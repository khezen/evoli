package darwin

// Crosser produces a new individual from two individuals. This operators provides convergence to a population.
type Crosser interface {
	Cross(individual1, individual2 Individual) (Individual, error)
}
