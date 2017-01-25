package darwin

import (
	"testing"

	"github.com/khezen/darwin/population"
	"github.com/khezen/darwin/population/individual"
	"github.com/khezen/darwin/population/selecter"
)

type crosserMock struct {
}

func (c crosserMock) Cross(individual1, individual2 individual.Interface) individual.Interface {
	return individual.New(individual1.Resilience() + individual2.Resilience()/2)
}

type evaluaterMock struct {
}

func (e evaluaterMock) Evaluate(individual individual.Interface) (resilience float32) {
	return individual.Resilience()
}

type mutaterMock struct {
}

func (m mutaterMock) Mutate(individual individual.Interface) individual.Interface {
	return individual
}

func TestNew(t *testing.T) {
	errorCases := []struct {
		s selecter.Interface
		c individual.Crosser
		m individual.Mutater
		e individual.Evaluater
	}{
		{nil, crosserMock{}, mutaterMock{}, evaluaterMock{}},
		{selecter.NewTruncationSelecter(), nil, mutaterMock{}, evaluaterMock{}},
		{selecter.NewTruncationSelecter(), crosserMock{}, nil, evaluaterMock{}},
		{selecter.NewTruncationSelecter(), crosserMock{}, mutaterMock{}, nil},
	}
	for _, c := range errorCases {
		_, err := New(c.s, c.c, c.m, c.e)
		if err == nil {
			t.Errorf("expected != nil")
		}
	}
	_, err := New(selecter.NewTruncationSelecter(), crosserMock{}, mutaterMock{}, evaluaterMock{})
	if err != nil {
		t.Errorf("expected == nil")
	}
}

func TestGeneration(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := individual.New(1), individual.New(-2), individual.New(3), individual.New(4), individual.New(5), individual.New(6)
	pop := population.Population{i1, i2, i3, i4, i5, i6}
	cpy, _ := population.New(pop.Cap())
	cpy.AppendAll(&pop)
	lifecycle, _ := New(selecter.NewTruncationSelecter(), crosserMock{}, mutaterMock{}, evaluaterMock{})
	newPop, _ := lifecycle.Generation(&pop, 5, 0.5)
	isNewPopDifferent := false
	for i := 0; i < newPop.Len(); i++ {
		indiv, _ := newPop.Get(i)
		if !cpy.Contains(indiv) {
			isNewPopDifferent = true
			break
		}
	}
	if !isNewPopDifferent {
		t.Errorf("the new Generation should be different from the previous one")
	}
	errorCases := []struct {
		pop          population.Interface
		survivorSize int
		mutationProb float32
	}{
		{nil, 2, 0.2},
		{&population.Population{i1, i2, i3}, -10, 0.2},
		{&population.Population{i1, i2, i3}, 2, 1.2},
		{&population.Population{i1, i2, i3}, 2, -0.2},
	}
	for _, c := range errorCases {
		_, err := lifecycle.Generation(c.pop, c.survivorSize, c.mutationProb)
		if err == nil {
			t.Errorf("expected != nil")
		}
	}
}
