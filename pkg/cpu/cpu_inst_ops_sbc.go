package cpu

import "github.com/nitwhiz/gameboy/pkg/types"

func addSBCHandlers() {
	// SBC A, B
	H.add(0x98, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.CPU().BC().Hi(), true)
	})

	// SBC A, C
	H.add(0x99, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.CPU().BC().Lo(), true)
	})

	// SBC A, D
	H.add(0x9A, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.CPU().DE().Hi(), true)
	})

	// SBC A, E
	H.add(0x9B, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.CPU().DE().Lo(), true)
	})

	// SBC A, H
	H.add(0x9C, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.CPU().HL().Hi(), true)
	})

	// SBC A, L
	H.add(0x9D, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.CPU().HL().Lo(), true)
	})

	// SBC A, [HL]
	H.add(0x9E, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.MMU().Read(g.CPU().HL().Val()), true) + 4
	})

	// SBC A, A
	H.add(0x9F, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.CPU().AF().Hi(), true)
	})

	// SBC A, n8
	H.add(0xDE, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.CPU().Fetch8(), true) + 4
	})
}
