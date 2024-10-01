package gb

import (
	"github.com/nitwhiz/gameboy/pkg/cpu"
)

func addRETHandlers() {
	// RET NZ
	h.add(0xC0, func(g *GameBoy) (ticks byte) {
		return instRetCond(g, cpu.Z, false)
	})

	// RET NC
	h.add(0xD0, func(g *GameBoy) (ticks byte) {
		return instRetCond(g, cpu.C, false)
	})

	// RET Z
	h.add(0xC8, func(g *GameBoy) (ticks byte) {
		return instRetCond(g, cpu.Z, true)
	})

	// RET C
	h.add(0xD8, func(g *GameBoy) (ticks byte) {
		return instRetCond(g, cpu.C, true)
	})

	// RET
	h.add(0xC9, func(g *GameBoy) (ticks byte) {
		return instRet(g) + 8
	})

	// RETI
	h.add(0xD9, func(g *GameBoy) (ticks byte) {
		g.CPU.IME = true
		return instRet(g) + 8
	})
}
