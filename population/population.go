package population

import (
	"math/rand"
	"sort"

	"fmt"

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
	PickCouple() (index1 int, indiv1 individual.Interface, index2 int, indiv2 individual.Interface)
	Contains(individual.Interface) bool
	IndexOf(individual.Interface) (int, error)
}

// Population is a set of individuals in population genetics.
type Population []individual.Interface

// New is population constructor
func New(capacity int) (*Population, error) {
	err := checkCap(capacity)
	if err != nil {
		return nil, err
	}
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
	indivi, erri := pop.Get(i)
	indivj, errj := pop.Get(j)
	if erri != nil || errj != nil {
		return false
	}
	return indivi.Resilience() >= indivj.Resilience()
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
	err := checkCap(newCap)
	if err != nil {
		return err
	}
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

// Truncate reduce population size to the given length
func (pop *Population) Truncate(length int) error {
	err := checkLength(length)
	if err != nil {
		return err
	}
	switch {
	case length == 0:
		newPop, _ := New(0)
		*pop = *newPop
	case length < pop.Len():
		*pop = (*pop)[0:length]
	case length > pop.Cap():
		newPop, _ := New(length)
		newPop.AppendAll(pop)
		*pop = *newPop
	}
	return nil
}

// Append adds an individual to a population. If the populagtion has already reached its capacity, capacity is incremented.
func (pop *Population) Append(indiv individual.Interface) error {
	err := checkIndivNotNil(indiv)
	if err != nil {
		return err
	}
	*pop = append(*pop, indiv)
	return nil
}

// AppendAll adds all individuals from a population to a population. If the populagtion has already reached its capacity, capacity is incremented.
func (pop *Population) AppendAll(externalPop *Population) error {
	err := CheckPopNotNil(externalPop)
	if err != nil {
		return err
	}
	for i := 0; i < externalPop.Len(); i++ {
		indiv, _ := externalPop.Get(i)
		err = checkIndivNotNil(indiv)
		if err != nil {
			return err
		}
	}
	*pop = append(*pop, *externalPop...)
	return nil
}

// Get returns the individual at index i
func (pop *Population) Get(i int) (individual.Interface, error) {
	err := checkIndex(i, pop.Len())
	if err != nil {
		return nil, err
	}
	return (*pop)[i], nil
}

// Remove removes and returns the individual at index i
func (pop *Population) Remove(i int) (individual.Interface, error) {
	err := checkIndex(i, pop.Len())
	if err != nil {
		return nil, err
	}
	removed, _ := pop.Get(i)
	new := (*pop)[0:i]
	*pop = append(new, (*pop)[i+1:pop.Len()]...)
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
func (pop *Population) PickCouple() (int, individual.Interface, int, individual.Interface, error) {
	if pop.Len() < 2 {
		return -1, nil, -1, nil, fmt.Errorf("pop must contains at least 2 individuals to pick a couple")
	}
	var i, j = rand.Intn(pop.Len() - 1), rand.Intn(pop.Len() - 1)
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
func (pop *Population) Contains(indiv individual.Interface) bool {
	_, err := pop.IndexOf(indiv)
	switch {
	case err != nil:
		return false
	default:
		return true
	}
}

// IndexOf returns the inde of the specified individual if it exists
func (pop *Population) IndexOf(indiv individual.Interface) (int, error) {
	err := checkIndivNotNil(indiv)
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

// CheckSuperior check that an int is positve
func CheckSuperior(value int, threshold int, label string) error {
	switch {
	case value < threshold:
		return fmt.Errorf("%v must be >= %v", label, threshold)
	default:
		return nil
	}
}

func checkCap(cap int) error {
	return CheckSuperior(cap, 0, "capcity")
}

func checkLength(length int) error {
	return CheckSuperior(length, 0, "length")
}

func checkIndex(index, length int) error {
	err := CheckSuperior(index, 0, "index")
	switch {
	case err != nil:
		return err
	default:
		if index >= length {
			return fmt.Errorf("index must be < length")
		}
		return nil
	}
}

//CheckPopNotNil checks that population is not nil
func CheckPopNotNil(pop *Population) error {
	if (*Population)(pop) == (*Population)(nil) {
		return fmt.Errorf("Nil pointer on population")
	}
	return nil
}

func checkIndivNotNil(indiv individual.Interface) error {
	if indiv == nil {
		return fmt.Errorf("Nil pointer on individual")
	}
	return nil
}
