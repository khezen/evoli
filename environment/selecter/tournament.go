package selecter

import (
	"errors"

	"github.com/khezen/darwin/environment/population"
)

type tournamentSelecter struct{}

func (s tournamentSelecter) Select(pop *population.Population, survivorsSize uint) (*population.Population, error) {
	err := checkArgs(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	return nil, errors.New("unsupported operation")
}

// NewTournamentSelecter is the constrctor for truncation selecter
func NewTournamentSelecter() Interface {
	return tournamentSelecter{}
}
