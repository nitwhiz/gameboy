package cpu

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/types"
)

func addControlHandlers() {
	// NOP
	H.add(0x00, func(g types.GameBoy) (ticks byte) {
		return 4
	})

	// STOP
	H.add(0x10, func(g types.GameBoy) (ticks byte) {
		g.CPU().SetHalt(true)

		g.CPU().Fetch8()

		g.MMU().Write(addr.DIV, 0x00)

		return 4
	})

	// HALT
	H.add(0x76, func(g types.GameBoy) (ticks byte) {
		haltBug := !g.CPU().IME() && ((g.MMU().Read(addr.IE) & g.MMU().Read(addr.IF) & 0x1F) != 0)

		// todo

		if haltBug {
			//g.HaltBug = 2
			g.CPU().SetHalt(false)
		} else {
			//g.HaltBug = 0
			g.CPU().SetHalt(true)
		}

		g.CPU().SetHalt(true)

		return 4
	})

	// DI
	H.add(0xF3, func(g types.GameBoy) (ticks byte) {
		g.CPU().SetIME(false)
		return 4
	})

	// EI
	H.add(0xFB, func(g types.GameBoy) (ticks byte) {
		g.CPU().SetIME(true)
		return 4
	})
}
