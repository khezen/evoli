package darwin

import (
	"testing"
)

type crosserMock struct {
}

func (c crosserMock) Cross(individual1, individual2 Individual) Individual {
	return NewIndividual(individual1.Fitness() + individual2.Fitness()/2)
}

type evaluaterMock struct {
}

func (e evaluaterMock) Evaluate(individual Individual) (Fitness float64) {
	return individual.Fitness()
}

type mutaterMock struct {
}

func (m mutaterMock) Mutate(individual Individual) Individual {
	return individual
}

func TestNew(t *testing.T) {
	errorCases := []struct {
		s Selecter
		c Crosser
		m Mutater
		e Evaluater
	}{
		{NewTruncationSelecter(), crosserMock{}, mutaterMock{}, evaluaterMock{}},
	}
	for _, c := range errorCases {
		_ = New(c.s, c.c, c.m, c.e)
	}
}

func TestGeneration(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := NewIndividual(1), NewIndividual(-2), NewIndividual(3), NewIndividual(4), NewIndividual(5), NewIndividual(6)
	pop := population{i1, i2, i3, i4, i5, i6}
	cpy := NewPopulation(pop.Cap())
	cpy.Add(pop...)
	cases := []struct {
		lifecycle Lifecycle
	}{
		{New(NewTruncationSelecter(), crosserMock{}, mutaterMock{}, evaluaterMock{})},
	}
	for _, c := range cases {
		newPop, _ := c.lifecycle.Iterate(&pop, 5, 1)
		isNewPopDifferent := false
		for i := 0; i < newPop.Len(); i++ {
			indiv, _ := newPop.Get(i)
			if !cpy.Has(indiv) {
				isNewPopDifferent = true
				break
			}
		}
		if !isNewPopDifferent {
			t.Errorf("the new Generation should be different from the previous one")
		}
		errorCases := []struct {
			pop          Population
			survivorSize int
			mutationProb float64
		}{
			{&population{i1, i2, i3}, -10, 0.2},
			{&population{i1, i2, i3}, 2, 1.2},
			{&population{i1, i2, i3}, 2, -0.2},
		}
		for _, edgeCase := range errorCases {
			_, err := c.lifecycle.Iterate(edgeCase.pop, edgeCase.survivorSize, edgeCase.mutationProb)
			if err == nil {
				t.Errorf("expected != nil")
			}
		}
	}
}
