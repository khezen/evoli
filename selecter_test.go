package evoli

import (
	"testing"
)

func TestTruncation(t *testing.T) {
	testSelecter(t, NewTruncationSelecter())
}

func TestTournament(t *testing.T) {
	testSelecter(t, NewTournamentSelecter())
}

func TestRandom(t *testing.T) {
	testSelecter(t, NewRandomSelecter())
}

func TestProportionalToRank(t *testing.T) {
	testSelecter(t, NewProportionalToRankSelecter())
}

func TestProportionalToFitness(t *testing.T) {
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
		{population{i1}, 3, 1, 3},
		{population{i3, i6}, 1, 1, 2},
		{population{}, 3, 0, 3},
	}
	for _, c := range cases {
		newPop, _ := s.Select(&c.in, c.survivalSize)
		length, capacity := newPop.Len(), newPop.Cap()
		if length != c.expectedLen {
			t.Errorf("s.Select(%v, %v) returned %v which has a length of %v instead of %v", c.in, c.survivalSize, newPop, length, c.expectedLen)
		}
		if capacity != c.expectedCap {
			t.Errorf("s.Select(%v, %v) returned %v which has a capacity of %v instead of %v", c.in, c.survivalSize, newPop, capacity, c.expectedCap)
		}
	}
	pop := population{i1, i2, i3}
	_, err := s.Select(&pop, -1)
	if err == nil {
		t.Errorf("expected != nil")
	}
}
