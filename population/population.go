package population

import (
	"math/rand"
	"sort"

	"github.com/khezen/darwin/population/individual"
)

// Interface is the population interface
type Interface interface {
	sort.Interface
	Sort()
	Cap() int
	SetCap(int) error
	Truncate(int) error
	Append(individual.Interface) error
	AppendAll(Interface) error
	Get(int) (individual.Interface, error)
	Remove(int) (individual.Interface, error)
	Max() individual.Interface
	Min() individual.Interface
	PickCouple() (index1, index2 int)
}

// Population is a set of individuals in population genetics.
type Population []individual.Interface

// New is population constructor
func New(capacity int) (*Population, error) {
	pop := Population(make([]individual.Interface, 0, capacity))
	return &pop, nil
}

// Len returns the current livings count of a population
func (pop *Population) Len() int {
	return len(*pop)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (pop *Population) Less(i, j int) bool {
	indivi, _ := pop.Get(i)
	indivj, _ := pop.Get(j)
	return indivi.Resilience() >= indivj.Resilience()
}

// Swap swaps the elements with indexes i and j.
func (pop *Population) Swap(i, j int) {
	tmp := (*pop)[i]
	(*pop)[i] = (*pop)[j]
	(*pop)[j] = tmp
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
	currentCap := pop.Cap()
	if newCap != currentCap {
		tmp := *pop
		switch {
		case newCap < currentCap:
			tmp = (*pop)[0:newCap]
			*pop = make([]individual.Interface, newCap, newCap)
		case newCap > currentCap:
			*pop = make([]individual.Interface, currentCap, newCap)
		}
		copy(tmp, *pop)
	}
	return nil
}

// Truncate rduce population size to the given length
func (pop *Population) Truncate(length int) error {
	if length < pop.Len() {
		*pop = (*pop)[0 : length-1]
	}
	return nil
}

// Append adds an individual to a population. If the populagtion has already reached its capacity, capacity is incremented.
func (pop *Population) Append(indiv individual.Interface) error {
	*pop = append(*pop, indiv)
	return nil
}

// AppendAll adds all individuals from a population to a population. If the populagtion has already reached its capacity, capacity is incremented.
func (pop *Population) AppendAll(externalPop Interface) error {
	for i := 0; i < externalPop.Len(); i++ {
		indiv, _ := externalPop.Get(i)
		pop.Append(indiv)
	}
	return nil
}

// Get returns the individual at index i
func (pop *Population) Get(i int) (individual.Interface, error) {
	return (*pop)[i], nil
}

// Remove removes and returns the individual at index i
func (pop *Population) Remove(i int) (individual.Interface, error) {
	removed, _ := pop.Get(i)
	new := (*pop)[0 : i-1]
	*pop = append(new, (*pop)[i+1:pop.Len()-1]...)
	return removed, nil
}

// Min returns the least Resilent individual
func (pop *Population) Min() individual.Interface {
	min, _ := pop.Get(0)
	length := pop.Len()
	for i := 1; i < length; i++ {
		indiv, _ := pop.Get(i)
		if indiv.Resilience() < min.Resilience() {
			min, _ = pop.Get(i)
		}
	}
	return min
}

// Max returns the most Resilent individual
func (pop *Population) Max() individual.Interface {
	max, _ := pop.Get(0)
	length := pop.Len()
	for i := 1; i < length; i++ {
		indiv, _ := pop.Get(i)
		if indiv.Resilience() > max.Resilience() {
			max, _ = pop.Get(i)
		}
	}
	return max
}

// PickCouple returns the index of two randomly choosen individuals
func (pop *Population) PickCouple() (index1, index2 int) {
	var i, j = rand.Intn(pop.Len() - 1), rand.Intn(pop.Len() - 1)
	if i == j {
		switch i {
		case pop.Len() - 1:
			j = i - 1
		default:
			j = i + 1
		}
	}
	return i, j
}
