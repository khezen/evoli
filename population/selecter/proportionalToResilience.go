package selecter

import (
	"math/rand"

	"github.com/khezen/darwin/population"
	"github.com/khezen/darwin/population/individual"
)

type proportionalToResilienceSelecter struct{}

func (s proportionalToResilienceSelecter) Select(pop *population.Population, survivorsSize int) (*population.Population, error) {
	err := checkParams(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	if survivorsSize >= pop.Len() {
		return pop, nil
	}
	newPop, _ := population.New(pop.Cap())
	minResilience := pop.Min().Resilience()
	totalScore := s.computeTotalScore(pop, minResilience)
	for newPop.Len() < survivorsSize {
		for i := 0; i < pop.Len(); i++ {
			indiv, _ := pop.Get(i)
			score := s.computeScore(indiv, minResilience)
			if rand.Float32() <= score/totalScore {
				pop.Remove(i)
				newPop.Append(indiv)
				totalScore -= score
			}
		}
	}
	return newPop, nil
}

func (s proportionalToResilienceSelecter) computeScore(indiv individual.Interface, minResilience float32) float32 {
	return indiv.Resilience() - minResilience + 1
}

func (s proportionalToResilienceSelecter) computeTotalScore(pop *population.Population, minResilience float32) float32 {
	var length, totalScore = pop.Len(), float32(0)
	for i := 0; i < length; i++ {
		indiv, _ := pop.Get(i)
		totalScore += s.computeScore(indiv, minResilience)
	}
	return totalScore
}

// NewProportionalToResilienceSelecter is the constrctor for truncation selecter
func NewProportionalToResilienceSelecter() Interface {
	return proportionalToResilienceSelecter{}
}
