package selecter

import (
	"math/rand"

	"github.com/khezen/darwin/population"
)

type proportionalToRankSelecter struct{}

func (s proportionalToRankSelecter) Select(pop *population.Population, survivorsSize int) *population.Population {
	newPop := population.New(pop.Cap())
	pop.Sort()
	totalScore := s.computeTotalScore(pop)
	for newPop.Len() < survivorsSize {
		for i := 0; i < pop.Len(); i++ {
			score := float32(pop.Len() - i)
			if rand.Float32() <= score/totalScore {
				newPop.Append(pop.Remove(i))
				totalScore -= score
			}
		}
	}
	return &newPop
}

func (s proportionalToRankSelecter) computeTotalScore(pop *population.Population) float32 {
	n := float32(pop.Len())
	return 1 / 2 * n * (n + 1) // 1+2+3+...+n
}

// NewProportionalToRankSelecter is the constrctor for truncation selecter
func NewProportionalToRankSelecter() Interface {
	return proportionalToRankSelecter{}
}
