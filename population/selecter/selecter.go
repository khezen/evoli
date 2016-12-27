package selecter

import "github.com/khezen/darwin/population"

// Interface is the selecter operator interface
type Interface interface {
	Select(pop *population.Population, survivorsSize int) (*population.Population, error)
}

func checkParams(pop *population.Population, survivorsSize int) error {
	err := population.CheckPopNotNil(pop)
	if err != nil {
		return err
	}
	err = population.CheckPositive(survivorsSize, "survivorSize must be >= 0")
	if err != nil {
		return err
	}
	if survivorsSize > pop.Cap() {
		pop.SetCap(survivorsSize)
	}
	return nil
}
