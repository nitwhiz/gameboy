package cpu

import (
	"github.com/nitwhiz/gameboy/pkg/types"
)

func addCALLHandlers() {
	// CALL NZ, a16
	H.add(0xC4, func(g types.GameBoy) (ticks byte) {
		return instCallCond(g, types.FlagZ, false)
	})

	// CALL NC, a16
	H.add(0xD4, func(g types.GameBoy) (ticks byte) {
		return instCallCond(g, types.FlagC, false)
	})

	// CALL Z, a16
	H.add(0xCC, func(g types.GameBoy) (ticks byte) {
		return instCallCond(g, types.FlagZ, true)
	})

	// CALL C, a16
	H.add(0xDC, func(g types.GameBoy) (ticks byte) {
		return instCallCond(g, types.FlagC, true)
	})

	// CALL a16
	H.add(0xCD, func(g types.GameBoy) (ticks byte) {
		return instCall(g, g.CPU().Fetch16()) + 8
	})
}
