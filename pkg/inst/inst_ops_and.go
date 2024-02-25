package inst

import (
	"github.com/nitwhiz/gameboy/pkg/gb"
)

func addANDHandlers() {
	// AND A, B
	h.add(0xA0, func(g *gb.GameBoy) (ticks byte) {
		return instAnd(g.CPU, g.CPU.BC.Hi())
	})

	// AND A, C
	h.add(0xA1, func(g *gb.GameBoy) (ticks byte) {
		return instAnd(g.CPU, g.CPU.BC.Lo())
	})

	// AND A, D
	h.add(0xA2, func(g *gb.GameBoy) (ticks byte) {
		return instAnd(g.CPU, g.CPU.DE.Hi())
	})

	// AND A, E
	h.add(0xA3, func(g *gb.GameBoy) (ticks byte) {
		return instAnd(g.CPU, g.CPU.DE.Lo())
	})

	// AND A, H
	h.add(0xA4, func(g *gb.GameBoy) (ticks byte) {
		return instAnd(g.CPU, g.CPU.HL.Hi())
	})

	// AND A, L
	h.add(0xA5, func(g *gb.GameBoy) (ticks byte) {
		return instAnd(g.CPU, g.CPU.HL.Lo())
	})

	// AND A, [HL]
	h.add(0xA6, func(g *gb.GameBoy) (ticks byte) {
		return instAnd(g.CPU, g.MMU.Read(g.CPU.HL.Val())) + 4
	})

	// AND A, A
	h.add(0xA7, func(g *gb.GameBoy) (ticks byte) {
		return instAnd(g.CPU, g.CPU.AF.Hi())
	})

	// AND A, n8
	h.add(0xE6, func(g *gb.GameBoy) (ticks byte) {
		return instAnd(g.CPU, g.Fetch8()) + 4
	})
}
