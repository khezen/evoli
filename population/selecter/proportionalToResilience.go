package selecter

import (
	"math/rand"

	"github.com/khezen/darwin/population"
	"github.com/khezen/darwin/population/individual"
)

type proportionalToResilienceSelecter struct{}

func (s proportionalToResilienceSelecter) Select(pop population.Interface, survivorsSize int) (population.Interface, error) {
	err := checkParams(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	if survivorsSize >= pop.Len() {
		return pop, nil
	}
	newPop, _ := population.New(pop.Cap())
	offset := s.computeOffset(pop)
	totalScore := s.computeTotalScore(pop, offset)
	for newPop.Len() < survivorsSize {
		for i := 0; i < pop.Len(); i++ {
			indiv, _ := pop.Get(i)
			score := s.computeScore(indiv, offset)
			if totalScore == 0 || rand.Float32() <= score/totalScore {
				pop.Remove(i)
				newPop.Append(indiv)
				totalScore -= score
			}
		}
	}
	return newPop, nil
}

func (s proportionalToResilienceSelecter) computeScore(indiv individual.Interface, offset float32) float32 {
	return indiv.Resilience() + offset
}

func (s proportionalToResilienceSelecter) computeTotalScore(pop population.Interface, offset float32) float32 {
	var length, totalScore = pop.Len(), float32(0)
	for i := 0; i < length; i++ {
		indiv, _ := pop.Get(i)
		totalScore += s.computeScore(indiv, offset)
	}
	return totalScore
}

func (s proportionalToResilienceSelecter) computeOffset(pop population.Interface) float32 {
	min := pop.Min().Resilience()
	max := pop.Max().Resilience()
	var offset float32
	switch {
	case min < 0:
		offset += -min
	case min > 0:
		offset += min
	}
	switch {
	case max < 0:
		offset += -max
	case max > 0:
		offset += max
	}
	return offset
}

// NewProportionalToResilienceSelecter is the constrctor for truncation selecter
func NewProportionalToResilienceSelecter() Interface {
	return proportionalToResilienceSelecter{}
}
