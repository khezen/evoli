package evoli

import "sync"

type poolTS struct {
	pool
	sync.RWMutex
}

// NewPoolTS creates a sync implementation of Pool
func NewPoolTS() Pool {
	return &poolTS{
		*NewPool().(*pool),
		sync.RWMutex{},
	}
}

func (p *poolTS) Put(e Evolution, pop Population) {
	p.Lock()
	defer p.Unlock()
	p.pool.Put(e, pop)
}

func (p *poolTS) Delete(e Evolution) {
	p.Lock()
	defer p.Unlock()
	p.pool.Delete(e)
}

func (p *poolTS) Has(e Evolution) bool {
	p.RLock()
	defer p.RUnlock()
	return p.pool.Has(e)
}

func (p *poolTS) Get(e Evolution) Population {
	p.RLock()
	defer p.RUnlock()
	return p.pool.Get(e)
}

func (p *poolTS) Max() Individual {
	p.RLock()
	defer p.RUnlock()
	return p.pool.Max()
}

func (p *poolTS) Min() Individual {
	p.RLock()
	defer p.RUnlock()
	return p.pool.Min()
}

func (p *poolTS) Shuffle() {
	p.Lock()
	defer p.Unlock()
	p.pool.Shuffle()
}

func (p *poolTS) Next() error {
	p.RLock()
	defer p.RUnlock()
	return p.pool.Next()
}
