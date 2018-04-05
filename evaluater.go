package evoli

// Evaluater calculate individual Fitness
type Evaluater interface {
	Evaluate(Individual) (Fitness float64, err error)
}
