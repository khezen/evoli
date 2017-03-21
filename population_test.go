package darwin

import (
	"github.com/khezen/check"
	"testing"
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
		var got IPopulation
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

func TestSort(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	cases := []struct {
		in, expected Population
	}{
		{Population{i1, i2, i3}, Population{i3, i2, i1}},
		{Population{i1, i3, i2}, Population{i3, i2, i1}},
		{Population{i3, i2, i1}, Population{i3, i2, i1}},
	}
	for _, c := range cases {
		c.in.Sort()
		for i := range c.expected {
			if c.in[i] != c.expected[i] {
				t.Errorf(".Sort() => %v; expected = %v", c.in, c.expected)
				break
			}
		}
	}
}

func TestCap(t *testing.T) {
	p1 := NewPopulation(7)
	p2 := NewPopulation(0)
	cases := []struct {
		in       *Population
		expected int
	}{
		{p1, 7},
		{p2, 0},
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
		in, expected int
	}{
		{0, 0},
		{1, 1},
		{7, 7},
	}
	pop := NewPopulation(0)
	for _, c := range cases {
		pop.SetCap(c.in)
		if pop.Cap() != c.expected {
			t.Errorf("expected  %v", c.expected)
		}
	}
	err := pop.SetCap(-1)
	if err == nil {
		t.Errorf("expected != nil")
	}
}

func TestTruncate(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	cases := []struct {
		in       Population
		size     int
		expected Population
	}{
		{Population{i1, i2, i3}, 3, Population{i1, i2, i3}},
		{Population{i1, i3, i2}, 4, Population{i3, i2, i1}},
		{Population{i3, i2, i1}, 2, Population{i3, i2}},
		{Population{i3, i2, i1}, 0, Population{}},
		{Population{i1}, 3, Population{i1}},
	}
	for _, c := range cases {
		c.in.Truncate(c.size)
		for i := range c.expected {
			if !c.expected.Has(c.in[i]) {
				t.Errorf(".Truncate(%v) => %v; expected = %v", c.size, c.in, c.expected)
				break
			}
		}
	}
	pop := Population{i1, i2, i3}
	err := pop.Truncate(-15)
	if err == nil {
		t.Errorf("expected != nil")
	}
}

func TestAppend(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	cases := []struct {
		in       Population
		indiv    IIndividual
		expected Population
	}{
		{Population{i1, i2}, i3, Population{i1, i2, i3}},
		{Population{}, i2, Population{i2}},
		{Population{i3, i2}, i1, Population{i3, i2, i1}},
	}
	for _, c := range cases {
		c.in.Append(c.indiv)
		for i := range c.expected {
			if c.in[i] != c.expected[i] {
				t.Errorf(".Append(%v) => %v; expected = %v", c.indiv, c.in, c.expected)
				break
			}
		}
	}
	pop := Population{i2, i1}
	pop.Append(nil)
}

func TestAppendAll(t *testing.T) {
	i1, i2, i3, i4 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1), NewIndividual(42.42)
	cases := []struct {
		in, toAp, expected Population
	}{
		{Population{i1, i2}, Population{i3}, Population{i1, i2, i3}},
		{Population{}, Population{i1, i3}, Population{i1, i3}},
		{Population{i1, i2}, Population{i3, i4}, Population{i1, i2, i3, i4}},
		{Population{i4, i1}, Population{}, Population{i4, i1}},
		{Population{}, Population{}, Population{}},
	}
	for _, c := range cases {
		c.in.Append(c.toAp...)
		for i := range c.expected {
			if c.in[i] != c.expected[i] {
				t.Errorf(".AppendAll(%v) => %v; expected = %v", c.toAp, c.in, c.expected)
				break
			}
		}
	}
	pop := &Population{i2, i1}
	toBeAppended := Population{i2, i1, nil}
	pop.Append(toBeAppended...)
}

func TestGet(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	pop := Population{i2, i1, i3}

	indiv, _ := pop.Get(1)
	if indiv != i1 {
		t.Errorf(".Get(%v) => %v; expected = %v", 1, indiv, i1)
	}
	_, err := pop.Get(-1000)
	if err == nil {
		t.Errorf("expected != nil")
	}
	_, err = pop.Get(pop.Len())
	if err == nil {
		t.Errorf("expected != nil")
	}
}

func TestRemove(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	cases := []struct {
		pop         *Population
		toBeRemoved IIndividual
		expected    []IIndividual
	}{
		{&Population{i1, i2, i3}, i2, []IIndividual{i1, i3}},
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
	pop := Population{i2, i1, i3}
	err := pop.RemoveAt(-1000)
	if err == nil {
		t.Errorf("expected != nil")
	}
	err = pop.RemoveAt(pop.Len())
	if err == nil {
		t.Errorf("expected != nil")
	}
	if pop.Len() != 3 {
		t.Errorf(".Remove(%v); pop.Len() => %v; expected = %v", 1, pop.Len(), 2)
	}
}

func TestReplace(t *testing.T) {
	i1, i2, i3, i4 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1), NewIndividual(10)
	pop := Population{i2, i1, i3}
	cases := []struct {
		index int
		indiv *Individual
		isErr bool
	}{
		{1, i4, false},
		{-1000, i4, true},
		{pop.Len(), i4, true},
		{-42, nil, true},
		{pop.Len(), nil, true},
	}
	for _, c := range cases {
		switch c.isErr {
		case true:
			err := pop.Replace(c.index, c.indiv)
			if err == nil {
				t.Errorf("expected != nil")
			}
		case false:
			pop.Replace(c.index, c.indiv)
			if pop.Len() != 3 {
				t.Errorf(".Replace(%v, %v); pop.Len() => %v; expected = %v", c.index, c.indiv, pop.Len(), 3)
			}
		}
	}
}

func TestMax(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	pop := Population{i2, i1, i3}
	max := pop.Max()
	if max != i3 {
		t.Errorf("%v.Max() returned %v instead of %v", pop, max, i3)
	}
}

func TestMin(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	pop := Population{i2, i1, i3}
	min := pop.Min()
	if min != i1 {
		t.Errorf("%v.Min() returned %v instead of %v", pop, min, i1)
	}
}

func TestExtremums(t *testing.T) {
	i1, i2, i3 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1)
	pop := Population{i2, i1, i3}
	min, max := pop.Extremums()
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
		{Population{i1, i2, i3}, 3, 3},
		{Population{i1, i3, i2}, 4, 3},
		{Population{i3, i1}, 12, 2},
		{Population{i3, i2, i1}, 2, 2},
		{Population{}, 2, 0},
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
		{Population{i1, i2, i3}, 0, 0, true},
		{Population{i1, i2, i3}, 0, 1, false},
		{Population{i1, i2, i3}, 0, 2, false},
		{Population{i1, i2, i3}, 1, 0, true},
		{Population{i1, i2, i3}, 1, 1, true},
		{Population{i1, i2, i3}, 1, 2, false},
		{Population{i1, i2, i3}, 2, 0, true},
		{Population{i1, i2, i3}, 2, 1, true},
		{Population{i1, i2, i3}, 2, 2, true},
		{Population{i1, i2, i3}, -1, 0, false},
		{Population{i1, i2, i3}, 0, -1, false},
		{Population{i1, i2, i3}, -1, -1, false},
		{Population{i1, i2, i3}, 1000, 0, false},
		{Population{i1, i2, i3}, 0, 1000, false},
		{Population{i1, i2, i3}, 1000, 1000, false},
		{Population{i1, i2, i3}, -1, 1000, false},
		{Population{i1, i2, i3}, 1000, -1, false},
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
		{Population{i1, i2, i3}, 0, 0, Population{i1, i2, i3}},
		{Population{i1, i2, i3}, 0, 1, Population{i2, i1, i3}},
		{Population{i1, i2, i3}, 0, 2, Population{i3, i2, i1}},
		{Population{i1, i2, i3}, 1, 0, Population{i2, i1, i3}},
		{Population{i1, i2, i3}, 1, 1, Population{i1, i2, i3}},
		{Population{i1, i2, i3}, 1, 2, Population{i1, i3, i2}},
		{Population{i1, i2, i3}, 2, 0, Population{i3, i2, i1}},
		{Population{i1, i2, i3}, 2, 1, Population{i1, i3, i2}},
		{Population{i1, i2, i3}, 2, 2, Population{i1, i2, i3}},
		{Population{i1, i2, i3}, -1, 0, Population{i1, i2, i3}},
		{Population{i1, i2, i3}, 0, -1, Population{i1, i2, i3}},
		{Population{i1, i2, i3}, -1, -1, Population{i1, i2, i3}},
		{Population{i1, i2, i3}, 1000, 0, Population{i1, i2, i3}},
		{Population{i1, i2, i3}, 0, 1000, Population{i1, i2, i3}},
		{Population{i1, i2, i3}, 1000, 1000, Population{i1, i2, i3}},
		{Population{i1, i2, i3}, -1, 1000, Population{i1, i2, i3}},
		{Population{i1, i2, i3}, 1000, -1, Population{i1, i2, i3}},
	}
	for _, c := range cases {
		pop := NewPopulation(c.in.Cap())
		pop.Append(c.in...)
		pop.Swap(c.i, c.j)
		for i := range *pop {
			indiv, _ := pop.Get(i)
			expected, _ := c.expected.Get(i)
			if indiv != expected {
				t.Errorf("%v.Swap(%v, %v) resulted in %v instead of %v", c.in, c.i, c.j, pop, c.expected)
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
		{Population{i1, i2, i3, i4, i5, i6}, false},
		{Population{i1, i2}, false},
		{Population{i1}, true},
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
		indiv    IIndividual
		expected bool
	}{
		{Population{i1, i2}, i1, true},
		{Population{i1, i2}, i3, false},
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
		indiv    IIndividual
		expected int
	}{
		{Population{i1, i2}, i1, 0},
		{Population{i1, i2}, i2, 1},
		{Population{i1, i2}, i3, -1},
	}
	for _, c := range cases {
		index, _ := c.in.IndexOf(c.indiv)
		if index != c.expected {
			t.Errorf("%v.Contains(%v) returned %v instead of %v", c.in, c.indiv, index, c.expected)
		}
	}
	pop := Population{i2, i1, i3}
	_, err := pop.IndexOf(nil)
	if err == nil {
		t.Errorf("expected err != nil")
	}

}
