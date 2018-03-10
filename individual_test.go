package evoli

import "testing"

func TestNewIndividual(t *testing.T) {
	cases := []struct {
		indiv    Individual
		expected float64
	}{
		{NewIndividual(0.7), 0.7},
		{NewIndividual(54.0), 54.0},
		{NewIndividual(0), 0},
		{NewIndividualTS(0.7), 0.7},
		{NewIndividualTS(54.0), 54.0},
		{NewIndividualTS(0), 0},
	}
	for _, c := range cases {
		got := c.indiv.Fitness()
		if got != c.expected {
			t.Errorf("expected %f got %f", c.expected, got)
		}
	}
}

func TestBest(t *testing.T) {
	i0, i10 := NewIndividual(0), NewIndividual(10)
	i0TS, i10TS := NewIndividualTS(0), NewIndividualTS(10)
	cases := []struct {
		indiv        Individual
		in, expected Individual
	}{
		{i0, i0, i0},
		{i0, i10, i10},
		{i10, i10, i10},
		{i10, i0, i0},
		{i0TS, i0TS, i0TS},
		{i0TS, i10TS, i10TS},
		{i10TS, i10TS, i10TS},
		{i10TS, i0TS, i0TS},
	}
	for _, c := range cases {
		c.indiv.SetBest(c.in)
		got := c.indiv.Best()
		if got != c.expected {
			t.Errorf("expected %v got %v", c.expected, got)
		}
	}
}

func TestFitness(t *testing.T) {
	cases := []struct {
		indiv        Individual
		in, expected float64
	}{
		{NewIndividual(0), 0.7, 0.7},
		{NewIndividual(0), 54.0, 54.0},
		{NewIndividual(10), 0, 0},
		{NewIndividualTS(0), 0.7, 0.7},
		{NewIndividualTS(0), 54.0, 54.0},
		{NewIndividualTS(10), 0, 0},
	}
	for _, c := range cases {
		c.indiv.SetFitness(c.in)
		got := c.indiv.Fitness()
		if got != c.expected {
			t.Errorf("indiv.SetFitness(%f); indiv.Fitness() == %f, expected %f", c.in, got, c.expected)
		}
	}
}
