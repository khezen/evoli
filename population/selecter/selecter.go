package selecter

import "github.com/khezen/darwin/population"

// Interface is the selecter operator interface
type Interface interface {
	Select(pop population.Interface, survivorsSize int) (population.Interface, error)
}

func checkParams(pop population.Interface, survivorsSize int) error {
	err := population.CheckPopNotNil(pop)
	if err != nil {
		return err
	}
	err = population.CheckSuperior(survivorsSize, 2, "survivorSize")
	if err != nil {
		return err
	}
	if survivorsSize > pop.Cap() {
		pop.SetCap(survivorsSize)
	}
	return nil
}
