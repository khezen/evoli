package evoli

import (
	"errors"
	"math/rand"
	"sync"
)

// genetic is a genetic algorithm implementation
type genetic struct {
	evolution
	selecter            Selecter
	SurvivorSize        int
	crosser             Crosser
	mutater             Mutater
	MutationProbability float64
}

var (
	// ErrSurvivorSize - survivorSize < 1
	ErrSurvivorSize = errors.New("ErrSurvivorSize - survivorSize must be >= 1")
	// ErrMutationProb - 0<= mutationProbability <= 1
	ErrMutationProb = errors.New("ErrMutationProb - mutation probability must be 0 <= mutationProbability <= 1")
)

// NewGenetic - constructor for Genetic Algorithm
func NewGenetic(pop Population, s Selecter, survivorSize int, c Crosser, m Mutater, mutationProbability float64, e Evaluater) Evolution {
	if survivorSize < 1 {
		panic(ErrSurvivorSize)
	}
	if mutationProbability < 0 || mutationProbability > 1 {
		panic(ErrMutationProb)
	}
	return &genetic{newEvolution(pop, e), s, survivorSize, c, m, mutationProbability}
}

// Next takes a population and produce a the new generation of this population
func (g *genetic) Next() error {
	newPop, err := g.evaluation(g.pop)
	if err != nil {
		return err
	}
	newPop, err = g.selecter.Select(newPop, g.SurvivorSize)
	if err != nil {
		return err
	}
	newPop, err = g.crossovers(newPop)
	if err != nil {
		return err
	}
	newPop, err = g.mutations(newPop)
	if err != nil {
		return err
	}
	g.pop = newPop
	return nil
}

func (g *genetic) evaluation(pop Population) (Population, error) {
	length := pop.Len()
	for i := 0; i < length; i++ {
		individual := pop.Get(i)
		fitness, err := g.evaluater.Evaluate(individual)
		if err != nil {
			return pop, err
		}
		individual.SetFitness(fitness)
	}
	return pop, nil
}

func (g *genetic) crossovers(pop Population) (Population, error) {
	newBorns := NewPopulation(pop.Cap() - pop.Len())
	capacity := newBorns.Cap()
	for newBorns.Len() < capacity {
		var i, j = rand.Intn(pop.Len()), rand.Intn(pop.Len())
		if i == j {
			switch i {
			case pop.Len() - 1:
				j = i - 1
			default:
				j = i + 1
			}
		}
		indiv1, indiv2 := pop.Get(i), pop.Get(j)
		newBorn, err := g.crosser.Cross(indiv1, indiv2)
		if err != nil {
			return nil, err
		}
		newBorns.Add(newBorn)
	}
	pop.Add(*newBorns.(*population)...)
	return pop, nil
}

func (g *genetic) mutations(pop Population) (Population, error) {
	for i := 0; i < pop.Len(); i++ {
		if rand.Float64() <= g.MutationProbability {
			indiv := pop.Get(i)
			mutant, err := g.mutater.Mutate(indiv)
			if err != nil {
				return nil, err
			}
			pop.Replace(i, mutant)
		}
	}
	return pop, nil
}

type geneticSync struct {
	genetic
	sync.RWMutex
}

// NewGeneticSync - constructor for Genetic Algorithm (sync impl)
func NewGeneticSync(pop Population, s Selecter, survivorSize int, c Crosser, m Mutater, mutationProbability float64, e Evaluater) Evolution {
	return &geneticSync{*NewGenetic(pop, s, survivorSize, c, m, mutationProbability, e).(*genetic), sync.RWMutex{}}
}

func (s *geneticSync) Next() error {
	s.Lock()
	defer s.Unlock()
	return s.genetic.Next()
}

func (s *geneticSync) Population() Population {
	s.RLock()
	defer s.RUnlock()
	return s.genetic.Population()
}

func (s *geneticSync) SetPopulation(pop Population) {
	s.Lock()
	defer s.Unlock()
	s.genetic.SetPopulation(pop)
}

func (s *geneticSync) Alpha() Individual {
	s.RLock()
	defer s.RUnlock()
	return s.genetic.Alpha()
}
