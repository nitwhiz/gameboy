package cpu

import (
	"github.com/nitwhiz/gameboy/pkg/types"
)

func addJRHandlers() {
	// JR NZ, e8
	H.add(0x20, func(g types.GameBoy) (ticks byte) {
		return instJrCond(g, types.FlagZ, false)
	})

	// JR NC, e8
	H.add(0x30, func(g types.GameBoy) (ticks byte) {
		return instJrCond(g, types.FlagC, false)
	})

	// JR Z, e8
	H.add(0x28, func(g types.GameBoy) (ticks byte) {
		return instJrCond(g, types.FlagZ, true)
	})

	// JR C, e8
	H.add(0x38, func(g types.GameBoy) (ticks byte) {
		return instJrCond(g, types.FlagC, true)
	})

	// JR e8
	H.add(0x18, func(g types.GameBoy) (ticks byte) {
		return instJr(g, g.CPU().Fetch8()) + 4
	})
}

func addJPHandlers() {
	// JP NZ, a16
	H.add(0xC2, func(g types.GameBoy) (ticks byte) {
		return instJpCond(g, types.FlagZ, false)
	})

	// JP NC, a16
	H.add(0xD2, func(g types.GameBoy) (ticks byte) {
		return instJpCond(g, types.FlagC, false)
	})

	// JP Z, a16
	H.add(0xCA, func(g types.GameBoy) (ticks byte) {
		return instJpCond(g, types.FlagZ, true)
	})

	// JP C, a16
	H.add(0xDA, func(g types.GameBoy) (ticks byte) {
		return instJpCond(g, types.FlagC, true)
	})

	// JP a18
	H.add(0xC3, func(g types.GameBoy) (ticks byte) {
		return instJp(g, g.CPU().Fetch16()) + 12
	})

	// JP HL
	H.add(0xE9, func(g types.GameBoy) (ticks byte) {
		return instJp(g, g.CPU().HL().Val())
	})
}
