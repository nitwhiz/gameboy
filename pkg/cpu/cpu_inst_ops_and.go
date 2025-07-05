package cpu

import "github.com/nitwhiz/gameboy/pkg/types"

func addANDHandlers() {
	// AND A, B
	H.add(0xA0, func(g types.GameBoy) {
		instAnd(g.CPU(), g.CPU().BC().Hi())
	})

	// AND A, C
	H.add(0xA1, func(g types.GameBoy) {
		instAnd(g.CPU(), g.CPU().BC().Lo())
	})

	// AND A, D
	H.add(0xA2, func(g types.GameBoy) {
		instAnd(g.CPU(), g.CPU().DE().Hi())
	})

	// AND A, E
	H.add(0xA3, func(g types.GameBoy) {
		instAnd(g.CPU(), g.CPU().DE().Lo())
	})

	// AND A, H
	H.add(0xA4, func(g types.GameBoy) {
		instAnd(g.CPU(), g.CPU().HL().Hi())
	})

	// AND A, L
	H.add(0xA5, func(g types.GameBoy) {
		instAnd(g.CPU(), g.CPU().HL().Lo())
	})

	// AND A, [HL]
	H.add(0xA6, func(g types.GameBoy) {
		instAnd(g.CPU(), g.Read8(g.CPU().HL().Val()))
	})

	// AND A, A
	H.add(0xA7, func(g types.GameBoy) {
		instAnd(g.CPU(), g.CPU().AF().Hi())
	})

	// AND A, n8
	H.add(0xE6, func(g types.GameBoy) {
		instAnd(g.CPU(), g.Fetch8())
	})
}
