package evoli

// ErrArbitration - must provide one or more individuals
var ErrArbitration = "ErrArbitration - must provide one or more individuals"

// Arbitrer - elect the winner between multiple participants
type Arbitrer interface {
	Abritrate(participants ...Individual) (winner Individual)
}

func checkArbitrersParams(participants []Individual) {
	if len(participants) < 1 {
		panic(ErrArbitration)
	}
}

type selecterBasedArbitrer struct {
	Selecter
}

func (a selecterBasedArbitrer) Abritrate(participants ...Individual) (winner Individual) {
	checkArbitrersParams(participants)
	pop := NewPopulation(len(participants))
	pop.Add(participants...)
	selected, _ := a.Select(pop, 1)
	return selected.Get(0)
}

// NewProportionalToFitnessArbitrer -  based on fitness value
func NewProportionalToFitnessArbitrer() Arbitrer {
	return selecterBasedArbitrer{NewProportionalToFitnessSelecter()}
}

// NewProportionalToRankArbitrer - based on rank
func NewProportionalToRankArbitrer() Arbitrer {
	return selecterBasedArbitrer{NewProportionalToRankSelecter()}
}

// NewTournamentArbitrer -  High Fitness increase chances to come out vcitorious from a duel
func NewTournamentArbitrer() Arbitrer {
	return selecterBasedArbitrer{NewTournamentSelecter()}
}

// NewTruncationArbitrer - take the highest fitness
func NewTruncationArbitrer() Arbitrer {
	return selecterBasedArbitrer{NewTruncationSelecter()}
}

// NewRandomArbitrer - choose randomly
func NewRandomArbitrer() Arbitrer {
	return selecterBasedArbitrer{NewRandomSelecter()}
}
