package population

import (
	"testing"

	"github.com/khezen/darwin/population/individual"
)

func TestNew(t *testing.T) {
	cases := []struct {
		in, expected int
	}{
		{0, 0},
		{1, 1},
		{7, 7},
	}
	for _, c := range cases {
		got, _ := New(c.in)
		if got.Cap() != c.expected {
			t.Errorf("expected  %v", c.expected)
		}
	}
	_, err := New(-1)
	if err == nil {
		t.Errorf("expected != nil")
	}
}

func TestSort(t *testing.T) {
	var i1, i2, i3 = individual.New(0.2), individual.New(0.7), individual.New(1)
	cases := []struct {
		in, expected Population
	}{
		{Population{i1, i2, i3}, Population{i3, i2, i1}},
		{Population{i1, i3, i2}, Population{i3, i2, i1}},
		{Population{i3, i2, i1}, Population{i3, i2, i1}},
	}
	for _, c := range cases {
		c.in.Sort()
		for i := range c.in {
			if c.in[i] != c.expected[i] {
				t.Errorf(".Sort() => %v; expected = %v", c.in, c.expected)
			}
		}
	}
}

func TestCap(t *testing.T) {
	p1, _ := New(7)
	p2, _ := New(0)
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
	pop, err := New(0)
	for _, c := range cases {
		pop.SetCap(c.in)
		if pop.Cap() != c.expected {
			t.Errorf("expected  %v", c.expected)
		}
	}
	err = pop.SetCap(-1)
	if err == nil {
		t.Errorf("expected != nil")
	}
}

func TestTruncate(t *testing.T) {
	var i1, i2, i3 = individual.New(0.2), individual.New(0.7), individual.New(1)
	cases := []struct {
		in       Population
		size     int
		expected Population
	}{
		{Population{i1, i2, i3}, 3, Population{i3, i2, i1}},
		{Population{i1, i3, i2}, 4, Population{i3, i2, i1, nil}},
		{Population{i3, i2, i1}, 2, Population{i3, i2}},
		{Population{i3, i2, i1}, 0, Population{}},
	}
	for _, c := range cases {
		c.in.Truncate(c.size)
		for i := range c.in {
			if c.in[i] != c.expected[i] {
				t.Errorf(".Truncate(%v) => %v; expected = %v", c.size, c.in, c.expected)
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
	var i1, i2, i3 = individual.New(0.2), individual.New(0.7), individual.New(1)
	pop := Population{i2, i1}
	pop.SetCap(10)
	cases := []struct {
		in       Population
		indiv    individual.Interface
		expected Population
	}{
		{Population{i1, i2}, i3, Population{i3, i2, i1}},
		{Population{}, i2, Population{i2}},
		{pop, i3, Population{i3, i2, i1}},
	}
	for _, c := range cases {
		c.in.Append(c.indiv)
		for i := range c.in {
			if c.in[i] != c.expected[i] {
				t.Errorf(".Append(%v) => %v; expected = %v", c.indiv, c.in, c.expected)
			}
		}
	}
	err := pop.Append(nil)
	if err == nil {
		t.Errorf("expected != nil")
	}
}

func TestAppendAll(t *testing.T) {
	var i1, i2, i3, i4 = individual.New(0.2), individual.New(0.7), individual.New(1), individual.New(42.42)
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
		c.in.AppendAll(&c.toAp)
		for i := range c.in {
			if c.in[i] != c.expected[i] {
				t.Errorf(".AppendAll(%v) => %v; expected = %v", c.toAp, c.in, c.expected)
			}
		}
	}
	pop := Population{i2, i1}
	err := pop.AppendAll(nil)
	if err == nil {
		t.Errorf("expected != nil")
	}
}

func TestGet(t *testing.T) {

}

func TestRemove(t *testing.T) {

}

func TestMax(t *testing.T) {

}

func TestMin(t *testing.T) {

}

func TestLen(t *testing.T) {

}

func TestLess(t *testing.T) {

}

func TestSwap(t *testing.T) {

}

func TestPickCouple(t *testing.T) {

}
