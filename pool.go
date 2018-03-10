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
	NextAsync() error
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
	populationSlice := make([]*Population, 0, len(p.populations))
	populationMap := make(map[*Population]Evolution)
	for pop, evolution := range p.populations {
		capacity := pop.Cap()
		individualsCap += capacity
		newPop := pop.New(capacity)
		populationSlice = append(populationSlice, &newPop)
		populationMap[&newPop] = evolution
	}
	individuals := make([]Individual, 0, individualsCap)
	for pop := range p.populations {
		individuals = append(individuals, pop.Slice()...)
	}
	for _, indiv := range individuals {
		populationSliceLen := len(populationSlice)
		i := rand.Intn(populationSliceLen)
		(*populationSlice[i]).Add(indiv)
		if (*populationSlice[i]).Len() == (*populationSlice[i]).Cap() {
			populationSlice[i] = populationSlice[populationSliceLen-1]
			populationSlice[populationSliceLen-1] = nil
			populationSlice = populationSlice[:populationSliceLen-1]
		}
	}
	populations := make(map[Population]Evolution)
	for pop, evolution := range populationMap {
		populations[*pop] = evolution
	}
	p.populations = populations
}

func (p *pool) Next() error {
	populations := make(map[Population]Evolution)
	for population, evolution := range p.populations {
		newPop, err := evolution.Next(population)
		if err != nil {
			return err
		}
		populations[newPop] = evolution
	}
	p.populations = populations
	return nil
}

func (p *pool) NextAsync() error {
	populationsLen := len(p.populations)
	type ResultSet struct {
		pop       Population
		evolution Evolution
		err       error
	}
	results := make([]chan ResultSet, 0, populationsLen)
	for i := 0; i < populationsLen; i++ {
		results = append(results, make(chan ResultSet))
	}
	i := 0
	for population, evolution := range p.populations {
		go func(population Population, evolution Evolution, resChan chan ResultSet) {
			newPop, err := evolution.Next(population)
			if err != nil {
				resChan <- ResultSet{nil, nil, err}
				return
			}
			resChan <- ResultSet{newPop, evolution, nil}
		}(population, evolution, results[i])
		i++
	}
	populations := make(map[Population]Evolution)
	for _, resChan := range results {
		res := <-resChan
		if res.err != nil {
			return res.err
		}
		populations[res.pop] = res.evolution
	}
	p.populations = populations
	return nil
}
