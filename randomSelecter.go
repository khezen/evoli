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
	newPop := NewPopulation(pop.Cap())
	newPop.Add(pop.Slice()...)
	size := newPop.Len() - survivorsSize
	for count := 0; count < size; {
		newPop.RemoveAt(rand.Intn(newPop.Len() - 1))
		count++
	}
	return newPop, nil
}

// NewRandomSelecter is the constructor for random selecter
func NewRandomSelecter() Selecter {
	return randomSelecter{}
}
