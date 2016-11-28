package main

import (
	"math/rand"

	"github.com/khezen/darwin/population"
	"github.com/khezen/darwin/population/individual"
	"github.com/khezen/darwin/population/selecter"
)

// Lifecycle is a genetic algorithm implementation
type Lifecycle struct {
	Selecter  selecter.Interface
	Crosser   individual.Crosser
	Mutater   individual.Mutater
	Evaluater individual.Evaluater
}

// New is the constructor for Lifecycle struct
func New(s selecter.Interface, c individual.Crosser, m individual.Mutater, e individual.Evaluater) Lifecycle {
	return Lifecycle{s, c, m, e}
}

// Generation takes a Population and produce a the new generation of this population
func (l Lifecycle) Generation(pop *population.Population, survivorSizeForSelection int, mutationProbability float32) *population.Population {
	l.evaluation(pop)
	newPop, err := l.Selecter.Select(pop, survivorSizeForSelection)
	newPop = l.crossovers(pop, mutationProbability)
	return newPop
}

func (l Lifecycle) evaluation(pop *population.Population) *population.Population {
	length := pop.Len()
	for i := 0; i < length; i++ {
		individual := pop.Get(i)
		individual.SetResilience(l.Evaluater.Evaluate(individual))
	}
}

func (l Lifecycle) crossovers(pop *population.Population, mutationProbability float32) *population.Population {
	length := pop.Len()
	newBorns := population.New(pop.Cap() - pop.Len())
	for newBorns.Len() <= newBorns.Cap() {
		var i, j = pop.PickCouple()
		newBorn := l.Crosser.Cross(pop.Get(i), pop.Get(j))
		if rand.Float32() <= mutationProbability {
			newBorn = l.Mutater.Mutate(newBorn)
		}
		newBorns.Append(newBorn)
	}
	pop.AppendAll(newBorns)
	return pop
}
