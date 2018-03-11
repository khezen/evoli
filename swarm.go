package evoli

var (
	// ErrLearningCoef - c1 & c2 must be > 0
	ErrLearningCoef = "ErrLearningCoef - c1 & c2 must be > 0"
)

type swarm struct {
	positioner Positioner
	c1, c2     float64
	evaluater  Evaluater
}

// NewSwarm - constructor for particles swarm optimization algorithm
// typical value for learning coef is c1 = c2 = 2
// the bigger are the coefficients the faster the population converge
func NewSwarm(positioner Positioner, c1, c2 float64, evaluater Evaluater) Evolution {
	if c1 <= 0 || c2 <= 0 {
		panic(ErrLearningCoef)
	}
	return &swarm{positioner, c1, c2, evaluater}
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
		if individual.Best() == nil || fitness > individual.Best().Fitness() {
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
		newIndiv, err := s.positioner.Position(indiv, pBest, gBest, s.c1, s.c2)
		if err != nil {
			return nil, err
		}
		newIndiv.SetBest(pBest)
		newPop.Add(newIndiv)
	}
	return newPop, nil
}

func (s *swarm) Evaluater() Evaluater {
	return s.evaluater
}
