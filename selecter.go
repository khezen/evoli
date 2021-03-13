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
			indiv := pop.Get(i)
			if newPop.Len() >= survivorsSize {
				leftovers.Add(indiv)
				continue
			}
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

type stochasticUniversalSampling struct{}

func (s stochasticUniversalSampling) Select(pop Population, survivorsSize int) (survivors, deads Population, err error) {
	checkSelectParams(survivorsSize)
	if survivorsSize >= pop.Len() {
		return pop, nil, nil
	}
	survivors = pop.New(pop.Cap())
	deads = pop.New(pop.Cap() - survivorsSize)
	var (
		minIndiv = pop.Min()
		minFit   = minIndiv.Fitness()
		offset   float64
		totalFit float64
	)
	if minFit < 0 {
		offset = -minFit
	}
	pop.Each(func(indiv Individual) bool {
		totalFit += indiv.Fitness() + offset
		return true
	})
	step := totalFit / float64(survivorsSize)
	start := rand.Float64() * step
	milestones := make([]float64, 0, survivorsSize)
	for i := 0; i < survivorsSize; i++ {
		milestones = append(milestones, start+float64(i)*step)
	}
	var (
		fitSum     float64
		indiv      Individual
		isSelected bool
		i          int
	)
	pop.Sort()
	for _, milestone := range milestones {
		isSelected = false
		for !isSelected {
			indiv = pop.Get(i)
			fitSum += indiv.Fitness() + offset
			isSelected = fitSum >= milestone
			if isSelected {
				survivors.Add(indiv)
			} else {
				deads.Add(indiv)
			}
			i++
		}
	}
	pop.Close()
	return survivors, deads, nil
}

// NewStochasticUniversalSamplingSelecter is the constructor for selecter based on fitness value
func NewStochasticUniversalSamplingSelecter() Selecter {
	return stochasticUniversalSampling{}
}

type tournamentSelecter struct {
	p float64
}

func (s tournamentSelecter) Select(pop Population, survivorsSize int) (survivors, deads Population, err error) {
	checkSelectParams(survivorsSize)
	if survivorsSize >= pop.Len() {
		return pop, nil, nil
	}
	survivors = pop.New(pop.Cap())
	var leftovers Population
	var popLen int
	pop.Sort()
	for survivors.Len() < survivorsSize {
		popLen = pop.Len()
		leftovers = pop.New(pop.Cap() - survivorsSize + survivors.Len())
		for i := 0; i < popLen; i++ {
			indiv := pop.Get(i)
			if survivors.Len() >= survivorsSize {
				leftovers.Add(indiv)
				continue
			}
			draw := rand.Float64()
			if draw <= s.p {
				survivors.Add(indiv)
			} else {
				leftovers.Add(indiv)
			}
		}
		pop.Close()
		pop = leftovers
	}
	return survivors, pop, nil
}

// NewTournamentSelecter is the constructor for tournament selecter. High rank increase chances to be selected
func NewTournamentSelecter(p float64) Selecter {
	return tournamentSelecter{p}
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
			indiv := pop.Get(i)
			if newPop.Len() >= survivorsSize {
				leftovers.Add(indiv)
				continue
			}
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
