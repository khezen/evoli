package evoli

import "testing"

func TestNewSwarm(t *testing.T) {
	errorCases := []struct {
		p      Positioner
		c1, c2 float64
		e      Evaluater
	}{
		{positionerMock{}, 2, 2, evaluaterMock{}},
	}
	for _, c := range errorCases {
		_ = NewSwarm(c.p, c.c1, c.c2, c.e)
	}
}

func TestSwarmNext(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := NewIndividual(1), NewIndividual(-2), NewIndividual(3), NewIndividual(4), NewIndividual(5), NewIndividual(6)
	pop := population{i1, i2, i3, i4, i5, i6}
	cpy := NewPopulation(pop.Cap())
	cpy.Add(pop...)
	cases := []struct {
		genetic Genetic
	}{
		{NewSwarm(positionerMock{}, 2, 2, evaluaterMock{})},
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
