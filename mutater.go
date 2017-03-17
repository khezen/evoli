package darwin

// IMutater randomly modify a individual. This operator maintain diversity in a population.
type IMutater interface {
	Mutate(IIndividual) IIndividual
}
