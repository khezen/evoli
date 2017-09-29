package darwin

import (
	"math/rand"
)

type tournamentSelecter struct{}

func (s tournamentSelecter) Select(pop Population, survivorsSize int) (Population, error) {
	err := checkParams(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	if survivorsSize >= pop.Len() {
		return pop, nil
	}
	newPop := pop.New(pop.Cap())
	for newPop.Len() < survivorsSize {
		var i, _, j, _, _ = pop.PickCouple()
		survivorIndex := s.fightForYourLives(pop, i, j)
		indiv, _ := pop.Get(survivorIndex)
		pop.RemoveAt(survivorIndex)
		newPop.Add(indiv)

	}
	return newPop, nil
}

func (s tournamentSelecter) fightForYourLives(pop Population, index1 int, index2 int) (survivorIndex int) {
	i1, _ := pop.Get(index1)
	i2, _ := pop.Get(index2)
	r1, r2 := i1.Fitness(), i2.Fitness()

	offset := s.computeOffset(r1, r2)
	r1 += offset
	r2 += offset
	total := r1 + r2

	switch {
	case total == 0 || rand.Float32() <= r1/total:
		return index1
	default:
		return index2
	}
}

func (s tournamentSelecter) computeOffset(r1, r2 float32) float32 {
	var offset float32
	switch {
	case r1 < 0:
		offset += -r1
	case r1 > 0:
		offset += r1
	}
	switch {
	case r2 < 0:
		offset += -r2
	case r2 > 0:
		offset += r2
	}
	return offset
}

// NewTournamentSelecter is the constructor for tournament selecter. High Fitness increase chances to come out vitorious from a duel
func NewTournamentSelecter() Selecter {
	return tournamentSelecter{}
}
