package broom

import (
	"sync"
)

type Broom struct {
	mu    sync.RWMutex
	funcs []func()
}

var broom = &Broom{}

func New() *Broom {
	return &Broom{}
}

func Gather(fn func()) {
	broom.Gather(fn)
}

func (b *Broom) Gather(fn func()) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.funcs = append(b.funcs, fn)
}

func Clean() {
	broom.Clean()
}

func (b *Broom) Clean() {
	b.mu.Lock()
	defer b.mu.Unlock()
	for i := len(b.funcs) - 1; i >= 0; i-- {
		b.funcs[i]()
	}
	b.funcs = nil
}
