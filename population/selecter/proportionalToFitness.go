package selecter

import (
	"math/rand"

	"github.com/khezen/darwin/population"
	"github.com/khezen/darwin/population/individual"
)

type proportionalToFitnessSelecter struct{}

func (s proportionalToFitnessSelecter) Select(pop population.Interface, survivorsSize int) (population.Interface, error) {
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

func (s proportionalToFitnessSelecter) computeScore(indiv individual.Interface, offset float32) float32 {
	return indiv.Fitness() + offset
}

func (s proportionalToFitnessSelecter) computeTotalScore(pop population.Interface, offset float32) float32 {
	var length, totalScore = pop.Len(), float32(0)
	for i := 0; i < length; i++ {
		indiv, _ := pop.Get(i)
		totalScore += s.computeScore(indiv, offset)
	}
	return totalScore
}

func (s proportionalToFitnessSelecter) computeOffset(pop population.Interface) float32 {
	minIndiv, maxIndiv := pop.Extremums()
	min, max := minIndiv.Fitness(), maxIndiv.Fitness()
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

// NewProportionalToFitnessSelecter is the constrctor for truncation selecter
func NewProportionalToFitnessSelecter() Interface {
	return proportionalToFitnessSelecter{}
}