package evoli

import (
	"testing"
)

func TestArbitrerTruncation(t *testing.T) {
	testArbitrer(t, NewTruncationArbitrer())
}

func TestArbitrerTournament(t *testing.T) {
	testArbitrer(t, NewTournamentArbitrer())
}

func TestArbitrerRandom(t *testing.T) {
	testArbitrer(t, NewRandomArbitrer())
}

func TestArbitrerProportionalToRank(t *testing.T) {
	testArbitrer(t, NewProportionalToRankArbitrer())
}

func TestArbitrerProportionalToFitness(t *testing.T) {
	testArbitrer(t, NewProportionalToFitnessArbitrer())
}

func testArbitrer(t *testing.T, a Arbitrer) {
	i1, i2, i3, i4, i5, i6 := NewIndividual(1), NewIndividual(2), NewIndividual(-3), NewIndividual(4), NewIndividual(5), NewIndividual(-6)
	cases := []struct {
		in []Individual
	}{
		{[]Individual{i1, i2, i3, i4, i5, i6}},
		{[]Individual{i1, i1, i1, i1, i1, i1, i1, i1, i1, i1, i1, i1, i1}},
		{[]Individual{i1, i2, i4, i5}},
		{[]Individual{i1}},
		{[]Individual{i3, i6}},
	}
	for _, c := range cases {
		indiv := a.Abritrate(c.in...)
		ok := false
		for _, inIndiv := range c.in {
			if indiv == inIndiv {
				ok = true
				break
			}
		}
		if !ok {
			t.Error("oups")
		}
	}
}
