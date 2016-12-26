package darwin

import (
	"math/rand"

	"github.com/khezen/darwin/population"
	"github.com/khezen/darwin/population/individual"
	"github.com/khezen/darwin/population/selecter"
)

// Interface for genetic algorithm step
type Interface interface {
	Generation(pop *population.Population, survivorSizeForSelection int, mutationProbability float32) (**population.Population, error)
}

// Lifecycle is a genetic algorithm implementation
type Lifecycle struct {
	Selecter  selecter.Interface
	Crosser   individual.Crosser
	Mutater   individual.Mutater
	Evaluater individual.Evaluater
}

// New is the constructor for Lifecycle struct
func New(s selecter.Interface, c individual.Crosser, m individual.Mutater, e individual.Evaluater) (Lifecycle, error) {
	return Lifecycle{s, c, m, e}, nil
}

// Generation takes a Population and produce a the new generation of this population
func (l Lifecycle) Generation(pop *population.Population, survivorSizeForSelection int, mutationProbability float32) (*population.Population, error) {
	newPop := l.evaluation(pop)
	newPop, _ = l.Selecter.Select(pop, survivorSizeForSelection)
	newPop, _ = l.crossovers(pop, mutationProbability)
	return newPop, nil
}

func (l Lifecycle) evaluation(pop *population.Population) *population.Population {
	length := pop.Len()
	for i := 0; i < length; i++ {
		individual, _ := pop.Get(i)
		individual.SetResilience(l.Evaluater.Evaluate(individual))
	}
	return pop
}

func (l Lifecycle) crossovers(pop *population.Population, mutationProbability float32) (*population.Population, error) {
	newBorns, _ := population.New(pop.Cap() - pop.Len())
	capacity := newBorns.Cap()
	for newBorns.Len() <= capacity {
		var _, indiv1, _, indiv2 = pop.PickCouple()
		newBorn := l.Crosser.Cross(indiv1, indiv2)
		if rand.Float32() <= mutationProbability {
			newBorn = l.Mutater.Mutate(newBorn)
		}
		newBorns.Append(newBorn)
	}
	pop.AppendAll(newBorns)
	return pop, nil
}
