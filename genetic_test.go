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
		i1 := NewIndividual(1)
		i2 := NewIndividual(2)
		popInit := NewPopulation(2)
		newPop := NewPopulation(1)
		popInit.Add(i1, i2)
		gen := NewGenetic(popInit, c.s, c.survivorSize, c.c, c.m, c.mutaionProb, c.e)
		pop := gen.Population()
		if pop != popInit {
			t.Errorf("expected %v got %v", popInit, pop)
		}
		alpha := gen.Alpha()
		if alpha != i2 {
			t.Errorf("expected %v got %v", i2, alpha)
		}
		gen.SetPopulation(newPop)
		pop = gen.Population()
		if pop != newPop {
			t.Errorf("expected %v got %v", newPop, pop)
		}
		gen = NewGeneticSync(popInit, c.s, c.survivorSize, c.c, c.m, c.mutaionProb, c.e)
		pop = gen.Population()
		if pop != popInit {
			t.Errorf("expected %v got %v", popInit, pop)
		}
		alpha = gen.Alpha()
		if alpha != i2 {
			t.Errorf("expected %v got %v", i2, alpha)
		}
		gen.SetPopulation(newPop)
		pop = gen.Population()
		if pop != newPop {
			t.Errorf("expected %v got %v", newPop, pop)
		}
	}
}

// To be completed
func TestGeneticNext(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := NewIndividual(1), NewIndividual(-2), NewIndividual(3), NewIndividual(4), NewIndividual(5), NewIndividual(6)
	pop1 := population{i1, i2, i3, i4, i5, i6}
	pop2 := population{i1, i2, i3, i4, i5, i6}
	cases := []struct {
		genetic Evolution
	}{
		{NewGenetic(&pop1, NewTruncationSelecter(), 5, crosserMock{}, mutaterMock{}, 1, evaluaterMock{})},
		{NewGeneticSync(&pop2, NewTruncationSelecter(), 5, crosserMock{}, mutaterMock{}, 1, evaluaterMock{})},
	}
	for _, c := range cases {
		c.genetic.Next()
	}
}
