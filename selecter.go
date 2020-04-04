package evoli

import (
	"math/rand"
)

// Selecter is the selecter operator interface
type Selecter interface {
	Select(pop Population, survivorsSize int) (survivors Population, deads Population, err error)
}

func checkSelectParams(survivorsSize int) {
	if survivorsSize < 1 {
		panic(ErrSurvivorSize)
	}
}

type proportionalToFitnessSelecter struct{}

func (s proportionalToFitnessSelecter) Select(pop Population, survivorsSize int) (survivors, deads Population, err error) {
	checkSelectParams(survivorsSize)
	if survivorsSize >= pop.Len() {
		return pop, nil, nil
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
	return newPop, pop, nil
}

// NewProportionalToFitnessSelecter is the constructor for selecter based on fitness value
func NewProportionalToFitnessSelecter() Selecter {
	return proportionalToFitnessSelecter{}
}

type tournamentSelecter struct{}

func (s tournamentSelecter) Select(pop Population, survivorsSize int) (survivors, deads Population, err error) {
	checkSelectParams(survivorsSize)
	if survivorsSize >= pop.Len() {
		return pop, nil, nil
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
	return newPop, pop, nil
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

func (s truncationSelecter) Select(pop Population, survivorsSize int) (survivors, deads Population, err error) {
	checkSelectParams(survivorsSize)
	if pop.Len() <= survivorsSize {
		return pop, nil, nil
	}
	pop.Sort()
	survivors = pop.New(pop.Cap())
	deads = pop.New(pop.Cap() - survivorsSize)
	survivors.Add(pop.Slice()[:survivorsSize]...)
	deads.Add(pop.Slice()[survivorsSize:]...)
	pop.Close()
	return survivors, deads, nil
}

// NewTruncationSelecter is the constructor for truncation selecter
func NewTruncationSelecter() Selecter {
	return truncationSelecter{}
}

type randomSelecter struct{}

func (s randomSelecter) Select(pop Population, survivorsSize int) (survivors, deads Population, err error) {
	checkSelectParams(survivorsSize)
	if pop.Len() <= survivorsSize {
		return pop, nil, nil
	}
	var (
		deathCount = pop.Cap() - survivorsSize
		count      int
		deadIndex  int
	)
	survivors = pop.New(pop.Cap())
	deads = pop.New(deathCount)
	survivors.Add(pop.Slice()...)
	for count = 0; count < deathCount; count++ {
		deadIndex = rand.Intn(survivors.Len() - 1)
		deads.Add(survivors.Get(deadIndex))
		survivors.RemoveAt(deadIndex)
	}
	pop.Close()
	return survivors, deads, nil
}

// NewRandomSelecter is the constructor for random selecter
func NewRandomSelecter() Selecter {
	return randomSelecter{}
}

type proportionalToRankSelecter struct{}

func (s proportionalToRankSelecter) Select(pop Population, survivorsSize int) (survivors, deads Population, err error) {
	checkSelectParams(survivorsSize)
	if survivorsSize >= pop.Len() {
		return pop, nil, nil
	}
	var (
		leftovers Population
		newPop    = pop.New(pop.Cap())
		benchmark float64
		score     float64
		popLen    int
		n         float64
		penalty   int
	)
	pop.Sort()
	for newPop.Len() < survivorsSize {
		popLen = pop.Len()
		leftovers = pop.New(popLen)
		penalty = 0
		benchmark = float64(popLen*(popLen+1)) / 2
		for i := 0; i < popLen; i++ {
			if newPop.Len() >= survivorsSize {
				break
			}
			indiv := pop.Get(i)
			n = float64(popLen - i - penalty)
			score = n * (n + 1) / 2
			if rand.Float64() <= 0.25+0.5*score/benchmark {
				newPop.Add(indiv)
				penalty++
			} else {
				leftovers.Add(indiv)
			}
		}
		pop.Close()
		pop = leftovers
	}
	return newPop, pop, nil
}

// NewProportionalToRankSelecter is the constructor for selecter based on ranking across the population
func NewProportionalToRankSelecter() Selecter {
	return proportionalToRankSelecter{}
}
