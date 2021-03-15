package evoli

import "math/rand"

type crosserMock struct {
}

func (c crosserMock) Cross(parent1, parent2 Individual) (child1, child2 Individual, err error) {
	w := 0.1 + 0.8*rand.Float64()
	return NewIndividual(w*parent1.Fitness() + (1-w)*parent2.Fitness()),
		NewIndividual((1-w)*parent1.Fitness() + w*parent2.Fitness()),
		nil
}

type evaluaterMock struct {
}

func (e evaluaterMock) Evaluate(individual Individual) (Fitness float64, err error) {
	return individual.Fitness(), nil
}

type mutaterMock struct {
}

func (m mutaterMock) Mutate(individual Individual, p float64) (Individual, error) {
	return individual, nil
}

type positionerMock struct {
}

func (p positionerMock) Position(indiv, pBest, gBest Individual, c1, c2 float64) (Individual, error) {
	return NewIndividual((indiv.Fitness() + pBest.Fitness() + gBest.Fitness()) / 3), nil
}
