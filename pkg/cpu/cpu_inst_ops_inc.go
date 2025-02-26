package cpu

import "github.com/nitwhiz/gameboy/pkg/types"

func addINCHandlers() {
	// INC BC
	H.add(0x03, func(g types.GameBoy) (ticks byte) {
		return instIncReg(g.CPU().BC())
	})

	// INC DE
	H.add(0x13, func(g types.GameBoy) (ticks byte) {
		return instIncReg(g.CPU().DE())
	})

	// INC HL
	H.add(0x23, func(g types.GameBoy) (ticks byte) {
		return instIncReg(g.CPU().HL())
	})

	// INC SP
	H.add(0x33, func(g types.GameBoy) (ticks byte) {
		return instIncReg(g.CPU().SP())
	})

	// INC B
	H.add(0x04, func(g types.GameBoy) (ticks byte) {
		return instIncRegHi(g.CPU(), g.CPU().BC())
	})

	// INC C
	H.add(0x0C, func(g types.GameBoy) (ticks byte) {
		return instIncRegLo(g.CPU(), g.CPU().BC())
	})

	// INC D
	H.add(0x14, func(g types.GameBoy) (ticks byte) {
		return instIncRegHi(g.CPU(), g.CPU().DE())
	})

	// INC E
	H.add(0x1C, func(g types.GameBoy) (ticks byte) {
		return instIncRegLo(g.CPU(), g.CPU().DE())
	})

	// INC H
	H.add(0x24, func(g types.GameBoy) (ticks byte) {
		return instIncRegHi(g.CPU(), g.CPU().HL())
	})

	// INC L
	H.add(0x2C, func(g types.GameBoy) (ticks byte) {
		return instIncRegLo(g.CPU(), g.CPU().HL())
	})

	// INC [HL]
	H.add(0x34, func(g types.GameBoy) (ticks byte) {
		t, r := instInc8(g.CPU(), g.MMU().Read(g.CPU().HL().Val()))

		g.MMU().Write(g.CPU().HL().Val(), r)

		return t + 8
	})

	// INC A
	H.add(0x3C, func(g types.GameBoy) (ticks byte) {
		return instIncRegHi(g.CPU(), g.CPU().AF())
	})
}
