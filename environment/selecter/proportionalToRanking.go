package selecter

import "github.com/khezen/darwin/environment/population"
import "errors"

type proportionalToRankingSelecter struct{}

func (s proportionalToRankingSelecter) Select(pop *population.Population, survivorsCount uint) (*population.Population, error) {
	err := checkArgs(pop, survivorsCount)
	if err != nil {
		return nil, err
	}
	return nil, errors.New("unsupported operation")
}

// NewProportionalToRankingSelecter is the constrctor for truncation selecter
func NewProportionalToRankingSelecter() Interface {
	return proportionalToRankingSelecter{}
}
