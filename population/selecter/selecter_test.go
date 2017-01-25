package selecter

import (
	"testing"

	"github.com/khezen/darwin/population"
	"github.com/khezen/darwin/population/individual"
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

func TestProportionalToResilience(t *testing.T) {
	testSelecter(t, NewProportionalToResilienceSelecter())
}

func testSelecter(t *testing.T, s Interface) {
	i1, i2, i3, i4, i5, i6 := individual.New(1), individual.New(2), individual.New(-3), individual.New(4), individual.New(5), individual.New(-6)
	cases := []struct {
		in           population.Population
		survivalSize int
		expectedLen  int
		expectedCap  int
	}{
		{population.Population{i1, i2, i3, i4, i5, i6}, 3, 3, 6},
		{population.Population{i1, i2, i4, i5}, 2, 2, 4},
		{population.Population{i1}, 3, 1, 3},
		{population.Population{i3, i6}, 1, 1, 2},
		{population.Population{}, 3, 0, 3},
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
	pop := population.Population{i1, i2, i3}
	_, err := s.Select(&pop, -1)
	if err == nil {
		t.Errorf("expected != nil")
	}
	_, err = s.Select(nil, 2)
	if err == nil {
		t.Errorf("expected != nil")
	}
	_, err = s.Select(nil, -1)
	if err == nil {
		t.Errorf("expected != nil")
	}
}
