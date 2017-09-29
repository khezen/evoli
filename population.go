package darwin

import (
	"math/rand"
	"sort"

	"fmt"

	"github.com/khezen/check"
)

// Population is the population interface
type Population interface {
	sort.Interface
	Sort()
	Cap() int
	SetCap(int) error
	Truncate(int) error
	Add(...Individual)
	Get(int) (Individual, error)
	RemoveAt(int) error
	Remove(...Individual)
	Replace(int, Individual) error
	Max() Individual
	Min() Individual
	Extremums() (Individual, Individual)
	PickCouple() (index1 int, indiv1 Individual, index2 int, indiv2 Individual, err error)
	Has(...Individual) bool
	IndexOf(Individual) (int, error)
	Each(func(item Individual) bool)
	Slice() []Individual
}

// population is a set of individuals in population genetics.
type population []Individual

// NewPopulation is population constructor
func NewPopulation(capacity int) Population {
	err := check.Cap(capacity)
	if err != nil {
		return nil
	}
	pop := population(make([]Individual, 0, capacity))
	return &pop
}

// Len returns the current livings count of a population
func (pop *population) Len() int {
	return len(*pop)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (pop *population) Less(i, j int) bool {
	indivi, err := pop.Get(i)
	if err != nil {
		return false
	}
	indivj, err := pop.Get(j)
	if err != nil {
		return false
	}
	return indivi.Fitness() >= indivj.Fitness()
}

// Swap swaps the elements with indexes i and j.
func (pop *population) Swap(i, j int) {
	if i != j && i >= 0 && j >= 0 && i < pop.Len() && j < pop.Len() {
		tmp := (*pop)[i]
		(*pop)[i] = (*pop)[j]
		(*pop)[j] = tmp
	}
}

// Sort sort the population
func (pop *population) Sort() {
	sort.Sort(pop)
}

// Cap returns the population capacity
func (pop *population) Cap() int {
	return cap(*pop)
}

// SetCap set the resize the population capacity
func (pop *population) SetCap(newCap int) error {
	err := check.Cap(newCap)
	if err != nil {
		return err
	}
	currentCap := pop.Cap()
	if newCap != currentCap {
		tmp := *pop
		switch {
		case newCap < currentCap:
			tmp = (*pop)[0:newCap]
			*pop = make([]Individual, newCap, newCap)
		case newCap > currentCap:
			*pop = make([]Individual, pop.Len(), newCap)
		}
		copy(tmp, *pop)
	}
	return nil
}

// Truncate reduce population size to the given length
func (pop *population) Truncate(length int) error {
	err := check.Length(length)
	if err != nil {
		return err
	}
	switch {
	case length == 0:
		*pop = *(NewPopulation(0).(*population))
	case length < pop.Len():
		*pop = (*pop)[0:length]
	case length > pop.Cap():
		newPop := NewPopulation(length)
		newPop.Add((*pop)...)
		*pop = *(newPop.(*population))
	}
	return nil
}

// Add adds an individual to a population. If the populagtion has already reached its capacity, capacity is incremented.
func (pop *population) Add(indiv ...Individual) {
	*pop = append(*pop, indiv...)
}

// Get returns the individual at index i
func (pop *population) Get(i int) (Individual, error) {
	err := check.Index(i, pop.Len())
	if err != nil {
		return nil, err
	}
	return (*pop)[i], nil
}

// RemoveAt removes and returns the individual at index i
func (pop *population) RemoveAt(i int) error {
	err := check.Index(i, pop.Len())
	if err != nil {
		return err
	}
	new := (*pop)[0:i]
	*pop = append(new, (*pop)[i+1:pop.Len()]...)
	return nil
}

// Remove removes all given individuals
func (pop *population) Remove(individuals ...Individual) {
	for _, indiv := range individuals {
		i, err := pop.IndexOf(indiv)
		if err == nil {
			pop.RemoveAt(i)
		}
	}
}

// Replace replaces and returns the individual at index i by the substitute
func (pop *population) Replace(i int, substitute Individual) error {
	err := check.Index(i, pop.Len())
	if err != nil {
		return err
	}
	(*pop)[i] = substitute
	return nil
}

// Min returns the least Resilent individual
func (pop *population) Min() Individual {
	return pop.extremum(false)
}

// Max returns the most Resilent individual
func (pop *population) Max() Individual {
	return pop.extremum(true)
}

func (pop *population) extremum(greaterThan bool) Individual {
	extremum, _ := pop.Get(0)
	length := pop.Len()
	for i := 1; i < length; i++ {
		indiv, _ := pop.Get(i)
		if (greaterThan && indiv.Fitness() > extremum.Fitness()) || (!greaterThan && indiv.Fitness() < extremum.Fitness()) {
			extremum = indiv
		}
	}
	return extremum
}

// Extremums returns the Min() & the Max() of the poplation
func (pop *population) Extremums() (min, max Individual) {
	return pop.Min(), pop.Max()
}

// PickCouple returns two randomly chosen individuals with their index
func (pop *population) PickCouple() (int, Individual, int, Individual, error) {
	if pop.Len() < 2 {
		return -1, nil, -1, nil, fmt.Errorf("pop must contains at least 2 individuals to pick a couple")
	}
	var i, j = rand.Intn(pop.Len()), rand.Intn(pop.Len())
	if i == j {
		switch i {
		case pop.Len() - 1:
			j = i - 1
		default:
			j = i + 1
		}
	}
	indivi, _ := pop.Get(i)
	indivj, _ := pop.Get(j)
	return i, indivi, j, indivj, nil
}

// Has return true if the specified individual is in the population
func (pop *population) Has(individuals ...Individual) bool {
	has := true
	for _, indiv := range individuals {
		_, err := pop.IndexOf(indiv)
		has = has && err == nil
	}
	return has
}

// IndexOf returns the inde of the specified individual if it exists
func (pop *population) IndexOf(indiv Individual) (int, error) {
	for i, current := range *pop {
		if current == indiv {
			return i, nil
		}
	}
	return -1, fmt.Errorf("individual %v not found in population %v", indiv, pop)
}

// Each traverse the population and execute given callback on each individual. Stops if the callbak return false.
func (pop *population) Each(f func(indiv Individual) bool) {
	for _, individual := range *pop {
		resume := f(individual)
		if !resume {
			break
		}
	}
}

// Slice returns the population as []Individual
func (pop *population) Slice() []Individual {
	return *pop
}
