package population

import (
	"math/rand"
	"sort"

	"github.com/khezen/darwin/evolution/individual"
)

// Interface is the population interface
type Interface interface {
	sort.Interface
	Sort()
	Cap() int
	SetCap(int)
	Truncate(int)
	Append(individual.Interface)
	Get(int) individual.Interface
	Remove(int) individual.Interface
	Max() individual.Interface
	Min() individual.Interface
}

// Population is a set of individuals in population genetics.
type Population struct {
	individuals []individual.Interface
}

// New is population constructor
func New(capacity int) Population {
	return Population{make([]individual.Interface, 0, capacity)}
}

// Len returns the current livings count of a population
func (pop *Population) Len() int {
	return len(pop.individuals)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (pop *Population) Less(i, j int) bool {
	return pop.Get(int(i)).Resilience() >= pop.Get(int(j)).Resilience()
}

// Swap swaps the elements with indexes i and j.
func (pop *Population) Swap(i, j int) {
	tmp := pop.Get(int(i))
	pop.individuals[i] = pop.Get(int(j))
	pop.individuals[j] = tmp
}

// Sort sort the population
func (pop *Population) Sort() {
	sort.Sort(pop)
}

// Cap returns the population capacity
func (pop *Population) Cap() int {
	return int(cap(pop.individuals))
}

// SetCap set the resize the population capacity
func (pop *Population) SetCap(newCap int) {
	currentCap := pop.Cap()
	if newCap != currentCap {
		tmp := pop.individuals
		switch {
		case newCap < currentCap:
			tmp = pop.individuals[0:newCap]
			pop.individuals = make([]individual.Interface, newCap, newCap)
		case newCap > currentCap:
			pop.individuals = make([]individual.Interface, currentCap, newCap)
		}
		copy(tmp, pop.individuals)
	}
}

// Truncate rduce population size to the given length
func (pop *Population) Truncate(length int) {
	if length < int(pop.Len()) {
		pop.individuals = pop.individuals[0 : length-1]
	}
}

// Append adds an individual to a population. If the populagtion has already reached its capacity, capacity is incremented.
func (pop *Population) Append(indiv individual.Interface) {
	pop.individuals = append(pop.individuals, indiv)
}

// Get returns the individual at index i
func (pop *Population) Get(i int) individual.Interface {
	return pop.individuals[i]
}

// Remove removes and returns the individual at index i
func (pop *Population) Remove(i int) individual.Interface {
	removed := pop.Get(i)
	new := pop.individuals[0 : i-1]
	pop.individuals = append(new, pop.individuals[i+1:pop.Len()-1]...)
	return removed
}

// Min returns the least Resilent individual
func (pop *Population) Min() individual.Interface {
	var min, length = pop.Get(0), int(pop.Len())
	for i := int(1); i < length; i++ {
		if pop.Get(i).Resilience() < min.Resilience() {
			min = pop.Get(i)
		}
	}
	return min
}

// Max returns the most Resilent individual
func (pop *Population) Max() individual.Interface {
	var max, length = pop.Get(0), int(pop.Len())
	for i := int(1); i < length; i++ {
		if pop.Get(i).Resilience() > max.Resilience() {
			max = pop.Get(i)
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
