package selecter

import (
	"errors"

	"github.com/khezen/darwin/environment/population"
)

type proportionalToResilienceSelecter struct{}

func (s proportionalToResilienceSelecter) Select(pop *population.Population, survivorsCount uint) (*population.Population, error) {
	err := checkArgs(pop, survivorsCount)
	if err != nil {
		return nil, err
	}
	return nil, errors.New("unsupported operation")
}

// NewProportionalToResilienceSelecter is the constrctor for truncation selecter
func NewProportionalToResilienceSelecter() Interface {
	return proportionalToResilienceSelecter{}
}
