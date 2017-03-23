package darwin

type truncationSelecter struct{}

func (s truncationSelecter) Select(pop Population, survivorsSize int) (Population, error) {
	err := checkParams(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	pop.Sort()
	pop.Truncate(survivorsSize)
	return pop, nil
}

// NewTruncationSelecter is the constructor for truncation selecter
func NewTruncationSelecter() ISelecter {
	return truncationSelecter{}
}
