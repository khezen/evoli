package darwin

import (
	"math/rand"
	"sort"

	"fmt"
	"github.com/khezen/check"
)

// IPopulation is the population interface
type IPopulation interface {
	sort.Interface
	Sort()
	Cap() int
	SetCap(int) error
	Truncate(int) error
	Append(IIndividual) error
	AppendAll(IPopulation) error
	Get(int) (IIndividual, error)
	Remove(int) (IIndividual, error)
	Replace(int, IIndividual) (IIndividual, error)
	Max() IIndividual
	Min() IIndividual
	Extremums() (IIndividual, IIndividual)
	PickCouple() (index1 int, indiv1 IIndividual, index2 int, indiv2 IIndividual, err error)
	Contains(IIndividual) bool
	IndexOf(IIndividual) (int, error)
}

// Population is a set of individuals in population genetics.
type Population []IIndividual

// NewPopulation is population constructor
func NewPopulation(capacity int) (*Population, error) {
	err := check.Cap(capacity)
	if err != nil {
		return nil, err
	}
	pop := Population(make([]IIndividual, 0, capacity))
	return &pop, nil
}

// Len returns the current livings count of a population
func (pop *Population) Len() int {
	return len(*pop)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (pop *Population) Less(i, j int) bool {
	indivi, erri := pop.Get(i)
	indivj, errj := pop.Get(j)
	if erri != nil || errj != nil {
		return false
	}
	return indivi.Fitness() >= indivj.Fitness()
}

// Swap swaps the elements with indexes i and j.
func (pop *Population) Swap(i, j int) {
	if i != j && i >= 0 && j >= 0 && i < pop.Len() && j < pop.Len() {
		tmp := (*pop)[i]
		(*pop)[i] = (*pop)[j]
		(*pop)[j] = tmp
	}
}

// Sort sort the population
func (pop *Population) Sort() {
	sort.Sort(pop)
}

// Cap returns the population capacity
func (pop *Population) Cap() int {
	return cap(*pop)
}

// SetCap set the resize the population capacity
func (pop *Population) SetCap(newCap int) error {
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
			*pop = make([]IIndividual, newCap, newCap)
		case newCap > currentCap:
			*pop = make([]IIndividual, currentCap, newCap)
		}
		copy(tmp, *pop)
	}
	return nil
}

// Truncate reduce population size to the given length
func (pop *Population) Truncate(length int) error {
	err := check.Length(length)
	if err != nil {
		return err
	}
	switch {
	case length == 0:
		newPop, _ := NewPopulation(0)
		*pop = *newPop
	case length < pop.Len():
		*pop = (*pop)[0:length]
	case length > pop.Cap():
		newPop, _ := NewPopulation(length)
		newPop.AppendAll(pop)
		*pop = *newPop
	}
	return nil
}

// Append adds an individual to a population. If the populagtion has already reached its capacity, capacity is incremented.
func (pop *Population) Append(indiv IIndividual) error {
	err := check.NotNil(indiv)
	if err != nil {
		return err
	}
	*pop = append(*pop, indiv)
	return nil
}

// AppendAll adds all individuals from a population to a population. If the populagtion has already reached its capacity, capacity is incremented. Nil individuals are not appended
func (pop *Population) AppendAll(externalPop IPopulation) error {
	err := check.NotNil(externalPop)
	if err != nil {
		return err
	}
	for i := 0; i < externalPop.Len(); i++ {
		indiv, _ := externalPop.Get(i)
		appendErr := pop.Append(indiv)
		if appendErr != nil { // indiv has not been appended
			err = appendErr
		}
	}
	return err
}

// Get returns the individual at index i
func (pop *Population) Get(i int) (IIndividual, error) {
	err := check.Index(i, pop.Len())
	if err != nil {
		return nil, err
	}
	return (*pop)[i], nil
}

// Remove removes and returns the individual at index i
func (pop *Population) Remove(i int) (IIndividual, error) {
	err := check.Index(i, pop.Len())
	if err != nil {
		return nil, err
	}
	removed, _ := pop.Get(i)
	new := (*pop)[0:i]
	*pop = append(new, (*pop)[i+1:pop.Len()]...)
	return removed, nil
}

// Replace replaces and returns the individual at index i by the substitute
func (pop *Population) Replace(i int, substitute IIndividual) (IIndividual, error) {
	err := check.Index(i, pop.Len())
	if err != nil {
		return nil, err
	}
	err = check.NotNil(substitute)
	if err != nil {
		return nil, err
	}
	removed, _ := pop.Get(i)
	(*pop)[i] = substitute
	return removed, nil
}

// Min returns the least Resilent individual
func (pop *Population) Min() IIndividual {
	return pop.extremum(false)
}

// Max returns the most Resilent individual
func (pop *Population) Max() IIndividual {
	return pop.extremum(true)
}

func (pop *Population) extremum(greaterThan bool) IIndividual {
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
func (pop *Population) Extremums() (min, max IIndividual) {
	return pop.Min(), pop.Max()
}

// PickCouple returns two randomly chosen individuals with their index
func (pop *Population) PickCouple() (int, IIndividual, int, IIndividual, error) {
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

// Contains return true if the specified individual is in the population
func (pop *Population) Contains(indiv IIndividual) bool {
	_, err := pop.IndexOf(indiv)
	switch {
	case err != nil:
		return false
	default:
		return true
	}
}

// IndexOf returns the inde of the specified individual if it exists
func (pop *Population) IndexOf(indiv IIndividual) (int, error) {
	err := check.NotNil(indiv)
	if err != nil {
		return -1, err
	}
	for i, current := range *pop {
		if current == indiv {
			return i, nil
		}
	}
	return -1, fmt.Errorf("individual %v not found in population %v", indiv, pop)
}
