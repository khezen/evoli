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
		_ = NewSwarm(NewPopulation(1), c.p, c.c1, c.c2, c.e)
		_ = NewSwarmTS(NewPopulation(1), c.p, c.c1, c.c2, c.e)
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
		{NewSwarmTS(&pop, positionerMock{}, 2, 2, evaluaterMock{})},
	}
	for _, c := range cases {
		_ = c.swarm.Next()
	}
}
