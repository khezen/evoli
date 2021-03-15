package evoli

import (
	"errors"
	"math/rand"
	"sync"
)

// genetic is a genetic algorithm implementation
type genetic struct {
	evolution
	selecter            Selecter
	SurvivorSize        int
	crosser             Crosser
	mutater             Mutater
	MutationProbability float64
}

var (
	// ErrSurvivorSize - survivorSize < 1
	ErrSurvivorSize = errors.New("ErrSurvivorSize - survivorSize must be >= 1")
	// ErrMutationProb - 0<= mutationProbability <= 1
	ErrMutationProb = errors.New("ErrMutationProb - mutation probability must be 0 <= mutationProbability <= 1")
)

// NewGenetic - constructor for Genetic Algorithm
func NewGenetic(pop Population, s Selecter, survivorSize int, c Crosser, m Mutater, mutationProbability float64, e Evaluater) Evolution {
	if survivorSize < 1 {
		panic(ErrSurvivorSize)
	}
	if mutationProbability < 0 || mutationProbability > 1 {
		panic(ErrMutationProb)
	}
	return &genetic{newEvolution(pop, e), s, survivorSize, c, m, mutationProbability}
}

// Next takes a population and produce a the new generation of this population
func (g *genetic) Next() error {
	err := g.evaluation(g.pop)
	if err != nil {
		return err
	}
	survivors, deads, err := g.selecter.Select(g.pop, g.SurvivorSize)
	if err != nil {
		return err
	}
	if deads != nil {
		deads.Close()
	}
	offsprings, err := g.crossovers(survivors)
	if err != nil {
		return err
	}
	offsprings, err = g.mutations(offsprings)
	if err != nil {
		return err
	}
	survivors.Add(offsprings.Slice()...)
	g.pop = survivors
	return nil
}

func (g *genetic) evaluation(pop Population) error {
	var (
		length     = pop.Len()
		wg         = sync.WaitGroup{}
		bubbledErr error
	)
	for i := 0; i < length; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			individual := pop.Get(i)
			fitness, err := g.evaluater.Evaluate(individual)
			if err != nil {
				bubbledErr = err
				return
			}
			individual.SetFitness(fitness)
		}(i)
	}
	wg.Wait()
	if bubbledErr != nil {
		return bubbledErr
	}
	return nil
}

func (g *genetic) crossovers(pop Population) (Population, error) {
	var (
		capacity   = pop.Cap() - pop.Len()
		offsprings = NewPopulation(capacity)
		mut        sync.Mutex
		wg         = sync.WaitGroup{}
		bubbledErr error
	)
	for index := 0; index < capacity; index += 2 {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			var i, j = rand.Intn(pop.Len()), rand.Intn(pop.Len())
			if i == j {
				switch i {
				case pop.Len() - 1:
					j = i - 1
				default:
					j = i + 1
				}
			}
			indiv1, indiv2 := pop.Get(i), pop.Get(j)
			child1, child2, err := g.crosser.Cross(indiv1, indiv2)
			if err != nil {
				bubbledErr = err
				return
			}
			mut.Lock()
			offsprings.Add(child1)
			if index+1 < capacity {
				offsprings.Add(child2)
			}
			mut.Unlock()
		}(index)
	}
	wg.Wait()
	if bubbledErr != nil {
		return nil, bubbledErr
	}
	return offsprings, nil
}

func (g *genetic) mutations(pop Population) (Population, error) {
	var (
		wg         = sync.WaitGroup{}
		bubbledErr error
	)
	for i := 0; i < pop.Len(); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			indiv := pop.Get(i)
			mutant, err := g.mutater.Mutate(indiv, g.MutationProbability)
			if err != nil {
				bubbledErr = err
				return
			}
			pop.Replace(i, mutant)
		}(i)
	}
	wg.Wait()
	if bubbledErr != nil {
		return nil, bubbledErr
	}
	return pop, nil
}

type geneticSync struct {
	genetic
	sync.RWMutex
}

// NewGeneticSync - constructor for Genetic Algorithm (sync impl)
func NewGeneticSync(pop Population, s Selecter, survivorSize int, c Crosser, m Mutater, mutationProbability float64, e Evaluater) Evolution {
	return &geneticSync{*NewGenetic(pop, s, survivorSize, c, m, mutationProbability, e).(*genetic), sync.RWMutex{}}
}

func (s *geneticSync) Next() error {
	s.Lock()
	defer s.Unlock()
	return s.genetic.Next()
}

func (s *geneticSync) Population() Population {
	s.RLock()
	defer s.RUnlock()
	return s.genetic.Population()
}

func (s *geneticSync) SetPopulation(pop Population) {
	s.Lock()
	defer s.Unlock()
	s.genetic.SetPopulation(pop)
}

func (s *geneticSync) Alpha() Individual {
	s.RLock()
	defer s.RUnlock()
	return s.genetic.Alpha()
}
