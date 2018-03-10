package evoli

var (
	// ErrLearningCoef - learningCoef1 & learningCoef2 must be > 0
	ErrLearningCoef = "ErrLearningCoef - learningCoef1 & learningCoef2 must be > 0"
)

type swarm struct {
	positioner                   Positioner
	learningCoef1, learningCoef2 float64
	evaluater                    Evaluater
}

// NewSwarm - constructor for particles swarm optimization algorithm
func NewSwarm(positioner Positioner, learningCoef1, learningCoef2 float64, evaluater Evaluater) Evolution {
	if learningCoef1 <= 0 || learningCoef2 <= 0 {
		panic(ErrLearningCoef)
	}
	return &swarm{positioner, learningCoef1, learningCoef2, evaluater}
}

func (s *swarm) Next(pop Population) (Population, error) {
	newPop, err := s.evaluation(pop)
	if err != nil {
		return pop, err
	}
	newPop, err = s.positioning(pop)
	if err != nil {
		return pop, err
	}
	return newPop, nil
}

func (s *swarm) evaluation(pop Population) (Population, error) {
	length := pop.Len()
	for i := 0; i < length; i++ {
		individual := pop.Get(i)
		fitness, err := s.evaluater.Evaluate(individual)
		if err != nil {
			return pop, err
		}
		if fitness >= individual.Fitness() {
			individual.SetBest(individual)
		}
		individual.SetFitness(fitness)
	}
	return pop, nil
}

func (s *swarm) positioning(pop Population) (Population, error) {
	newPop := NewPopulation(pop.Cap())
	gBest := pop.Max()
	individuals := pop.Slice()
	for _, indiv := range individuals {
		pBest := indiv.Best()
		newIndiv, err := s.positioner.Position(indiv, pBest, gBest, s.learningCoef1, s.learningCoef2)
		if err != nil {
			return nil, err
		}
		newPop.Add(newIndiv)
	}
	return newPop, nil
}

func (s *swarm) Evaluater() Evaluater {
	return s.evaluater
}
