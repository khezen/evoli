package darwin

// IEvaluater computes and set individual Fitness
type IEvaluater interface {
	Evaluate(IIndividual) (Fitness float32)
}
