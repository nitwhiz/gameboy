package cpu

import "github.com/nitwhiz/gameboy/pkg/types"

func addORHandlers() {
	// OR A, B
	H.add(0xB0, func(g types.GameBoy) (ticks byte) {
		return instOr(g.CPU(), g.CPU().BC().Hi())
	})

	// OR A, C
	H.add(0xB1, func(g types.GameBoy) (ticks byte) {
		return instOr(g.CPU(), g.CPU().BC().Lo())
	})

	// OR A, D
	H.add(0xB2, func(g types.GameBoy) (ticks byte) {
		return instOr(g.CPU(), g.CPU().DE().Hi())
	})

	// OR A, E
	H.add(0xB3, func(g types.GameBoy) (ticks byte) {
		return instOr(g.CPU(), g.CPU().DE().Lo())
	})

	// OR A, H
	H.add(0xB4, func(g types.GameBoy) (ticks byte) {
		return instOr(g.CPU(), g.CPU().HL().Hi())
	})

	// OR A, L
	H.add(0xB5, func(g types.GameBoy) (ticks byte) {
		return instOr(g.CPU(), g.CPU().HL().Lo())
	})

	// OR A, [HL]
	H.add(0xB6, func(g types.GameBoy) (ticks byte) {
		return instOr(g.CPU(), g.MMU().Read(g.CPU().HL().Val())) + 4
	})

	// OR A, A
	H.add(0xB7, func(g types.GameBoy) (ticks byte) {
		return instOr(g.CPU(), g.CPU().AF().Hi())
	})

	// OR A, n8
	H.add(0xF6, func(g types.GameBoy) (ticks byte) {
		return instOr(g.CPU(), g.CPU().Fetch8()) + 4
	})
}
