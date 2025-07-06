package cpu

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/types"
)

func addControlHandlers() {
	// NOP
	H.add(0x00, func(g types.GameBoy) {})

	// STOP
	H.add(0x10, func(g types.GameBoy) {
		g.FlushPendingTicks()

		g.Read8(g.CPU().PC().Val())

		joyp := g.MMU().Read(addr.JOYP)

		if joyp&0x30 != 0x30 {
			g.Input().SetAccessed(true)
		}

		exitByJoyp := joyp&0xF != 0xF
		immediateExit := exitByJoyp
		interruptPending := g.InterruptController().IE()&g.InterruptController().IF()&0x1F != 0

		if !exitByJoyp {
			if !immediateExit {
				// todo: dma run
			}

			g.MMU().Write(addr.DIV, 0)

			if !g.CPU().IME() {
				g.Timer().SetDivCycles(g.Timer().DivCycles() - 4)
			}

			g.SetStopped(true)
			// todo: block oam, vram and such
		}

		if !interruptPending {
			g.Fetch8()
		}

		if immediateExit {
			g.SetStopped(false)

			// todo: dma_cycles = 4
			// todo: dma_run

			// todo: unblock oam, vram and such

			if !interruptPending {
				// todo: dma_run

				g.SetHalted(true)
				g.SetJustHalted(true)
			}
		}
	})

	// HALT
	H.add(0x76, func(g types.GameBoy) {
		g.Read8(g.CPU().PC().Val())

		g.SetPendingTicks(0)

		iEnable := g.InterruptController().IE()
		iFlag := g.InterruptController().IF()

		// todo: 0x1F is ^unused bits
		if (iEnable & iFlag & 0x1F) != 0 {
			if g.CPU().IME() {
				g.SetHalted(false)
				g.CPU().PC().Set(g.CPU().PC().Val() - 1)
			} else {
				g.SetHalted(false)
				g.SetHaltBug(true)
			}
		} else {
			g.SetHalted(true)
		}

		g.SetJustHalted(true)
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
