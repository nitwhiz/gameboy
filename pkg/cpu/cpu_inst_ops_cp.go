package cpu

import "github.com/nitwhiz/gameboy/pkg/types"

func addCPHandlers() {
	// CP A, B
	H.add(0xB8, func(g types.GameBoy) (ticks byte) {
		return instCp(g.CPU(), g.CPU().BC().Hi())
	})

	// CP A, C
	H.add(0xB9, func(g types.GameBoy) (ticks byte) {
		return instCp(g.CPU(), g.CPU().BC().Lo())
	})

	// CP A, D
	H.add(0xBA, func(g types.GameBoy) (ticks byte) {
		return instCp(g.CPU(), g.CPU().DE().Hi())
	})

	// CP A, E
	H.add(0xBB, func(g types.GameBoy) (ticks byte) {
		return instCp(g.CPU(), g.CPU().DE().Lo())
	})

	// CP A, H
	H.add(0xBC, func(g types.GameBoy) (ticks byte) {
		return instCp(g.CPU(), g.CPU().HL().Hi())
	})

	// CP A, L
	H.add(0xBD, func(g types.GameBoy) (ticks byte) {
		return instCp(g.CPU(), g.CPU().HL().Lo())
	})

	// CP A, [HL]
	H.add(0xBE, func(g types.GameBoy) (ticks byte) {
		return instCp(g.CPU(), g.MMU().Read(g.CPU().HL().Val())) + 4
	})

	// CP A, A
	H.add(0xBF, func(g types.GameBoy) (ticks byte) {
		return instCp(g.CPU(), g.CPU().AF().Hi())
	})

	// CP A, n8
	H.add(0xFE, func(g types.GameBoy) (ticks byte) {
		return instCp(g.CPU(), g.CPU().Fetch8()) + 4
	})
}
