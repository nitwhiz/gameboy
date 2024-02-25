package inst

import (
	"github.com/nitwhiz/gameboy/pkg/cpu"
	"github.com/nitwhiz/gameboy/pkg/gb"
)

func addCALLHandlers() {
	// CALL NZ, a16
	h.add(0xC4, func(g *gb.GameBoy) (ticks byte) {
		return instCallCond(g, cpu.Z, false)
	})

	// CALL NC, a16
	h.add(0xD4, func(g *gb.GameBoy) (ticks byte) {
		return instCallCond(g, cpu.C, false)
	})

	// CALL Z, a16
	h.add(0xCC, func(g *gb.GameBoy) (ticks byte) {
		return instCallCond(g, cpu.Z, true)
	})

	// CALL C, a16
	h.add(0xDC, func(g *gb.GameBoy) (ticks byte) {
		return instCallCond(g, cpu.C, true)
	})

	// CALL a16
	h.add(0xCD, func(g *gb.GameBoy) (ticks byte) {
		return instCall(g, g.Fetch16()) + 8
	})
}
