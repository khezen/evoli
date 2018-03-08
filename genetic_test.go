package evoli

import (
	"testing"
)

type crosserMock struct {
}

func (c crosserMock) Cross(individual1, individual2 Individual) (Individual, error) {
	return NewIndividual(individual1.Fitness() + individual2.Fitness()/2), nil
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

func TestNew(t *testing.T) {
	errorCases := []struct {
		s            Selecter
		survivorSize int
		c            Crosser
		m            Mutater
		mutaionProb  float64
		e            Evaluater
	}{
		{NewTruncationSelecter(), 10, crosserMock{}, mutaterMock{}, 0.01, evaluaterMock{}},
	}
	for _, c := range errorCases {
		_ = NewGenetic(c.s, c.survivorSize, c.c, c.m, c.mutaionProb, c.e)
	}
}

func TestGeneration(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := NewIndividual(1), NewIndividual(-2), NewIndividual(3), NewIndividual(4), NewIndividual(5), NewIndividual(6)
	pop := population{i1, i2, i3, i4, i5, i6}
	cpy := NewPopulation(pop.Cap())
	cpy.Add(pop...)
	cases := []struct {
		genetic Genetic
	}{
		{NewGenetic(NewTruncationSelecter(), 5, crosserMock{}, mutaterMock{}, 1, evaluaterMock{})},
	}
	for _, c := range cases {
		newPop, _ := c.genetic.Next(&pop)
		isNewPopDifferent := false
		for i := 0; i < newPop.Len(); i++ {
			indiv := newPop.Get(i)
			if !cpy.Has(indiv) {
				isNewPopDifferent = true
				break
			}
		}
		if !isNewPopDifferent {
			t.Errorf("the new Generation should be different from the previous one")
		}
	}
}
