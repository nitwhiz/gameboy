package interrupt

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/types"
)

func GetISR(i types.InterruptType) uint16 {
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
		return 0x00
	}
}

type Manager struct {
	cpu   types.CPU
	mmu   types.MMU
	stack types.Stack
}

func NewManager(cpu types.CPU, mmu types.MMU, stack types.Stack) *Manager {
	return &Manager{
		cpu:   cpu,
		mmu:   mmu,
		stack: stack,
	}
}

func (m *Manager) Service() (ticks int) {
	requested := m.mmu.Read(addr.IF)
	enabled := m.mmu.Read(addr.IE)

	if requested&enabled == 0 {
		return
	} else {
		m.cpu.SetHalt(false)
	}

	if m.cpu.IME() {
		m.cpu.SetIME(false)

		for i := 0; i < 5; i++ {
			t := types.InterruptType(i)

			if bits.Test(requested, byte(t)) && bits.Test(enabled, byte(t)) {
				requested = bits.Reset(requested, byte(t))
				m.mmu.Write(addr.IF, requested)

				m.stack.Push(m.cpu.PC().Val())
				m.cpu.PC().Set(GetISR(t))

				return 20
			}
		}
	}

	return 0
}
