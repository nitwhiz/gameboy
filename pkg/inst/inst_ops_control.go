package inst

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/gb"
)

func addControlHandlers() {
	// NOP
	h.add(0x00, func(g *gb.GameBoy) (ticks byte) {
		return 4
	})

	// STOP
	h.add(0x10, func(g *gb.GameBoy) (ticks byte) {
		g.CPU.Halt = true

		g.Fetch8()

		g.MMU.ResetTimer()

		return 4
	})

	// HALT
	h.add(0x76, func(g *gb.GameBoy) (ticks byte) {
		haltBug := !g.CPU.IME && ((g.MMU.Read(addr.IE) & g.MMU.Read(addr.IF) & 0x1F) != 0)

		if haltBug {
			g.HaltBug = 2
			g.CPU.Halt = false
		} else {
			g.HaltBug = 0
			g.CPU.Halt = true
		}

		g.CPU.Halt = true

		return 4
	})

	// DI
	h.add(0xF3, func(g *gb.GameBoy) (ticks byte) {
		g.CPU.IME = false
		return 4
	})

	// EI
	h.add(0xFB, func(g *gb.GameBoy) (ticks byte) {
		g.CPU.IME = true
		return 4
	})
}
