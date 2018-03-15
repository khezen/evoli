package evoli

import (
	"math/rand"
)

// Selecter is the selecter operator interface
type Selecter interface {
	Select(pop Population, survivorsSize int) (Population, error)
}

func checkSelectParams(survivorsSize int) {
	if survivorsSize < 1 {
		panic(ErrSurvivorSize)
	}
}

type proportionalToFitnessSelecter struct{}

func (s proportionalToFitnessSelecter) Select(pop Population, survivorsSize int) (Population, error) {
	checkSelectParams(survivorsSize)
	if survivorsSize >= pop.Len() {
		return pop, nil
	}
	newPop := pop.New(pop.Cap())
	offset := s.computeOffset(pop)
	totalScore := s.computeTotalScore(pop, offset)
	for newPop.Len() < survivorsSize {
		for i := 0; i < pop.Len(); i++ {
			if newPop.Len() >= survivorsSize {
				break
			}
			indiv := pop.Get(i)
			score := s.computeScore(indiv, offset)
			if totalScore == 0 || rand.Float64() <= score/totalScore {
				pop.RemoveAt(i)
				newPop.Add(indiv)
				totalScore -= score
			}
		}
	}
	return newPop, nil
}

func (s proportionalToFitnessSelecter) computeScore(indiv Individual, offset float64) float64 {
	return indiv.Fitness() + offset
}

func (s proportionalToFitnessSelecter) computeTotalScore(pop Population, offset float64) float64 {
	var length, totalScore = pop.Len(), float64(0)
	for i := 0; i < length; i++ {
		indiv := pop.Get(i)
		totalScore += s.computeScore(indiv, offset)
	}
	return totalScore
}

func (s proportionalToFitnessSelecter) computeOffset(pop Population) float64 {
	minIndiv, maxIndiv := pop.Min(), pop.Max()
	min, max := minIndiv.Fitness(), maxIndiv.Fitness()
	var offset float64
	switch {
	case min < 0:
		offset += -min
	case min > 0:
		offset += min
	}
	switch {
	case max < 0:
		offset += -max
	case max > 0:
		offset += max
	}
	return offset
}

// NewProportionalToFitnessSelecter is the constructor for selecter based on fitness value
func NewProportionalToFitnessSelecter() Selecter {
	return proportionalToFitnessSelecter{}
}

type tournamentSelecter struct{}

func (s tournamentSelecter) Select(pop Population, survivorsSize int) (Population, error) {
	checkSelectParams(survivorsSize)
	if survivorsSize >= pop.Len() {
		return pop, nil
	}
	newPop := pop.New(pop.Cap())
	for newPop.Len() < survivorsSize {
		var i, j = rand.Intn(pop.Len()), rand.Intn(pop.Len())
		if i == j {
			switch i {
			case pop.Len() - 1:
				j = i - 1
			default:
				j = i + 1
			}
		}
		survivorIndex := s.fightForYourLives(pop, i, j)
		indiv := pop.Get(survivorIndex)
		pop.RemoveAt(survivorIndex)
		newPop.Add(indiv)

	}
	return newPop, nil
}

func (s tournamentSelecter) fightForYourLives(pop Population, index1 int, index2 int) (survivorIndex int) {
	i1, i2 := pop.Get(index1), pop.Get(index2)
	r1, r2 := i1.Fitness(), i2.Fitness()

	offset := s.computeOffset(r1, r2)
	r1 += offset
	r2 += offset
	total := r1 + r2

	switch {
	case total == 0 || rand.Float64() <= r1/total:
		return index1
	default:
		return index2
	}
}

func (s tournamentSelecter) computeOffset(r1, r2 float64) float64 {
	var offset float64
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

type truncationSelecter struct{}

func (s truncationSelecter) Select(pop Population, survivorsSize int) (Population, error) {
	checkSelectParams(survivorsSize)
	pop.Sort()
	individuals := pop.Slice()
	var survivors []Individual
	if len(individuals) < survivorsSize {
		survivors = individuals
	} else {
		survivors = individuals[:survivorsSize]
	}
	newPop := pop.New(pop.Cap())
	newPop.Add(survivors...)
	return newPop, nil
}

// NewTruncationSelecter is the constructor for truncation selecter
func NewTruncationSelecter() Selecter {
	return truncationSelecter{}
}

type randomSelecter struct{}

func (s randomSelecter) Select(pop Population, survivorsSize int) (Population, error) {
	checkSelectParams(survivorsSize)
	newPop := pop.New(pop.Cap())
	newPop.Add(pop.Slice()...)
	size := newPop.Len() - survivorsSize
	for count := 0; count < size; count++ {
		newPop.RemoveAt(rand.Intn(newPop.Len() - 1))
	}
	return newPop, nil
}

// NewRandomSelecter is the constructor for random selecter
func NewRandomSelecter() Selecter {
	return randomSelecter{}
}

type proportionalToRankSelecter struct{}

func (s proportionalToRankSelecter) Select(pop Population, survivorsSize int) (Population, error) {
	checkSelectParams(survivorsSize)
	if survivorsSize >= pop.Len() {
		return pop, nil
	}
	newPop := pop.New(pop.Cap())
	totalScore := s.computeTotalScore(pop)
	pop.Sort()
	for newPop.Len() < survivorsSize {
		for i := 0; i < pop.Len(); i++ {
			if newPop.Len() >= survivorsSize {
				break
			}
			score := float64(pop.Len() - i)
			if rand.Float64() <= score/totalScore {
				indiv := pop.Get(i)
				pop.RemoveAt(i)
				newPop.Add(indiv)
				totalScore -= score
			}
		}
	}
	return newPop, nil
}

func (s proportionalToRankSelecter) computeTotalScore(pop Population) float64 {
	n := float64(pop.Len())
	return n * (n + 1) / 2 // 1+2+3+...+n
}

// NewProportionalToRankSelecter is the constructor for selecter based on ranking across the population
func NewProportionalToRankSelecter() Selecter {
	return proportionalToRankSelecter{}
}
