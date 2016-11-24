package selecter

import (
	"math/rand"

	"github.com/khezen/darwin/environment/population"
)

type randomSelecter struct{}

func (s randomSelecter) Select(pop *population.Population, survivorsSize uint) (*population.Population, error) {
	err := checkArgs(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	count := uint(0)
	size := pop.Len() - survivorsSize
	for count < size {
		pop.Remove(rand.Intn(pop.Len() - 1))
		count++
	}
	return pop, nil
}

// NewRandomSelecter is the constrctor for truncation selecter
func NewRandomSelecter() Interface {
	return randomSelecter{}
}
