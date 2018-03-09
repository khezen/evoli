package evoli

import "math/rand"

// ErrPoolEvaluater - all evolution of he pool must shqre the same evaluater operator
var ErrPoolEvaluater = "ErrPoolEvaluater - all evolution of he pool must shqre the same evaluater operator"

// Pool -
type Pool interface {
	Put(Evolution, Population)
	Delete(Evolution)
	Has(Evolution) bool
	Get(Evolution) Population
	Max() Individual
	Min() Individual
	Shuffle()
	Next() error
}

type pool struct {
	evaluater   Evaluater
	populations map[Evolution]Population
}

func (p *pool) Put(e Evolution, pop Population) {
	switch p.evaluater {
	case nil:
		p.evaluater = e.Evaluater()
		break
	case e.Evaluater():
		break
	default:
		panic(ErrPoolEvaluater)
	}
	p.populations[e] = pop
}

func (p *pool) delete(e Evolution) {
	delete(p.populations, e)
	if len(p.populations) == 0 {
		p.evaluater = nil
	}
}

func (p *pool) Has(e Evolution) bool {
	_, ok := p.populations[e]
	return ok
}

func (p *pool) Get(e Evolution) Population {
	return p.populations[e]
}

func (p *pool) Max() Individual {
	var max Individual
	for _, pop := range p.populations {
		outsider := pop.Max()
		if max == nil || max.Fitness() < outsider.Fitness() {
			max = outsider
		}
	}
	return max
}

func (p *pool) Min() Individual {
	var min Individual
	for _, pop := range p.populations {
		outsider := pop.Min()
		if min == nil || min.Fitness() > outsider.Fitness() {
			min = outsider
		}
	}
	return min
}

func (p *pool) Shuffle() {
	allIndivCap := 0
	populationSlice := make([]Population, 0, len(p.populations))
	for _, pop := range p.populations {
		capacity := pop.Cap()
		populationSlice = append(populationSlice, pop.New(capacity))
		allIndivCap += capacity
	}
	allIndiv := make([]Individual, 0, allIndivCap)
	for _, pop := range p.populations {
		allIndiv = append(allIndiv, pop.Slice()...)
	}
	for _, indiv := range allIndiv {
		populationSliceLen := len(populationSlice)
		i := rand.Intn(populationSliceLen)
		pop := populationSlice[i]
		pop.Add(indiv)
		if pop.Len() == pop.Cap() {
			populationSlice[i] = populationSlice[populationSliceLen-1]
			populationSlice[populationSliceLen-1] = nil
			populationSlice = populationSlice[:populationSliceLen-1]
		}
	}
}

func (p *pool) Next() error {
	for evolution, population := range p.populations {
		newPop, err := evolution.Next(population)
		if err != nil {
			return err
		}
		p.populations[evolution] = newPop
	}
	return nil
}
