package cpu

import "github.com/nitwhiz/gameboy/pkg/types"

func addXORHandlers() {
	// XOR A, B
	H.add(0xA8, func(g types.GameBoy) {
		instXor(g.CPU(), g.CPU().BC().Hi())
	})

	// XOR A, C
	H.add(0xA9, func(g types.GameBoy) {
		instXor(g.CPU(), g.CPU().BC().Lo())
	})

	// XOR A, D
	H.add(0xAA, func(g types.GameBoy) {
		instXor(g.CPU(), g.CPU().DE().Hi())
	})

	// XOR A, E
	H.add(0xAB, func(g types.GameBoy) {
		instXor(g.CPU(), g.CPU().DE().Lo())
	})

	// XOR A, H
	H.add(0xAC, func(g types.GameBoy) {
		instXor(g.CPU(), g.CPU().HL().Hi())
	})

	// XOR A, L
	H.add(0xAD, func(g types.GameBoy) {
		instXor(g.CPU(), g.CPU().HL().Lo())
	})

	// XOR A, [HL]
	H.add(0xAE, func(g types.GameBoy) {
		instXor(g.CPU(), g.Read8(g.CPU().HL().Val()))
	})

	// XOR A, A
	H.add(0xAF, func(g types.GameBoy) {
		instXor(g.CPU(), g.CPU().AF().Hi())
	})

	// XOR A, n8
	H.add(0xEE, func(g types.GameBoy) {
		instXor(g.CPU(), g.Fetch8())
	})
}
