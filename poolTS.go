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

func (p *poolTS) Put(pop Population, e Evolution) {
	p.Lock()
	defer p.Unlock()
	p.pool.Put(pop, e)
}

func (p *poolTS) Delete(pop Population) {
	p.Lock()
	defer p.Unlock()
	p.pool.Delete(pop)
}

func (p *poolTS) Has(pop Population) bool {
	p.RLock()
	defer p.RUnlock()
	return p.pool.Has(pop)
}

func (p *poolTS) Evolution(pop Population) Evolution {
	p.RLock()
	defer p.RUnlock()
	return p.pool.Evolution(pop)
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
