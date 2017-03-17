package darwin

import "testing"

func TestNewIndividual(t *testing.T) {
	cases := []struct {
		in, expected float32
	}{
		{0.7, 0.7},
		{54.0, 54.0},
		{0, 0},
	}
	for _, c := range cases {
		indiv := NewIndividual(c.in)
		got := indiv.Fitness()
		if got != c.expected {
			t.Errorf("NewIndividual(%f); indiv.Fitness() == %f, expected %f", c.in, got, c.expected)
		}
	}
}

func TestFitness(t *testing.T) {
	cases := []struct {
		in, expected float32
	}{
		{0.7, 0.7},
		{54.0, 54.0},
		{0, 0},
	}
	indiv := NewIndividual(0)
	for _, c := range cases {
		indiv.SetFitness(c.in)
		got := indiv.Fitness()
		if got != c.expected {
			t.Errorf("indiv.SetFitness(%f); indiv.Fitness() == %f, expected %f", c.in, got, c.expected)
		}
	}
}
