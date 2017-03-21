package darwin

import (
	"math/rand"
)

type proportionalToRankSelecter struct{}

func (s proportionalToRankSelecter) Select(pop IPopulation, survivorsSize int) (IPopulation, error) {
	err := checkParams(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	if survivorsSize >= pop.Len() {
		return pop, nil
	}
	newPop := NewPopulation(pop.Cap())
	totalScore := s.computeTotalScore(pop)
	pop.Sort()
	for newPop.Len() < survivorsSize {
		for i := 0; i < pop.Len(); i++ {
			if newPop.Len() >= survivorsSize {
				break
			}
			score := float32(pop.Len() - i)
			if rand.Float32() <= score/totalScore {
				indiv, _ := pop.Get(i)
				pop.RemoveAt(i)
				newPop.Append(indiv)
				totalScore -= score
			}
		}
	}
	return newPop, nil
}

func (s proportionalToRankSelecter) computeTotalScore(pop IPopulation) float32 {
	n := float32(pop.Len())
	return n * (n + 1) / 2 // 1+2+3+...+n
}

// NewProportionalToRankSelecter is the constructor for selecter based on ranking across the population
func NewProportionalToRankSelecter() ISelecter {
	return proportionalToRankSelecter{}
}
