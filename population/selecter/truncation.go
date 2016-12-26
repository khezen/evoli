package selecter

import "github.com/khezen/darwin/population"

type truncationSelecter struct{}

func (s truncationSelecter) Select(pop *population.Population, survivorsSize int) (*population.Population, error) {
	err := checkParams(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	pop.Sort()
	pop.Truncate(survivorsSize)
	return pop, nil
}

// NewTruncationSelecter is the constrctor for truncation selecter
func NewTruncationSelecter() Interface {
	return truncationSelecter{}
}
