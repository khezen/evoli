package individual

import (
	"testing"

	"github.com/khezen/darwin/population/individual"
)

func testResilience(t *testing.T) {
	cases := []struct {
		in, want float32
	}{
		{0.7, 0.7},
		{54.0, 54.0},
		{0, 0},
	}
	indiv = individual.New(0)
	for _, c := range cases {
		got := indiv.Resilience()
		if got != c.want {
			t.Errorf("indiv.SetResilience(%float32); indiv.Resilience() == %float32, want %float32", c.in, got, c.want)
		}
	}

	err := indiv.SetResilience(-1)
	if err == nil {
		t.Errorf("indiv.SetResilience(%float32); must return an error", -1.0)
	}
}
