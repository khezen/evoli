package individual

// Phenotype https://en.wikipedia.org/wiki/Phenotype
type Phenotype struct {
	resilience float32
}

// Resilience returns the strength of a Phenotype regarding to its environement. Higher is stronger.
func (g *Phenotype) Resilience() float32 {
	return g.resilience
}

// SetResilience set the strength of a Phenotype regarding to its environement. Higher is stronger.
func (g *Phenotype) SetResilience(resilience float32) {
	g.resilience = resilience
}
