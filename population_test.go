package evoli

import (
	"sync"
	"testing"
)

func TestNewPopulation(t *testing.T) {
	cases := []struct {
		in, expected int
	}{
		{1, 1},
		{7, 7},
	}
	for _, c := range cases {
		var got Population
		got = NewPopulation(c.in)
		if got.Cap() != c.expected {
			t.Errorf("expected  %v", c.expected)
		}
		got = got.New(c.in)
		if got.Cap() != c.expected {
			t.Errorf("expected  %v", c.expected)
		}
	}
}

func TestNewPopulationSync(t *testing.T) {
	cases := []struct {
		in, expected int
	}{
		{1, 1},
		{7, 7},
	}
	for _, c := range cases {
		var got Population
		got = NewPopulationSync(c.in)
		if got.Cap() != c.expected {
			t.Errorf("expected  %v", c.expected)
		}
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
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, &populationSync{population{i3, i2, i1}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, &populationSync{population{i3, i2, i1}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, &populationSync{population{i3, i2, i1}, sync.RWMutex{}}},
	}
	for _, c := range cases {
		c.in.Sort()
		for i := 0; i < c.in.Len(); i++ {
			exp := c.expected.Get(i)
			in := c.in.Get(i)
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
		{NewPopulationSync(7), 7},
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
	}{
		{NewPopulation(2), 1, 1},
		{NewPopulation(2), 7, 7},
		{NewPopulationSync(2), 1, 1},
		{NewPopulationSync(2), 7, 7},
	}
	for _, c := range cases {
		c.pop.SetCap(c.in)
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
		{&populationSync{population{i1, i2}, sync.RWMutex{}}, i3, &populationSync{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationSync{population{}, sync.RWMutex{}}, i2, &populationSync{population{i2}, sync.RWMutex{}}},
		{&populationSync{population{i3, i2}, sync.RWMutex{}}, i1, &populationSync{population{i3, i2, i1}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2}, sync.RWMutex{}}, nil, &populationSync{population{i1, i2, nil}, sync.RWMutex{}}},
	}
	for _, c := range cases {
		c.in.Add(c.indiv)
		for i := 0; i < c.in.Len(); i++ {
			in := c.in.Get(i)
			exp := c.expected.Get(i)
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
		{&populationSync{population{i2, i1, i3}, sync.RWMutex{}}},
	}
	for _, c := range cases {
		indiv := c.pop.Get(1)
		if indiv != i1 {
			t.Errorf(".Get(%v) => %v; expected = %v", 1, indiv, i1)
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
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, i2, []Individual{i1, i3}},
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
		pop              Population
		indexToBeRemoved int
		expected         []Individual
	}{
		{&population{i1, i2, i3}, 1, []Individual{i1, i3}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 1, []Individual{i1, i3}},
	}
	for _, c := range cases {
		c.pop.RemoveAt(c.indexToBeRemoved)
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

func TestReplace(t *testing.T) {
	i1, i2, i3, i4 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1), NewIndividual(10)
	cases := []struct {
		pop   Population
		index int
		indiv Individual
	}{
		{&population{i2, i1, i3}, 1, i4},
		{&populationSync{population{i2, i1, i3}, sync.RWMutex{}}, 1, i4},
	}
	for _, c := range cases {
		c.pop.Replace(c.index, c.indiv)
		if c.pop.Len() != 3 {
			t.Errorf(".Replace(%v, %v); pop.Len() => %v; expected = %v", c.index, c.indiv, c.pop.Len(), 3)
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
	pop = &populationSync{population{i2, i1, i3}, sync.RWMutex{}}
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
	pop = &populationSync{population{i2, i1, i3}, sync.RWMutex{}}
	min = pop.Min()
	if min != i1 {
		t.Errorf("%v.Min() returned %v instead of %v", pop, min, i1)
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
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 3, 3},
		{&populationSync{population{i1, i3, i2}, sync.RWMutex{}}, 4, 3},
		{&populationSync{population{i3, i1}, sync.RWMutex{}}, 12, 2},
		{&populationSync{population{i3, i2, i1}, sync.RWMutex{}}, 2, 2},
		{&populationSync{population{}, sync.RWMutex{}}, 2, 0},
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
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 0, 0, true},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 0, 1, false},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 0, 2, false},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 1, 0, true},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 1, 1, true},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 1, 2, false},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 2, 0, true},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 2, 1, true},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 2, 2, true},
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
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 0, 0, &populationSync{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 0, 1, &populationSync{population{i2, i1, i3}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 0, 2, &populationSync{population{i3, i2, i1}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 1, 0, &populationSync{population{i2, i1, i3}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 1, 1, &populationSync{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 1, 2, &populationSync{population{i1, i3, i2}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 2, 0, &populationSync{population{i3, i2, i1}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 2, 1, &populationSync{population{i1, i3, i2}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 2, 2, &populationSync{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, -1, 0, &populationSync{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 0, -1, &populationSync{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, -1, -1, &populationSync{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 1000, 0, &populationSync{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 0, 1000, &populationSync{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 1000, 1000, &populationSync{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, -1, 1000, &populationSync{population{i1, i2, i3}, sync.RWMutex{}}},
		{&populationSync{population{i1, i2, i3}, sync.RWMutex{}}, 1000, -1, &populationSync{population{i1, i2, i3}, sync.RWMutex{}}},
	}
	for _, c := range cases {
		c.in.Swap(c.i, c.j)
		for i := 0; i < c.in.Len(); i++ {
			indiv := c.in.Get(i)
			expected := c.expected.Get(i)
			if indiv != expected {
				t.Errorf("%v.Swap(%v, %v) resulted in %v instead of %v", c.in, c.i, c.j, c.in, c.expected)
				break
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
		{&populationSync{population{i1, i2}, sync.RWMutex{}}, i1, true},
		{&populationSync{population{i1, i2}, sync.RWMutex{}}, i3, false},
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
		{&populationSync{population{i1, i2}, sync.RWMutex{}}, i1, 0},
		{&populationSync{population{i1, i2}, sync.RWMutex{}}, i2, 1},
		{&populationSync{population{i1, i2}, sync.RWMutex{}}, i3, -1},
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
		{NewPopulationSync(4), []Individual{i1, i5, i6, i4}},
		{NewPopulationSync(5), []Individual{i1, i3, i5, i2, i6}},
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
		{NewPopulationSync(3), []Individual{i1, i2, i3}},
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
