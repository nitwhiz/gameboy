package cpu

import "github.com/nitwhiz/gameboy/pkg/types"

func addPOPHandlers() {
	// POP BC
	H.add(0xC1, func(g types.GameBoy) {
		instPopReg(g, g.CPU().BC())
	})

	// POP DE
	H.add(0xD1, func(g types.GameBoy) {
		instPopReg(g, g.CPU().DE())
	})

	// POP HL
	H.add(0xE1, func(g types.GameBoy) {
		instPopReg(g, g.CPU().HL())
	})

	// POP AF
	H.add(0xF1, func(g types.GameBoy) {
		instPopReg(g, g.CPU().AF())
	})
}

func addPUSHHandlers() {
	// PUSH BC
	H.add(0xC5, func(g types.GameBoy) {
		instPushReg(g, g.CPU().BC())
	})

	// PUSH DE
	H.add(0xD5, func(g types.GameBoy) {
		instPushReg(g, g.CPU().DE())
	})

	// PUSH HL
	H.add(0xE5, func(g types.GameBoy) {
		instPushReg(g, g.CPU().HL())
	})

	// PUSH AF
	H.add(0xF5, func(g types.GameBoy) {
		instPushReg(g, g.CPU().AF())
	})
}
