package individual

// Evaluater computes and set individual Fitness
type Evaluater interface {
	Evaluate(Interface) (Fitness float32)
}
