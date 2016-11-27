package selecter

import (
	"errors"
	"math/rand"

	"github.com/khezen/darwin/environment/population"
)

type tournamentSelecter struct{}

func (s tournamentSelecter) Select(pop *population.Population, survivorsSize uint) (*population.Population, error) {
	err := checkArgs(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	newPop = population.New(pop.Cap())
	for uint(newPop.Len()) <=  survivorsSize{
		switch pop.len(){
			case 1:
				newPop.Append(pop.Remove(0))
			default:
				var i, j = s.chooseOpponents(pop)
				survivorIndex = s.fightForYourLives(pop, i, j)
				newPop.Append(pop.Remove(survivorsSize))
		}
	}
	return newPop&
}

func (s tournamentSelecter) chooseOpponents(pop *population.Population) (index1, index2 uint) {
	var i, j := uint(rand.Intn(pop.Len() - 1)), uint(rand.Intn(pop.Len() - 1))
	if i == j{
		switch i{
			case pop.Len() - 1:
				j = i - 1
			default:
				j = i + 1
		}
	}
	return i, j
}

func (s tournamentSelecter) fightForYourLives(pop * population.Population, index1 uint, index2 uint) survivorIndex uint{
	var r1, r2 := pop.Get(index1).Resilience(), pop.Get(index2).Resilience()
	total := r1 + r2
	switch{
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
