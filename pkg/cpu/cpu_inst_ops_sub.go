package cpu

import "github.com/nitwhiz/gameboy/pkg/types"

func addSUBHandlers() {
	// SUB A, B
	H.add(0x90, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.CPU().BC().Hi(), false)
	})

	// SUB A, C
	H.add(0x91, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.CPU().BC().Lo(), false)
	})

	// SUB A, D
	H.add(0x92, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.CPU().DE().Hi(), false)
	})

	// SUB A, E
	H.add(0x93, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.CPU().DE().Lo(), false)
	})

	// SUB A, H
	H.add(0x94, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.CPU().HL().Hi(), false)
	})

	// SUB A, L
	H.add(0x95, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.CPU().HL().Lo(), false)
	})

	// SUB A, [HL]
	H.add(0x96, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.MMU().Read(g.CPU().HL().Val()), false) + 4
	})

	// SUB A, A
	H.add(0x97, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.CPU().AF().Hi(), false)
	})

	// SUB A, n8
	H.add(0xD6, func(g types.GameBoy) (ticks byte) {
		return instSub(g.CPU(), g.CPU().Fetch8(), false) + 4
	})
}
