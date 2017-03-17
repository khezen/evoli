package darwin

// ICrosser produces a new individual from two individuals. This operators provides convergence to a population.
type ICrosser interface {
	Cross(individual1, individual2 IIndividual) IIndividual
}
