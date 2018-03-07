package darwin

// Mutater randomly modify a individual. This operator maintain diversity in a population.
type Mutater interface {
	Mutate(Individual) (Individual, error)
}
