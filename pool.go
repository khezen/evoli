package evoli

import "math/rand"

var (
	// ErrPoolEvaluater - all evolution of a pool must share the same evaluater operator
	ErrPoolEvaluater = "ErrPoolEvaluater - all evolution of a pool must share the same evaluater operator"
)

// Pool -
type Pool interface {
	Put(Population, Evolution)
	Delete(Population)
	Has(Population) bool
	Evolution(Population) Evolution
	Populations() []Population
	Individuals() []Individual
	Max() Individual
	Min() Individual
	Shuffle()
	Next() error
}

type pool struct {
	evaluater   Evaluater
	populations map[Population]Evolution
}

// NewPool - creates a Pool
func NewPool() Pool {
	return &pool{nil, make(map[Population]Evolution)}
}

func (p *pool) Put(pop Population, e Evolution) {
	switch p.evaluater {
	case nil:
		p.evaluater = e.Evaluater()
		break
	case e.Evaluater():
		break
	default:
		panic(ErrPoolEvaluater)
	}
	p.populations[pop] = e
}

func (p *pool) Delete(pop Population) {
	delete(p.populations, pop)
	if len(p.populations) == 0 {
		p.evaluater = nil
	}
}

func (p *pool) Has(pop Population) bool {
	_, ok := p.populations[pop]
	return ok
}

func (p *pool) Evolution(pop Population) Evolution {
	return p.populations[pop]
}

func (p *pool) Populations() []Population {
	populations := make([]Population, 0, len(p.populations))
	for pop := range p.populations {
		populations = append(populations, pop)
	}
	return populations
}

func (p *pool) Individuals() []Individual {
	individualsLen := 0
	for pop := range p.populations {
		individualsLen += pop.Len()
	}
	individuals := make([]Individual, 0, individualsLen)
	for pop := range p.populations {
		individuals = append(individuals, pop.Slice()...)
	}
	return individuals
}

func (p *pool) Max() Individual {
	var max Individual
	for pop := range p.populations {
		outsider := pop.Max()
		if max == nil || max.Fitness() < outsider.Fitness() {
			max = outsider
		}
	}
	return max
}

func (p *pool) Min() Individual {
	var min Individual
	for pop := range p.populations {
		outsider := pop.Min()
		if min == nil || min.Fitness() > outsider.Fitness() {
			min = outsider
		}
	}
	return min
}

func (p *pool) Shuffle() {
	individualsCap := 0
	populationSlice := make([]Population, 0, len(p.populations))
	for pop := range p.populations {
		capacity := pop.Cap()
		populationSlice = append(populationSlice, pop.New(capacity))
		individualsCap += capacity
	}
	individuals := make([]Individual, 0, individualsCap)
	for pop := range p.populations {
		individuals = append(individuals, pop.Slice()...)
	}
	for _, indiv := range individuals {
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
	populationsLen := len(p.populations)
	results := make([]chan error, 0, populationsLen)
	for i := 0; i < populationsLen; i++ {
		results = append(results, make(chan error))
	}
	i := 0
	for population, evolution := range p.populations {
		go func(errChan chan error) {
			newPop, err := evolution.Next(population)
			if err != nil {
				errChan <- err
				return
			}
			delete(p.populations, population)
			p.populations[newPop] = evolution
		}(results[i])
		i++
	}
	for _, errChan := range results {
		err := <-errChan
		if err != nil {
			return err
		}
	}
	return nil
}
