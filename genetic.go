package evoli

import (
	"errors"
	"math/rand"
)

// genetic is a genetic algorithm implementation
type genetic struct {
	selecter            Selecter
	SurvivorSize        int
	crosser             Crosser
	mutater             Mutater
	MutationProbability float64
	evaluater           Evaluater
}

var (
	// ErrSurvivorSize - survivorSize < 1
	ErrSurvivorSize = errors.New("ErrSurvivorSize - survivorSize must be >= 1")
	// ErrMutationProb - 0<= mutationProbability <= 1
	ErrMutationProb = errors.New("ErrMutationProb - mutation probability must be 0 <= mutationProbability <= 1")
)

// NewGenetic - constructor for Genetic Algorithm
func NewGenetic(s Selecter, survivorSize int, c Crosser, m Mutater, mutationProbability float64, e Evaluater) Evolution {
	if survivorSize < 1 {
		panic(ErrSurvivorSize)
	}
	if mutationProbability < 0 || mutationProbability > 1 {
		panic(ErrMutationProb)
	}
	return &genetic{s, survivorSize, c, m, mutationProbability, e}
}

// Next takes a population and produce a the new generation of this population
func (l *genetic) Next(pop Population) (Population, error) {
	newPop, err := l.evaluation(pop)
	if err != nil {
		return pop, err
	}
	newPop, err = l.selecter.Select(newPop, l.SurvivorSize)
	if err != nil {
		return pop, err
	}
	newPop, err = l.crossovers(newPop)
	if err != nil {
		return pop, err
	}
	newPop, err = l.mutations(newPop)
	if err != nil {
		return pop, err
	}
	return newPop, nil
}

func (l *genetic) evaluation(pop Population) (Population, error) {
	length := pop.Len()
	for i := 0; i < length; i++ {
		individual := pop.Get(i)
		fitness, err := l.evaluater.Evaluate(individual)
		if err != nil {
			return pop, err
		}
		individual.SetFitness(fitness)
	}
	return pop, nil
}

func (l *genetic) crossovers(pop Population) (Population, error) {
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
		newBorn, err := l.crosser.Cross(indiv1, indiv2)
		if err != nil {
			return nil, err
		}
		newBorns.Add(newBorn)
	}
	pop.Add(*newBorns.(*population)...)
	return pop, nil
}

func (l *genetic) mutations(pop Population) (Population, error) {
	for i := 0; i < pop.Len(); i++ {
		if rand.Float64() <= l.MutationProbability {
			indiv := pop.Get(i)
			mutant, err := l.mutater.Mutate(indiv)
			if err != nil {
				return nil, err
			}
			pop.Replace(i, mutant)
		}
	}
	return pop, nil
}

func (l *genetic) Evaluater() Evaluater {
	return l.evaluater
}
