package individual

// Interface is an interface for an individual
type Interface interface {
	Resilience() float32
	SetResilience(float32)
}

// Individual https://en.wikipedia.org/wiki/individual
type Individual struct {
	resilience float32
}

// New is the constructor for individuals
func New() Individual {
	return Individual{0}
}

// Resilience returns the strength of a individual regarding to its environement. Higher is stronger.
func (g *Individual) Resilience() float32 {
	return g.resilience
}

// SetResilience set the strength of a individual regarding to its environement. Higher is stronger.
func (g *Individual) SetResilience(resilience float32) {
	g.resilience = resilience
}
