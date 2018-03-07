package darwin

import (
	"fmt"
	"math/rand"
)

// Lifecycle for genetic algorithm step
type Lifecycle interface {
	Iterate(pop Population, survivorSizeForSelection int, mutationProbability float64) (Population, error)
}

// lifecycle is a genetic algorithm implementation
type lifecycle struct {
	Selecter  Selecter
	Crosser   Crosser
	Mutater   Mutater
	Evaluater Evaluater
}

// New is the constructor for Lifecycle
func New(s Selecter, c Crosser, m Mutater, e Evaluater) Lifecycle {
	return &lifecycle{s, c, m, e}
}

// Iterate takes a population and produce a the new generation of this population
func (l lifecycle) Iterate(pop Population, survivorSizeForSelection int, mutationProbability float64) (Population, error) {
	newPop := l.evaluation(pop)
	newPop, err := l.Selecter.Select(newPop, survivorSizeForSelection)
	if err != nil {
		return pop, err
	}
	newPop, err = l.crossovers(newPop)
	if err != nil {
		return pop, err
	}
	newPop, err = l.mutations(newPop, mutationProbability)
	if err != nil {
		return pop, err
	}
	return newPop, nil
}

func (l lifecycle) evaluation(pop Population) Population {
	length := pop.Len()
	for i := 0; i < length; i++ {
		individual, _ := pop.Get(i)
		individual.SetFitness(l.Evaluater.Evaluate(individual))
	}
	return pop
}

func (l lifecycle) crossovers(pop Population) (Population, error) {
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

func (l lifecycle) mutations(pop Population, mutationProbability float64) (Population, error) {
	if mutationProbability < 0 || mutationProbability > 1 {
		return pop, fmt.Errorf("mutation probability = %v. Expected: 0 <= probability <= 1", mutationProbability)
	}
	for i := 0; i < pop.Len(); i++ {
		if rand.Float64() <= mutationProbability {
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
