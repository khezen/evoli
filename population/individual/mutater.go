package individual

// Mutater randomly modify a individual. This operator maintain diversity in a population.
type Mutater interface {
	Mutate(Interface) Interface
}
