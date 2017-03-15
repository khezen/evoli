package selecter

import "github.com/khezen/darwin/population"

type truncationSelecter struct{}

func (s truncationSelecter) Select(pop population.Interface, survivorsSize int) (population.Interface, error) {
	err := checkParams(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	pop.Sort()
	pop.Truncate(survivorsSize)
	return pop, nil
}

// NewTruncationSelecter is the constructor for truncation selecter
func NewTruncationSelecter() Interface {
	return truncationSelecter{}
}
