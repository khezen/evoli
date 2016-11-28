package selecter

import (
	"errors"

	"github.com/khezen/darwin/environment/population"
)

// Interface is the selecter operator interface
type Interface interface {
	Select(pop *population.Population, survivorsSize int) (*population.Population, error)
}

func checkArgs(pop *population.Population, survivorsSize int) error {
	switch {
	case pop == nil:
		return errors.New("pop cannot be nil")
	case pop.Len() <= 0:
		return errors.New("pop must contains at least one individual")
	case survivorsSize > pop.Len():
		return errors.New("survivalSize cannot be >= pop.Size()")
	case survivorsSize < 0:
		return errors.New("survivalSize cannot be < 0")
	default:
		return nil
	}
}
