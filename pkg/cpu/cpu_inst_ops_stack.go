package cpu

import "github.com/nitwhiz/gameboy/pkg/types"

func addPOPHandlers() {
	// POP BC
	H.add(0xC1, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().Set(g.Stack().Pop())
		return 12
	})

	// POP DE
	H.add(0xD1, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().Set(g.Stack().Pop())
		return 12
	})

	// POP HL
	H.add(0xE1, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().Set(g.Stack().Pop())
		return 12
	})

	// POP AF
	H.add(0xF1, func(g types.GameBoy) (ticks byte) {
		g.CPU().AF().Set(g.Stack().Pop())
		return 12
	})
}

func addPUSHHandlers() {
	// PUSH BC
	H.add(0xC5, func(g types.GameBoy) (ticks byte) {
		g.Stack().Push(g.CPU().BC().Val())
		return 16
	})

	// PUSH DE
	H.add(0xD5, func(g types.GameBoy) (ticks byte) {
		g.Stack().Push(g.CPU().DE().Val())
		return 16
	})

	// PUSH HL
	H.add(0xE5, func(g types.GameBoy) (ticks byte) {
		g.Stack().Push(g.CPU().HL().Val())
		return 16
	})

	// PUSH AF
	H.add(0xF5, func(g types.GameBoy) (ticks byte) {
		g.Stack().Push(g.CPU().AF().Val())
		return 16
	})
}
