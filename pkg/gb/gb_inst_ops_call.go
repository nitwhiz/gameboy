package gb

import (
	"github.com/nitwhiz/gameboy/pkg/types"
)

func addCALLHandlers() {
	// CALL NZ, a16
	h.add(0xC4, func(g *GameBoy) (ticks byte) {
		return instCallCond(g, types.FlagZ, false)
	})

	// CALL NC, a16
	h.add(0xD4, func(g *GameBoy) (ticks byte) {
		return instCallCond(g, types.FlagC, false)
	})

	// CALL Z, a16
	h.add(0xCC, func(g *GameBoy) (ticks byte) {
		return instCallCond(g, types.FlagZ, true)
	})

	// CALL C, a16
	h.add(0xDC, func(g *GameBoy) (ticks byte) {
		return instCallCond(g, types.FlagC, true)
	})

	// CALL a16
	h.add(0xCD, func(g *GameBoy) (ticks byte) {
		return instCall(g, g.Fetch16()) + 8
	})
}
