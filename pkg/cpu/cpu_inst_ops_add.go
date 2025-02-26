package cpu

import "github.com/nitwhiz/gameboy/pkg/types"

func addADDHandlers() {
	// ADD A, B
	H.add(0x80, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().BC().Hi(), false)
	})

	// ADD A, C
	H.add(0x81, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().BC().Lo(), false)
	})

	// ADD A, D
	H.add(0x82, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().DE().Hi(), false)
	})

	// ADD A, E
	H.add(0x83, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().DE().Lo(), false)
	})

	// ADD A, H
	H.add(0x84, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().HL().Hi(), false)
	})

	// ADD A, L
	H.add(0x85, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().HL().Lo(), false)
	})

	// ADD A, [HL]
	H.add(0x86, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.MMU().Read(g.CPU().HL().Val()), false) + 4
	})

	// ADD A, A
	H.add(0x87, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().AF().Hi(), false)
	})

	// ADD HL, BC
	H.add(0x09, func(g types.GameBoy) (ticks byte) {
		return instAdd16HL(g.CPU(), g.CPU().BC().Val())
	})

	// ADD HL, DE
	H.add(0x19, func(g types.GameBoy) (ticks byte) {
		return instAdd16HL(g.CPU(), g.CPU().DE().Val())
	})

	// ADD HL, HL
	H.add(0x29, func(g types.GameBoy) (ticks byte) {
		return instAdd16HL(g.CPU(), g.CPU().HL().Val())
	})

	// ADD HL, SP
	H.add(0x39, func(g types.GameBoy) (ticks byte) {
		return instAdd16HL(g.CPU(), g.CPU().SP().Val())
	})

	// ADD A, n8
	H.add(0xC6, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().Fetch8(), false) + 4
	})

	// ADD SP, e8
	H.add(0xE8, func(g types.GameBoy) (ticks byte) {
		return instAdd16Signed(g.CPU(), g.CPU().SP(), g.CPU().SP(), int8(g.CPU().Fetch8())) + 4
	})
}
