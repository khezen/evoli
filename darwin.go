package darwin

import (
	"fmt"
	"math/rand"

	"github.com/khezen/check"
)

// Lifecycle for genetic algorithm step
type Lifecycle interface {
	Generation(pop Population, survivorSizeForSelection int, mutationProbability float32) (Population, error)
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

// Generation takes a population and produce a the new generation of this population
func (l lifecycle) Generation(pop Population, survivorSizeForSelection int, mutationProbability float32) (Population, error) {
	err := check.NotNil(pop)
	if err != nil {
		return pop, err
	}
	newPop := l.evaluation(pop)
	newPop, err = l.Selecter.Select(newPop, survivorSizeForSelection)
	if err != nil {
		return pop, err
	}
	newPop = l.crossovers(newPop)
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

func (l lifecycle) crossovers(pop Population) Population {
	newBorns := NewPopulation(pop.Cap() - pop.Len())
	capacity := newBorns.Cap()
	for newBorns.Len() < capacity {
		var _, indiv1, _, indiv2, _ = pop.PickCouple()
		newBorn := l.Crosser.Cross(indiv1, indiv2)
		newBorns.Add(newBorn)
	}
	pop.Add(*newBorns.(*population)...)
	return pop
}

func (l lifecycle) mutations(pop Population, mutationProbability float32) (Population, error) {
	if mutationProbability < 0 || mutationProbability > 1 {
		return pop, fmt.Errorf("mutation probability = %v. Expected: 0 <= probability <= 1", mutationProbability)
	}
	for i := 0; i < pop.Len(); i++ {
		if rand.Float32() <= mutationProbability {
			indiv, _ := pop.Get(i)
			mutant := l.Mutater.Mutate(indiv)
			pop.Replace(i, mutant)
		}
	}
	return pop, nil
}
