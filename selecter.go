package darwin

import "fmt"

// Selecter is the selecter operator interface
type Selecter interface {
	Select(pop Population, survivorsSize int) (Population, error)
}

func checkParams(pop Population, survivorsSize int) error {
	if survivorsSize < 1 {
		return fmt.Errorf("%d must be >= 1", survivorsSize)
	}
	if survivorsSize > pop.Cap() {
		pop.SetCap(survivorsSize)
	}
	return nil
}
