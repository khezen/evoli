package individual

import "testing"

func TestNew(t *testing.T) {
	cases := []struct {
		in, expected float32
	}{
		{0.7, 0.7},
		{54.0, 54.0},
		{0, 0},
	}
	for _, c := range cases {
		indiv := New(c.in)
		got := indiv.Resilience()
		if got != c.expected {
			t.Errorf("individual.New(%f); indiv.Resilience() == %f, expected %f", c.in, got, c.expected)
		}
	}
}

func TestResilience(t *testing.T) {
	cases := []struct {
		in, expected float32
	}{
		{0.7, 0.7},
		{54.0, 54.0},
		{0, 0},
	}
	indiv := New(0)
	for _, c := range cases {
		indiv.SetResilience(c.in)
		got := indiv.Resilience()
		if got != c.expected {
			t.Errorf("indiv.SetResilience(%f); indiv.Resilience() == %f, expected %f", c.in, got, c.expected)
		}
	}
}

func TestCheckNotNil(t *testing.T) {
	err := CheckIndivNotNil(New(0))
	if err != nil {
		t.Errorf("expected err == nil")
	}
	var indiv Interface
	err = CheckIndivNotNil(indiv)
	if err == nil {
		t.Error("expected err != nil")
	}
}
