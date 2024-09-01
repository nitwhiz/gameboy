package interrupt

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/cpu"
	"github.com/nitwhiz/gameboy/pkg/interrupt_bus"
	"github.com/nitwhiz/gameboy/pkg/mmu"
	"github.com/nitwhiz/gameboy/pkg/stack"
)

func GetISR(i interrupt_bus.Type) uint16 {
	switch i {
	case interrupt_bus.VBlank:
		return addr.ISRVBlank
	case interrupt_bus.LCD:
		return addr.ISRLCD
	case interrupt_bus.Timer:
		return addr.ISRTimer
	case interrupt_bus.Serial:
		return addr.ISRSerial
	case interrupt_bus.Joypad:
		return addr.ISRJoypad
	default:
		panic("missing interrupt type")
	}
}

type Manager struct {
	CPU   *cpu.CPU
	MMU   *mmu.MMU
	IMBus *interrupt_bus.Bus

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
			t := interrupt_bus.Type(i)

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
