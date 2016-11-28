package main

import (
	"math/rand"

	"github.com/khezen/darwin/environment/population"
	"github.com/khezen/darwin/environment/selecter"
	"github.com/khezen/darwin/evolution/operators"
)

// Lifecycle is a genetic algorithm implementation
type Lifecycle struct {
	Selecter  selecter.Interface
	Crosser   operators.Crosser
	Mutater   operators.Mutater
	Evaluater operators.Evaluater
}

// New is the constructor for Lifecycle struct
func New(s selecter.Interface, c operators.Crosser, m operators.Mutater, e operators.Evaluater) Lifecycle {
	return Lifecycle{s, c, m, e}
}

// Generation takes a Population and produce a the new generation of this population
func (l Lifecycle) Generation(pop *population.Population, survivorSizeForSelection int, mutationProb float32) *population.Population {
	l.evaluation(pop)
	newPop, err := l.Selecter.Select(pop, survivorSizeForSelection)
	newPop = l.crossovers(pop, mutationProb)
	return newPop
}

func (l Lifecycle) evaluation(pop *population.Population) *population.Population {
	length := pop.Len()
	for i := 0; i < length; i++ {
		individual := pop.Get(i)
		individual.SetResilience(l.Evaluater.Evaluate(individual))
	}
}

func (l Lifecycle) crossovers(pop *population.Population, mutationProb float32) *population.Population {
	length := pop.Len()
	newBorns := population.New(pop.Cap() - pop.Len())
	for newBorns.Len() <= newBorns.Cap() {
		var i, j = pop.PickCouple()
		newBorn := l.Crosser.Cross(pop.Get(i), pop.Get(j))
		if rand.Float32() <= mutationProb {
			newBorn = l.Mutater.Mutate(newBorn)
		}
		newBorns.Append(newBorn)
	}
	return pop
}
