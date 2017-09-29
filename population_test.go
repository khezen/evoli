package darwin

import (
	"sync"
	"testing"

	"github.com/khezen/check"
)

func TestNewPopulation(t *testing.T) {
	cases := []struct {
		in, expected int
	}{
		{0, 0},
		{1, 1},
		{7, 7},
	}
	for _, c := range cases {
		var got Population
		got = NewPopulation(c.in)
		if got.Cap() != c.expected {
			t.Errorf("expected  %v", c.expected)
		}
	}
	fail := NewPopulation(-1)
	if fail != nil {
		t.Errorf("Expected nil, got %v", fail)
	}
}

func TestNewPopulationTS(t *testing.T) {
	cases := []struct {
		in, expected int
	}{
		{0, 0},
		{1, 1},
		{7, 7},
	}
	for _, c := range cases {
		var got Population
		got = NewPopulationTS(c.in)
		if got.Cap() != c.expected {
			t.Errorf("expected  %v", c.expected)
		}
	}
	fail := NewPopulation(-1)
	if fail != nil {
		t.Errorf("Expected nil, got %v", fail)
	}
}

func TestSort(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	cases := []struct {
		in, expected Population
	}{
		{&population{i1, i2, i3}, &population{i3, i2, i1}},
		{&population{i1, i3, i2}, &population{i3, i2, i1}},
		{&population{i3, i2, i1}, &population{i3, i2, i1}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, &populationTS{population{i3, i2, i1}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, &populationTS{population{i3, i2, i1}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, &populationTS{population{i3, i2, i1}, sync.RWMutex{}}},
	}
	for _, c := range cases {
		c.in.Sort()
		for i := 0; i < c.in.Len(); i++ {
			exp, err := c.expected.Get(i)
			if err != nil {
				panic(err)
			}
			in, err := c.in.Get(i)
			if err != nil {
				panic(err)
			}
			if in != exp {
				t.Errorf(".Sort() => %v; expected = %v", c.in, c.expected)
				break
			}
		}
	}
}

func TestCap(t *testing.T) {
	cases := []struct {
		in       Population
		expected int
	}{
		{NewPopulation(7), 7},
		{NewPopulation(0), 0},
		{NewPopulationTS(7), 7},
		{NewPopulationTS(0), 0},
	}
	for _, c := range cases {
		got := c.in.Cap()
		if got != c.expected {
			t.Errorf("%v.Cap() => %v; expected = %v", c.in, got, c.expected)
		}
	}
}

func TestSetCap(t *testing.T) {
	cases := []struct {
		pop          Population
		in, expected int
		expectErr    bool
	}{
		{NewPopulation(0), 0, 0, false},
		{NewPopulation(0), 1, 1, false},
		{NewPopulation(0), 7, 7, false},
		{NewPopulation(0), -1, 0, true},
		{NewPopulationTS(0), 0, 0, false},
		{NewPopulationTS(0), 1, 1, false},
		{NewPopulationTS(0), 7, 7, false},
		{NewPopulationTS(0), -1, 0, true},
	}
	for _, c := range cases {
		err := c.pop.SetCap(c.in)
		if c.expectErr && err == nil {
			t.Error("expected error got nil")
		}
		if !c.expectErr && err != nil {
			panic(err)
		}
		if c.pop.Cap() != c.expected {
			t.Errorf("expected  %v", c.expected)
		}
	}
}

func TestAdd(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	cases := []struct {
		in       Population
		indiv    Individual
		expected Population
	}{
		{&population{i1, i2}, i3, &population{i1, i2, i3}},
		{&population{}, i2, &population{i2}},
		{&population{i3, i2}, i1, &population{i3, i2, i1}},
		{&population{i3, i2}, nil, &population{i3, i2, nil}},
		{&populationTS{population{i1, i2}, sync.RWMutex{}}, i3, &populationTS{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationTS{population{}, sync.RWMutex{}}, i2, &populationTS{population{i2}, sync.RWMutex{}}},
		{&populationTS{population{i3, i2}, sync.RWMutex{}}, i1, &populationTS{population{i3, i2, i1}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2}, sync.RWMutex{}}, nil, &populationTS{population{i1, i2, nil}, sync.RWMutex{}}},
	}
	for _, c := range cases {
		c.in.Add(c.indiv)
		for i := 0; i < c.in.Len(); i++ {
			in, err := c.in.Get(i)
			if err != nil {
				panic(err)
			}
			exp, err := c.expected.Get(i)
			if err != nil {
				panic(err)
			}
			if in != exp {
				t.Errorf(".Add(%v) => %v; expected = %v", c.indiv, c.in, c.expected)
				break
			}
		}
	}
}

func TestGet(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	cases := []struct {
		pop Population
	}{
		{&population{i2, i1, i3}},
		{&populationTS{population{i2, i1, i3}, sync.RWMutex{}}},
	}
	for _, c := range cases {
		indiv, _ := c.pop.Get(1)
		if indiv != i1 {
			t.Errorf(".Get(%v) => %v; expected = %v", 1, indiv, i1)
		}
		_, err := c.pop.Get(-1000)
		if err == nil {
			t.Errorf("expected != nil")
		}
		_, err = c.pop.Get(c.pop.Len())
		if err == nil {
			t.Errorf("expected != nil")
		}
	}
}

func TestRemove(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	cases := []struct {
		pop         Population
		toBeRemoved Individual
		expected    []Individual
	}{
		{&population{i1, i2, i3}, i2, []Individual{i1, i3}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, i2, []Individual{i1, i3}},
	}
	for _, c := range cases {
		c.pop.Remove(c.toBeRemoved)
		for _, indiv := range c.expected {
			if !c.pop.Has(indiv) {
				t.Errorf("unexpected")
			}
			if c.pop.Len() != len(c.expected) {
				t.Errorf("unexpected")
			}
		}
	}
}

func TestRemoveAt(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	cases := []struct {
		pop Population
	}{
		{&population{i1, i2, i3}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}},
	}
	for _, c := range cases {
		err := c.pop.RemoveAt(-1000)
		if err == nil {
			t.Errorf("expected != nil")
		}
		err = c.pop.RemoveAt(c.pop.Len())
		if err == nil {
			t.Errorf("expected != nil")
		}
		if c.pop.Len() != 3 {
			t.Errorf(".Remove(%v); pop.Len() => %v; expected = %v", 1, c.pop.Len(), 2)
		}
	}
}

func TestReplace(t *testing.T) {
	i1, i2, i3, i4 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1), NewIndividual(10)
	cases := []struct {
		pop   Population
		index int
		indiv Individual
		isErr bool
	}{
		{&population{i2, i1, i3}, 1, i4, false},
		{&population{i2, i1, i3}, -1000, i4, true},
		{&population{i2, i1, i3}, 3, i4, true},
		{&population{i2, i1, i3}, -42, nil, true},
		{&population{i2, i1, i3}, 3, nil, true},
		{&populationTS{population{i2, i1, i3}, sync.RWMutex{}}, 1, i4, false},
		{&populationTS{population{i2, i1, i3}, sync.RWMutex{}}, -1000, i4, true},
		{&populationTS{population{i2, i1, i3}, sync.RWMutex{}}, 3, i4, true},
		{&populationTS{population{i2, i1, i3}, sync.RWMutex{}}, -42, nil, true},
		{&populationTS{population{i2, i1, i3}, sync.RWMutex{}}, 3, nil, true},
	}
	for _, c := range cases {
		switch c.isErr {
		case true:
			err := c.pop.Replace(c.index, c.indiv)
			if err == nil {
				t.Errorf("expected != nil")
			}
		case false:
			c.pop.Replace(c.index, c.indiv)
			if c.pop.Len() != 3 {
				t.Errorf(".Replace(%v, %v); pop.Len() => %v; expected = %v", c.index, c.indiv, c.pop.Len(), 3)
			}
		}
	}
}

func TestMax(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	var pop Population
	pop = &population{i2, i1, i3}
	max := pop.Max()
	if max != i3 {
		t.Errorf("%v.Max() returned %v instead of %v", pop, max, i3)
	}
	pop = &populationTS{population{i2, i1, i3}, sync.RWMutex{}}
	max = pop.Max()
	if max != i3 {
		t.Errorf("%v.Max() returned %v instead of %v", pop, max, i3)
	}
}

func TestMin(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	var pop Population
	pop = &population{i2, i1, i3}
	min := pop.Min()
	if min != i1 {
		t.Errorf("%v.Min() returned %v instead of %v", pop, min, i1)
	}
	pop = &populationTS{population{i2, i1, i3}, sync.RWMutex{}}
	min = pop.Min()
	if min != i1 {
		t.Errorf("%v.Min() returned %v instead of %v", pop, min, i1)
	}
}

func TestExtremums(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	var pop Population
	pop = &population{i2, i1, i3}
	min, max := pop.Extremums()
	if min != i1 || max != i3 {
		t.Errorf("%v.Extremums() returned (%v, %v) instead of (%v, %v)", pop, min, max, i1, i3)
	}
	pop = &populationTS{population{i2, i1, i3}, sync.RWMutex{}}
	min, max = pop.Extremums()
	if min != i1 || max != i3 {
		t.Errorf("%v.Extremums() returned (%v, %v) instead of (%v, %v)", pop, min, max, i1, i3)
	}
}

func TestLen(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	cases := []struct {
		in       Population
		cap      int
		expected int
	}{
		{&population{i1, i2, i3}, 3, 3},
		{&population{i1, i3, i2}, 4, 3},
		{&population{i3, i1}, 12, 2},
		{&population{i3, i2, i1}, 2, 2},
		{&population{}, 2, 0},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 3, 3},
		{&populationTS{population{i1, i3, i2}, sync.RWMutex{}}, 4, 3},
		{&populationTS{population{i3, i1}, sync.RWMutex{}}, 12, 2},
		{&populationTS{population{i3, i2, i1}, sync.RWMutex{}}, 2, 2},
		{&populationTS{population{}, sync.RWMutex{}}, 2, 0},
	}
	for _, c := range cases {
		c.in.SetCap(c.cap)
		length := c.in.Len()
		if length != c.expected {
			t.Errorf("%v.Len() returned %v instead of %v", c.in, length, c.expected)
		}
	}
}

func TestLess(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	cases := []struct {
		in       Population
		i        int
		j        int
		expected bool
	}{
		{&population{i1, i2, i3}, 0, 0, true},
		{&population{i1, i2, i3}, 0, 1, false},
		{&population{i1, i2, i3}, 0, 2, false},
		{&population{i1, i2, i3}, 1, 0, true},
		{&population{i1, i2, i3}, 1, 1, true},
		{&population{i1, i2, i3}, 1, 2, false},
		{&population{i1, i2, i3}, 2, 0, true},
		{&population{i1, i2, i3}, 2, 1, true},
		{&population{i1, i2, i3}, 2, 2, true},
		{&population{i1, i2, i3}, -1, 0, false},
		{&population{i1, i2, i3}, 0, -1, false},
		{&population{i1, i2, i3}, -1, -1, false},
		{&population{i1, i2, i3}, 1000, 0, false},
		{&population{i1, i2, i3}, 0, 1000, false},
		{&population{i1, i2, i3}, 1000, 1000, false},
		{&population{i1, i2, i3}, -1, 1000, false},
		{&population{i1, i2, i3}, 1000, -1, false},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 0, 0, true},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 0, 1, false},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 0, 2, false},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 1, 0, true},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 1, 1, true},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 1, 2, false},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 2, 0, true},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 2, 1, true},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 2, 2, true},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, -1, 0, false},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 0, -1, false},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, -1, -1, false},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 1000, 0, false},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 0, 1000, false},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 1000, 1000, false},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, -1, 1000, false},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 1000, -1, false},
	}
	for _, c := range cases {
		less := c.in.Less(c.i, c.j)
		if less != c.expected {
			t.Errorf("%v.Less(%v, %v) returned %v instead of %v", c.in, c.i, c.j, less, c.expected)
		}
	}
}

func TestSwap(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	cases := []struct {
		in       Population
		i        int
		j        int
		expected Population
	}{
		{&population{i1, i2, i3}, 0, 0, &population{i1, i2, i3}},
		{&population{i1, i2, i3}, 0, 1, &population{i2, i1, i3}},
		{&population{i1, i2, i3}, 0, 2, &population{i3, i2, i1}},
		{&population{i1, i2, i3}, 1, 0, &population{i2, i1, i3}},
		{&population{i1, i2, i3}, 1, 1, &population{i1, i2, i3}},
		{&population{i1, i2, i3}, 1, 2, &population{i1, i3, i2}},
		{&population{i1, i2, i3}, 2, 0, &population{i3, i2, i1}},
		{&population{i1, i2, i3}, 2, 1, &population{i1, i3, i2}},
		{&population{i1, i2, i3}, 2, 2, &population{i1, i2, i3}},
		{&population{i1, i2, i3}, -1, 0, &population{i1, i2, i3}},
		{&population{i1, i2, i3}, 0, -1, &population{i1, i2, i3}},
		{&population{i1, i2, i3}, -1, -1, &population{i1, i2, i3}},
		{&population{i1, i2, i3}, 1000, 0, &population{i1, i2, i3}},
		{&population{i1, i2, i3}, 0, 1000, &population{i1, i2, i3}},
		{&population{i1, i2, i3}, 1000, 1000, &population{i1, i2, i3}},
		{&population{i1, i2, i3}, -1, 1000, &population{i1, i2, i3}},
		{&population{i1, i2, i3}, 1000, -1, &population{i1, i2, i3}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 0, 0, &populationTS{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 0, 1, &populationTS{population{i2, i1, i3}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 0, 2, &populationTS{population{i3, i2, i1}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 1, 0, &populationTS{population{i2, i1, i3}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 1, 1, &populationTS{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 1, 2, &populationTS{population{i1, i3, i2}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 2, 0, &populationTS{population{i3, i2, i1}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 2, 1, &populationTS{population{i1, i3, i2}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 2, 2, &populationTS{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, -1, 0, &populationTS{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 0, -1, &populationTS{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, -1, -1, &populationTS{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 1000, 0, &populationTS{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 0, 1000, &populationTS{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 1000, 1000, &populationTS{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, -1, 1000, &populationTS{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationTS{population{i1, i2, i3}, sync.RWMutex{}}, 1000, -1, &populationTS{population{i1, i2, i3}, sync.RWMutex{}}},
	}
	for _, c := range cases {
		c.in.Swap(c.i, c.j)
		for i := 0; i < c.in.Len(); i++ {
			indiv, err := c.in.Get(i)
			if err != nil {
				panic(err)
			}
			expected, err := c.expected.Get(i)
			if err != nil {
				panic(err)
			}
			if indiv != expected {
				t.Errorf("%v.Swap(%v, %v) resulted in %v instead of %v", c.in, c.i, c.j, c.in, c.expected)
				break
			}
		}
	}
}

func TestPickCouple(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := NewIndividual(1), NewIndividual(2), NewIndividual(3), NewIndividual(4), NewIndividual(5), NewIndividual(6)

	cases := []struct {
		pop       Population
		expectErr bool
	}{
		{&population{i1, i2, i3, i4, i5, i6}, false},
		{&population{i1, i2}, false},
		{&population{i1}, true},
		{&populationTS{population{i1, i2, i3, i4, i5, i6}, sync.RWMutex{}}, false},
		{&populationTS{population{i1, i2}, sync.RWMutex{}}, false},
		{&populationTS{population{i1}, sync.RWMutex{}}, true},
	}
	for _, c := range cases {
		for i := 0; i < 32; i++ {
			index1, indiv1, index2, indiv2, err := c.pop.PickCouple()
			if !check.ErrorExpectation(err, c.expectErr) {
				t.Errorf("unexpected output %v", err)
			}
			if !c.expectErr {
				if index1 < 0 || index1 >= c.pop.Len() || index2 < 0 || index2 >= c.pop.Len() {
					t.Errorf("%v.PickCouple() returned indexes %v, %v which are out of bounds", c.pop, index1, index2)
				}
				if index1 == index2 {
					t.Errorf("%v.PickCouple() returned indexes %v, %v which are equals", c.pop, index1, index2)
				}
				if indiv1 == nil || indiv2 == nil || indiv1 == indiv2 {
					t.Errorf("%v.PickCouple() returned individuals %v, %v which are nils", c.pop, indiv1, indiv2)
				}
				if err != nil {
					t.Errorf("expected err == nil")
				}
			}
		}
	}
}

func TestContains(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	cases := []struct {
		in       Population
		indiv    Individual
		expected bool
	}{
		{&population{i1, i2}, i1, true},
		{&population{i1, i2}, i3, false},
		{&populationTS{population{i1, i2}, sync.RWMutex{}}, i1, true},
		{&populationTS{population{i1, i2}, sync.RWMutex{}}, i3, false},
	}
	for _, c := range cases {
		contains := c.in.Has(c.indiv)
		if contains != c.expected {
			t.Errorf("%v.Contains(%v) returned %v instead of %v", c.in, c.indiv, contains, c.expected)
		}
	}
}

func TestIndexOf(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	cases := []struct {
		in       Population
		indiv    Individual
		expected int
	}{
		{&population{i1, i2}, i1, 0},
		{&population{i1, i2}, i2, 1},
		{&population{i1, i2}, i3, -1},
		{&populationTS{population{i1, i2}, sync.RWMutex{}}, i1, 0},
		{&populationTS{population{i1, i2}, sync.RWMutex{}}, i2, 1},
		{&populationTS{population{i1, i2}, sync.RWMutex{}}, i3, -1},
	}
	for _, c := range cases {
		index, _ := c.in.IndexOf(c.indiv)
		if index != c.expected {
			t.Errorf("%v.Contains(%v) returned %v instead of %v", c.in, c.indiv, index, c.expected)
		}
	}
	pop := population{i2, i1, i3}
	_, err := pop.IndexOf(nil)
	if err == nil {
		panic(err)
	}
}

func TestEach(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := NewIndividual(1), NewIndividual(2), NewIndividual(3), NewIndividual(4), NewIndividual(5), NewIndividual(6)
	cases := []struct {
		pop         Population
		individuals []Individual
	}{
		{NewPopulation(4), []Individual{i1, i5, i6, i4}},
		{NewPopulation(5), []Individual{i1, i3, i5, i2, i6}},
		{NewPopulationTS(4), []Individual{i1, i5, i6, i4}},
		{NewPopulationTS(5), []Individual{i1, i3, i5, i2, i6}},
	}
	for _, c := range cases {
		c.pop.Add(c.individuals...)
		c.pop.Each(func(indiv Individual) bool {
			has := false
			for _, current := range c.individuals {
				if current == indiv {
					has = true
					break
				}
			}
			if !has {
				t.Errorf("Each traversal %v which is not found in %v", indiv, c.individuals)
			}
			return true
		})
		c.pop.Each(func(indiv Individual) bool {
			return false
		})
	}
}

func TestSlice(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	cases := []struct {
		pop   Population
		slice []Individual
	}{
		{NewPopulation(3), []Individual{i1, i2, i3}},
		{NewPopulationTS(3), []Individual{i1, i2, i3}},
	}
	for _, c := range cases {
		c.pop.Add(c.slice...)
		slice := c.pop.Slice()
		for i := range c.slice {
			indiv := slice[i]
			if indiv != c.slice[i] {
				t.Errorf("expected %v, got %v", c.slice[i], indiv)
			}
		}
	}
}
