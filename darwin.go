package darwin

import (
	"fmt"
	"math/rand"

	"github.com/khezen/cop"
	"github.com/khezen/darwin/population"
	"github.com/khezen/darwin/population/individual"
	"github.com/khezen/darwin/population/selecter"
)

// Interface for genetic algorithm step
type Interface interface {
	Generation(pop population.Interface, survivorSizeForSelection int, mutationProbability float32) (*population.Interface, error)
}

// Lifecycle is a genetic algorithm implementation
type Lifecycle struct {
	Selecter  selecter.Interface
	Crosser   individual.Crosser
	Mutater   individual.Mutater
	Evaluater individual.Evaluater
}

// New is the constructor for Lifecycle struct
func New(s selecter.Interface, c individual.Crosser, m individual.Mutater, e individual.Evaluater) (*Lifecycle, error) {
	if s == nil {
		return nil, fmt.Errorf("selecter is nil")
	}
	if c == nil {
		return nil, fmt.Errorf("crosser is nil")
	}
	if m == nil {
		return nil, fmt.Errorf("mutater is nil")
	}
	if e == nil {
		return nil, fmt.Errorf("evaluater is nil")
	}
	return &Lifecycle{s, c, m, e}, nil
}

// Generation takes a Population and produce a the new generation of this population
func (l Lifecycle) Generation(pop population.Interface, survivorSizeForSelection int, mutationProbability float32) (population.Interface, error) {
	err := cop.CheckNotNil(pop)
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

func (l Lifecycle) evaluation(pop population.Interface) population.Interface {
	length := pop.Len()
	for i := 0; i < length; i++ {
		individual, _ := pop.Get(i)
		individual.SetFitness(l.Evaluater.Evaluate(individual))
	}
	return pop
}

func (l Lifecycle) crossovers(pop population.Interface) population.Interface {
	newBorns, _ := population.New(pop.Cap() - pop.Len())
	capacity := newBorns.Cap()
	for newBorns.Len() < capacity {
		var _, indiv1, _, indiv2, _ = pop.PickCouple()
		newBorn := l.Crosser.Cross(indiv1, indiv2)
		newBorns.Append(newBorn)
	}
	pop.AppendAll(newBorns)
	return pop
}

func (l Lifecycle) mutations(pop population.Interface, mutationProbability float32) (population.Interface, error) {
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
