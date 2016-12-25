package selecter

import (
	"math/rand"

	"github.com/khezen/darwin/population"
)

type tournamentSelecter struct{}

func (s tournamentSelecter) Select(pop *population.Population, survivorsSize int) *population.Population {
	newPop, _ := population.New(pop.Cap())
	for newPop.Len() <= survivorsSize {
		switch pop.Len() {
		case 1:
			indiv, _ := pop.Remove(0)
			pop.Append(indiv)
		default:
			var i, _, j, _ = pop.PickCouple()
			survivorIndex := s.fightForYourLives(pop, i, j)
			indiv, _ := pop.Remove(survivorIndex)
			newPop.Append(indiv)
		}
	}
	return newPop
}

func (s tournamentSelecter) fightForYourLives(pop *population.Population, index1 int, index2 int) (survivorIndex int) {
	i1, _ := pop.Get(index1)
	i2, _ := pop.Get(index2)
	r1, r2 := i1.Resilience(), i2.Resilience()
	total := r1 + r2
	switch {
	case rand.Float32() <= r1/total:
		return index1
	default:
		return index2
	}
}

// NewTournamentSelecter is the constrctor for truncation selecter
func NewTournamentSelecter() Interface {
	return tournamentSelecter{}
}
