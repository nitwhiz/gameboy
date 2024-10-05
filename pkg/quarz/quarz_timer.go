package quarz

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/types"
)

type Timer struct {
	LastTACEnabled        bool
	TriggerOnFallingEdge  bool
	PostTIMAOverflowTicks int
	mmu                   types.MMU
}

func NewTimer(mmu types.MMU) *Timer {
	return &Timer{
		LastTACEnabled:        false,
		TriggerOnFallingEdge:  false,
		PostTIMAOverflowTicks: -1,
		mmu:                   mmu,
	}
}

func (t *Timer) Tick(ticks int) {
	tac := t.mmu.Read(addr.TAC)
	tima := t.mmu.Read(addr.TIMA)
	tma := t.mmu.Read(addr.TMA)

	nextTima := int(tima)

	if t.mmu.Memory().TimerCounter() == 0 && t.TriggerOnFallingEdge {
		t.TriggerOnFallingEdge = false
		nextTima++
	}

	tacEnabled := bits.IsTACEnabled(tac)
	clockSelect := bits.GetTACClockSelect(tac)
	tacMask := GetTACMask(clockSelect)

	if !tacEnabled && t.LastTACEnabled && t.mmu.Memory().TimerCounter()&tacMask != 0 {
		nextTima++
	}

	for range ticks {
		t.mmu.Memory().IncTimerCounter()

		if t.PostTIMAOverflowTicks > -1 {
			t.PostTIMAOverflowTicks++
		}

		if t.PostTIMAOverflowTicks == 4 {
			nextTima = int(tma)
			t.mmu.RequestInterrupt(addr.InterruptTimer)
		}

		if t.PostTIMAOverflowTicks == 5 {
			nextTima = int(tma)
			t.PostTIMAOverflowTicks = -1
		}

		if tacEnabled {
			if t.mmu.Memory().TimerCounter()&tacMask != 0 {
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

	t.mmu.Memory().WriteIO(t.mmu.Memory().AddrIO(addr.TIMA), byte(nextTima))
}
