package quarz

import (
	"github.com/nitwhiz/gameboy/pkg/bits"
	"log/slog"
)

type Timer struct {
	LastTACEnabled bool

	ForceTIMAIncrease    bool
	TriggerOnFallingEdge bool

	Counter uint16
}

func NewTimer() *Timer {
	return &Timer{
		LastTACEnabled:       false,
		ForceTIMAIncrease:    false,
		TriggerOnFallingEdge: false,
		Counter:              0xABCC,
	}
}

func (t *Timer) Update(ticks int, tac byte) int {
	increaseTima := 0

	if t.ForceTIMAIncrease {
		t.ForceTIMAIncrease = false
		increaseTima++
	}

	tacEnabled := bits.IsTACEnabled(tac)
	clockSelect := bits.GetTACClockSelect(tac)
	tacMask, ok := TACMask[bits.GetTACClockSelect(clockSelect)]

	if !ok {
		slog.Error("missing clock speed", "clockSelect", clockSelect)
	}

	for range ticks {
		t.Counter++

		if !tacEnabled && t.LastTACEnabled && t.Counter&tacMask > 0 {
			increaseTima++
		}

		if tacEnabled {
			if t.Counter&tacMask != 0 {
				t.TriggerOnFallingEdge = true
			} else if t.TriggerOnFallingEdge {
				t.TriggerOnFallingEdge = false
				increaseTima++
			}
		}
	}

	t.LastTACEnabled = tacEnabled

	return increaseTima
}

func (t *Timer) Reset() {
	t.Counter = 0

	if t.TriggerOnFallingEdge {
		t.TriggerOnFallingEdge = false
		t.ForceTIMAIncrease = true
	}
}

func (t *Timer) Div() byte {
	return byte((t.Counter & 0xFF00) >> 8)
}
