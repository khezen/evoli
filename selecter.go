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
	var (
		leftovers          Population
		newPop             = pop.New(pop.Cap())
		minIndiv, maxIndiv Individual
		min, max           float64
		pivot              float64
		score              float64
		benchmark          float64
		penaltyStep        float64
		penalty            float64
	)
	for newPop.Len() < survivorsSize {
		leftovers = pop.New(pop.Len())
		minIndiv, maxIndiv = pop.Min(), pop.Max()
		min, max = minIndiv.Fitness(), maxIndiv.Fitness()
		if min < 0 {
			pivot += -min
		} else {
			pivot += min
		}
		benchmark = max + pivot
		penalty = 0
		penaltyStep = benchmark / float64(pop.Len())
		for i := 0; i < pop.Len(); i++ {
			if newPop.Len() >= survivorsSize {
				break
			}
			indiv := pop.Get(i)
			score = (indiv.Fitness() + pivot - penalty) / benchmark
			if rand.Float64() <= 0.1+0.8*score {
				newPop.Add(indiv)
				penalty += penaltyStep
			} else {
				leftovers.Add(indiv)
			}
		}
		pop.Close()
		pop = leftovers
	}
	pop.Close()
	return newPop, nil
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
	pop.Close()
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
	case total == 0, rand.Float64() <= r1/total:
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
	newPop := pop.New(pop.Cap())
	if pop.Len() < survivorsSize {
		newPop.Add(pop.Slice()...)
	} else {
		newPop.Add(pop.Slice()[:survivorsSize]...)
	}
	pop.Close()
	return newPop, nil
}

// NewTruncationSelecter is the constructor for truncation selecter
func NewTruncationSelecter() Selecter {
	return truncationSelecter{}
}

type randomSelecter struct{}

func (s randomSelecter) Select(pop Population, survivorsSize int) (Population, error) {
	checkSelectParams(survivorsSize)
	var (
		newPop = pop.New(pop.Cap())
		count  int
	)
	newPop.Add(pop.Slice()...)
	size := newPop.Len() - survivorsSize
	for count = 0; count < size; count++ {
		newPop.RemoveAt(rand.Intn(newPop.Len() - 1))
	}
	pop.Close()
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
	var (
		leftovers Population
		newPop    = pop.New(pop.Cap())
		benchmark float64
		score     float64
		popLen    int
		n         float64
	)
	pop.Sort()
	for newPop.Len() < survivorsSize {
		popLen = pop.Len()
		leftovers = pop.New(popLen)
		benchmark = float64(popLen*(popLen+1)) / 2
		for i := 0; i < popLen; i++ {
			if newPop.Len() >= survivorsSize {
				break
			}
			indiv := pop.Get(i)
			n = float64(popLen - i)
			score = n * (n + 1) / 2
			if rand.Float64() <= 0.25+0.5*score/benchmark {
				newPop.Add(indiv)
			} else {
				leftovers.Add(indiv)
			}
		}
		pop.Close()
		pop = leftovers
	}
	pop.Close()
	return newPop, nil
}

// NewProportionalToRankSelecter is the constructor for selecter based on ranking across the population
func NewProportionalToRankSelecter() Selecter {
	return proportionalToRankSelecter{}
}
