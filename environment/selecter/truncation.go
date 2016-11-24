package selecter

import "github.com/khezen/darwin/environment/population"

type truncationSelecter struct{}

func (s truncationSelecter) Select(pop *population.Population, survivorsCount uint) (*population.Population, error) {
	err := checkArgs(pop, survivorsCount)
	if err != nil {
		return nil, err
	}
	pop.Sort()
	pop.Truncate(survivorsCount - 1)
	return pop, nil
}

// NewTruncationSelecter is the constrctor for truncation selecter
func NewTruncationSelecter() Interface {
	return truncationSelecter{}
}
