package evoli

// Mutater randomly modify a individual. This operator maintain diversity in a population.
type Mutater interface {
	Mutate(indiv Individual, p float64) (Individual, error)
}
