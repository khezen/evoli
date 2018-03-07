package evoli

// Evaluater computes and set individual Fitness
type Evaluater interface {
	Evaluate(Individual) (Fitness float64)
}
