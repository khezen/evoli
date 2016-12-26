package individual

import "fmt"

// Interface is an interface for an individual
type Interface interface {
	Resilience() float32
	SetResilience(float32)
}

// Individual https://en.wikipedia.org/wiki/individual
type Individual float32

// New is the constructor for individuals
func New(resilience float32) *Individual {
	indiv := Individual(resilience)
	return &indiv
}

// Resilience returns the strength of a individual regarding to its environement. Higher is stronger.
func (indiv *Individual) Resilience() float32 {
	return float32(*indiv)
}

// SetResilience set the strength of a individual regarding to its environement. Higher is stronger.
func (indiv *Individual) SetResilience(resilience float32) {
	*indiv = Individual(resilience)
}

// CheckIndivNotNil checks that individual is not nil
func CheckIndivNotNil(indiv Interface) error {
	if indiv == nil {
		return fmt.Errorf("Nil pointer on individual")
	}
	return nil
}
