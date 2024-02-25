package inst

import (
	"github.com/nitwhiz/gameboy/pkg/gb"
)

func addSBCHandlers() {
	// SBC A, B
	h.add(0x98, func(g *gb.GameBoy) (ticks byte) {
		return instSub(g.CPU, g.CPU.BC.Hi(), true)
	})

	// SBC A, C
	h.add(0x99, func(g *gb.GameBoy) (ticks byte) {
		return instSub(g.CPU, g.CPU.BC.Lo(), true)
	})

	// SBC A, D
	h.add(0x9A, func(g *gb.GameBoy) (ticks byte) {
		return instSub(g.CPU, g.CPU.DE.Hi(), true)
	})

	// SBC A, E
	h.add(0x9B, func(g *gb.GameBoy) (ticks byte) {
		return instSub(g.CPU, g.CPU.DE.Lo(), true)
	})

	// SBC A, H
	h.add(0x9C, func(g *gb.GameBoy) (ticks byte) {
		return instSub(g.CPU, g.CPU.HL.Hi(), true)
	})

	// SBC A, L
	h.add(0x9D, func(g *gb.GameBoy) (ticks byte) {
		return instSub(g.CPU, g.CPU.HL.Lo(), true)
	})

	// SBC A, [HL]
	h.add(0x9E, func(g *gb.GameBoy) (ticks byte) {
		return instSub(g.CPU, g.MMU.Read(g.CPU.HL.Val()), true) + 4
	})

	// SBC A, A
	h.add(0x9F, func(g *gb.GameBoy) (ticks byte) {
		return instSub(g.CPU, g.CPU.AF.Hi(), true)
	})

	// SBC A, n8
	h.add(0xDE, func(g *gb.GameBoy) (ticks byte) {
		return instSub(g.CPU, g.Fetch8(), true) + 4
	})
}
