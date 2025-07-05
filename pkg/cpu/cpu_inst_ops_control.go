package cpu

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/types"
)

func addControlHandlers() {
	// NOP
	H.add(0x00, func(g types.GameBoy) {
	})

	// STOP
	H.add(0x10, func(g types.GameBoy) {
		g.CPU().SetHalt(true)

		// todo: is that correct?
		g.Fetch8()

		g.Write(addr.DIV, 0x00)
	})

	// HALT
	H.add(0x76, func(g types.GameBoy) {
		haltBug := !g.CPU().IME() && ((g.Read8(addr.IE) & g.Read8(addr.IF) & 0x1F) != 0)

		// todo: implement halt bug

		if haltBug {
			//g.HaltBug = 2
			g.CPU().SetHalt(false)
		} else {
			//g.HaltBug = 0
			g.CPU().SetHalt(true)
		}

		g.CPU().SetHalt(true)
	})

	// DI
	H.add(0xF3, func(g types.GameBoy) {
		g.CPU().SetIME(false)
	})

	// EI
	H.add(0xFB, func(g types.GameBoy) {
		c := g.CPU()

		if !c.IME() && !c.IMEToggle() {
			c.SetIMEToggle(true)
		}
	})
}
