package gb

import (
	"github.com/nitwhiz/gameboy/pkg/types"
)

func addRETHandlers() {
	// RET NZ
	h.add(0xC0, func(g *GameBoy) (ticks byte) {
		return instRetCond(g, types.FlagZ, false)
	})

	// RET NC
	h.add(0xD0, func(g *GameBoy) (ticks byte) {
		return instRetCond(g, types.FlagC, false)
	})

	// RET Z
	h.add(0xC8, func(g *GameBoy) (ticks byte) {
		return instRetCond(g, types.FlagZ, true)
	})

	// RET C
	h.add(0xD8, func(g *GameBoy) (ticks byte) {
		return instRetCond(g, types.FlagC, true)
	})

	// RET
	h.add(0xC9, func(g *GameBoy) (ticks byte) {
		return instRet(g) + 8
	})

	// RETI
	h.add(0xD9, func(g *GameBoy) (ticks byte) {
		g.CPU.SetIME(true)
		return instRet(g) + 8
	})
}
