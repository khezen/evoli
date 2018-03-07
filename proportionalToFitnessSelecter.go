package darwin

import (
	"math/rand"
)

type proportionalToFitnessSelecter struct{}

func (s proportionalToFitnessSelecter) Select(pop Population, survivorsSize int) (Population, error) {
	err := checkParams(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	if survivorsSize >= pop.Len() {
		return pop, nil
	}
	newPop := pop.New(pop.Cap())
	offset := s.computeOffset(pop)
	totalScore := s.computeTotalScore(pop, offset)
	for newPop.Len() < survivorsSize {
		for i := 0; i < pop.Len(); i++ {
			if newPop.Len() >= survivorsSize {
				break
			}
			indiv, _ := pop.Get(i)
			score := s.computeScore(indiv, offset)
			if totalScore == 0 || rand.Float64() <= score/totalScore {
				pop.RemoveAt(i)
				newPop.Add(indiv)
				totalScore -= score
			}
		}
	}
	return newPop, nil
}

func (s proportionalToFitnessSelecter) computeScore(indiv Individual, offset float64) float64 {
	return indiv.Fitness() + offset
}

func (s proportionalToFitnessSelecter) computeTotalScore(pop Population, offset float64) float64 {
	var length, totalScore = pop.Len(), float64(0)
	for i := 0; i < length; i++ {
		indiv, _ := pop.Get(i)
		totalScore += s.computeScore(indiv, offset)
	}
	return totalScore
}

func (s proportionalToFitnessSelecter) computeOffset(pop Population) float64 {
	minIndiv, maxIndiv := pop.Min(), pop.Max()
	min, max := minIndiv.Fitness(), maxIndiv.Fitness()
	var offset float64
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

// NewProportionalToFitnessSelecter is the constructor for selecter based on fitness value
func NewProportionalToFitnessSelecter() Selecter {
	return proportionalToFitnessSelecter{}
}
