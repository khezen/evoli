package darwin

import (
	"github.com/khezen/check"
)

// Selecter is the selecter operator interface
type Selecter interface {
	Select(pop Population, survivorsSize int) (Population, error)
}

func checkParams(pop Population, survivorsSize int) error {
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
