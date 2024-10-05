package gb

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
)

func addControlHandlers() {
	// NOP
	h.add(0x00, func(g *GameBoy) (ticks byte) {
		return 4
	})

	// STOP
	h.add(0x10, func(g *GameBoy) (ticks byte) {
		g.CPU.SetHalt(true)

		g.Fetch8()

		g.MMU.Memory().ResetTimerCounter()

		return 4
	})

	// HALT
	h.add(0x76, func(g *GameBoy) (ticks byte) {
		haltBug := !g.CPU.IME() && ((g.MMU.Read(addr.IE) & g.MMU.Read(addr.IF) & 0x1F) != 0)

		if haltBug {
			g.HaltBug = 2
			g.CPU.SetHalt(false)
		} else {
			g.HaltBug = 0
			g.CPU.SetHalt(true)
		}

		g.CPU.SetHalt(true)

		return 4
	})

	// DI
	h.add(0xF3, func(g *GameBoy) (ticks byte) {
		g.CPU.SetIME(false)
		return 4
	})

	// EI
	h.add(0xFB, func(g *GameBoy) (ticks byte) {
		g.CPU.SetIME(true)
		return 4
	})
}
