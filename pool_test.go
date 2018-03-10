package evoli

import "testing"

func TestPoolCRUD(t *testing.T) {
	gen, pop := NewGenetic(NewTruncationSelecter(), 5, crosserMock{}, mutaterMock{}, 1, evaluaterMock{}), NewPopulationTS(10)
	cases := []struct {
		pool Pool
		e    Evolution
		p    Population
	}{
		{NewPool(), gen, pop},
		{NewPoolTS(), gen, pop},
	}
	for _, c := range cases {
		c.pool.Put(c.p, c.e)
		has := c.pool.Has(c.p)
		if !has {
			t.Error("expected true, got false")
		}
		e := c.pool.Evolution(c.p)
		if e != c.e {
			t.Errorf("expected %v, got %v", c.e, e)
		}
		c.pool.Delete(c.p)
		has = c.pool.Has(c.p)
		if has {
			t.Errorf("expected false, got true")
		}
		e = c.pool.Evolution(c.p)
		if e != nil {
			t.Errorf("expected nil, got %v", e)
		}
	}
}

func TestMinMax(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1), NewIndividual(0), NewIndividual(100), NewIndividual(42)
	pop1, pop2 := &population{i1, i2, i3}, &population{i4, i5, i6}
	gen := NewGenetic(NewTruncationSelecter(), 5, crosserMock{}, mutaterMock{}, 1, evaluaterMock{})
	cases := []struct {
		pool        Pool
		populations map[Population]Evolution
		min, max    Individual
	}{
		{NewPool(), map[Population]Evolution{pop1: gen, pop2: gen}, i1, i5},
		{NewPoolTS(), map[Population]Evolution{pop1: gen, pop2: gen}, i1, i5},
	}
	for _, c := range cases {
		for pop, evolution := range c.populations {
			c.pool.Put(pop, evolution)
		}
		max := c.pool.Max()
		if max != c.max {
			t.Errorf("expected %v got %v", c.max, max)
		}
		min := c.pool.Min()
		if max != c.max {
			t.Errorf("expected %v got %v", c.min, min)
		}
	}
}

func TestCollections(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1), NewIndividual(0), NewIndividual(100), NewIndividual(42)
	pop1, pop2 := &population{i1, i2, i3}, &population{i4, i5, i6}
	gen := NewGenetic(NewTruncationSelecter(), 5, crosserMock{}, mutaterMock{}, 1, evaluaterMock{})
	cases := []struct {
		pool        Pool
		populations map[Population]Evolution
	}{
		{NewPool(), map[Population]Evolution{pop1: gen, pop2: gen}},
		{NewPoolTS(), map[Population]Evolution{pop1: gen, pop2: gen}},
	}
	for _, c := range cases {
		for pop, evolution := range c.populations {
			c.pool.Put(pop, evolution)
		}
		populations := c.pool.Populations()
		if len(populations) != len(c.populations) {
			t.Errorf("expected %v, got %v", len(c.populations), len(populations))
		}
		for _, pop := range populations {
			_, ok := c.populations[pop]
			if !ok {
				t.Error("unexpected result")
			}
		}
		individuals := c.pool.Individuals()
		for _, indiv := range individuals {
			has := false
			for pop := range c.populations {
				has = pop.Has(indiv)
				if has {
					break
				}
			}
			if !has {
				t.Error("indiv not found")
			}
		}
	}
}

func testShuffle(t *testing.T) {

}

func testNext(t *testing.T) {

}
