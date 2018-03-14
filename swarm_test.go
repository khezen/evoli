package evoli

import "testing"

// To be completed
func TestNewSwarm(t *testing.T) {
	errorCases := []struct {
		p      Positioner
		c1, c2 float64
		e      Evaluater
	}{
		{positionerMock{}, 2, 2, evaluaterMock{}},
	}
	for _, c := range errorCases {
		i1 := NewIndividual(1)
		i2 := NewIndividual(2)
		popInit := NewPopulation(2)
		newPop := NewPopulation(1)
		popInit.Add(i1, i2)
		sw := NewSwarm(popInit, c.p, c.c1, c.c2, c.e)
		pop := sw.Population()
		if pop != popInit {
			t.Errorf("expected %v got %v", popInit, pop)
		}
		alpha := sw.Alpha()
		if alpha != i2 {
			t.Errorf("expected %v got %v", i2, alpha)
		}
		sw.SetPopulation(newPop)
		pop = sw.Population()
		if pop != newPop {
			t.Errorf("expected %v got %v", newPop, pop)
		}
		sw = NewSwarmSync(popInit, c.p, c.c1, c.c2, c.e)
		pop = sw.Population()
		if pop != popInit {
			t.Errorf("expected %v got %v", popInit, pop)
		}
		alpha = sw.Alpha()
		if alpha != i2 {
			t.Errorf("expected %v got %v", i2, alpha)
		}
		sw.SetPopulation(newPop)
		pop = sw.Population()
		if pop != newPop {
			t.Errorf("expected %v got %v", newPop, pop)
		}
	}
}

// To be completed
func TestSwarmNext(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := NewIndividual(1), NewIndividual(-2), NewIndividual(3), NewIndividual(4), NewIndividual(5), NewIndividual(6)
	pop := population{i1, i2, i3, i4, i5, i6}
	cpy := NewPopulation(pop.Cap())
	cpy.Add(pop...)
	cases := []struct {
		swarm Evolution
	}{
		{NewSwarm(&pop, positionerMock{}, 2, 2, evaluaterMock{})},
		{NewSwarmSync(&pop, positionerMock{}, 2, 2, evaluaterMock{})},
	}
	for _, c := range cases {
		_ = c.swarm.Next()
	}
}
