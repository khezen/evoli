package evoli

import (
	"sync"
	"testing"
)

func TestPoolCRUD(t *testing.T) {
	gen := NewGenetic(NewPopulationSync(10), NewTruncationSelecter(), 5, crosserMock{}, mutaterMock{}, 1, evaluaterMock{})
	cases := []struct {
		pool Pool
		e    Evolution
	}{
		{NewPool(1), gen},
		{NewPoolSync(1), gen},
	}
	for _, c := range cases {
		c.pool.Add(c.e)
		has := c.pool.Has(c.e)
		if !has {
			t.Error("expected true, got false")
		}
		c.pool.Delete(c.e)
		has = c.pool.Has(c.e)
		if has {
			t.Errorf("expected false, got true")
		}
	}
}

func TestAlpha(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1), NewIndividual(0), NewIndividual(100), NewIndividual(42)
	pop1, pop2 := &population{i1, i2, i3}, &population{i4, i5, i6}
	gen1, gen2 := NewGenetic(pop1, NewTruncationSelecter(), 5, crosserMock{}, mutaterMock{}, 1, evaluaterMock{}), NewGenetic(pop2, NewTruncationSelecter(), 5, crosserMock{}, mutaterMock{}, 1, evaluaterMock{})
	cases := []struct {
		pool        Pool
		populations []Evolution
		alpha       Individual
	}{
		{NewPool(2), []Evolution{gen1, gen2}, i5},
		{NewPoolSync(2), []Evolution{gen1, gen2}, i5},
	}
	for _, c := range cases {
		for _, evolution := range c.populations {
			c.pool.Add(evolution)
		}
		alpha := c.pool.Alpha()
		if alpha != c.alpha {
			t.Errorf("expected %v got %v", c.alpha, alpha)
		}
	}
}

func TestCollections(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1), NewIndividual(0), NewIndividual(100), NewIndividual(42)
	pop1, pop2 := &population{i1, i2, i3}, &population{i4, i5, i6}
	gen1, gen2 := NewGenetic(pop1, NewTruncationSelecter(), 5, crosserMock{}, mutaterMock{}, 1, evaluaterMock{}), NewGenetic(pop2, NewTruncationSelecter(), 5, crosserMock{}, mutaterMock{}, 1, evaluaterMock{})
	cases := []struct {
		pool       Pool
		evolutions []Evolution
	}{
		{NewPool(2), []Evolution{gen1, gen2}},
		{NewPoolSync(2), []Evolution{gen1, gen2}},
	}
	for _, c := range cases {
		for _, evolution := range c.evolutions {
			c.pool.Add(evolution)
		}
		evolutions := c.pool.Evolutions()
		if len(evolutions) != len(c.evolutions) {
			t.Errorf("expected %v, got %v", len(c.evolutions), len(evolutions))
		}
		for _, e := range evolutions {
			found := false
			for _, ce := range c.evolutions {
				if e == ce {
					found = true
					break
				}
			}
			if !found {
				t.Error("unexpected result")
			}
		}

		populations := c.pool.Populations()
		if len(populations) != len(c.evolutions) {
			t.Errorf("expected %v, got %v", len(c.evolutions), len(populations))
		}
		for _, pop := range populations {
			found := false
			for _, e := range c.evolutions {
				if e.Population() == pop {
					found = true
					break
				}
			}
			if !found {
				t.Error("unexpected result")
			}
		}
		individuals := c.pool.Individuals()
		for _, indiv := range individuals {
			has := false
			for _, e := range c.evolutions {
				pop := e.Population()
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

func TestShuffle(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1), NewIndividual(0), NewIndividual(100), NewIndividual(42)
	i7, i8, i9, i10, i11, i12 := NewIndividual(0.2), NewIndividual(0.7), NewIndividual(1), NewIndividual(0), NewIndividual(100), NewIndividual(42)
	pop1, pop2 := &population{i1, i2, i3, i4, i5, i6}, &population{i7, i8, i9, i10, i11, i12}
	cases := []struct {
		pool        Pool
		populations []Population
	}{
		{NewPool(2), []Population{pop1, pop2}},
		{NewPoolSync(2), []Population{pop1, pop2}},
	}
	for _, c := range cases {
		for _, cpop := range c.populations {
			pop := cpop.New(cpop.Cap())
			pop.Add(cpop.Slice()...)
			gen := NewGenetic(pop, NewTruncationSelecter(), 5, crosserMock{}, mutaterMock{}, 1, evaluaterMock{})
			c.pool.Add(gen)
		}
		c.pool.Shuffle()
		populations := c.pool.Populations()
		for _, pop := range populations {
			if pop.Len() != 6 {
				t.Errorf("expected %v got %v", 6, pop.Len())
			}
			individuals := pop.Slice()
			for _, cpop := range c.populations {
				different := false
				cindividuals := cpop.Slice()
				for _, indiv := range individuals {
					has := false
					for _, cindiv := range cindividuals {
						if indiv == cindiv {
							has = true
							break
						}
					}
					if !has {
						different = true
						break
					}
				}
				if !different {
					t.Error("shuffle produced the same populations")
				}
			}
		}
	}
}

func TestPoolNext(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := NewIndividualSync(0.2), NewIndividualSync(0.7), NewIndividualSync(1), NewIndividualSync(0), NewIndividualSync(100), NewIndividualSync(42)
	i7, i8, i9, i10, i11, i12 := NewIndividualSync(0.2), NewIndividualSync(0.7), NewIndividualSync(1), NewIndividualSync(0), NewIndividualSync(100), NewIndividualSync(42)
	pop1, pop2 := &populationSync{population{i1, i2, i3, i4, i5, i6}, sync.RWMutex{}}, &populationSync{population{i7, i8, i9, i10, i11, i12}, sync.RWMutex{}}
	cases := []struct {
		pool        Pool
		populations []Population
	}{
		{NewPool(2), []Population{pop1, pop2}},
		{NewPoolSync(2), []Population{pop1, pop2}},
	}
	for _, c := range cases {
		for _, cpop := range c.populations {
			pop := cpop.New(cpop.Cap())
			pop.Add(cpop.Slice()...)
			gen := NewGenetic(pop, NewTruncationSelecter(), 2, crosserMock{}, mutaterMock{}, 0.05, evaluaterMock{})
			c.pool.Add(gen)
		}
		err := c.pool.Next()
		if err != nil {
			t.Error(err)
		}
		populations := c.pool.Populations()
		for _, pop := range populations {
			if pop.Len() != 6 {
				t.Errorf("expected %v got %v", 6, pop.Len())
			}
			individuals := pop.Slice()
			for _, cpop := range c.populations {
				different := false
				cindividuals := cpop.Slice()
				for _, indiv := range individuals {
					has := false
					for _, cindiv := range cindividuals {
						if indiv == cindiv {
							has = true
							break
						}
					}
					if !has {
						different = true
						break
					}
				}
				if !different {
					t.Error("Next produced the same populations")
				}
			}
		}
	}
}

func TestPoolNextAsync(t *testing.T) {
	i1, i2, i3, i4, i5, i6 := NewIndividualSync(0.2), NewIndividualSync(0.7), NewIndividualSync(1), NewIndividualSync(0), NewIndividualSync(100), NewIndividualSync(42)
	i7, i8, i9, i10, i11, i12 := NewIndividualSync(0.2), NewIndividualSync(0.7), NewIndividualSync(1), NewIndividualSync(0), NewIndividualSync(100), NewIndividualSync(42)
	pop1, pop2 := &populationSync{population{i1, i2, i3, i4, i5, i6}, sync.RWMutex{}}, &populationSync{population{i7, i8, i9, i10, i11, i12}, sync.RWMutex{}}
	cases := []struct {
		pool        Pool
		populations []Population
	}{
		{NewPool(2), []Population{pop1, pop2}},
		{NewPoolSync(2), []Population{pop1, pop2}},
	}
	for _, c := range cases {
		for _, cpop := range c.populations {
			pop := cpop.New(cpop.Cap())
			pop.Add(cpop.Slice()...)
			gen := NewGenetic(pop, NewTruncationSelecter(), 2, crosserMock{}, mutaterMock{}, 0.05, evaluaterMock{})
			c.pool.Add(gen)
		}
		err := c.pool.NextAsync()
		if err != nil {
			t.Error(err)
		}
		populations := c.pool.Populations()
		for _, pop := range populations {
			if pop.Len() != 6 {
				t.Errorf("expected %v got %v", 6, pop.Len())
			}
			individuals := pop.Slice()
			for _, cpop := range c.populations {
				different := false
				cindividuals := cpop.Slice()
				for _, indiv := range individuals {
					has := false
					for _, cindiv := range cindividuals {
						if indiv == cindiv {
							has = true
							break
						}
					}
					if !has {
						different = true
						break
					}
				}
				if !different {
					t.Error("Next produced the same populations")
				}
			}

		}
	}
}
