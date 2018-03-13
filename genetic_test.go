package evoli

import (
	"testing"
)

// To be completed
func TestNewGenetic(t *testing.T) {
	cases := []struct {
		s            Selecter
		survivorSize int
		c            Crosser
		m            Mutater
		mutaionProb  float64
		e            Evaluater
	}{
		{NewTruncationSelecter(), 10, crosserMock{}, mutaterMock{}, 0.01, evaluaterMock{}},
	}
	for _, c := range cases {
		_ = NewGenetic(NewPopulation(1), c.s, c.survivorSize, c.c, c.m, c.mutaionProb, c.e)
		_ = NewGeneticSync(NewPopulation(1), c.s, c.survivorSize, c.c, c.m, c.mutaionProb, c.e)
	}
}

// To be completed
func TestGeneticNext(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := NewIndividual(1), NewIndividual(-2), NewIndividual(3), NewIndividual(4), NewIndividual(5), NewIndividual(6)
	pop := population{i1, i2, i3, i4, i5, i6}
	cpy := NewPopulation(pop.Cap())
	cpy.Add(pop...)
	cases := []struct {
		genetic Evolution
	}{
		{NewGenetic(&pop, NewTruncationSelecter(), 5, crosserMock{}, mutaterMock{}, 1, evaluaterMock{})},
		{NewGeneticSync(&pop, NewTruncationSelecter(), 5, crosserMock{}, mutaterMock{}, 1, evaluaterMock{})},
	}
	for _, c := range cases {
		c.genetic.Next()
	}
}
