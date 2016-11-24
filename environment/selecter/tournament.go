package selecter

import (
	"errors"

	"github.com/khezen/darwin/environment/population"
)

type tournamentSelecter struct{}

func (s tournamentSelecter) Select(pop *population.Population, survivorsCount uint) (*population.Population, error) {
	err := checkArgs(pop, survivorsCount)
	if err != nil {
		return nil, err
	}
	return nil, errors.New("unsupported operation")
}

// NewTournamentSelecter is the constrctor for truncation selecter
func NewTournamentSelecter() Interface {
	return tournamentSelecter{}
}
