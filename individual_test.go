package darwin

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
