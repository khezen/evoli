package evoli

import "sync"

var (
	// ErrLearningCoef - c1 & c2 must be > 0
	ErrLearningCoef = "ErrLearningCoef - c1 & c2 must be > 0"
)

type swarm struct {
	evolution
	bests      map[Individual]Individual
	positioner Positioner
	c1, c2     float64
}

// NewSwarm - constructor for particles swarm optimization algorithm
// typical value for learning coef is c1 = c2 = 2
// the bigger are the coefficients the faster the population converge
func NewSwarm(pop Population, positioner Positioner, c1, c2 float64, evaluater Evaluater) Evolution {
	if c1 <= 0 || c2 <= 0 {
		panic(ErrLearningCoef)
	}
	return &swarm{newEvolution(pop, evaluater), make(map[Individual]Individual), positioner, c1, c2}
}

func (s *swarm) Next() error {
	newPop, err := s.evaluation(s.pop)
	if err != nil {
		return err
	}
	newPop, err = s.positioning(newPop)
	if err != nil {
		return err
	}
	s.pop = newPop
	return nil
}

func (s *swarm) evaluation(pop Population) (Population, error) {
	length := pop.Len()
	for i := 0; i < length; i++ {
		individual := pop.Get(i)
		fitness, err := s.evaluater.Evaluate(individual)
		if err != nil {
			return pop, err
		}
		best, exists := s.bests[individual]
		if !exists || fitness > best.Fitness() {
			s.bests[individual] = individual
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
		pBest := s.bests[indiv]
		newIndiv, err := s.positioner.Position(indiv, pBest, gBest, s.c1, s.c2)
		if err != nil {
			return nil, err
		}
		delete(s.bests, indiv)
		s.bests[newIndiv] = pBest
		newPop.Add(newIndiv)
	}
	return newPop, nil
}

func (s *swarm) SetPopulation(pop Population) {
	s.bests = make(map[Individual]Individual)
	s.evolution.SetPopulation(pop)
}

type swarmTS struct {
	swarm
	sync.RWMutex
}

// NewSwarmTS - constructor for particles swarm optimization algorithm (sync impl)
// typical value for learning coef is c1 = c2 = 2
// the bigger are the coefficients the faster the population converge
func NewSwarmTS(pop Population, positioner Positioner, c1, c2 float64, evaluater Evaluater) Evolution {
	return &swarmTS{*NewSwarm(pop, positioner, c1, c2, evaluater).(*swarm), sync.RWMutex{}}
}

func (s *swarmTS) Next() error {
	s.Lock()
	defer s.Unlock()
	return s.swarm.Next()
}

func (s *swarmTS) Population() Population {
	s.RLock()
	defer s.RUnlock()
	return s.swarm.Population()
}

func (s *swarmTS) SetPopulation(pop Population) {
	s.Lock()
	defer s.Unlock()
	s.swarm.SetPopulation(pop)
}

func (s *swarmTS) Alpha() Individual {
	s.RLock()
	defer s.RUnlock()
	return s.swarm.Alpha()
}
