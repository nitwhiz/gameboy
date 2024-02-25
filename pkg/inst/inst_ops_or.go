package inst

import (
	"github.com/nitwhiz/gameboy/pkg/gb"
)

func addORHandlers() {
	// OR A, B
	h.add(0xB0, func(g *gb.GameBoy) (ticks byte) {
		return instOr(g.CPU, g.CPU.BC.Hi())
	})

	// OR A, C
	h.add(0xB1, func(g *gb.GameBoy) (ticks byte) {
		return instOr(g.CPU, g.CPU.BC.Lo())
	})

	// OR A, D
	h.add(0xB2, func(g *gb.GameBoy) (ticks byte) {
		return instOr(g.CPU, g.CPU.DE.Hi())
	})

	// OR A, E
	h.add(0xB3, func(g *gb.GameBoy) (ticks byte) {
		return instOr(g.CPU, g.CPU.DE.Lo())
	})

	// OR A, H
	h.add(0xB4, func(g *gb.GameBoy) (ticks byte) {
		return instOr(g.CPU, g.CPU.HL.Hi())
	})

	// OR A, L
	h.add(0xB5, func(g *gb.GameBoy) (ticks byte) {
		return instOr(g.CPU, g.CPU.HL.Lo())
	})

	// OR A, [HL]
	h.add(0xB6, func(g *gb.GameBoy) (ticks byte) {
		return instOr(g.CPU, g.MMU.Read(g.CPU.HL.Val())) + 4
	})

	// OR A, A
	h.add(0xB7, func(g *gb.GameBoy) (ticks byte) {
		return instOr(g.CPU, g.CPU.AF.Hi())
	})

	// OR A, n8
	h.add(0xF6, func(g *gb.GameBoy) (ticks byte) {
		return instOr(g.CPU, g.Fetch8()) + 4
	})
}
