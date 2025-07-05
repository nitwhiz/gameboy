package cpu

import "github.com/nitwhiz/gameboy/pkg/types"

func addINCHandlers() {
	// INC BC
	H.add(0x03, func(g types.GameBoy) {
		instIncReg(g.CPU().BC())
	})

	// INC DE
	H.add(0x13, func(g types.GameBoy) {
		instIncReg(g.CPU().DE())
	})

	// INC HL
	H.add(0x23, func(g types.GameBoy) {
		instIncReg(g.CPU().HL())
	})

	// INC SP
	H.add(0x33, func(g types.GameBoy) {
		instIncReg(g.CPU().SP())
	})

	// INC B
	H.add(0x04, func(g types.GameBoy) {
		instIncRegHi(g.CPU(), g.CPU().BC())
	})

	// INC C
	H.add(0x0C, func(g types.GameBoy) {
		instIncRegLo(g.CPU(), g.CPU().BC())
	})

	// INC D
	H.add(0x14, func(g types.GameBoy) {
		instIncRegHi(g.CPU(), g.CPU().DE())
	})

	// INC E
	H.add(0x1C, func(g types.GameBoy) {
		instIncRegLo(g.CPU(), g.CPU().DE())
	})

	// INC H
	H.add(0x24, func(g types.GameBoy) {
		instIncRegHi(g.CPU(), g.CPU().HL())
	})

	// INC L
	H.add(0x2C, func(g types.GameBoy) {
		instIncRegLo(g.CPU(), g.CPU().HL())
	})

	// INC [HL]
	H.add(0x34, func(g types.GameBoy) {
		g.Write(g.CPU().HL().Val(), instInc8(g.CPU(), g.Read8(g.CPU().HL().Val())))
	})

	// INC A
	H.add(0x3C, func(g types.GameBoy) {
		instIncRegHi(g.CPU(), g.CPU().AF())
	})
}
