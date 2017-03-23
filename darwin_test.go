package darwin

import (
	"testing"
)

type crosserMock struct {
}

func (c crosserMock) Cross(individual1, individual2 IIndividual) IIndividual {
	return NewIndividual(individual1.Fitness() + individual2.Fitness()/2)
}

type evaluaterMock struct {
}

func (e evaluaterMock) Evaluate(individual IIndividual) (Fitness float32) {
	return individual.Fitness()
}

type mutaterMock struct {
}

func (m mutaterMock) Mutate(individual IIndividual) IIndividual {
	return individual
}

func TestNew(t *testing.T) {
	errorCases := []struct {
		s ISelecter
		c ICrosser
		m IMutater
		e IEvaluater
	}{
		{nil, crosserMock{}, mutaterMock{}, evaluaterMock{}},
		{NewTruncationSelecter(), nil, mutaterMock{}, evaluaterMock{}},
		{NewTruncationSelecter(), crosserMock{}, nil, evaluaterMock{}},
		{NewTruncationSelecter(), crosserMock{}, mutaterMock{}, nil},
	}
	for _, c := range errorCases {
		_, err := New(c.s, c.c, c.m, c.e)
		if err == nil {
			t.Errorf("expected != nil")
		}
	}
	_, err := New(NewTruncationSelecter(), crosserMock{}, mutaterMock{}, evaluaterMock{})
	if err != nil {
		t.Errorf("expected == nil")
	}
}

func TestGeneration(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := NewIndividual(1), NewIndividual(-2), NewIndividual(3), NewIndividual(4), NewIndividual(5), NewIndividual(6)
	pop := population{i1, i2, i3, i4, i5, i6}
	cpy := NewPopulation(pop.Cap())
	cpy.Append(pop...)
	lifecycle, _ := New(NewTruncationSelecter(), crosserMock{}, mutaterMock{}, evaluaterMock{})
	newPop, _ := lifecycle.Generation(&pop, 5, 1)
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
		mutationProb float32
	}{
		{nil, 2, 0.2},
		{&population{i1, i2, i3}, -10, 0.2},
		{&population{i1, i2, i3}, 2, 1.2},
		{&population{i1, i2, i3}, 2, -0.2},
	}
	for _, c := range errorCases {
		_, err := lifecycle.Generation(c.pop, c.survivorSize, c.mutationProb)
		if err == nil {
			t.Errorf("expected != nil")
		}
	}
}
