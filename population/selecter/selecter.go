package selecter

import "github.com/khezen/darwin/population"

// Interface is the selecter operator interface
type Interface interface {
	Select(pop *population.Population, survivorsSize int) (*population.Population, error)
}
