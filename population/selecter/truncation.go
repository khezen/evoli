package selecter

import "github.com/khezen/darwin/population"

type truncationSelecter struct{}

func (s truncationSelecter) Select(pop *population.Population, survivorsSize int) (*population.Population, error) {
	pop.Sort()
	pop.Truncate(survivorsSize - 1)
	return pop, nil
}

// NewTruncationSelecter is the constrctor for truncation selecter
func NewTruncationSelecter() Interface {
	return truncationSelecter{}
}
