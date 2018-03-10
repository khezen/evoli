package evoli

type crosserMock struct {
}

func (c crosserMock) Cross(individual1, individual2 Individual) (Individual, error) {
	return NewIndividual((individual1.Fitness() + individual2.Fitness()) / 2), nil
}

type evaluaterMock struct {
}

func (e evaluaterMock) Evaluate(individual Individual) (Fitness float64, err error) {
	return individual.Fitness(), nil
}

type mutaterMock struct {
}

func (m mutaterMock) Mutate(individual Individual) (Individual, error) {
	return individual, nil
}

type positionerMock struct {
}

func (p positionerMock) Position(indiv, pBest, gBest Individual, c1, c2 float64) (Individual, error) {
	return NewIndividual((indiv.Fitness() + pBest.Fitness() + gBest.Fitness()) / 3), nil
}
