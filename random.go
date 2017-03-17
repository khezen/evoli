package darwin

import (
	"math/rand"
)

type randomSelecter struct{}

func (s randomSelecter) Select(pop IPopulation, survivorsSize int) (IPopulation, error) {
	err := checkParams(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	size := pop.Len() - survivorsSize
	for count := 0; count < size; {
		pop.Remove(rand.Intn(pop.Len() - 1))
		count++
	}
	return pop, nil
}

// NewRandomSelecter is the constructor for random selecter
func NewRandomSelecter() ISelecter {
	return randomSelecter{}
}
