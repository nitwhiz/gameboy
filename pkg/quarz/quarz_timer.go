package quarz

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/interrupt"
	"github.com/nitwhiz/gameboy/pkg/mmu"
	"log/slog"
)

type Timer struct {
	LastTACEnabled       bool
	TriggerOnFallingEdge bool
	LoadTMAIn            int
	TriggerInterrupt     bool
	MMU                  *mmu.MMU
	IM                   *interrupt.Manager
}

func NewTimer(m *mmu.MMU, i *interrupt.Manager) *Timer {
	return &Timer{
		LastTACEnabled:       false,
		TriggerOnFallingEdge: false,
		LoadTMAIn:            -1,
		TriggerInterrupt:     false,
		MMU:                  m,
		IM:                   i,
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

	if t.LoadTMAIn == -1 {
		t.MMU.TimerLock = false
	}

	tacEnabled := bits.IsTACEnabled(tac)
	clockSelect := bits.GetTACClockSelect(tac)
	tacMask, ok := TACMask[clockSelect]

	if !ok {
		slog.Error("missing clock speed", "clockSelect", clockSelect)
	}

	if !tacEnabled && t.LastTACEnabled && t.MMU.TimerCounter&tacMask != 0 {
		nextTima++
	}

	for range ticks {
		t.MMU.TimerCounter++

		if t.LoadTMAIn > -1 {
			t.LoadTMAIn--
		}

		if t.LoadTMAIn == 0 {
			t.LoadTMAIn = -1
			t.TriggerInterrupt = true
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
		t.MMU.TimerLock = true
		t.LoadTMAIn = 4
		nextTima = nextTima - 0x100
	}

	if t.TriggerInterrupt {
		t.TriggerInterrupt = false
		nextTima = int(tma)
		t.IM.Request(interrupt.Timer)
	}

	t.MMU.SetTIMA(byte(nextTima))
}
