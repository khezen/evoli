package darwin

import (
	"fmt"
	"math/rand"

	"github.com/khezen/check"
)

// lifecycle is a genetic algorithm implementation
type lifecycleAsync struct {
	lifecycle
}

// NewAsync is the constructor for Lifecycle
func NewAsync(s Selecter, c Crosser, m Mutater, e Evaluater) Lifecycle {
	return &lifecycleAsync{lifecycle{s, c, m, e}}
}

// Generation takes a population and produce a the new generation of this population
func (l lifecycleAsync) Generation(pop Population, survivorSizeForSelection int, mutationProbability float64) (Population, error) {
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

func (l lifecycleAsync) evaluation(pop Population) Population {
	length := pop.Len()
	dones := make([]chan bool, 0, length)
	for i := 0; i < length; i++ {
		dones = append(dones, make(chan bool))
	}
	for i := 0; i < length; i++ {
		go func(done chan bool, index int) {
			individual, err := pop.Get(index)
			fmt.Println(index, length, err)
			individual.SetFitness(l.Evaluater.Evaluate(individual))
			done <- true
		}(dones[i], i)
	}
	for _, done := range dones {
		<-done
	}
	return pop
}

func (l lifecycleAsync) crossovers(pop Population) Population {
	capacity := pop.Cap() - pop.Len()
	dones := make([]chan bool, 0, capacity)
	for i := 0; i < capacity; i++ {
		dones = append(dones, make(chan bool))
	}
	for i := 0; i < capacity; i++ {
		go func(done chan bool) {
			var _, indiv1, _, indiv2, _ = pop.PickCouple()
			newBorn := l.Crosser.Cross(indiv1, indiv2)
			pop.Add(newBorn)
			done <- true
		}(dones[i])
	}
	for _, done := range dones {
		<-done
	}
	return pop
}

func (l lifecycleAsync) mutations(pop Population, mutationProbability float64) (Population, error) {
	if mutationProbability < 0 || mutationProbability > 1 {
		return pop, fmt.Errorf("mutation probability = %v. Expected: 0 <= probability <= 1", mutationProbability)
	}
	length := pop.Len()
	dones := make([]chan bool, 0, length)
	for i := 0; i < length; i++ {
		dones = append(dones, make(chan bool))
	}
	for i := 0; i < length; i++ {
		go func(done chan bool, index int) {
			if rand.Float64() <= mutationProbability {
				indiv, _ := pop.Get(index)
				mutant := l.Mutater.Mutate(indiv)
				pop.Replace(index, mutant)
			}
			done <- true
		}(dones[i], i)
	}
	for _, done := range dones {
		<-done
	}
	return pop, nil
}
