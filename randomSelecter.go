package darwin

import (
	"math/rand"
)

type randomSelecter struct{}

func (s randomSelecter) Select(pop Population, survivorsSize int) (Population, error) {
	err := checkParams(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	size := pop.Len() - survivorsSize
	for count := 0; count < size; {
		pop.RemoveAt(rand.Intn(pop.Len() - 1))
		count++
	}
	return pop, nil
}

// NewRandomSelecter is the constructor for random selecter
func NewRandomSelecter() Selecter {
	return randomSelecter{}
}
