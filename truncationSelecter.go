package darwin

type truncationSelecter struct{}

func (s truncationSelecter) Select(pop Population, survivorsSize int) (Population, error) {
	err := checkParams(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	newPop := NewPopulation(pop.Cap())
	newPop.Add(pop.Slice()...)
	newPop.Sort()
	newPop.Truncate(survivorsSize)
	return newPop, nil
}

// NewTruncationSelecter is the constructor for truncation selecter
func NewTruncationSelecter() Selecter {
	return truncationSelecter{}
}
