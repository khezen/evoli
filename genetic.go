package evoli

import (
	"math/rand"
)

// Genetic for genetic algorithm step
type Genetic interface {
	Next(pop Population) (Population, error)
}

// genetic is a genetic algorithm implementation
type genetic struct {
	Selecter            Selecter
	SurvivorSize        int
	Crosser             Crosser
	Mutater             Mutater
	MutationProbability float64
	Evaluater           Evaluater
}

// NewGenetic is the constructor for Genetic
func NewGenetic(s Selecter, survivorSize int, c Crosser, m Mutater, mutationProbability float64, e Evaluater) Evolution {
	return &genetic{s, survivorSize, c, m, mutationProbability, e}
}

// Next takes a population and produce a the new generation of this population
func (l genetic) Next(pop Population) (Population, error) {
	newPop, err := l.evaluation(pop)
	if err != nil {
		return pop, err
	}
	newPop, err = l.Selecter.Select(newPop, l.SurvivorSize)
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

func (l genetic) evaluation(pop Population) (Population, error) {
	length := pop.Len()
	for i := 0; i < length; i++ {
		individual, _ := pop.Get(i)
		fitness, err := l.Evaluater.Evaluate(individual)
		if err != nil {
			return pop, err
		}
		individual.SetFitness(fitness)
	}
	return pop, nil
}

func (l genetic) crossovers(pop Population) (Population, error) {
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
		indiv1, _ := pop.Get(i)
		indiv2, _ := pop.Get(j)
		newBorn, err := l.Crosser.Cross(indiv1, indiv2)
		if err != nil {
			return nil, err
		}
		newBorns.Add(newBorn)
	}
	pop.Add(*newBorns.(*population)...)
	return pop, nil
}

func (l genetic) mutations(pop Population) (Population, error) {
	for i := 0; i < pop.Len(); i++ {
		if rand.Float64() <= l.MutationProbability {
			indiv, _ := pop.Get(i)
			mutant, err := l.Mutater.Mutate(indiv)
			if err != nil {
				return nil, err
			}
			pop.Replace(i, mutant)
		}
	}
	return pop, nil
}
