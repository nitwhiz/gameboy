package quarz

import "sync"

type TickCounter[T int | float64] struct {
	Value T
	Wrap  T
	mu    *sync.RWMutex
}

func NewTickCounter[T int | float64](wrap T) *TickCounter[T] {
	return &TickCounter[T]{
		Value: 0,
		Wrap:  wrap,
		mu:    &sync.RWMutex{},
	}
}

func (t *TickCounter[T]) Increase(v T) (wrapped bool) {
	t.mu.Lock()
	defer t.mu.Unlock()

	wrapped = false
	t.Value += v

	if t.Value >= t.Wrap {
		wrapped = true
		t.Value -= t.Wrap
	}

	return wrapped
}

func (t *TickCounter[T]) Reset() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.Value = 0
}

func (t *TickCounter[T]) SetValue(v T) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.Value = v
}

func (t *TickCounter[T]) GetValue() T {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.Value
}
