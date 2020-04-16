package evoli

import (
	"math/rand"
	"sync"
)

var (
	// ErrPoolEvaluater - all evolutions of a pool must share the same evaluater operator
	ErrPoolEvaluater = "ErrPoolEvaluater - all evolution of a pool must share the same evaluater operator"
)

// Pool - solve one problem with different algorithms(with different pros/cons)
// It also make sense to have a pool running multiple instances of the same algorithm asynchronously with smaller populations.
type Pool interface {
	Add(Evolution)
	Delete(Evolution)
	Has(Evolution) bool
	Alpha() Individual
	Individuals() []Individual
	Populations() []Population
	Evolutions() []Evolution
	Shuffle()
	Next() error
	NextAsync() error
}

type pool struct {
	evaluater  Evaluater
	evolutions []Evolution
}

// NewPool - creates a Pool
func NewPool(length int) Pool {
	return &pool{nil, make([]Evolution, 0, length)}
}

func (p *pool) Add(e Evolution) {
	switch p.evaluater {
	case nil:
		p.evaluater = e.Evaluater()
	case e.Evaluater():
		break
	default:
		panic(ErrPoolEvaluater)
	}
	p.evolutions = append(p.evolutions, e)
}

func (p *pool) Delete(e Evolution) {
	length := len(p.evolutions)
	for i := range p.evolutions {
		if p.evolutions[i] == e {
			p.evolutions[i] = p.evolutions[length-1]
			p.evolutions[length-1] = nil
			p.evolutions = p.evolutions[:length-1]
			break
		}
	}
	if len(p.evolutions) == 0 {
		p.evaluater = nil
	}
}

func (p *pool) Has(e Evolution) bool {
	for i := range p.evolutions {
		if p.evolutions[i] == e {
			return true
		}
	}
	return false
}

func (p *pool) Evolutions() []Evolution {
	return p.evolutions
}

func (p *pool) Populations() []Population {
	populations := make([]Population, 0, len(p.evolutions))
	for _, e := range p.evolutions {
		populations = append(populations, e.Population())
	}
	return populations
}

func (p *pool) Individuals() []Individual {
	individualsLen := 0
	for _, e := range p.evolutions {
		individualsLen += e.Population().Len()
	}
	individuals := make([]Individual, 0, individualsLen)
	for _, e := range p.evolutions {
		individuals = append(individuals, e.Population().Slice()...)
	}
	return individuals
}

func (p *pool) Alpha() Individual {
	var alpha Individual
	for _, e := range p.evolutions {
		outsider := e.Alpha()
		if alpha == nil || alpha.Fitness() < outsider.Fitness() {
			alpha = outsider
		}
	}
	return alpha
}

func (p *pool) Shuffle() {
	individualsCap := 0
	prevPopulations := make([]Population, 0, len(p.evolutions))
	nextPopulations := make([]Population, 0, len(p.evolutions))
	for _, e := range p.evolutions {
		pop := e.Population()
		capacity := pop.Cap()
		individualsCap += capacity
		newPop := pop.New(capacity)
		prevPopulations = append(prevPopulations, newPop)
	}
	individuals := make([]Individual, 0, individualsCap)
	for _, e := range p.evolutions {
		individuals = append(individuals, e.Population().Slice()...)
	}
	for _, indiv := range individuals {
		prevPopulationsLen := len(prevPopulations)
		i := rand.Intn(prevPopulationsLen)
		prevPopulations[i].Add(indiv)
		if prevPopulations[i].Len() == prevPopulations[i].Cap() {
			nextPopulations = append(nextPopulations, prevPopulations[i])
			prevPopulations[i] = prevPopulations[prevPopulationsLen-1]
			prevPopulations[prevPopulationsLen-1] = nil
			prevPopulations = prevPopulations[:prevPopulationsLen-1]
		}
	}
	for i := range p.evolutions {
		p.evolutions[i].SetPopulation(nextPopulations[i])
	}
}

func (p *pool) Next() error {
	for _, e := range p.evolutions {
		err := e.Next()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *pool) NextAsync() error {
	evolutionsLen := len(p.evolutions)
	wg := sync.WaitGroup{}
	wg.Add(evolutionsLen)
	var bubbledErr error
	for _, e := range p.evolutions {
		go func(e Evolution) {
			err := e.Next()
			if err != nil {
				bubbledErr = err
			}
			wg.Done()
		}(e)
	}
	wg.Wait()
	return bubbledErr
}
