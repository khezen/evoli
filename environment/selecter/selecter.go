package selecter

import (
	"errors"

	"github.com/khezen/darwin/environment/population"
)

// Interface is the selecter operator interface
type Interface interface {
	Select(pop *population.Population, survivorsCount uint) (*population.Population, error)
}

func checkArgs(pop *population.Population, survivorsCount uint) error {
	switch {
	case pop == nil:
		return errors.New("pop cannot be nil")
	case survivorsCount > uint(pop.Len()):
		return errors.New("survivalSize cannot be >= pop.Size()")
	}
}
