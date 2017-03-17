package darwin

import (
	"github.com/khezen/check"
)

// ISelecter is the selecter operator interface
type ISelecter interface {
	Select(pop IPopulation, survivorsSize int) (IPopulation, error)
}

func checkParams(pop IPopulation, survivorsSize int) error {
	err := check.NotNil(pop)
	if err != nil {
		return err
	}
	err = check.Superior(survivorsSize, 1, "survivorSize")
	if err != nil {
		return err
	}
	if survivorsSize > pop.Cap() {
		pop.SetCap(survivorsSize)
	}
	return nil
}
