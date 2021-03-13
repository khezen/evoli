package evoli

import (
	"testing"
)

func TestSelecterTruncation(t *testing.T) {
	testSelecter(t, NewTruncationSelecter())
}

func TestSelecterTournament(t *testing.T) {
	testSelecter(t, NewTournamentSelecter(1))
}

func TestSelecterRandom(t *testing.T) {
	testSelecter(t, NewRandomSelecter())
}

func TestSelecterProportionalToRank(t *testing.T) {
	testSelecter(t, NewProportionalToRankSelecter())
}

func TestSelecterProportionalToFitness(t *testing.T) {
	testSelecter(t, NewProportionalToFitnessSelecter())
}

func testSelecter(t *testing.T, s Selecter) {
	i1, i2, i3, i4, i5, i6 := NewIndividual(1), NewIndividual(2), NewIndividual(-3), NewIndividual(4), NewIndividual(5), NewIndividual(-6)
	cases := []struct {
		in           population
		survivalSize int
		expectedLen  int
		expectedCap  int
	}{
		{population{i1, i2, i3, i4, i5, i6}, 3, 3, 6},
		{population{i1, i1, i1, i1, i1, i1, i1, i1, i1, i1, i1, i1, i1}, 1, 1, 13},
		{population{i1, i2, i4, i5}, 2, 2, 4},
		{population{i1}, 3, 1, 1},
		{population{i3, i6}, 1, 1, 2},
	}
	for _, c := range cases {
		newPop, _, _ := s.Select(&c.in, c.survivalSize)
		length, capacity := newPop.Len(), newPop.Cap()
		if length != c.expectedLen {
			t.Errorf("s.Select(%v, %v) returned %v which has a length of %v instead of %v", c.in, c.survivalSize, newPop, length, c.expectedLen)
		}
		if capacity != c.expectedCap {
			t.Errorf("s.Select(%v, %v) returned %v which has a capacity of %v instead of %v", c.in, c.survivalSize, newPop, capacity, c.expectedCap)
		}
	}
}
