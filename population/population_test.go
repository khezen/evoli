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

func testSetCap(t *testing.T) {
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

func testTruncate(t *testing.T) {

}

func testAppend(t *testing.T) {

}

func testGet(t *testing.T) {

}

func testRemove(t *testing.T) {

}

func testMax(t *testing.T) {

}

func testMin(t *testing.T) {

}

func testLen(t *testing.T) {

}

func testLess(t *testing.T) {

}

func testSwap(t *testing.T) {

}
