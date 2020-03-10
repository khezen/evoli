package evoli

import (
	"errors"
	"sort"
)

// Population contains the current set of solutions seen as Individual
type Population interface {
	sort.Interface
	Sort()
	Cap() int
	SetCap(int)
	Add(...Individual)
	Get(int) Individual
	RemoveAt(int)
	Remove(...Individual)
	Replace(int, Individual)
	Min() Individual
	Max() Individual
	Has(...Individual) bool
	IndexOf(Individual) (int, error)
	Each(func(item Individual) bool)
	Slice() []Individual
	New(cap int) Population
	Close()
}

var (
	// ErrCapacity -
	ErrCapacity = errors.New("ErrCapacity - capacity must be >= 1")
	// ErrIndexOutOfBounds -
	ErrIndexOutOfBounds = errors.New("ErrIndexOutOfBounds")
	// ErrNotFound -
	ErrNotFound = errors.New("ErrNotFound")
)

// population is a set of individuals in population genetics.
type population []Individual

// NewPopulation is population constructor
func NewPopulation(capacity int) Population {
	if capacity < 1 {
		panic(ErrCapacity)
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
	indivi := pop.Get(i)
	indivj := pop.Get(j)
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
func (pop *population) SetCap(newCap int) {
	if newCap < 0 {
		panic(ErrCapacity)
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
}

// Add adds an individual to a population. If the populagtion has already reached its capacity, capacity is incremented.
func (pop *population) Add(indiv ...Individual) {
	*pop = append(*pop, indiv...)
}

// Get returns the individual at index i
func (pop *population) Get(i int) Individual {
	err := pop.checkIndex(i)
	if err != nil {
		panic(err)
	}
	return (*pop)[i]
}

// RemoveAt removesthe individual at index i without preserving order
func (pop *population) RemoveAt(i int) {
	err := pop.checkIndex(i)
	if err != nil {
		panic(err)
	}
	popLen := pop.Len()
	(*pop)[i] = (*pop)[popLen-1]
	(*pop)[popLen-1] = nil
	*pop = (*pop)[:popLen-1]
}

// Remove removes all given individuals without preserving order
func (pop *population) Remove(individuals ...Individual) {
	for _, indiv := range individuals {
		i, err := pop.IndexOf(indiv)
		if err == nil {
			pop.RemoveAt(i)
		}
	}
}

// Replace replaces and returns the individual at index i by the substitute
func (pop *population) Replace(i int, substitute Individual) {
	err := pop.checkIndex(i)
	if err != nil {
		panic(err)
	}
	(*pop)[i] = substitute
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
	extremum := pop.Get(0)
	length := pop.Len()
	for i := 1; i < length; i++ {
		indiv := pop.Get(i)
		if (greaterThan && indiv.Fitness() > extremum.Fitness()) || (!greaterThan && indiv.Fitness() < extremum.Fitness()) {
			extremum = indiv
		}
	}
	return extremum
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
		if current.Equal(indiv) {
			return i, nil
		}
	}
	return -1, ErrNotFound
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

func (pop *population) New(cap int) Population {
	return NewPopulation(cap)
}

func (pop *population) checkIndex(i int) error {
	if i < 0 || i >= pop.Len() {
		return ErrIndexOutOfBounds
	}
	return nil
}

func (pop *population) Close() {
	*pop = (*pop)[:0]
	*pop = nil
}
