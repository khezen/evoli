package selecter

import (
	"math/rand"

	"github.com/khezen/darwin/environment/population"
)

type randomSelecter struct{}

func (s randomSelecter) Select(pop *population.Population, survivorsSize int) *population.Population {
	size := pop.Len() - survivorsSize
	for count := 0; count < size; {
		pop.Remove(rand.Intn(pop.Len() - 1))
		count++
	}
	return pop
}

// NewRandomSelecter is the constrctor for truncation selecter
func NewRandomSelecter() Interface {
	return randomSelecter{}
}
