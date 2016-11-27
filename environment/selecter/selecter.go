package selecter

import (
	"errors"

	"github.com/khezen/darwin/environment/population"
)

// Interface is the selecter operator interface
type Interface interface {
	Select(pop *population.Population, survivorsSize uint) (*population.Population, error)
}

func checkArgs(pop *population.Population, survivorsSize uint) error {
	switch {
	case pop == nil:
		return errors.New("pop cannot be nil")
	case pop.Len() <= 0:
		return errors.New("pop must contains at least one individual")
	case survivorsSize > uint(pop.Len()):
		return errors.New("survivalSize cannot be >= pop.Size()")
	default:
		return nil
	}
}
