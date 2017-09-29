package darwin

type truncationSelecter struct{}

func (s truncationSelecter) Select(pop Population, survivorsSize int) (Population, error) {
	err := checkParams(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	pop.Sort()
	individuals := pop.Slice()
	var survivors []Individual
	if len(individuals) < survivorsSize {
		survivors = individuals
	} else {
		survivors = individuals[:survivorsSize]
	}
	newPop := NewPopulation(pop.Cap())
	newPop.Add(survivors...)
	return newPop, nil
}

// NewTruncationSelecter is the constructor for truncation selecter
func NewTruncationSelecter() Selecter {
	return truncationSelecter{}
}
