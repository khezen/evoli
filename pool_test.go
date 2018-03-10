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
		c.pool.Put(c.e, c.p)
		has := c.pool.Has(c.e)
		if !has {
			t.Error("expected true, got false")
		}
		pop := c.pool.Get(c.e)
		if pop != c.p {
			t.Errorf("expected %v, got %v", c.p, pop)
		}
		c.pool.Delete(c.e)
		has = c.pool.Has(c.e)
		if has {
			t.Errorf("expected false, got true")
		}
		pop = c.pool.Get(c.e)
		if pop != nil {
			t.Errorf("expected nil, got %v", pop)
		}
	}
}

func testMax(t *testing.T) {

}

func testMin(t *testing.T) {

}

func testShuffle(t *testing.T) {

}

func testNext(t *testing.T) {

}
