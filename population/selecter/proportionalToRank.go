package selecter

import (
	"math/rand"

	"github.com/khezen/darwin/population"
)

type proportionalToRankSelecter struct{}

func (s proportionalToRankSelecter) Select(pop population.Interface, survivorsSize int) (population.Interface, error) {
	err := checkParams(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	if survivorsSize >= pop.Len() {
		return pop, nil
	}
	newPop, _ := population.New(pop.Cap())
	totalScore := s.computeTotalScore(pop)
	pop.Sort()
	for newPop.Len() < survivorsSize {
		for i := 0; i < pop.Len(); i++ {
			score := float32(pop.Len() - i)
			if rand.Float32() <= score/totalScore {
				indiv, _ := pop.Remove(i)
				newPop.Append(indiv)
				totalScore -= score
			}
		}
	}
	return newPop, nil
}

func (s proportionalToRankSelecter) computeTotalScore(pop population.Interface) float32 {
	n := float32(pop.Len())
	return n * (n + 1) / 2 // 1+2+3+...+n
}

// NewProportionalToRankSelecter is the constrctor for truncation selecter
func NewProportionalToRankSelecter() Interface {
	return proportionalToRankSelecter{}
}
