package selecter

import (
	"math/rand"

	"github.com/khezen/darwin/environment/population"
)

type randomSelecter struct{}

func (s randomSelecter) Select(pop *population.Population, survivorsCount uint) (*population.Population, error) {
	err := checkArgs(pop, survivorsCount)
	if err != nil {
		return nil, err
	}
	count := uint(0)
	survivors := population.New(pop.Cap())
	for count < survivorsCount {
		survivor := pop.Remove(rand.Intn(pop.Len() - 1))
		survivors.Append(survivor)
		count++
	}
	return &survivors, nil
}

// NewRandomSelecter is the constrctor for truncation selecter
func NewRandomSelecter() Interface {
	return randomSelecter{}
}
