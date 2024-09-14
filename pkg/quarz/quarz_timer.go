package quarz

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/mmu"
)

type Timer struct {
	LastTACEnabled        bool
	TriggerOnFallingEdge  bool
	PostTIMAOverflowTicks int
	MMU                   *mmu.MMU
}

func NewTimer(m *mmu.MMU) *Timer {
	return &Timer{
		LastTACEnabled:        false,
		TriggerOnFallingEdge:  false,
		PostTIMAOverflowTicks: -1,
		MMU:                   m,
	}
}

func (t *Timer) Tick(ticks int) {
	tac := t.MMU.Read(addr.TAC)
	tima := t.MMU.Read(addr.TIMA)
	tma := t.MMU.Read(addr.TMA)

	nextTima := int(tima)

	if t.MMU.TimerCounter == 0 && t.TriggerOnFallingEdge {
		t.TriggerOnFallingEdge = false
		nextTima++
	}

	tacEnabled := bits.IsTACEnabled(tac)
	clockSelect := bits.GetTACClockSelect(tac)
	tacMask := GetTACMask(clockSelect)

	if !tacEnabled && t.LastTACEnabled && t.MMU.TimerCounter&tacMask != 0 {
		nextTima++
	}

	for range ticks {
		t.MMU.TimerCounter++

		if t.PostTIMAOverflowTicks > -1 {
			t.PostTIMAOverflowTicks++
		}

		if t.PostTIMAOverflowTicks == 4 {
			nextTima = int(tma)
			t.MMU.RequestInterrupt(addr.InterruptTimer)
		}

		if t.PostTIMAOverflowTicks == 5 {
			nextTima = int(tma)
			t.PostTIMAOverflowTicks = -1
		}

		if tacEnabled {
			if t.MMU.TimerCounter&tacMask != 0 {
				t.TriggerOnFallingEdge = true
			} else if t.TriggerOnFallingEdge {
				t.TriggerOnFallingEdge = false
				nextTima++
			}
		}
	}

	t.LastTACEnabled = tacEnabled

	if nextTima > 0xFF {
		t.PostTIMAOverflowTicks = 0
		nextTima = nextTima - 0x100
	}

	t.MMU.SetTIMA(byte(nextTima))
}
