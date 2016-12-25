package selecter

import (
	"math/rand"

	"github.com/khezen/darwin/population"
)

type proportionalToRankSelecter struct{}

func (s proportionalToRankSelecter) Select(pop *population.Population, survivorsSize int) (*population.Population, error) {
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

func (s proportionalToRankSelecter) computeTotalScore(pop *population.Population) float32 {
	n := float32(pop.Len())
	return 1 / 2 * n * (n + 1) // 1+2+3+...+n
}

// NewProportionalToRankSelecter is the constrctor for truncation selecter
func NewProportionalToRankSelecter() Interface {
	return proportionalToRankSelecter{}
}
