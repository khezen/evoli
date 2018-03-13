package evoli

import "sync"

type poolSync struct {
	pool
	sync.RWMutex
}

// NewPoolSync creates a sync implementation of Pool
func NewPoolSync(length int) Pool {
	return &poolSync{
		*NewPool(length).(*pool),
		sync.RWMutex{},
	}
}

func (p *poolSync) Add(e Evolution) {
	p.Lock()
	defer p.Unlock()
	p.pool.Add(e)
}

func (p *poolSync) Delete(e Evolution) {
	p.Lock()
	defer p.Unlock()
	p.pool.Delete(e)
}

func (p *poolSync) Has(e Evolution) bool {
	p.RLock()
	defer p.RUnlock()
	return p.pool.Has(e)
}

func (p *poolSync) Evolutions() []Evolution {
	p.RLock()
	defer p.RUnlock()
	return p.pool.Evolutions()
}

func (p *poolSync) Populations() []Population {
	p.RLock()
	defer p.RUnlock()
	return p.pool.Populations()
}

func (p *poolSync) Individuals() []Individual {
	p.RLock()
	defer p.RUnlock()
	return p.pool.Individuals()
}

func (p *poolSync) Alpha() Individual {
	p.RLock()
	defer p.RUnlock()
	return p.pool.Alpha()
}

func (p *poolSync) Shuffle() {
	p.Lock()
	defer p.Unlock()
	p.pool.Shuffle()
}

func (p *poolSync) Next() error {
	p.Lock()
	defer p.Unlock()
	return p.pool.Next()
}

func (p *poolSync) NextAsync() error {
	p.Lock()
	defer p.Unlock()
	return p.pool.NextAsync()
}
