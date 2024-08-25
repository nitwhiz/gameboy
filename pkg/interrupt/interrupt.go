package interrupt

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/cpu"
	"github.com/nitwhiz/gameboy/pkg/mmu"
	"github.com/nitwhiz/gameboy/pkg/stack"
)

type Type byte

const (
	VBlank = Type(0)
	LCD    = Type(1)
	Timer  = Type(2)
	Serial = Type(3)
	Joypad = Type(4)
)

func GetISR(i Type) uint16 {
	switch i {
	case VBlank:
		return addr.ISRVBlank
	case LCD:
		return addr.ISRLCD
	case Timer:
		return addr.ISRTimer
	case Serial:
		return addr.ISRSerial
	case Joypad:
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

func (m *Manager) Request(t Type) {
	m.MMU.Write(addr.IF, bits.Set(m.MMU.Read(addr.IF), byte(t)))
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
			t := Type(i)

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
