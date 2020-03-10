package evoli

import "sync"

type populationSync struct {
	population
	sync.RWMutex
}

// NewPopulationSync creates a threadsafe population
func NewPopulationSync(capacity int) Population {
	pop := NewPopulation(capacity)
	return &populationSync{
		*pop.(*population),
		sync.RWMutex{},
	}
}

// Len returns the current livings count of a population
func (p *populationSync) Len() int {
	p.RLock()
	defer p.RUnlock()
	return p.population.Len()
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p *populationSync) Less(i, j int) bool {
	p.RLock()
	defer p.RUnlock()
	return p.population.Less(i, j)
}

// Swap swaps the elements with indexes i and j.
func (p *populationSync) Swap(i, j int) {
	p.Lock()
	defer p.Unlock()
	p.population.Swap(i, j)
}

// Sort sort the population
func (p *populationSync) Sort() {
	p.Lock()
	defer p.Unlock()
	p.population.Sort()
}

// SetCap set the resize the population capacity
func (p *populationSync) SetCap(newCap int) {
	p.Lock()
	defer p.Unlock()
	p.population.SetCap(newCap)
}

// Add adds an individual to a population. If the populagtion has already reached its capacity, capacity is incremented.
func (p *populationSync) Add(indiv ...Individual) {
	p.Lock()
	defer p.Unlock()
	p.population.Add(indiv...)
}

// Get returns the individual at index i
func (p *populationSync) Get(i int) Individual {
	p.RLock()
	defer p.RUnlock()
	return p.population.Get(i)
}

// RemoveAt removes and returns the individual at index i
func (p *populationSync) RemoveAt(i int) {
	p.Lock()
	defer p.Unlock()
	p.population.RemoveAt(i)
}

// Remove removes all given individuals
func (p *populationSync) Remove(individuals ...Individual) {
	p.Lock()
	defer p.Unlock()
	p.population.Remove(individuals...)
}

// Replace replaces and returns the individual at index i by the substitute
func (p *populationSync) Replace(i int, substitute Individual) {
	p.Lock()
	defer p.Unlock()
	p.population.Replace(i, substitute)
}

// Min returns the least Resilent individual
func (p *populationSync) Min() Individual {
	p.RLock()
	defer p.RUnlock()
	return p.population.Min()
}

// Max returns the most Resilent individual
func (p *populationSync) Max() Individual {
	p.RLock()
	defer p.RUnlock()
	return p.population.Max()
}

// Has return true if the specified individual is in the population
func (p *populationSync) Has(individuals ...Individual) bool {
	has := true
	for _, indiv := range individuals {
		_, err := p.IndexOf(indiv)
		has = has && err == nil
	}
	return has
}

// IndexOf returns the inde of the specified individual if it exists
func (p *populationSync) IndexOf(indiv Individual) (int, error) {
	p.RLock()
	defer p.RUnlock()
	return p.population.IndexOf(indiv)
}

// Each traverse the population and execute given callback on each individual. Stops if the callbak return false.
func (p *populationSync) Each(f func(indiv Individual) bool) {
	p.RLock()
	defer p.RUnlock()
	p.population.Each(f)
}

// Slice returns the population as []Individual
func (p *populationSync) Slice() []Individual {
	p.RLock()
	defer p.RUnlock()
	return p.population.Slice()
}

func (p *populationSync) New(cap int) Population {
	return NewPopulationSync(cap)
}

func (p *populationSync) Close() {
	p.Lock()
	defer p.Unlock()
	p.population.Close()
}
