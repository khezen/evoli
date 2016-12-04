package individual

import (
	"fmt"
)

// Interface is an interface for an individual
type Interface interface {
	Resilience() float32
	SetResilience(float32)
}

// Individual https://en.wikipedia.org/wiki/individual
type Individual float32

// New is the constructor for individuals
func New(resilience float32) Individual {
	return Individual(resilience)
}

// Resilience returns the strength of a individual regarding to its environement. Higher is stronger.
func (indiv *Individual) Resilience() float32 {
	return float32(*indiv)
}

// SetResilience set the strength of a individual regarding to its environement. Higher is stronger.
func (indiv *Individual) SetResilience(resilience float32) error {
	if resilience < 0 {
		return fmt.Errorf("SetResilience(%float32): argument must be > 0", resilience)
	}
	*indiv = Individual(resilience)
	return nil
}
