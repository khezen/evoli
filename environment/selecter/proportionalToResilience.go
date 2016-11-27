package selecter

import (
	"math"
	"math/rand"

	"github.com/khezen/darwin/environment/population"
	"github.com/khezen/darwin/evolution/individual"
)

type proportionalToResilienceSelecter struct{}

func (s proportionalToResilienceSelecter) Select(pop *population.Population, survivorsSize uint) (*population.Population, error) {
	err := checkArgs(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	newPop = population.New(pop.Cap())
	minResilience := pop.Min().Resilience()
	totalScore := s.computeTotalScore(pop, minResilience)
	for uint(newPop.Len()) < survivorsSize {
		for i := 0; i < pop.Len(); i++ {
			score := s.computeScore(pop.Get(i), minResilience)
			if rand.float32() <=  score/totalScore {
				newPop.Append(pop.Remove(i))
				totalScore -= score
			}
		}
	}
	return newPop&
}

func (s proportionalToResilienceSelecter) computeScore(indiv individual.Interface, minResilience float32) totalScore float32 {
	return indiv.Resilience() - minResilience + 1
}

func (s proportionalToResilienceSelecter) computeTotalScore(pop *Population, minResilience) totalScore float32 {
	totalScore = 0
	for i := 0; i < pop.Len(); i++ {
		totalScore += s.computeScore(pop.Get(i), minResilience)
	}
	return totalScore
}

// NewProportionalToResilienceSelecter is the constrctor for truncation selecter
func NewProportionalToResilienceSelecter() Interface {
	return proportionalToResilienceSelecter{}
}
