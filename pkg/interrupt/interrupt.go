package interrupt

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/cpu"
	"github.com/nitwhiz/gameboy/pkg/mmu"
	"github.com/nitwhiz/gameboy/pkg/stack"
)

func GetISR(i addr.InterruptType) uint16 {
	switch i {
	case addr.InterruptVBlank:
		return addr.ISRVBlank
	case addr.InterruptLCD:
		return addr.ISRLCD
	case addr.InterruptTimer:
		return addr.ISRTimer
	case addr.InterruptSerial:
		return addr.ISRSerial
	case addr.InterruptJoypad:
		return addr.ISRJoypad
	default:
		panic("missing interrupt type")
	}
}

type Manager struct {
	CPU *cpu.CPU
	MMU *mmu.MMU

	Stack *stack.Stack
}

func (m *Manager) Service() (ticks int) {
	requested := m.MMU.Read(addr.IF)
	enabled := m.MMU.Read(addr.IE)

	if requested&enabled == 0 {
		return
	} else {
		m.CPU.Halt = false
	}

	if m.CPU.IME {
		m.CPU.IME = false

		for i := 0; i < 5; i++ {
			t := addr.InterruptType(i)

			if bits.Test(requested, byte(t)) && bits.Test(enabled, byte(t)) {
				requested = bits.Reset(requested, byte(t))
				m.MMU.Write(addr.IF, requested)

				m.Stack.Push(m.CPU.PC.Val())
				m.CPU.PC.Set(GetISR(t))

				return 20
			}
		}
	}

	return 0
}
