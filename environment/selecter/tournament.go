package selecter

import (
	"math/rand"

	"github.com/khezen/darwin/environment/population"
)

type tournamentSelecter struct{}

func (s tournamentSelecter) Select(pop *population.Population, survivorsSize int) *population.Population {
	newPop := population.New(pop.Cap())
	for newPop.Len() <= survivorsSize {
		switch pop.Len() {
		case 1:
			newPop.Append(pop.Remove(0))
		default:
			var i, j = pop.PickCouple()
			survivorIndex := s.fightForYourLives(pop, i, j)
			newPop.Append(pop.Remove(survivorIndex))
		}
	}
	return &newPop
}

func (s tournamentSelecter) fightForYourLives(pop *population.Population, index1 int, index2 int) (survivorIndex int) {
	var r1, r2 = pop.Get(index1).Resilience(), pop.Get(index2).Resilience()
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
