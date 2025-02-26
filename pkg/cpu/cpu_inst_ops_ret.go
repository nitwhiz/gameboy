package cpu

import (
	"github.com/nitwhiz/gameboy/pkg/types"
)

func addRETHandlers() {
	// RET NZ
	H.add(0xC0, func(g types.GameBoy) (ticks byte) {
		return instRetCond(g, types.FlagZ, false)
	})

	// RET NC
	H.add(0xD0, func(g types.GameBoy) (ticks byte) {
		return instRetCond(g, types.FlagC, false)
	})

	// RET Z
	H.add(0xC8, func(g types.GameBoy) (ticks byte) {
		return instRetCond(g, types.FlagZ, true)
	})

	// RET C
	H.add(0xD8, func(g types.GameBoy) (ticks byte) {
		return instRetCond(g, types.FlagC, true)
	})

	// RET
	H.add(0xC9, func(g types.GameBoy) (ticks byte) {
		return instRet(g) + 8
	})

	// RETI
	H.add(0xD9, func(g types.GameBoy) (ticks byte) {
		g.CPU().SetIME(true)
		return instRet(g) + 8
	})
}
