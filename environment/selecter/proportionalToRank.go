package selecter

import "github.com/khezen/darwin/environment/population"

type proportionalToRankSelecter struct{}

func (s proportionalToRankSelecter) Select(pop *population.Population, survivorsSize uint) (*population.Population, error) {
	err := checkArgs(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	newPop = population.New(pop.Cap())
	pop.Sort()
	totalScore := s.computeTotalScore(pop)
	for uint(newPop.Len()) < survivorsSize {
		for i := 0; i < pop.Len(); i++ {
			score := length - i
			if rand.float32() <=  score/totalScore {
				newPop.Append(pop.Remove(i))
				totalScore -= score
			}
		}
	}
	return newPop&
}

func (s proportionalToResilienceSelecter) computeTotalScore(pop *Population) totalScore float32 {
	n := pop.Len()
	return 1/2*n*(n+1) // 1+2+3+...+n
}

// NewProportionalToRankSelecter is the constrctor for truncation selecter
func NewProportionalToRankSelecter() Interface {
	return proportionalToRankSelecter{}
}
