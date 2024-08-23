package quarz

import "sync"

type TickCounter[T int | float64] struct {
	Value T
	Wrap  T
	mu    *sync.Mutex
}

func NewTickCounter[T int | float64](wrap T) *TickCounter[T] {
	return &TickCounter[T]{
		Value: 0,
		Wrap:  wrap,
		mu:    &sync.Mutex{},
	}
}

func (t *TickCounter[T]) Increase(v T) int {
	t.mu.Lock()
	defer t.mu.Unlock()

	wrapCount := 0
	t.Value += v

	if t.Value >= t.Wrap {
		wrapCount = (int)(t.Value / t.Wrap)
		t.Value -= t.Wrap * T(wrapCount)
	}

	return wrapCount
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
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.Value
}
