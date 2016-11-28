package selecter

import "github.com/khezen/darwin/environment/population"

type truncationSelecter struct{}

func (s truncationSelecter) Select(pop *population.Population, survivorsSize int) (*population.Population, error) {
	err := checkArgs(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	pop.Sort()
	pop.Truncate(survivorsSize - 1)
	return pop, nil
}

// NewTruncationSelecter is the constrctor for truncation selecter
func NewTruncationSelecter() Interface {
	return truncationSelecter{}
}
