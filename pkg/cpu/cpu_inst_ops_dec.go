package cpu

import "github.com/nitwhiz/gameboy/pkg/types"

func addDECHandlers() {
	// DEC BC
	H.add(0x0B, func(g types.GameBoy) (ticks byte) {
		return instDecReg(g.CPU().BC())
	})

	// DEC DE
	H.add(0x1B, func(g types.GameBoy) (ticks byte) {
		return instDecReg(g.CPU().DE())
	})

	// DEC HL
	H.add(0x2B, func(g types.GameBoy) (ticks byte) {
		return instDecReg(g.CPU().HL())
	})

	// DEC SP
	H.add(0x3B, func(g types.GameBoy) (ticks byte) {
		return instDecReg(g.CPU().SP())
	})

	// DEC B
	H.add(0x05, func(g types.GameBoy) (ticks byte) {
		return instDecRegHi(g.CPU(), g.CPU().BC())
	})

	// DEC C
	H.add(0x0D, func(g types.GameBoy) (ticks byte) {
		return instDecRegLo(g.CPU(), g.CPU().BC())
	})

	// DEC D
	H.add(0x15, func(g types.GameBoy) (ticks byte) {
		return instDecRegHi(g.CPU(), g.CPU().DE())
	})

	// DEC E
	H.add(0x1D, func(g types.GameBoy) (ticks byte) {
		return instDecRegLo(g.CPU(), g.CPU().DE())
	})

	// DEC H
	H.add(0x25, func(g types.GameBoy) (ticks byte) {
		return instDecRegHi(g.CPU(), g.CPU().HL())
	})

	// DEC L
	H.add(0x2D, func(g types.GameBoy) (ticks byte) {
		return instDecRegLo(g.CPU(), g.CPU().HL())
	})

	// DEC [HL]
	H.add(0x35, func(g types.GameBoy) (ticks byte) {
		t, r := instDec8(g.CPU(), g.MMU().Read(g.CPU().HL().Val()))

		g.MMU().Write(g.CPU().HL().Val(), r)

		return t + 8
	})

	// DEC A
	H.add(0x3D, func(g types.GameBoy) (ticks byte) {
		return instDecRegHi(g.CPU(), g.CPU().AF())
	})
}
